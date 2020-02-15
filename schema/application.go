package schema

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/mszsgo/hconfig"
	"github.com/mszsgo/hgraph"
	"github.com/mszsgo/hjson"
	"github.com/mszsgo/hmgdb"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	App = &hconfig.App{Name: "config", Version: "0.0.1", Port: 80}
	// Mongodb连接
	DB *mongo.Database
)

// main方法第一行加载配置信息
func Main() {
	var mcs string
	// 读取命令参数
	flag.StringVar(&mcs, "mongo", "", "Mongodb connection string ")
	flag.Parse()
	if mcs == "" {
		mcs = os.Getenv("MONGO")
		log.Print("MONGO=" + mcs)
	}

	// Mongodb 连接
	DB = hmgdb.Connect(mcs)

	App.Start(preFunc)
}

func preFunc(app *hconfig.App) {
	// Graphql
	hgraph.HttpHandle(&Query{}, &Mutation{})

	// 读取配置 http://config/get?name=xxxxxx
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		defer func() {
			if err := recover(); err != nil {
				log.Print(err)
			}
		}()
		name := r.FormValue("name")
		conf := NewConfig().FindOne(name)
		if conf != nil {
			w.Write(hjson.MapToBytes(conf.Value))
		}
	})

}
