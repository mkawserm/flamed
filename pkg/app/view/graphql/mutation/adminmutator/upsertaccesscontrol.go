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

		"permission": &graphql.ArgumentConfig{
			Description: "Access permission",
			Type:        graphql.NewNonNull(types.GQLPermissionInputType),
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

		permission := p.Args["permission"].(map[string]interface{})
		readPermission := permission["read"].(bool)
		writePermission := permission["write"].(bool)
		updatePermission := permission["update"].(bool)
		deletePermission := permission["delete"].(bool)

		globalSearchPermission := permission["globalSearch"].(bool)
		globalIteratePermission := permission["globalIterate"].(bool)
		globalRetrievePermission := permission["globalRetrieve"].(bool)
		globalCRUDPermission := permission["globalCRUD"].(bool)

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
		permissionNumber := utility.NewPermission(
			readPermission,
			writePermission,
			updatePermission,
			deletePermission,
			globalSearchPermission,
			globalIteratePermission,
			globalRetrievePermission,
			globalCRUDPermission)

		accessControl := &pb.AccessControl{
			Username:   username,
			Namespace:  namespace,
			Permission: permissionNumber,
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
}
