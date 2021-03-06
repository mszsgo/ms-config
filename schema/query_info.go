package schema

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/mszsgo/hjson"
)

type InfoQuery struct {
	Name      string
	Value     string
	Remark    string
	UpdatedAt string
	CreatedAt string
}

func (*InfoQuery) Description() string {
	return "查询单个配置信息"
}

type InfoQueryArgs struct {
	Name string `graphql:"!" description:"配置名"`
}

func (*InfoQuery) Args() *InfoQueryArgs {
	return &InfoQueryArgs{}
}

func (*InfoQuery) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args *InfoQueryArgs
		hjson.MapToStruct(p.Args, &args)
		conf := NewConfig().FindOne(args.Name)
		return &InfoQuery{
			Name:      conf.Name,
			Value:     hjson.MapToJson(conf.Value),
			Remark:    conf.Remark,
			UpdatedAt: conf.UpdatedAt.Local().Format(time.RFC3339),
			CreatedAt: conf.CreatedAt.Local().Format(time.RFC3339),
		}, err
	}
}
