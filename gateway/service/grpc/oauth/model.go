package oauth

type signUpResponse struct {
	UUID   string
	Tokens *tokens
}
type signInResponse struct {
	UUID   string
	Tokens *tokens
}
type tokens struct {
	AccessExp   int64
	AccessToken string

	RefreshExp   int64
	RefreshToken string
}
