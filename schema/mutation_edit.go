package schema

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/mszsgo/hjson"
)

type EditMutation struct {
	Name      string    `description:"配置名"`
	UpdatedAt time.Time `description:"更新时间"`
}

func (*EditMutation) Description() string {
	return "编辑"
}

type EditMutationArgs struct {
	Name   string `graphql:"!" description:"配置名"`
	Value  string `graphql:"!" description:"配置值"`
	Remark string `graphql:"" description:"备注信息"`
}

func (*EditMutation) Args() *EditMutationArgs {
	return &EditMutationArgs{}
}

func (*EditMutation) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args *EditMutationArgs
		hjson.MapToStruct(p.Args, &args)
		NewConfig().Update(args.Name, args.Value, args.Remark)
		return &EditMutation{
			Name:      args.Name,
			UpdatedAt: time.Now(),
		}, err
	}
}
