package main

import (
	"context"
	"time"

	"github.com/maestre3d/cinephilia/watch-list-service/pkg/app/api"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	httpApi := api.InitHTTP(ctx)
	httpApi.Run()
}
