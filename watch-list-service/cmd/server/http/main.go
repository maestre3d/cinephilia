package main

import (
	"context"

	"github.com/maestre3d/cinephilia/watch-list-service/pkg/app/api"
)

func main() {
	httpApi := api.InitHTTP(context.Background())
	httpApi.Run()

	select {
	case <-httpApi.Done():
	}
}
