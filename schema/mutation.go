package schema

import (
	"github.com/graphql-go/graphql"
)

type Mutation struct {
	Config *ConfigMutation `description:"配置更新"`
}

type ConfigMutation struct {
	Add  *AddMutation  `description:"添加配置"`
	Edit *EditMutation `description:"编辑配置"`
}

func (*ConfigMutation) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		return "", err
	}
}
