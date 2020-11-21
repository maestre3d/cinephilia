package movie

// CreateCommand requests a Movie creation
//	@DTO
//	@Command
type CreateCommand struct {
	Id          string
	UserId      string
	CategoryId  string
	DirectorId  string
	DisplayName string
	Description string
	Year        int
	Picture     string
	WatchUrl    string
	CrawlUrl    string
}
