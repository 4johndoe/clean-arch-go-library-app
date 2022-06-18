package v1

import (
	"ca-library-app/internal/controller/http/dto"
	"ca-library-app/internal/domain/entity"
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	genreURL  = "/genres/:genre_id"
	genresURL = "/genres"
)

type GenreUsecase interface {
	GetByID(ctx context.Context, id string) entity.Genre
	GetAll(ctx context.Context, limit, offset int) []entity.Genre
	Create(ctx context.Context, dto dto.CreateGenreDto) *entity.Genre
}

type genreHandler struct {
	genreUsecase GenreUsecase
}

func NewGenreHandler(genreUsecase GenreUsecase) *genreHandler {
	return &genreHandler{genreUsecase: genreUsecase}
}

func (h genreHandler) Register(router *httprouter.Router) {
	router.GET(genresURL, h.GetAllAuthors)
}

func (h genreHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//books := h.bookService.GetAllBooks(context.Background(), 0, 0)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("genres"))
}
