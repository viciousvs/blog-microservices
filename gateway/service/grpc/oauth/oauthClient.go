package oauth

import (
	"context"
	oauthPb "github.com/viciousvs/OAuth-services/proto/oauthService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OAuthService struct {
	addr string
}

func NewOAuthService(addr string) *OAuthService {
	return &OAuthService{addr: addr}
}

func (r OAuthService) AuthenticateAccessToken(ctx context.Context, accessToken string) (string, error) {
	conn, err := grpc.Dial(r.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := oauthPb.NewOAuthServiceClient(conn)
	resp, err := client.Access(ctx, &oauthPb.AccessRequest{AccessToken: accessToken})
	if err != nil {
		return "", err
	}
	return resp.GetUserUuid(), nil
}
