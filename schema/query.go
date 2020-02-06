package schema

import (
	"github.com/graphql-go/graphql"
)

type Query struct {
	Config *ConfigQuery `description:"配置查询"`
}

type ConfigQuery struct {
	Info  string `description:"配置单个信息"`
	List  string `description:"配置列表信息"`
	Total int64  `description:"配置个数统计"`
}

func (*ConfigQuery) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		return "", err
	}
}
