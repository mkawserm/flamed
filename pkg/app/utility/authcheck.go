package utility

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/types"
	fContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/pb"
	"strings"
)

func AuthCheck(p graphql.ResolveParams,
	flamedContext *fContext.FlamedContext) (clusterID uint64,
	namespace string,
	accessControl *pb.AccessControl,
	err error) {
	clusterID = p.Args["clusterID"].(*types.UInt64).Value()
	namespace = p.Args["namespace"].(string)

	if !flamedContext.Flamed().IsClusterIDAvailable(clusterID) {
		err = gqlerrors.NewFormattedError(
			fmt.Sprintf("clusterID [%d] is not available", clusterID))
		return
	}

	if strings.EqualFold(namespace, "meta") {
		err = gqlerrors.NewFormattedError("meta namespace is reserved")
		return
	}

	//if p.Context.Value("GraphQLContext") == nil {
	//	return nil, nil
	//}

	gqlContext := p.Context.Value("GraphQLContext").(*fContext.GraphQLContext)
	admin := flamedContext.Flamed().NewAdmin(clusterID, flamedContext.GlobalRequestTimeout())

	// Authenticate user
	username, password := gqlContext.GetUsernameAndPasswordFromAuth()
	if len(username) == 0 || len(password) == 0 {
		err = gqlerrors.NewFormattedError("Access denied." +
			" Only authenticated user can access")
		return
	}

	if !gqlContext.IsUserPasswordValid(admin, username, password) {
		err = gqlerrors.NewFormattedError("Access denied." +
			" Only authenticated user can access")
		return
	}

	accessControl, err = admin.GetAccessControl(username, []byte(namespace))
	if err != nil {
		accessControl, err = admin.GetAccessControl(username, []byte("*"))
		if err != nil {
			err = gqlerrors.NewFormattedError("namespace access control not found")
			return
		}
	}
	return
}
