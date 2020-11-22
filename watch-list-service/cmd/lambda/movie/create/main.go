package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/webscrap"

	"github.com/google/uuid"
	movieapp "github.com/maestre3d/cinephilia/watch-list-service/internal/application/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/domain/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/bus"
	movieinfra "github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie"
	"github.com/maestre3d/cinephilia/watch-list-service/internal/infrastructure/tracker/movie/persistence"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/neutrinocorp/ddderr"
)

func main() {
	ctx := context.Background()
	movieRepo := persistence.NewInMemoryMovieRepository()
	commandBus := initSyncCommandBus(movieRepo)

	go execCreate(ctx, commandBus)
	movieId := execCrawl(ctx, commandBus)
	execFind(ctx, initSyncQueryBus(movieRepo), movieId)
}

func initSyncCommandBus(repo movie.Repository) domain.CommandBus {
	creator := movieapp.NewCreator(repo)
	commandBus := bus.NewInMemorySyncCommand()
	imdbCollector, err := webscrap.NewCollyImdbCollector()
	if err != nil {
		log.Fatal(err)
	}

	err = commandBus.RegisterHandler(movieapp.CreateByCrawlCommand{}, movieapp.NewCreateByCrawlCommandHandler(creator,
		movieinfra.NewImdbMovieCrawler(movieinfra.NewImdbWebScrapper(imdbCollector))))
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}

	err = commandBus.RegisterHandler(movieapp.CreateCommand{}, movieapp.NewCreateCommandHandler(creator))
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}

	return commandBus
}

func initSyncQueryBus(repo movie.Repository) domain.QueryBus {
	queryBus := bus.NewInMemorySyncQuery()
	err := queryBus.RegisterHandler(movieapp.FindQuery{},
		movieapp.NewFindQueryHandler(movieapp.NewFinder(repo)))
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}

	return queryBus
}

func execCreate(ctx context.Context, commandBus domain.CommandBus) {
	movieId := uuid.New()
	userId, _ := gonanoid.ID(16)
	categoryId, _ := gonanoid.ID(16)

	err := commandBus.Dispatch(ctx, movieapp.CreateCommand{
		MovieId:     movieId.String(),
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
	err = commandBus.Dispatch(ctx, movieapp.CreateCommand{
		MovieId:     movieId.String(),
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
}

func execCrawl(ctx context.Context, commandBus domain.CommandBus) string {
	movieId, _ := gonanoid.ID(16)
	userId, _ := gonanoid.ID(16)
	err := commandBus.Dispatch(ctx, movieapp.CreateByCrawlCommand{
		MovieId:  movieId,
		UserId:   userId,
		CrawlUrl: "https://www.imdb.com/title/tt5727208/?ref_=ttls_li_tt", // Uncut Gems by Safdie Brothers
	})
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}

	return movieId
}

func execFind(ctx context.Context, queryBus domain.QueryBus, id string) {
	mov, err := queryBus.Ask(ctx, movieapp.FindQuery{MovieId: id})
	if err != nil {
		log.Fatal(ddderr.GetDescription(err))
	}
	res := mov.(*movieapp.MovieResponse)
	log.Print(res.Id)

	movieJSON, _ := json.Marshal(mov)
	log.Print(string(movieJSON))
}
