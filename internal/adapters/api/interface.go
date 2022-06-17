package api

import "github.com/julienschmidt/httprouter"

type Handler interface {
	Registry(router *httprouter.Router)
}
