package rest

import (
	"context"
	"net/http"

	"github.com/DexScen/ApiLinkShortener/internal/domain"
	"github.com/gorilla/mux"
)

type Links interface {
	GetLongFromShort(ctx context.Context, link domain.Link) (domain.Link, error)
	GetShortFromLong(ctx context.Context, link domain.Link) (domain.Link, error)
}

type Handler struct{
	linksService Links
}

func NewHandler(links Links) *Handler {
	return &Handler{
		linksService: links,
	}
}

func (h *Handler) InitRouter() *mux.Router{
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	links := r.PathPrefix("/links").Subrouter()
	{
		links.HandleFunc("", h.getLongFromShort).Methods(http.MethodGet)
		links.HandleFunc("", h.getShortFromLong).Methods(http.MethodGet)
	}

	return r
}

func (h *Handler) getShortFromLong(w http.ResponseWriter, r *http.Request) {

	link := &domain.Link{
		ID: 0,
		LongLink: ,
		Shortlink: nil,
		Created: nil,
	}

	link, err := h.linksService.GetShortFromLong(context.TODO(), )
}

func (h *Handler) getLongFromShort(w http.ResponseWriter, r *http.Request) {
	//todo
}

func getLinkFromRequest(r *http.Request) (string, error){
	
}