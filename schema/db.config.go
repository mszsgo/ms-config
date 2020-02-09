package schema

import (
	"log"
	"time"

	"github.com/mszsgo/hmgdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	hmgdb.UpdateOne(nil, o.Collection(), bson.M{"name": name}, bson.M{"$set": bson.M{"value": value, "remark": remark, "updatedAt": time.Now()}})
}

func (o *Config) Find(name string, skip, limit int64) []*Config {
	var list []*Config
	m := bson.M{}
	if name != "" {
		m["name"] = name
	}
	hmgdb.Find(nil, o.Collection(), m, &list, options.Find().SetSkip(skip).SetLimit(limit))
	return list
}

func (o *Config) FindOne(name string) *Config {
	var conf *Config
	hmgdb.FindOne(nil, o.Collection(), bson.M{"name": name}, &conf)
	return conf
}

func (o *Config) Exists(name string) bool {
	return hmgdb.Exists(nil, o.Collection(), bson.M{"name": name})
}

func (o *Config) Total(name string) int64 {
	m := bson.M{}
	if name != "" {
		m["name"] = name
	}
	count, err := o.Collection().CountDocuments(nil, m)
	if err != nil {
		log.Print(err)
	}
	return count
}
