package deleteById

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
func (h *handler) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(utils.ContentType, utils.ApplJson)
	uuid := r.URL.Query().Get("id")
	if uuid == "" {
		utils.ErrorHandler(w, utils.ErrEmptyUUID, http.StatusBadRequest)
		return
	}
	if !utils.IsValidUUID(uuid) {
		utils.ErrorHandler(w, utils.ErrInvalidUUID, http.StatusBadRequest)
		return
	}
	dr, err := h.PostService.DeleteByID(r.Context(), &pbPost.DeleteByIdRequest{UUID: uuid})
	if err != nil {
		utils.ErrorHandler(w, utils.ErrNotFound, http.StatusUnprocessableEntity)
		return
	}

	resp := make(map[string]string)
	resp["uuid"] = dr.GetUUID()
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		utils.ErrorHandler(w, err, http.StatusUnprocessableEntity)
		return
	}
}
