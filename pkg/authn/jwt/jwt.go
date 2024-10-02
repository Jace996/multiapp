package jwt

import (
	kitdi "github.com/jace996/multiapp/pkg/di"
)

var ProviderSet = kitdi.NewSet(NewTokenizer, NewTokenizerConfig)

const (
	AuthorizationHeader = "Authorization"
	BearerTokenType     = "Bearer"
	AuthorizationQuery  = "access_token"
)
