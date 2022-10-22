package routes

import (
	"github.com/gorilla/mux"
	"github.com/viciousvs/blog-microservices/gateway/server/http/handler/post/add"
	"github.com/viciousvs/blog-microservices/gateway/server/http/handler/post/deleteById"
	"github.com/viciousvs/blog-microservices/gateway/server/http/handler/post/getAll"
	"github.com/viciousvs/blog-microservices/gateway/server/http/handler/post/getById"
	"github.com/viciousvs/blog-microservices/gateway/server/http/handler/post/update"
	"github.com/viciousvs/blog-microservices/gateway/server/http/middlewares"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
	"net/http"
)

const oAuthServerAddr = "localhost:50054"

type Routes struct {
	PostService pbPost.PostServiceClient
}

func NewMux(postService pbPost.PostServiceClient) Routes {
	return Routes{PostService: postService}
}
func (rt Routes) InitPostRoutes() *mux.Router {
	r := mux.NewRouter()
	apiv1 := r.PathPrefix("/v1").Subrouter()
	apiv1.Use(middlewares.JsonMiddleware)
	postsRouter := apiv1.PathPrefix("/posts").Subrouter()
	middleware := middlewares.NewMiddleware(oAuthServerAddr)
	postsRouter.Use(middleware.EnsureAuth)

	getByIdH := getById.NewHandler(rt.PostService)
	postsRouter.HandleFunc("", getByIdH.GetById).Methods(http.MethodGet)

	delH := deleteById.NewHandler(rt.PostService)
	postsRouter.HandleFunc("", delH.DeleteById).Methods(http.MethodDelete)

	getAllH := getAll.NewHandler(rt.PostService)
	postsRouter.HandleFunc("/all", getAllH.GetAll).Methods(http.MethodGet)

	addH := add.NewHandler(rt.PostService)
	postsRouter.HandleFunc("/add", addH.Create).Methods(http.MethodPost)

	updateH := update.NewHandler(rt.PostService)
	postsRouter.HandleFunc("/update", updateH.UpdateById).Methods(http.MethodPost)

	return r
}
