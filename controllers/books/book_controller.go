package books

import (
	"Pinjem/businesses/books"
	"Pinjem/controllers"
	"Pinjem/controllers/books/requests"
	"Pinjem/controllers/books/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	Usecase books.Usecase
}

type Header struct {
	Cookie string `json:"cookie"`
}

func NewBookController(u books.Usecase) *BookController {
	return &BookController{
		Usecase: u,
	}
}

func (b *BookController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	books, err := b.Usecase.GetAll(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := make([]responses.BookResponse, len(books))
	for i, book := range books {
		response[i] = responses.BookResponse{
			ID:     book.Id,
			BookId: book.BookId,
			// WorkId:        book.WorkId,
			ISBN:          book.ISBN,
			Publisher:     book.Publisher,
			PublishDate:   book.PublishDate,
			Title:         book.Title,
			Description:   book.Description,
			Language:      book.Language,
			Picture:       book.Picture,
			MinDeposit:    book.MinDeposit,
			NumberOfPages: book.NumberOfPages,
			Status:        book.Status,
			CreatedAt:     book.CreatedAt,
			UpdatedAt:     book.UpdatedAt,
		}
	}
	return controllers.SuccessResponse(c, response)
}

func (u *BookController) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	bookId := c.Param("bookId")
	// bookIdParam := c.Param("bookId")
	// bookIdInt, _ := (strconv.Atoi(bookIdParam))
	// bookId := uint(bookIdInt)
	user, err := u.Usecase.GetById(ctx, bookId)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := responses.BookResponse{
		ID:     user.Id,
		BookId: user.BookId,
		// WorkId:        user.WorkId,
		ISBN:          user.ISBN,
		Publisher:     user.Publisher,
		PublishDate:   user.PublishDate,
		Title:         user.Title,
		Description:   user.Description,
		Language:      user.Language,
		Picture:       user.Picture,
		NumberOfPages: user.NumberOfPages,
		MinDeposit:    user.MinDeposit,
		Status:        user.Status,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}

	return controllers.SuccessResponse(c, response)
}

func (b *BookController) Create(c echo.Context) error {
	ctx := c.Request().Context()

	createdBook := requests.CreateBook{}
	c.Bind(&createdBook)

	// check if book already exist
	dbBook, err := b.Usecase.GetByISBN(ctx, createdBook.ISBN)
	if err != nil && err.Error() != "record not found" {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	if dbBook.ISBN != "" {
		return controllers.ErrorResponse(c, http.StatusForbidden, fmt.Errorf("ISBN already exist"))
	}

	// get book from google api by isbn
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=+isbn:%s", createdBook.ISBN)
	response, err := http.Get(url)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}
	responseData, _ := ioutil.ReadAll(response.Body)
	var bookReq requests.GetGoogleBookByISBN
	json.Unmarshal(responseData, &bookReq)

	if len(bookReq.Items) == 0 {
		return controllers.ErrorResponse(c, http.StatusNotFound, fmt.Errorf("Book not found"))
	}

	bookDomain := books.Domain{
		BookId:        bookReq.Items[0].Id,
		ISBN:          createdBook.ISBN,
		Publisher:     bookReq.Items[0].VolumeInfo.Publisher,
		PublishDate:   bookReq.Items[0].VolumeInfo.PublishedDate,
		Title:         bookReq.Items[0].VolumeInfo.Title,
		Description:   bookReq.Items[0].VolumeInfo.Description,
		Language:      bookReq.Items[0].VolumeInfo.Language,
		Picture:       bookReq.Items[0].VolumeInfo.ImageLinks.Thumbnail,
		NumberOfPages: bookReq.Items[0].VolumeInfo.NumberOfPages,
		MinDeposit:    createdBook.MinDeposit,
		Status:        createdBook.Status,
	}

	book, err := b.Usecase.Create(ctx, bookDomain)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
	}

	bookResponse := responses.BookResponse{
		ID:            book.Id,
		BookId:        book.BookId,
		ISBN:          book.ISBN,
		Publisher:     book.Publisher,
		PublishDate:   book.PublishDate,
		Title:         book.Title,
		Description:   book.Description,
		Language:      book.Language,
		Picture:       book.Picture,
		NumberOfPages: book.NumberOfPages,
		MinDeposit:    book.MinDeposit,
		Status:        book.Status,
		CreatedAt:     book.CreatedAt,
		UpdatedAt:     book.UpdatedAt,
	}

	return controllers.SuccessResponse(c, bookResponse)
}

