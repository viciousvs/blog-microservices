package grpc

import (
	"context"
	"github.com/viciousvs/blog-microservices/post/model/post/handler/add"
	"github.com/viciousvs/blog-microservices/post/model/post/handler/deleteById"
	"github.com/viciousvs/blog-microservices/post/model/post/handler/getAll"
	"github.com/viciousvs/blog-microservices/post/model/post/handler/getById"
	"github.com/viciousvs/blog-microservices/post/model/post/handler/update"
	"github.com/viciousvs/blog-microservices/post/utils"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAll(ctx context.Context, req *pbPost.GetAllRequest) (*pbPost.Posts, error) {
	_ = req
	h := getAll.NewHandler(s.repo)
	posts, err := h.Handle(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot get all from DB: %v", err)
	}
	resultPosts := make([]*pbPost.Post, 0, len(posts))
	for _, p := range posts {
		newPost := &pbPost.Post{
			UUID:      p.UUID,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		}
		resultPosts = append(resultPosts, newPost)
	}

	return &pbPost.Posts{Posts: resultPosts}, nil
}
func (s *Server) GetByID(ctx context.Context, req *pbPost.GetByIdRequest) (*pbPost.Post, error) {
	h := getById.NewHandler(s.repo)
	p, err := h.Handle(ctx, req)
	if err != nil {
		if err == utils.ErrNotExist {
			return nil, status.Errorf(codes.NotFound, "data with uuid:%s not found:%v", req.GetUUID(), err)
		}
		return nil, status.Errorf(codes.Internal, "cannot getById from DB: %v", err)
	}

	res := &pbPost.Post{
		UUID:      p.UUID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
	return res, nil
}
func (s *Server) Update(ctx context.Context, req *pbPost.UpdateRequest) (*pbPost.UpdateResponse, error) {
	h := update.NewHandler(s.repo)
	uuid, err := h.Handle(ctx, req)
	if err != nil {
		if err == utils.ErrNotExist {
			return nil, status.Errorf(codes.NotFound, "data with uuid:%s not found:%v", req.GetUUID(), err)
		}
		return nil, status.Errorf(codes.Internal, "cannot updateById: %v", err)
	}

	return &pbPost.UpdateResponse{UUID: uuid}, nil
}
func (s *Server) DeleteByID(ctx context.Context, req *pbPost.DeleteByIdRequest) (*pbPost.DeleteByIDResponse, error) {
	h := deleteById.NewHandler(s.repo)
	uuid, err := h.Handle(ctx, req)
	if err != nil {
		if err == utils.ErrNotExist {
			return nil, status.Errorf(codes.NotFound, "data with uuid:%s not found:%v", req.GetUUID(), err)
		}
		return nil, status.Errorf(codes.Internal, "cannot delete from DB: %v", err)
	}

	return &pbPost.DeleteByIDResponse{UUID: uuid}, nil
}
func (s *Server) Create(ctx context.Context, req *pbPost.CreateRequest) (*pbPost.CreateResponse, error) {
	h := add.NewHandler(s.repo)
	uuid, err := h.Handle(ctx, req)
	if err != nil {
		return nil, err
	}

	return &pbPost.CreateResponse{UUID: uuid}, nil
}
