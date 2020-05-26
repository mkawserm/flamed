package adminmutator

import (
	"bytes"
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

var GQLAdminMutatorType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "AdminMutator",
	Description: "`AdminMutator`",
	Fields: graphql.Fields{
		"upsertAccessControl": &graphql.Field{
			Name:        "UpsertAccessControl",
			Description: "",
			Type:        types.ProposalResponseType,
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Description: "Username",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"namespace": &graphql.ArgumentConfig{
					Description: "Namespace",
					Type:        graphql.NewNonNull(graphql.String),
				},

				"readAccess": &graphql.ArgumentConfig{
					Description: "Read access",
					Type:        graphql.NewNonNull(graphql.Boolean),
				},
				"writeAccess": &graphql.ArgumentConfig{
					Description: "Write access",
					Type:        graphql.NewNonNull(graphql.Boolean),
				},
				"updateAccess": &graphql.ArgumentConfig{
					Description: "Update access",
					Type:        graphql.NewNonNull(graphql.Boolean),
				},
				"deleteAccess": &graphql.ArgumentConfig{
					Description: "Delete access",
					Type:        graphql.NewNonNull(graphql.Boolean),
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
				namespace := []byte(p.Args["namespace"].(string))

				readAccess := p.Args["readAccess"].(bool)
				writeAccess := p.Args["writeAccess"].(bool)
				updateAccess := p.Args["updateAccess"].(bool)
				deleteAccess := p.Args["deleteAccess"].(bool)

				if !bytes.Equal(namespace, []byte("*")) {
					if !utility.IsNamespaceValid(namespace) {
						return nil, gqlerrors.NewFormattedError("invalid namespace")
					}
				}

				if !utility.IsUsernameValid(username) {
					return nil, gqlerrors.NewFormattedError("invalid username")
				}

				admin, ok := p.Source.(*flamed.Admin)
				if !ok {
					return nil, gqlerrors.NewFormattedError("Unknown source type." +
						" FlamedContext required")
				}

				//if !admin.IsUserAvailable(username) {
				//	return nil, gqlerrors.NewFormattedError(username+" is not available")
				//}

				accessControl := &pb.AccessControl{
					Username:   username,
					Namespace:  namespace,
					Permission: utility.NewPermission(readAccess, writeAccess, updateAccess, deleteAccess),
					CreatedAt:  uint64(time.Now().UnixNano()),
					UpdatedAt:  uint64(time.Now().UnixNano()),
					Data:       nil,
					Meta:       nil,
				}

				if p.Args["data"] != nil {
					data, e := base64.StdEncoding.DecodeString(p.Args["data"].(string))
					if e == nil {
						accessControl.Data = data
					}
				}

				if p.Args["meta"] != nil {
					meta, e := base64.StdEncoding.DecodeString(p.Args["meta"].(string))
					if e == nil {
						accessControl.Meta = meta
					}
				}

				pr, err := admin.UpsertAccessControl(accessControl)
				if err != nil {
					return nil, err
				}

				return pr, nil
			},
		},

		"deleteAccessControl": &graphql.Field{
			Name:        "DeleteAccessControl",
			Description: "",
			Type:        types.ProposalResponseType,
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Description: "Username",
					Type:        graphql.NewNonNull(graphql.String),
				},
				"namespace": &graphql.ArgumentConfig{
					Description: "Namespace",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				username := p.Args["username"].(string)
				namespace := []byte(p.Args["namespace"].(string))
				admin, ok := p.Source.(*flamed.Admin)
				if !ok {
					return nil, gqlerrors.NewFormattedError("Unknown source type." +
						" FlamedContext required")
				}

				pr, err := admin.DeleteAccessControl(namespace, username)
				if err != nil {
					return nil, err
				}

				return pr, nil
			},
		},

		"deleteUser": &graphql.Field{
			Name:        "DeleteUser",
			Description: "",
			Type:        types.ProposalResponseType,
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Description: "Username",
					Type:        graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				username := p.Args["username"].(string)
				admin, ok := p.Source.(*flamed.Admin)

				if username == "admin" {
					return nil, gqlerrors.NewFormattedError("delete operation is not" +
						" allowed on admin user")
				}

				if !ok {
					return nil, gqlerrors.NewFormattedError("Unknown source type." +
						" FlamedContext required")
				}

				pr, err := admin.DeleteUser(username)
				if err != nil {
					return nil, err
				}

				return pr, nil
			},
		},

		"upsertUser": &graphql.Field{
			Name:        "UpsertUser",
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
					return nil, gqlerrors.NewFormattedError("upsert operation is not" +
						" allowed on admin user")
				}

				//TODO: check password validity

				admin, ok := p.Source.(*flamed.Admin)
				if !ok {
					return nil, gqlerrors.NewFormattedError("Unknown source type." +
						" FlamedContext required")
				}

				//if admin.IsUserAvailable(username) {
				//	return nil, gqlerrors.NewFormattedError(username+" is already available")
				//}

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
		},

		"updateUser": &graphql.Field{
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
		},
	},
})
