package context

import (
	"bytes"
	"encoding/base64"
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/variable"
	"strings"
)

type AuthContext struct {
	URL        string
	Host       string
	RequestURI string
	RemoteAddr string

	Protocol string
	Data     map[string]interface{}
	KVPair   map[string][]string
}

func (c *AuthContext) AddData(key string, value interface{}) {
	if c.Data == nil {
		c.Data = make(map[string]interface{})
	}
	c.Data[key] = value
}

func (c *AuthContext) GetBasicUsernameAndPassword(authData string) (string, string) {
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

func (c *AuthContext) GetUsernameFromAuth() string {
	bearer, authData := c.GetAuthorizationData()
	if strings.EqualFold(bearer, "Basic") {
		username, _ := c.GetBasicUsernameAndPassword(authData)
		return username
	}

	return ""
}

func (c *AuthContext) GetUsernameAndPasswordFromAuth() (string, string) {
	bearer, authData := c.GetAuthorizationData()
	if strings.EqualFold(bearer, "Basic") {
		username, password := c.GetBasicUsernameAndPassword(authData)
		return username, password
	}

	return "", ""
}

func (c *AuthContext) GetAuthorizationData() (string, string) {
	authorizationValue := ""
	if a, ok := c.KVPair["authorization"]; ok {
		authorizationValue = a[0]
	} else if a, ok := c.KVPair["Authorization"]; ok {
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

func (c *AuthContext) AuthenticateSuperUser(admin *flamed.Admin) bool {
	username, password := c.GetUsernameAndPasswordFromAuth()
	if len(username) == 0 || len(password) == 0 {
		return false
	}

	return c.IsSuperUserPasswordValid(admin, username, password)
}

func (c *AuthContext) Authenticate(admin *flamed.Admin) bool {
	username, password := c.GetUsernameAndPasswordFromAuth()
	if len(username) == 0 || len(password) == 0 {
		return false
	}

	return c.IsUserPasswordValid(admin, username, password)
}

func (c *AuthContext) IsSuperUserPasswordValid(admin *flamed.Admin,
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

func (c *AuthContext) IsUserPasswordValid(admin *flamed.Admin,
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
