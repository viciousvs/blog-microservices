package add

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
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(utils.ContentType, utils.ApplJson)
	//request validation

	dto := DTO{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dto); err != nil {
		utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := dto.Validate(); err != nil {
		utils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	p := pbPost.CreateRequest{
		Title:   dto.Title,
		Content: dto.Content,
	}

	cr, err := h.PostService.Create(r.Context(), &p)
	if err != nil {
		utils.NewErrorResponse(w, http.StatusUnprocessableEntity, utils.ErrDoesntCreated.Error())
		return
	}

	resp := make(map[string]string)
	resp["uuid"] = cr.GetUUID()
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		utils.NewErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
}
