package utility

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	fContext "github.com/mkawserm/flamed/pkg/context"
	"github.com/mkawserm/flamed/pkg/pb"
	"github.com/mkawserm/flamed/pkg/x"
	"strings"
)

func GraphQLAuthCheck(p graphql.ResolveParams,
	flamedContext *fContext.FlamedContext) (clusterID uint64,
	namespace string,
	accessControl *pb.AccessControl,
	err error) {
	clusterID = p.Args["clusterID"].(*kind.UInt64).Value()
	namespace = p.Args["namespace"].(string)

	if !flamedContext.Flamed().IsClusterIDAvailable(clusterID) {
		err = gqlerrors.NewFormattedError(x.ErrClusterIsNotAvailable.Error())
		return
	}

	if strings.EqualFold(namespace, "meta") {
		err = gqlerrors.NewFormattedError(x.ErrMetaNamespaceIsReserved.Error())
		return
	}

	//if p.Context.Value("GraphQLContext") == nil {
	//	return nil, nil
	//}

	gqlContext := p.Context.Value("GraphQLContext").(*fContext.AuthContext)
	admin := flamedContext.Flamed().NewAdmin(clusterID, flamedContext.GlobalRequestTimeout())

	// Authenticate user
	username, password := gqlContext.GetUsernameAndPasswordFromAuth()
	if len(username) == 0 || len(password) == 0 {
		err = gqlerrors.NewFormattedError(x.ErrAccessDenied.Error())
		return
	}

	if !gqlContext.IsUserPasswordValid(admin, username, password) {
		err = gqlerrors.NewFormattedError(x.ErrAccessDenied.Error())
		return
	}

	accessControl, err = admin.GetAccessControl(username, []byte(namespace))
	if err != nil {
		accessControl, err = admin.GetAccessControl(username, []byte("*"))
		if err != nil {
			err = gqlerrors.NewFormattedError(x.ErrAccessControlNotFound.Error())
			return
		}
	}
	return
}
