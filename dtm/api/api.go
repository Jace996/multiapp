package api

import (
	"context"
	"fmt"
	sapi "github.com/jace996/multiapp/pkg/api"
	"github.com/jace996/multiapp/pkg/authn/jwt"
	"github.com/jace996/multiapp/pkg/conf"
)

const ServiceName = "dtmservice"

var ClientConf = &conf.Client{
	ClientId: ServiceName,
}

func MustAddBranchHeader(ctx context.Context, tokenMgr sapi.TokenManager) map[string]string {
	t, err := tokenMgr.GetOrGenerateToken(ctx, ClientConf)
	if err != nil {
		panic(err)
	}
	return map[string]string{
		jwt.AuthorizationHeader: fmt.Sprintf("%s %s", jwt.BearerTokenType, t),
	}
}
