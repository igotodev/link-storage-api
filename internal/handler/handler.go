package handler

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"link-storage-api/internal/service"
	"link-storage-api/internal/storage/model"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	router  *chi.Mux
	service service.ServiceImpl
}

type answer struct {
	ID int `json:"id"`
}

type answerError struct {
	Message string `json:"message"`
}

func NewHandler(router *chi.Mux, service service.ServiceImpl) *Handler {
	return &Handler{router: router, service: service}
}

func (h *Handler) RegisterRouting() *chi.Mux {
	h.router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-in", nil)
		r.Post("/sign-up", nil)
	})

	h.router.Route("/api/v1/link/", func(r chi.Router) {
		//r.Use() Auth jwt

		r.Get("/{id}", h.addContentType(h.link()))
		r.Get("/", h.addContentType(h.allLinks()))
		r.Post("/", h.addContentType(h.addLink()))
		r.Put("/{id}", h.addContentType(h.updateLink()))
		r.Delete("/{id}", h.addContentType(h.deleteLink()))
	})

	return h.router
}

func (h *Handler) link() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		idStr := chi.URLParam(req, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			errorBadRequest(w)
			log.Println(err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		link, err := h.service.Link(ctx, id)
		if err != nil {
			errorNotFound(w)
			log.Println(err)
			return
		}

		err = json.NewEncoder(w).Encode(link)
		if err != nil {
			errorInternalError(w)
			log.Println(err)
			return
		}

		w.WriteHeader(200)
	}
}

func (h *Handler) allLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		links, err := h.service.AllLinks(ctx)
		if err != nil {
			errorNotFound(w)
			log.Println(err)
			return
		}

		err = json.NewEncoder(w).Encode(links)
		if err != nil {
			errorInternalError(w)
			log.Println(err)
			return
		}

		w.WriteHeader(200)
	}
}

func (h *Handler) addLink() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var link model.Link

		err := json.NewDecoder(req.Body).Decode(&link)
		if err != nil {
			errorBadRequest(w)
			log.Println(err)
			return
		}

		linkID, err := h.service.AddLink(ctx, link)
		if err != nil {
			errorConflict(w)
			log.Println(err)
			return
		}

		answ := answer{ID: linkID}

		err = json.NewEncoder(w).Encode(answ)
		if err != nil {
			errorInternalError(w)
			log.Println(err)
			return
		}

		w.WriteHeader(200)
	}
}

func (h *Handler) updateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		idStr := chi.URLParam(req, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			errorBadRequest(w)
			log.Println(err)
			return
		}

		var link model.Link

		err = json.NewDecoder(req.Body).Decode(&link)
		if err != nil {
			errorBadRequest(w)
			log.Println(err)
			return
		}

		link.ID = id

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		idFromDB, err := h.service.UpdateLink(ctx, link)
		if err != nil {
			errorConflict(w)
			log.Println(err)
			return
		}

		answ := answer{ID: idFromDB}

		err = json.NewEncoder(w).Encode(answ)
		if err != nil {
			errorInternalError(w)
			log.Println(err)
			return
		}

		w.WriteHeader(200)
	}
}

func (h *Handler) deleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		idStr := chi.URLParam(req, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			errorBadRequest(w)
			log.Println(err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = h.service.DeleteLink(ctx, id)
		if err != nil {
			errorNotFound(w)
			log.Println(err)
			return
		}

		w.WriteHeader(204)
	}
}

func errorBadRequest(w http.ResponseWriter) {
	w.WriteHeader(400)

	answ := answerError{
		Message: "bad request",
	}

	err := json.NewEncoder(w).Encode(answ)
	if err != nil {
		log.Println(err)
		return
	}
}

func errorNotFound(w http.ResponseWriter) {
	w.WriteHeader(404)

	answ := answerError{
		Message: "not found",
	}

	err := json.NewEncoder(w).Encode(answ)
	if err != nil {
		log.Println(err)
		return
	}
}

func errorInternalError(w http.ResponseWriter) {
	w.WriteHeader(500)

	answ := answerError{
		Message: "internal error",
	}

	err := json.NewEncoder(w).Encode(answ)
	if err != nil {
		log.Println(err)
		return
	}
}

func errorConflict(w http.ResponseWriter) {

	w.WriteHeader(409)

	answ := answerError{
		Message: "conflict",
	}

	err := json.NewEncoder(w).Encode(answ)
	if err != nil {
		log.Println(err)
		return
	}
}
