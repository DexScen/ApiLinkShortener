package rest

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/DexScen/ApiLinkShortener/internal/domain"
	"github.com/gorilla/mux"
)

type Links interface {
	GetByShortLink(ctx context.Context, link *domain.Link) error
	GetByLongLink(ctx context.Context, link *domain.Link) error
}

type Handler struct {
	linksService Links
}

func NewHandler(links Links) *Handler {
	return &Handler{
		linksService: links,
	}
}

func (h *Handler) InitRouter() *mux.Router {
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
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp string
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	link := &domain.Link{
		ID:        0,
		LongLink:  inp,
		ShortLink: "",
		Created:   time.Time{},
	}

	err = h.linksService.GetByLongLink(context.TODO(), link)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	final, err := json.Marshal(*link)
	if err != nil {
		log.Println("getShortFromLong() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(final)
}

func (h *Handler) getLongFromShort(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var inp string
	if err = json.Unmarshal(reqBytes, &inp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	link := &domain.Link{
		ID:        0,
		LongLink:  "",
		ShortLink: inp,
		Created:   time.Time{},
	}

	err = h.linksService.GetByShortLink(context.TODO(), link)
	if err != nil {
		log.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	final, err := json.Marshal(*link)
	if err != nil {
		log.Println("getLongFromShort() error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(final)
}
