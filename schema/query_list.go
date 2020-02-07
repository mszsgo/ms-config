package schema

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/mszsgo/hjson"
)

type ListQuery struct {
	Name      string
	Value     string
	Remark    string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (*ListQuery) Description() string {
	return "查询单个配置信息"
}

type ListQueryArgs struct {
	Name string `graphql:"!" description:"配置名"`
}

func (*ListQuery) Args() *ListQueryArgs {
	return &ListQueryArgs{}
}

func (*ListQuery) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args *ListQueryArgs
		hjson.MapToStruct(p.Args, &args)
		conf := NewConfig().FindOne(args.Name)
		return &conf, err
	}
}
