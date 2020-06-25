package adminmutator

import (
	"encoding/base64"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/variable"
	"github.com/mkawserm/flamed/pkg/x"
	"go.uber.org/zap"
	"time"
)

var UpsertUser = &graphql.Field{
	Name:        "UpsertUser",
	Description: "",
	Type:        kind.GQLProposalResponseType,
	Args: graphql.FieldConfigArgument{
		"userType": &graphql.ArgumentConfig{
			Description: "User type",
			Type:        kind.GQLUserTypeEnum,
		},
		"roles": &graphql.ArgumentConfig{
			Description: "Roles",
			Type:        graphql.String,
		},
		"username": &graphql.ArgumentConfig{
			Description: "Username",
			Type:        graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Description: "Password",
			Type:        graphql.NewNonNull(graphql.String),
		},
		"data": &graphql.ArgumentConfig{
			Description: "Data in base64 encoded string",
			Type:        graphql.String,
		},
		"meta": &graphql.ArgumentConfig{
			Description: "Meta in base64 encoded string",
			Type:        graphql.String,
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		username := p.Args["username"].(string)
		password := p.Args["password"].(string)
		if !utility.IsUsernameValid(username) {
			return nil, gqlerrors.NewFormattedError("invalid username")
		}

		if username == "admin" {
			return nil, gqlerrors.NewFormattedError(x.ErrInvalidOperation.Error())
		}

		//TODO: check password validity

		admin, ok := p.Source.(*flamed.Admin)
		if !ok {
			return nil, gqlerrors.NewFormattedError(x.ErrInvalidSourceType.Error())
		}

		//if admin.IsUserAvailable(username) {
		//	return nil, gqlerrors.NewFormattedError(username+" is already available")
		//}

		pha := variable.DefaultPasswordHashAlgorithmFactory
		if !pha.IsAlgorithmAvailable(variable.DefaultPasswordHashAlgorithm) {
			logger.L("app").Error(variable.DefaultPasswordHashAlgorithm +
				" password hash algorithm is to available")
			return nil, gqlerrors.NewFormattedError(x.ErrPasswordHashAlgorithmIsNotAvailable.Error())
		}

		encoded, err := pha.MakePassword(password,
			crypto.GetRandomString(12),
			variable.DefaultPasswordHashAlgorithm)

		if err != nil {
			logger.L("app").Error("make password returned error", zap.Error(err))
			return nil, gqlerrors.NewFormattedError(x.ErrFailedToGeneratePassword.Error())
		}

		user := &pb.User{
			UserType:  pb.UserType_NORMAL_USER,
			Roles:     "",
			Username:  username,
			Password:  encoded,
			CreatedAt: uint64(time.Now().UnixNano()),
			UpdatedAt: uint64(time.Now().UnixNano()),
			Data:      nil,
			Meta:      nil,
		}

		if p.Args["userType"] != nil {
			user.UserType = pb.UserType(p.Args["userType"].(int))
		}

		if p.Args["roles"] != nil {
			user.Roles = p.Args["roles"].(string)
		}

		if p.Args["data"] != nil {
			data, e := base64.StdEncoding.DecodeString(p.Args["data"].(string))
			if e == nil {
				user.Data = data
			}
		}

		if p.Args["meta"] != nil {
			meta, e := base64.StdEncoding.DecodeString(p.Args["meta"].(string))
			if e == nil {
				user.Meta = meta
			}
		}

		pr, err := admin.UpsertUser(user)
		if err != nil {
			return nil, err
		}

		return pr, nil
	},
}
