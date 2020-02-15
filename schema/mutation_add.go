package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/mszsgo/hjson"
)

type AddMutation int64

func (*AddMutation) Description() string {
	return "新增"
}

type AddMutationArgs struct {
	Name   string `graphql:"!" description:"配置名"`
	Value  string `graphql:"!" description:"Value值"`
	Remark string `graphql:"" description:"备注信息"`
}

func (*AddMutation) Args() *AddMutationArgs {
	return &AddMutationArgs{}
}

func (*AddMutation) Resolve() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (i interface{}, err error) {
		var args *AddMutationArgs
		hjson.MapToStruct(p.Args, &args)
		err = NewConfig().Add(args.Name, args.Value, args.Remark)
		i = 0
		if err == nil {
			i = 1
		}
		return i, err
	}
}