// func (b *BookController) Create(c echo.Context) error {
// 	// ctx := c.Request().Context()

// 	minDepositBody := c.FormValue("minDeposit")
// 	statusBody := c.FormValue("status")
// 	minDeposit, err := strconv.Atoi(minDepositBody)
// 	if err != nil {
// 		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
// 	}
// 	status, err := strconv.ParseBool(statusBody)
// 	if err != nil {
// 		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	idParam := c.Param("isbn")
// 	// sbn := c.Param("userId")
// 	isbn := "9780140328721"
// 	log.Println(idParam)
// 	url := fmt.Sprintf("https://openlibrary.org/isbn/%s.json", isbn)
// 	log.Println(url)
// 	response, err := http.Get(url)
// 	// response, err := http.Get("https://openlibrary.org/books/OL7353617M.json")
// 	// response, err := http.Get("https://api.ipify.org?format=json")
// 	log.Println("---------------------")
// 	log.Println(err)
// 	log.Println(response)
// 	if err != nil {
// 		log.Println("---------------------")
// 		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
// 	}
// 	responseData, _ := ioutil.ReadAll(response.Body)
// 	var bookReq requests.GetBookByISBN
// 	json.Unmarshal(responseData, &bookReq)

// 	// parse authors and works id to array
// 	authorArr := []string{}
// 	for _, author := range bookReq.AuthorId {
// 		// author.Key = author.Key[:len(author.Key)-1]
// 		authorKeySplit := strings.Split(author.Key, "/")
// 		authorArr = append(authorArr, authorKeySplit[len(authorKeySplit)-1])
// 	}
// 	workArr := []string{}
// 	for _, work := range bookReq.WorkId {
// 		workKeySplit := strings.Split(work.Key, "/")
// 		workArr = append(workArr, workKeySplit[len(workKeySplit)-1])
// 	}
// 	bookKeySplit := strings.Split(bookReq.BookId, "/")
// 	bookReq.BookId = bookKeySplit[len(bookKeySplit)-1]

// 	// get book data by workId
// 	// getBookByWorkUrl := fmt.Sprintf("https://openlibrary.org/api/books?bibkeys=ISBN:%s&jscmd=data&format=json", bookReq.ISBN)
// 	getBookByWorkUrl := fmt.Sprintf("https://openlibrary.org/works/%s.json", workArr[0])
// 	log.Println("---------------------")
// 	response, err = http.Get(getBookByWorkUrl)
// 	if err != nil {
// 		return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
// 	}
// 	responseData, _ = ioutil.ReadAll(response.Body)
// 	var bookByWorkReq requests.GetBookByWorkId
// 	json.Unmarshal(responseData, &bookByWorkReq)
// 	log.Println(bookByWorkReq)

// 	bookDomain := books.Domain{
// 		BookId:        bookReq.BookId,
// 		WorkId:        workArr[0],
// 		ISBN:          isbn,
// 		Publisher:     bookReq.Publisher,
// 		PublishDate:   bookReq.PublishDate,
// 		Title:         bookReq.Title,
// 		Description:   bookByWorkReq.Description,
// 		NumberOfPages: bookReq.NumberOfPages,
// 		MinDeposit:    uint(minDeposit),
// 		Status:        status,
// 	}

// 	// book := new(books.Book)
// 	// if err := c.Bind(book); err != nil {
// 	// 	return controllers.ErrorResponse(c, http.StatusBadRequest, err)
// 	// }

// 	// if err := b.Usecase.Create(ctx, book); err != nil {
// 	// 	return controllers.ErrorResponse(c, http.StatusInternalServerError, err)
// 	// }

// 	// return controllers.SuccessResponse(c, string(responseData))
// 	// return controllers.SuccessResponse(c, responseData)
// 	// return controllers.SuccessResponse(c, bookReq)
// 	// return controllers.SuccessResponse(c, bookByWorkReq)
// 	return controllers.SuccessResponse(c, bookDomain)
// }
