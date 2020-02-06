package schema

import (
	"time"

	"github.com/mszsgo/hmgdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 配置集合，name为服务名，value为服务配置
// Collection: ms_config
type Config struct {
	Name      string    `bson:"name"`
	Value     string    `bson:"value"`
	Remark    string    `bson:"remark"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func NewConfig() *Config {
	return &Config{}
}

func (o *Config) Collection() *mongo.Collection {
	return hmgdb.Db().Collection("ms_config")
}

func (o *Config) Insert(name, value, remark string) {
	hmgdb.InsertOne(nil, o.Collection(), &Config{
		Name:      name,
		Value:     value,
		Remark:    remark,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}

func (o *Config) Update(name, value, remark string) {
	hmgdb.UpdateOne(nil, o.Collection(), bson.M{"name": name}, &Config{
		Name:      name,
		Value:     value,
		Remark:    remark,
		UpdatedAt: time.Now(),
	})
}

func (o *Config) FindOne(name string) *Config {
	var conf *Config
	hmgdb.FindOne(nil, o.Collection(), bson.M{"name": name}, &conf)
	return conf
}

func (o *Config) Exists(name string) bool {
	return hmgdb.Exists(nil, o.Collection(), bson.M{"name": name})
}
