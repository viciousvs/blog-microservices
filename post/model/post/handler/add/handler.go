package add

import "github.com/viciousvs/blog-microservices/post/model/post"

type handler struct {
	repo post.Repository
}

func NewHandler(repo post.Repository) *handler {
	return &handler{repo: repo}
}

func (h handler) Handle() (string, error) {
	return h.repo.Create()
}
