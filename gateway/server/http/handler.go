package http

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
	"net/http"
)

const JSON = "application/json"

func errorHandler(w http.ResponseWriter, err string, status int) {
	http.Error(w, err, status)
}

type Handler struct {
	PostService pbPost.PostServiceClient
}

func NewHandler(ps pbPost.PostServiceClient) *Handler {
	return &Handler{ps}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h.Home)
	apiRouter := r.PathPrefix("/api").Subrouter()

	postsRouter := apiRouter.PathPrefix("/posts").Subrouter()
	postsRouter.Handle("", handlers.ContentTypeHandler(http.HandlerFunc(h.GetByID), JSON)).Methods(http.MethodGet)
	postsRouter.Handle("", handlers.ContentTypeHandler(http.HandlerFunc(h.DeleteById), JSON)).Methods(http.MethodDelete)
	postsRouter.Handle("/all", handlers.ContentTypeHandler(http.HandlerFunc(h.GetAll), JSON)).Methods(http.MethodGet)
	postsRouter.Handle("/add", handlers.ContentTypeHandler(http.HandlerFunc(h.Create), JSON)).Methods(http.MethodPost)
	postsRouter.Handle("/update", handlers.ContentTypeHandler(http.HandlerFunc(h.Update), JSON)).Methods(http.MethodPost)

	return r
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>home</h1>"))
}
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("id")
	if uuid == "" {
		errorHandler(w, "emty uuid", http.StatusBadRequest)
		return
	}
	p, err := h.PostService.GetByID(r.Context(), &pbPost.GetByIdRequest{UUID: uuid})
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&p)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := h.PostService.GetAll(r.Context(), &pbPost.GetAllRequest{})
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&posts)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var p pbPost.CreateRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := make(map[string]string)
	cr, err := h.PostService.Create(r.Context(), &p)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp["uuid"] = cr.GetUUID()
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	var p pbPost.UpdateRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		errorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp := make(map[string]string)
	ur, err := h.PostService.Update(r.Context(), &p)
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp["uuid"] = ur.GetUUID()
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (h *Handler) DeleteById(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("id")
	if uuid == "" {
		errorHandler(w, "emty uuid", http.StatusBadRequest)
		return
	}
	dr, err := h.PostService.DeleteByID(r.Context(), &pbPost.DeleteByIdRequest{UUID: uuid})
	if err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := make(map[string]string)
	resp["uuid"] = dr.GetUUID()
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		errorHandler(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
