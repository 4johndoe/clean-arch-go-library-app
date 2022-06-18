package v1

import (
	"ca-library-app/internal/controller/http/dto"
	"ca-library-app/internal/domain/entity"
	book2 "ca-library-app/internal/domain/usecase/book"
	"context"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "/books"
)

type BookUsecase interface {
	GetByID(ctx context.Context, id string) entity.Book
	GetAll(ctx context.Context, limit, offset int) []entity.Book
	Create(ctx context.Context, dto book2.CreateBookDto) *entity.Book
}

type bookHandler struct {
	bookUsecase BookUsecase
}

func NewBookHandler(bookUsecase BookUsecase) *bookHandler {
	return &bookHandler{bookUsecase: bookUsecase}
}

func (h *bookHandler) Register(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBooks)
}

func (h *bookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//books := h.bookService.GetAllBooks(context.Background(), 0, 0)
	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)
}

func (h *bookHandler) CreateBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var d dto.CreateBookDto
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return //err
	}

	// validate

	// mapping

	useCaseDto := book2.CreateBookDto{
		Name:     "",
		Year:     0,
		AuthorID: "",
		GenreID:  "",
	}

	book := h.bookUsecase.Create(r.Context(), useCaseDto)
	//if err != nil {
	//	// 200
	//	return
	//}

	fmt.Println(book)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("book"))
}
