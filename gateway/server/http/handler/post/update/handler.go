package update

import (
	"encoding/json"
	"errors"
	"github.com/viciousvs/blog-microservices/gateway/utils"
	utils2 "github.com/viciousvs/blog-microservices/post/utils"
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
	w.Header().Set(utils.ContentType, utils.ApplJson)
	dto := DTO{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dto); err != nil {
		utils.ErrorHandler(w, err, http.StatusBadRequest)
		return
	}
	p := pbPost.UpdateRequest{
		UUID:    dto.UUID,
		Title:   dto.Title,
		Content: dto.Content,
	}
	if !utils.IsValidUUID(dto.UUID) {
		utils.ErrorHandler(w, utils.ErrInvalidUUID, http.StatusBadRequest)
		return
	}
	if dto.EmptyBody() {
		utils.ErrorHandler(w, utils.ErrEmptyBody, http.StatusBadRequest)
		return
	}
	if err := dto.Validate(); err != nil {
		utils.ErrorHandler(w, utils.ErrNotValid, http.StatusBadRequest)
		return
	}
	ur, err := h.PostService.Update(r.Context(), &p)
	if err != nil {
		if errors.Is(err, utils2.ErrNotFound) {
			utils.ErrorHandler(w, utils.ErrNotFound, http.StatusBadRequest)
			return
		}
		utils.ErrorHandler(w, utils.ErrDoesntUpdated, http.StatusUnprocessableEntity)
		return
	}
	resp := make(map[string]string)
	resp["uuid"] = ur.GetUUID()
	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		utils.ErrorHandler(w, err, http.StatusUnprocessableEntity)
		return
	}
}
