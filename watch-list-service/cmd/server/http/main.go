package main

import (
	"context"
	"log"

	"github.com/maestre3d/cinephilia/watch-list-service/pkg/app/api"
)

func main() {
	ctx := context.Background()
	httpApi := api.InitHTTP(ctx)
	err := httpApi.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}

	select {
	case <-ctx.Done():
	}
}
