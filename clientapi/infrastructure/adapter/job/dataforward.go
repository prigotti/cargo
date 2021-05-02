package job

import (
	"context"

	"github.com/prigotti/cargo/clientapi/application"
	"github.com/prigotti/cargo/pkg/dataloader"
)

// RunJSONFileDataForwarderJob receives a JSON file path and
// an application service and forwards JSON data from the file
// to it.
func RunJSONFileDataForwarderJob(
	ctx context.Context,
	portService application.PortService,
	path string,
) {
	dataCh := make(chan map[string]interface{})
	stopCh := make(chan struct{})

	go func() {
		portService.StreamCreateOrUpdate(ctx, dataCh, stopCh)
	}()

	go func() {
		dataloader.LoadAndForwardJSONData(ctx, path, dataCh, stopCh, 10)
	}()
}
