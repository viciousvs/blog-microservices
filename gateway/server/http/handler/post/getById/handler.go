package getById

import (
	"encoding/json"
	"github.com/viciousvs/blog-microservices/gateway/utils"
	"github.com/viciousvs/blog-microservices/post/model/post"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
	"net/http"
)

type handler struct {
	PostService pbPost.PostServiceClient
}

func NewHandler(ps pbPost.PostServiceClient) *handler {
	return &handler{ps}
}

func (h *handler) GetById(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("id")
	if uuid == "" {
		utils.ErrorHandler(w, "empty uuid", http.StatusBadRequest)
		return
	}
	p, err := h.PostService.GetByID(r.Context(), &pbPost.GetByIdRequest{UUID: uuid})
	if err != nil {
		utils.ErrorHandler(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	res := &post.Post{
		UUID:      p.GetUUID(),
		Title:     p.GetTitle(),
		Content:   p.GetContent(),
		UpdatedAt: p.GetUpdatedAt(),
		CreatedAt: p.GetCreatedAt(),
	}
	err = json.NewEncoder(w).Encode(&res)
	w.Header().Set(utils.ContentType, utils.ApplJson)
	if err != nil {
		utils.ErrorHandler(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
}
