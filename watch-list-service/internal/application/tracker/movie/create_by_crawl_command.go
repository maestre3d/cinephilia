package movie

// CreateByCrawlCommand requests a movie creation by crawling technique
//	@Command
//	@Async
//	@DTO
type CreateByCrawlCommand struct {
	Id       string
	UserId   string
	CrawlUrl string
}
