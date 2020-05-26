package context

import (
	"bytes"
	"encoding/base64"
	"fmt"
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

func (g *GraphQLContext) IsSuperUser(ctx *FlamedContext) bool {
	fmt.Println(g.Header)
	authorizationValue := ""
	if a, ok := g.Header["Authorization"]; ok {
		authorizationValue = a[0]
	}
	authorizationValue = strings.TrimSpace(authorizationValue)

	authData := strings.Split(authorizationValue, " ")

	if len(authData) != 2 {
		return false
	}

	if len(authData[0]) == 0 {
		return false
	}

	if len(authData[1]) == 0 {
		return false
	}

	if strings.EqualFold(authData[0], "Basic") {
		userAndPassBytes, err := base64.StdEncoding.DecodeString(authData[1])
		if err != nil {
			return false
		}
		userAndPass := bytes.Split(userAndPassBytes, []byte(":"))
		if len(userAndPass) != 2 {
			return false
		}

		username := string(userAndPass[0])
		password := string(userAndPass[1])

		admin := ctx.Flamed.NewAdmin(1, ctx.GlobalRequestTimeout)

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

	return false
}
