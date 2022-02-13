package handler

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"link-storage-api/internal/service"
	"link-storage-api/internal/storage/model"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	router  *chi.Mux
	service service.ServiceImpl
}

type answer struct {
	id int
}

type answerErr struct {
	error string
}

func NewHandler(router *chi.Mux, service service.ServiceImpl) *Handler {
	return &Handler{router: router, service: service}
}

func (h *Handler) RegisterRouting() *chi.Mux {
	h.router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-in", nil)
		r.Post("/sign-up", nil)
	})

	h.router.Route("/api/v1/this/", func(r chi.Router) {
		//r.Use() Auth jwt

		r.Get("/{id}", nil)
		r.Get("/", h.addContentType(h.allLinks()))
		r.Post("/", h.addContentType(h.addLink()))
		r.Put("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	return h.router
}

func (h *Handler) allLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		links, err := h.service.AllLinks(ctx)
		if err != nil {
			log.Println(err)
			return
		}

		jb, err := json.Marshal(links)
		if err != nil {
			log.Println(err)
			return
		}

		_, err = w.Write(jb)
		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(204)
	}
}

func (h *Handler) addLink() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var link model.Link

		err := json.NewDecoder(req.Body).Decode(&link)
		if err != nil {
			return
		}

		linkID, err := h.service.AddLink(ctx, link)
		if err != nil {
			log.Println(err)
			return
		}

		answ := answer{id: linkID}

		jb, err := json.Marshal(&answ)
		if err != nil {
			log.Println(err)
			return
		}

		_, err = w.Write(jb)
		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(200)
	}
}
