package dataloader

import (
	"context"
	"fmt"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/prigotti/cargo/common/pb"
)

// This test is not intended for automation
func TestLoadAndForwardJSONData(t *testing.T) {
	ctx := context.Background()

	dataCh := make(chan map[string]interface{})
	stopCh := make(chan struct{})

	go func() {
		handler(ctx, dataCh, stopCh)
	}()

	err := LoadAndForwardJSONData(ctx, "../../ports.json", dataCh, stopCh, 5)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func handler(
	ctx context.Context,
	data <-chan map[string]interface{},
	stop <-chan struct{},
) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("finishing handler")

			return
		case <-stop:
			return
		case d := <-data:
			o := &pb.Port{}
			mapstructure.Decode(d, o)
			fmt.Println(o)
		}
	}
}
