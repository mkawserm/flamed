package adminmutator

import (
	"encoding/base64"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/variable"
	"go.uber.org/zap"
	"time"
)

var UpdateUser = &graphql.Field{
	Name:        "UpdateUser",
	Description: "",
	Type:        types.ProposalResponseType,
	Args: graphql.FieldConfigArgument{
		"userType": &graphql.ArgumentConfig{
			Description: "User type",
			Type:        types.UserTypeEnum,
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
			Type:        graphql.String,
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
		if !utility.IsUsernameValid(username) {
			return nil, gqlerrors.NewFormattedError("invalid username")
		}

		if username == "admin" {
			if p.Args["userType"] != nil {
				if pb.UserType(p.Args["userType"].(int)) != pb.UserType_SUPER_USER {
					return nil, gqlerrors.NewFormattedError("changing user type is not" +
						" allowed on admin user")
				}
			}
		}

		admin, ok := p.Source.(*flamed.Admin)
		if !ok {
			return nil, gqlerrors.NewFormattedError("Unknown source type." +
				" FlamedContext required")
		}

		user, err := admin.GetUser(username)

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		user.UpdatedAt = uint64(time.Now().UnixNano())

		//if admin.IsUserAvailable(username) {
		//	return nil, gqlerrors.NewFormattedError(username+" is already available")
		//}

		if p.Args["password"] != nil {
			password := p.Args["password"].(string)
			//TODO: check password validity

			pha := variable.DefaultPasswordHashAlgorithmFactory
			if !pha.IsAlgorithmAvailable(variable.DefaultPasswordHashAlgorithm) {
				logger.L("app").Error(variable.DefaultPasswordHashAlgorithm +
					" password hash algorithm is to available")
				return nil, gqlerrors.NewFormattedError(variable.DefaultPasswordHashAlgorithm +
					" is not available")
			}

			encoded, err := pha.MakePassword(password,
				crypto.GetRandomString(12),
				variable.DefaultPasswordHashAlgorithm)

			if err != nil {
				logger.L("app").Error("make password returned error", zap.Error(err))
				return nil, gqlerrors.NewFormattedError(err.Error())
			}
			user.Password = encoded
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
