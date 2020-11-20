package main

import (
	"context"
	"encoding/json"
	"log"

	movieapp "github.com/maestre3d/cinephilia/watch-list-service/internal/application/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie/persistence"
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
	categoryId, _ := gonanoid.ID(16)

	err := movieHandler.Invoke(ctx, movieapp.CreateCommand{
		Id:          movieId,
		DisplayName: "There will be blood",
		Description: "Directed by Paul Thomas Anderson. Lead actor: Daniel Day Lewis.",
		CategoryId:  categoryId,
		UserId:      userId,
	})
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}

	queryHandler := movieapp.NewFindQueryHandler(movieapp.NewFinder(movieRepo))
	mov, err := queryHandler.Invoke(ctx, movieapp.FindQuery{Id: movieId})
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}

	movieJSON, _ := json.Marshal(mov)
	log.Print(string(movieJSON))
	/*
		err = movieHandler.Invoke(ctx, movieapp.CreateCommand{
			Id:          movieId,
			DisplayName: "Thurman Show",
			Description: "",
			UserId:      userId,
		})
		if err != nil {
			log.Fatal(ddderr.GetDescription(err))
		}*/
}
