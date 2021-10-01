package requests

type KeyStruct struct {
	Key string `json:"key"`
}

// type GetBookByISBN struct {
// 	Publisher     []string    `json:"publishers"`
// 	Title         string      `json:"title"`
// 	NumberOfPages uint        `json:"number_of_pages"`
// 	PublishDate   string      `json:"publish_date"`
// 	AuthorId      []KeyStruct `json:"authors"`
// 	WorkId        []KeyStruct `json:"works"`
// 	BookId        string      `json:"key"`
// }

// type GetGoogleBookByISBN struct {
// 	Items []struct {
// 		Id         string `json:"id"`
// 		VolumeInfo struct {
// 			Title         string   `json:"title"`
// 			Authors       []string `json:"authors"`
// 			Publisher     string   `json:"publisher"`
// 			PublishedDate string   `json:"publishedDate"`
// 			Description   string   `json:"description"`
// 			NumberOfPages uint     `json:"pageCount"`
// 			Language      string   `json:"language"`
// 			ImageLinks    struct {
// 				Thumbnail string `json:"thumbnail"`
// 			} `json:"imageLinks"`
// 		} `json:"volumeInfo"`
// 	} `json:"items"`
// }

// type GetBookByWorkId struct {
// 	Description string `json:"description"`
// 	Title       string `json:"title"`
// }

type CreateOrder struct {
	// Books []string `json:"books"`
	Books []int `json:"books"`
}
