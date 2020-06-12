package nodeadminmutator

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/mkawserm/flamed/pkg/app/graphql/kind"
	"github.com/mkawserm/flamed/pkg/constant"
	"github.com/mkawserm/flamed/pkg/flamed"
	"github.com/mkawserm/flamed/pkg/utility"
	"github.com/spf13/viper"
)

var RequestSnapshot = &graphql.Field{
	Type:        kind.GQLUInt64Type,
	Description: "Add new node to the cluster",
	Args: graphql.FieldConfigArgument{
		"compactionOverhead": &graphql.ArgumentConfig{
			Description: "Compaction overhead",
			Type:        graphql.NewNonNull(kind.GQLUInt64Type),
		},
		"exported": &graphql.ArgumentConfig{
			Description: "Will the snapshot be exported?",
			Type:        graphql.NewNonNull(graphql.Boolean),
		},
		"overrideCompactionOverhead": &graphql.ArgumentConfig{
			Description: "Compaction overhead override flag",
			Type:        graphql.NewNonNull(graphql.Boolean),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		compactionOverhead := p.Args["compactionOverhead"].(*kind.UInt64)
		exported := p.Args["exported"].(bool)
		overrideCompactionOverhead := p.Args["overrideCompactionOverhead"].(bool)
		exportPath := viper.GetString(constant.StoragePath) + "/snapshot"
		utility.MkPath(exportPath)
		nodeAdmin, ok := p.Source.(*flamed.NodeAdmin)
		if !ok {
			return nil, nil
		}
		n, err := nodeAdmin.RequestSnapshot(compactionOverhead.Value(),
			exportPath,
			exported,
			overrideCompactionOverhead)

		if err != nil {
			return nil, gqlerrors.NewFormattedError(err.Error())
		}

		return kind.NewUInt64FromUInt64(n), nil
	},
}
