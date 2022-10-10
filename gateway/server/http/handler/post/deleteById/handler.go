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
	uuid := r.URL.Query().Get("id")
	if uuid == "" {
		utils.ErrorHandler(w, "empty uuid", http.StatusBadRequest)
		return
	}
	dr, err := h.PostService.DeleteByID(r.Context(), &pbPost.DeleteByIdRequest{UUID: uuid})
	if err != nil {
		utils.ErrorHandler(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	resp := make(map[string]string)
	resp["uuid"] = dr.GetUUID()
	w.Header().Set(utils.ContentType, utils.ApplJson)
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		utils.ErrorHandler(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
}
