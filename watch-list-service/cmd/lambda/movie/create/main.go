package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"

	movieapp "github.com/maestre3d/cinephilia/watch-list-service/internal/application/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie/persistence"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/neutrinocorp/ddderr"
)

func main() {
	ctx := context.Background()
	movieRepo := persistence.NewInMemoryMovieRepository()
	movieHandler := movieapp.NewCreateCommandHandler(movieapp.NewCreator(movieRepo))

	movieId := uuid.New()
	userId, _ := gonanoid.ID(16)
	categoryId, _ := gonanoid.ID(16)

	err := movieHandler.Invoke(ctx, movieapp.CreateCommand{
		Id:          movieId.String(),
		DisplayName: "There will be blood",
		Description: "Directed by Paul Thomas Anderson. Lead actor: Daniel Day Lewis.",
		CategoryId:  categoryId,
		UserId:      userId,
	})
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}

	movieId = uuid.New()

	err = movieHandler.Invoke(ctx, movieapp.CreateCommand{
		Id:          movieId.String(),
		DisplayName: "Blade Runner 2049",
		Description: "Young Blade Runner K's discovery of a long-buried secret leads him to track down former Blade " +
			"Runner Rick Deckard, who's been missing for thirty years.",
		UserId: userId,
	})
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}

	queryHandler := movieapp.NewFindQueryHandler(movieapp.NewFinder(movieRepo))
	mov, err := queryHandler.Invoke(ctx, movieapp.FindQuery{Id: movieId.String()})
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}

	movieJSON, _ := json.Marshal(mov)
	log.Print(string(movieJSON))
}
