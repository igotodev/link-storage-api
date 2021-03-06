package handler

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"link-storage-api/internal/service"
	"link-storage-api/internal/storage/model"
	"link-storage-api/pkg/logger"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	router    *chi.Mux
	service   service.ServiceImpl
	appLogger *logger.Logger
}

type answer struct {
	ID int `json:"id"`
}

type tokenAnswer struct {
	Token string `json:"token"`
}

type answerError struct {
	Message string `json:"message"`
}

type userReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func NewHandler(router *chi.Mux, service service.ServiceImpl, appLogger *logger.Logger) *Handler {
	return &Handler{
		router:    router,
		service:   service,
		appLogger: appLogger,
	}
}

func (h *Handler) RegisterRouting() *chi.Mux {
	h.router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-in", h.addContentType(h.User()))
		r.Post("/sign-up", h.addContentType(h.AddUser()))
	})

	h.router.Route("/api/v1/link/", func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)

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
			errorBadRequest(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		link, err := h.service.Link(ctx, id)
		if err != nil {
			errorNotFound(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}

		w.WriteHeader(200)

		err = json.NewEncoder(w).Encode(link)
		if err != nil {
			errorInternalError(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}
	}
}

func (h *Handler) allLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		links, err := h.service.AllLinks(ctx)
		if err != nil {
			errorNotFound(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}

		w.WriteHeader(200)

		err = json.NewEncoder(w).Encode(links)
		if err != nil {
			errorInternalError(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}
	}
}

func (h *Handler) addLink() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		var link model.Link

		err := json.NewDecoder(req.Body).Decode(&link)
		if err != nil {
			errorBadRequest(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}

		linkID, err := h.service.AddLink(ctx, link)
		if err != nil {
			errorConflict(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}

		answ := answer{ID: linkID}

		w.WriteHeader(200)

		err = json.NewEncoder(w).Encode(answ)
		if err != nil {
			errorInternalError(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}
	}
}

func (h *Handler) updateLink() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		idStr := chi.URLParam(req, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			errorBadRequest(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}

		var link model.Link

		err = json.NewDecoder(req.Body).Decode(&link)
		if err != nil {
			errorBadRequest(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}

		link.ID = id

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		idFromDB, err := h.service.UpdateLink(ctx, link)
		if err != nil {
			errorConflict(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}

		answ := answer{ID: idFromDB}

		w.WriteHeader(200)

		err = json.NewEncoder(w).Encode(answ)
		if err != nil {
			errorInternalError(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}
	}
}

func (h *Handler) deleteLink() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		idStr := chi.URLParam(req, "id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			errorBadRequest(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		w.WriteHeader(204)

		err = h.service.DeleteLink(ctx, id)
		if err != nil {
			errorNotFound(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}
	}
}

func (h *Handler) AddUser() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var userReq userReq

		err := json.NewDecoder(req.Body).Decode(&userReq)
		if err != nil {
			errorBadRequest(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}

		var user model.User

		user.Username = userReq.Username
		user.PasswordHash, err = generatePasswordHash(userReq.Password)
		if err != nil {
			errorInternalError(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}
		user.Email = userReq.Email
		user.Active = true

		defer req.Body.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		id, err := h.service.AddUser(ctx, user)
		if err != nil || id == 0 {
			errorNotFound(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}

		w.WriteHeader(200)

		jwt := generateJWT(id)
		ta := tokenAnswer{Token: jwt}

		err = json.NewEncoder(w).Encode(ta)
		if err != nil {
			errorInternalError(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}
	}
}

func (h *Handler) User() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var userReq userReq

		err := json.NewDecoder(req.Body).Decode(&userReq)
		if err != nil {
			errorBadRequest(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}

		defer req.Body.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		user, err := h.service.User(ctx, userReq.Username, userReq.Password)
		if err != nil || user.ID != 0 {
			errorNotFound(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}

		w.WriteHeader(200)

		jwt := generateJWT(user.ID)
		ta := tokenAnswer{Token: jwt}

		err = json.NewEncoder(w).Encode(ta)
		if err != nil {
			errorInternalError(w, h.appLogger)
			h.appLogger.Info(err)
			return
		}
	}
}

func errorBadRequest(w http.ResponseWriter, appLogger *logger.Logger) {
	w.WriteHeader(400)

	answ := answerError{
		Message: "bad request",
	}

	err := json.NewEncoder(w).Encode(answ)
	if err != nil {
		appLogger.Info(err)
		return
	}
}

func errorNotFound(w http.ResponseWriter, appLogger *logger.Logger) {
	w.WriteHeader(404)

	answ := answerError{
		Message: "not found",
	}

	err := json.NewEncoder(w).Encode(answ)
	if err != nil {
		appLogger.Info(err)
		return
	}
}

func errorInternalError(w http.ResponseWriter, appLogger *logger.Logger) {
	w.WriteHeader(500)

	answ := answerError{
		Message: "internal error",
	}

	err := json.NewEncoder(w).Encode(answ)
	if err != nil {
		appLogger.Info(err)
		return
	}
}

func errorConflict(w http.ResponseWriter, appLogger *logger.Logger) {

	w.WriteHeader(409)

	answ := answerError{
		Message: "conflict",
	}

	err := json.NewEncoder(w).Encode(answ)
	if err != nil {
		appLogger.Info(err)
		return
	}
}
