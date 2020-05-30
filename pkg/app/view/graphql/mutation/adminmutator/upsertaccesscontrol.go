package adminmutator

import (
	"bytes"
	"encoding/base64"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/view/graphql/types"
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/utility"
	"time"
)

var UpsertAccessControl = &graphql.Field{
	Name:        "UpsertAccessControl",
	Description: "",
	Type:        types.GQLProposalResponseType,
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
		"globalSearchAccess": &graphql.ArgumentConfig{
			Description: "GlobalOperation globaloperation access",
			Type:        graphql.NewNonNull(graphql.Boolean),
		},
		"globalIterateAccess": &graphql.ArgumentConfig{
			Description: "GlobalOperation iterate access",
			Type:        graphql.NewNonNull(graphql.Boolean),
		},
		"globalRetrieveAccess": &graphql.ArgumentConfig{
			Description: "GlobalOperation retrieve access",
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

		globalSearchAccess := p.Args["globalSearchAccess"].(bool)
		globalIterateAccess := p.Args["globalIterateAccess"].(bool)
		globalRetrieveAccess := p.Args["globalRetrieveAccess"].(bool)

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
			Username:  username,
			Namespace: namespace,
			Permission: utility.NewPermission(readAccess,
				writeAccess,
				updateAccess,
				deleteAccess,
				globalSearchAccess,
				globalIterateAccess,
				globalRetrieveAccess),
			CreatedAt: uint64(time.Now().UnixNano()),
			UpdatedAt: uint64(time.Now().UnixNano()),
			Data:      nil,
			Meta:      nil,
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
}
