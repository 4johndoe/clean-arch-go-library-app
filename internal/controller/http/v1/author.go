package v1

import (
	"ca-library-app/internal/controller/http/dto"
	"ca-library-app/internal/domain/entity"
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	authorURL  = "/authors/:author_id"
	authorsURL = "/authors"
)

type AuthorUsecase interface {
	GetByID(ctx context.Context, id string) entity.Author
	GetAll(ctx context.Context, limit, offset int) []entity.Author
	Create(ctx context.Context, dto dto.CreateAuthorDto) *entity.Author
}

type authorHandler struct {
	authorUsecase AuthorUsecase
}

func NewAuthorHandler(authorUsecase AuthorUsecase) *authorHandler {
	return &authorHandler{authorUsecase: authorUsecase}
}

func (h authorHandler) Register(router *httprouter.Router) {
	router.GET(authorsURL, h.GetAll)
}

func (h authorHandler) GetAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//books := h.bookService.GetAllBooks(context.Background(), 0, 0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("authors"))
}
