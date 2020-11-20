package main

import (
	"context"
	"log"

	movieapp "github.com/maestre3d/cinephilia/internal/application/tracker/movie"
	"github.com/maestre3d/cinephilia/internal/infrastructure/tracker/movie/persistence"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/neutrinocorp/ddderr"
)

func main() {
	ctx := context.Background()
	movieRepo := persistence.NewInMemoryMovieRepository()
	creator := movieapp.NewCreator(movieRepo)
	movieHandler := movieapp.NewCreateCommandHandler(creator)

	movieId, _ := gonanoid.ID(16)
	userId, _ := gonanoid.ID(16)

	err := movieHandler.Invoke(ctx, movieapp.CreateCommand{
		Id:          movieId,
		DisplayName: "There will be blood",
		Description: "",
		UserId:      userId,
	})
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}

	err = movieHandler.Invoke(ctx, movieapp.CreateCommand{
		Id:          movieId,
		DisplayName: "Thurman Show",
		Description: "",
		UserId:      userId,
	})
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}
}
