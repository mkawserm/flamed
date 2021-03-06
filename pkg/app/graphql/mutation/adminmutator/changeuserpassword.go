package adminmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/crypto"
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/logger"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/mkawserm/flamed/pkg/variable"
	"github.com/mkawserm/flamed/pkg/x"
	"go.uber.org/zap"
	"time"
)

var ChangeUserPassword = &graphql.Field{
	Name:        "ChangeUserPassword",
	Description: "",
	Type:        kind.GQLProposalResponseType,
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Description: "Username",
			Type:        graphql.NewNonNull(graphql.String),
		},
		"password": &graphql.ArgumentConfig{
			Description: "Password",
			Type:        graphql.NewNonNull(graphql.String),
		},
	},

	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		username := p.Args["username"].(string)
		if !utility.IsUsernameValid(username) {
			return nil, gqlerrors.NewFormattedError(x.ErrInvalidUsername.Error())
		}

		admin, ok := p.Source.(*flamed.Admin)
		if !ok {
			return nil, gqlerrors.NewFormattedError(x.ErrInvalidSourceType.Error())
		}

		user, err := admin.GetUser(username)

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		user.UpdatedAt = uint64(time.Now().UnixNano())
		password := p.Args["password"].(string)

		//TODO: check password policy
		if len(password) == 0 {
			return nil, gqlerrors.NewFormattedError(x.ErrPasswordCanNotBeEmpty.Error())
		}

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
		user.Password = encoded

		pr, err := admin.UpsertUser(user)
		if err != nil {
			return nil, err
		}

		return pr, nil
	},
}
