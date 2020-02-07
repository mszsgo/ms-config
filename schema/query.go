package schema

import (
	"github.com/graphql-go/graphql"
)

type Query struct {
	Config *ConfigQuery `description:"配置查询"`
}

type ConfigQuery struct {
	Info  InfoQuery   `description:"配置信息"`
	List  []ListQuery `description:"配置列表"`
	Total TotalQuery  `description:"记录数"`
}

func (*ConfigQuery) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		return "", err
	}
}
