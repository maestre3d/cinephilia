package movie

// CreateByCrawlCommand requests a movie creation by crawling technique
//	@Command
//	@Async
//	@DTO
type CreateByCrawlCommand struct {
	CrawlUrl string
}
