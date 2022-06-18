package author

import (
	"ca-library-app/internal/___backup/adapters/api"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	authorURL  = "/authors/:author_id"
	authorsURL = "/authors"
)

type handler struct {
	authorService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{authorService: service}
}

func (h handler) Register(router *httprouter.Router) {
	router.GET(authorsURL, h.GetAllAuthors)
}

func (h handler) GetAllAuthors(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//books := h.bookService.GetAllBooks(context.Background(), 0, 0)
	w.Write([]byte("authors"))
	w.WriteHeader(http.StatusOK)
}
