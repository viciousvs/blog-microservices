package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/viciousvs/blog-microservices/gateway/server/http/handler/post/add"
	"github.com/viciousvs/blog-microservices/gateway/server/http/handler/post/deleteById"
	"github.com/viciousvs/blog-microservices/gateway/server/http/handler/post/getAll"
	"github.com/viciousvs/blog-microservices/gateway/server/http/handler/post/getById"
	"github.com/viciousvs/blog-microservices/gateway/server/http/handler/post/update"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
	"net/http"
)

type Routes struct {
	PostService pbPost.PostServiceClient
}

func NewMux(postService pbPost.PostServiceClient) Routes {
	return Routes{PostService: postService}
}
func (rt Routes) InitPostRoutes() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	postsRouter := api.PathPrefix("/posts").Subrouter()

	getByIdH := getById.NewHandler(rt.PostService)
	postsRouter.Handle("", JsonMiddleware(getByIdH.GetById)).Methods(http.MethodGet)

	delH := deleteById.NewHandler(rt.PostService)
	postsRouter.Handle("", JsonMiddleware(delH.DeleteById)).Methods(http.MethodDelete)

	getAllH := getAll.NewHandler(rt.PostService)
	postsRouter.Handle("/all", JsonMiddleware(getAllH.GetAll)).Methods(http.MethodGet)

	addH := add.NewHandler(rt.PostService)
	postsRouter.Handle("/add", JsonMiddleware(addH.Create)).Methods(http.MethodPost)

	updateH := update.NewHandler(rt.PostService)
	postsRouter.Handle("/update", JsonMiddleware(updateH.UpdateById)).Methods(http.MethodPost)

	return r
}

func JsonMiddleware(handler http.HandlerFunc) http.Handler {
	return handlers.ContentTypeHandler(handler, "application/json")
}
