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
	w.Header().Set(utils.ContentType, utils.ApplJson)
	uuid := r.URL.Query().Get("id")
	if uuid == "" {
		utils.NewErrorResponse(w, http.StatusBadRequest, utils.ErrEmptyUUID.Error())
		return
	}
	if !utils.IsValidUUID(uuid) {
		utils.NewErrorResponse(w, http.StatusBadRequest, utils.ErrInvalidUUID.Error())
		return
	}
	p, err := h.PostService.GetByID(r.Context(), &pbPost.GetByIdRequest{UUID: uuid})
	if err != nil {
		utils.NewErrorResponse(w, http.StatusUnprocessableEntity, utils.ErrNotFound.Error())
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
	if err != nil {
		utils.NewErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
}
