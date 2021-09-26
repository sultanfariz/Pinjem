package requests

type KeyStruct struct {
	Key string `json:"key"`
}

type GetBookByISBN struct {
	Publisher     []string    `json:"publishers"`
	Title         string      `json:"title"`
	NumberOfPages uint        `json:"number_of_pages"`
	PublishDate   string      `json:"publish_date"`
	AuthorId      []KeyStruct `json:"authors"`
	WorkId        []KeyStruct `json:"works"`
	BookId        string      `json:"key"`
}

type GetBookByWorkId struct {
	Description string `json:"description"`
	Title       string `json:"title"`
}
