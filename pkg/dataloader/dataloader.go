package dataloader

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

var ErrDataForwardingTimeout = errors.New("json data forwarding timeout")

// LoadJSONData takes a path to a file containing JSON data
// and forwards the value of each token as a map[string]interface{}
// to the designated data channel until it is finished, when it
// sends a stop signal to the stop channel.
//
// The function also stores the token key in the map with key "id".
func LoadAndForwardJSONData(
	ctx context.Context,
	path string,
	data chan<- map[string]interface{},
	stop chan<- struct{},
	timeout int,
) error {
	fmt.Println("starting JSON data loader job")

	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)

		return err
	}

	defer jsonFile.Close()

	jsonDecoder := json.NewDecoder(jsonFile)
	_, err = jsonDecoder.Token()
	if err != nil {
		fmt.Println(err)

		return err
	}

	for jsonDecoder.More() {
		select {
		case <-ctx.Done():
			fmt.Println("terminating JSON forwarding")

			return nil
		default:
			token, err := jsonDecoder.Token()
			if err != nil {
				return err
			}

			m := make(map[string]interface{})
			err = jsonDecoder.Decode(&m)
			if err != nil {
				return err
			}

			m["id"] = token

			// Just in case the consumer fails
			select {
			case data <- m:
				continue
			case <-time.After(time.Duration(timeout) * time.Second):
				fmt.Println("JSON data forwarding timeout")

				return ErrDataForwardingTimeout
			}
		}
	}

	fmt.Println("finished loading JSON data from document")

	stop <- struct{}{}

	return nil
}
