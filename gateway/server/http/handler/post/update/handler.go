package update

import (
	"encoding/json"
	"github.com/viciousvs/blog-microservices/gateway/utils"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
	"net/http"
)

type handler struct {
	PostService pbPost.PostServiceClient
}

func NewHandler(ps pbPost.PostServiceClient) *handler {
	return &handler{ps}
}
func (h *handler) UpdateById(w http.ResponseWriter, r *http.Request) {
	dto := DTO{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dto); err != nil {
		utils.ErrorHandler(w, err.Error(), http.StatusBadRequest)
		return
	}
	p := pbPost.UpdateRequest{
		UUID:    dto.UUID,
		Title:   dto.Title,
		Content: dto.Content,
	}
	ur, err := h.PostService.Update(r.Context(), &p)
	if err != nil {
		utils.ErrorHandler(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	resp := make(map[string]string)
	resp["uuid"] = ur.GetUUID()
	w.Header().Set(utils.ContentType, utils.ApplJson)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		utils.ErrorHandler(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
}
