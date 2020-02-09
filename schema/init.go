package schema

import (
	"log"
	"os"

	"github.com/mszsgo/hmgdb"
)

func init() {
	log.Print(os.Getenv("MS_CONFIG_MONGO"))
	// 项目加载设置Mongodb连接字符串
	hmgdb.SetConnectString(hmgdb.DEFAULT, os.Getenv("MS_CONFIG_MONGO"))
}
