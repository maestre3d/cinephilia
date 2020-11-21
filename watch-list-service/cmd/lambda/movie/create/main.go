package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/eefret/gomdb"
	"github.com/google/uuid"
	movieapp "github.com/maestre3d/cinephilia/watch-list-service/internal/application/tracker/movie"
	movieinfra "github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie/persistence"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/neutrinocorp/ddderr"
)

func main() {
	ctx := context.Background()
	movieRepo := persistence.NewInMemoryMovieRepository()
	movieCreator := movieapp.NewCreator(movieRepo)
	movieHandler := movieapp.NewCreateCommandHandler(movieCreator)

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
	directorId, _ := gonanoid.ID(16)
	err = movieHandler.Invoke(ctx, movieapp.CreateCommand{
		Id:          movieId.String(),
		DisplayName: "Blade Runner 2049",
		UserId:      userId,
		CategoryId:  categoryId,
		DirectorId:  directorId,
		Description: "Thirty years after the events of the first film, a new blade runner, LAPD Officer K " +
			"(Ryan Gosling), unearths a long-buried secret that has the potential to plunge what's left " +
			"of society into chaos. K's discovery leads him on a quest to find Rick Deckard (Harrison Ford), " +
			"a former LAPD blade runner who has been missing for 30 years.",
		Year:     2017,
		WatchUrl: "https://www.netflix.com/us/title/80185760",
		Picture:  "https://m.media-amazon.com/images/M/MV5BNzA1Njg4NzYxOV5BMl5BanBnXkFtZTgwODk5NjU3MzI@._V1_SX300.jpg",
	})
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}

	queryHandler := movieapp.NewFindQueryHandler(movieapp.NewFinder(movieRepo))
	mov, err := queryHandler.Invoke(ctx, movieapp.FindQuery{Id: execCrawl(ctx, movieCreator)})
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}

	movieJSON, _ := json.Marshal(mov)
	log.Print(string(movieJSON))
}

func execCrawl(ctx context.Context, creator *movieapp.Creator) string {
	handler := movieapp.NewCreateByCrawlCommandHandler(creator,
		movieinfra.NewImdbMovieCrawler(gomdb.Init("XXXX")))

	movieId := uuid.New()
	userId, _ := gonanoid.ID(16)
	err := handler.Invoke(ctx, movieapp.CreateByCrawlCommand{
		Id:       movieId.String(),
		UserId:   userId,
		CrawlUrl: "https://www.imdb.com/title/tt5727208/?ref_=ttls_li_tt", // Uncut Gems by Safdie Brothers
	})
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}

	return movieId.String()
}
