package getAll

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
func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	posts, err := h.PostService.GetAll(r.Context(), &pbPost.GetAllRequest{})
	if err != nil {
		utils.ErrorHandler(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	var resPosts = make([]*post.Post, 0, len(posts.Posts))
	for _, p := range posts.Posts {
		rp := &post.Post{
			UUID:      p.GetUUID(),
			Title:     p.GetTitle(),
			Content:   p.GetContent(),
			UpdatedAt: p.GetUpdatedAt(),
			CreatedAt: p.GetCreatedAt(),
		}
		resPosts = append(resPosts, rp)
	}
	err = json.NewEncoder(w).Encode(&resPosts)
	w.Header().Set(utils.ContentType, utils.ApplJson)
	if err != nil {
		utils.ErrorHandler(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
}
