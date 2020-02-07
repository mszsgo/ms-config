package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/mszsgo/hjson"
)

type TotalQuery int64

func (*TotalQuery) Description() string {
	return "查询单个配置信息"
}

type TotalQueryArgs struct {
	Name string `graphql:"!" description:"配置名"`
}

func (*TotalQuery) Args() *TotalQueryArgs {
	return &TotalQueryArgs{}
}

func (*TotalQuery) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args *TotalQueryArgs
		hjson.MapToStruct(p.Args, &args)
		conf := NewConfig().FindOne(args.Name)
		return &conf, err
	}
}
