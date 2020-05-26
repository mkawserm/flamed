package context

import (
	"bytes"
	"encoding/base64"
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variable"
	"net/http"
	"strings"
)

type GraphQLContext struct {
	URL        string
	Host       string
	RequestURI string
	Header     http.Header
	RemoteAddr string

	Data interface{}
}

func (g *GraphQLContext) GetBasicUsernameAndPassword(authData string) (string, string) {
	userAndPassBytes, err := base64.StdEncoding.DecodeString(authData)
	if err != nil {
		return "", ""
	}
	userAndPass := bytes.Split(userAndPassBytes, []byte(":"))
	if len(userAndPass) != 2 {
		return "", ""
	}
	username := string(userAndPass[0])
	password := string(userAndPass[1])
	return username, password

}

func (g *GraphQLContext) GetAuthorizationData() (string, string) {
	authorizationValue := ""
	if a, ok := g.Header["Authorization"]; ok {
		authorizationValue = a[0]
	}
	authorizationValue = strings.TrimSpace(authorizationValue)

	authData := strings.Split(authorizationValue, " ")

	if len(authData) != 2 {
		return "", ""
	}

	if len(authData[0]) == 0 {
		return "", ""
	}

	if len(authData[1]) == 0 {
		return "", ""
	}

	return authData[0], authData[1]
}

func (g *GraphQLContext) IsSuperUser(ctx *FlamedContext) bool {
	bearer, authData := g.GetAuthorizationData()
	if strings.EqualFold(bearer, "Basic") {
		username, password := g.GetBasicUsernameAndPassword(authData)
		admin := ctx.Flamed.NewAdmin(1, ctx.GlobalRequestTimeout)
		return g.IsSuperUserPasswordValid(admin, username, password)
	}

	return false
}

func (g *GraphQLContext) IsSuperUserPasswordValid(admin *flamed.Admin,
	username string,
	password string) bool {
	user, err := admin.GetUser(username)

	if err != nil {
		return false
	}

	if user.UserType != pb.UserType_SUPER_USER {
		return false
	}

	b, err := variable.DefaultPasswordHashAlgorithmFactory.CheckPassword(password, user.Password)
	if err != nil {
		return false
	}

	return b
}

func (g *GraphQLContext) IsUserPasswordValid(admin *flamed.Admin,
	username string,
	password string) bool {
	user, err := admin.GetUser(username)

	if err != nil {
		return false
	}

	b, err := variable.DefaultPasswordHashAlgorithmFactory.CheckPassword(password, user.Password)
	if err != nil {
		return false
	}

	return b
}
