package context

import (
	"bytes"
	"encoding/base64"
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variable"
	"strings"
)

type GraphQLContext struct {
	URL        string
	Host       string
	RequestURI string
	RemoteAddr string

	Protocol string
	Data     map[string]interface{}
	Header   map[string][]string
}

func (g *GraphQLContext) AddData(key string, value interface{}) {
	if g.Data == nil {
		g.Data = make(map[string]interface{})
	}
	g.Data[key] = value
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

func (g *GraphQLContext) GetUsernameFromAuth() string {
	bearer, authData := g.GetAuthorizationData()
	if strings.EqualFold(bearer, "Basic") {
		username, _ := g.GetBasicUsernameAndPassword(authData)
		return username
	}

	return ""
}

func (g *GraphQLContext) GetUsernameAndPasswordFromAuth() (string, string) {
	bearer, authData := g.GetAuthorizationData()
	if strings.EqualFold(bearer, "Basic") {
		username, password := g.GetBasicUsernameAndPassword(authData)
		return username, password
	}

	return "", ""
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

func (g *GraphQLContext) AuthenticateSuperUser(admin *flamed.Admin) bool {
	username, password := g.GetUsernameAndPasswordFromAuth()
	if len(username) == 0 || len(password) == 0 {
		return false
	}

	return g.IsSuperUserPasswordValid(admin, username, password)
}

func (g *GraphQLContext) Authenticate(admin *flamed.Admin) bool {
	username, password := g.GetUsernameAndPasswordFromAuth()
	if len(username) == 0 || len(password) == 0 {
		return false
	}

	return g.IsUserPasswordValid(admin, username, password)
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
