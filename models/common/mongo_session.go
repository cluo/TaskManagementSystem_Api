package common

import (
	"errors"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2"
)

// MongoSessionStruct 定义
type MongoSessionStruct struct {
	session    *mgo.Session
	db         *mgo.Database
	collection *mgo.Collection
}

var rootSession *mgo.Session

// GetMongoSession 定义
func GetMongoSession() (mongo *MongoSessionStruct, err error) {
	if rootSession == nil {
		dbURL := beego.AppConfig.String("mongo")
		rootSession, err = mgo.Dial(dbURL)
		if err != nil {
			return nil, err
		}
	}
	mongo = &MongoSessionStruct{}
	mongo.session = rootSession.Copy()
	return mongo, nil
}

// CloseSession 定义
func (mongo *MongoSessionStruct) CloseSession() error {
	if mongo == nil {
		return errors.New("MongoSessionStruct类型空指针错误。")
	}
	if mongo.session == nil {
		return errors.New("Session类型空指针错误。")
	}
	mongo.collection = nil
	mongo.db = nil

	mongo.session.Close()
	mongo.session = nil

	return nil
}

// UseDB 定义
func (mongo *MongoSessionStruct) UseDB(dbName string) error {
	if mongo == nil {
		return errors.New("MongoSessionStruct类型空指针错误。")
	}
	if mongo.session == nil {
		return errors.New("Session类型空指针错误。")
	}
	mongo.db = mongo.session.DB(dbName)
	mongo.collection = nil

	return nil
}

// UseCollection 定义
func (mongo *MongoSessionStruct) UseCollection(collectionName string) error {
	if mongo == nil {
		return errors.New("MongoSessionStruct类型空指针错误。")
	}
	if mongo.session == nil {
		return errors.New("Session类型空指针错误。")
	}
	if mongo.db == nil {
		return errors.New("Database类型空指针错误。")
	}
	mongo.collection = mongo.db.C(collectionName)

	return nil
}

// type Person struct {
// 	NAME  string
// 	PHONE string
// }
// type Men struct {
// 	Persons []Person
// }

// func test() {

// 	session, err := mgo.Dial("192.168.2.175:27017") //连接数据库
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer session.Close()
// 	session.SetMode(mgo.Monotonic, true)

// 	db := session.DB("mydb")     //数据库名称
// 	collection := db.C("person") //如果该集合已经存在的话，则直接返回

// 	//*****集合中元素数目********
// 	countNum, err := collection.Count()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Things objects count: ", countNum)

// 	//*******插入元素*******
// 	temp := &Person{
// 		NAME:  "zhangzheHero",
// 		PHONE: "18811577546"}

// 	//一次可以插入多个对象 插入两个Person对象
// 	err = collection.Insert(&Person{"Ale", "+55 53 8116 9639"}, temp)
// 	if err != nil {
// 		panic(err)
// 	}

// 	//*****查询单条数据*******
// 	result := Person{}
// 	err = collection.Find(bson.M{"phone": "456"}).One(&result)
// 	fmt.Println("Phone:", result.NAME, result.PHONE)

// 	//*****查询多条数据*******
// 	var personAll Men //存放结果
// 	iter := collection.Find(nil).Iter()
// 	for iter.Next(&result) {
// 		fmt.Printf("Result: %v\n", result.NAME)
// 		personAll.Persons = append(personAll.Persons, result)
// 	}

// 	//*******更新数据**********
// 	err = collection.Update(bson.M{"name": "ccc"}, bson.M{"$set": bson.M{"name": "ddd"}})
// 	err = collection.Update(bson.M{"name": "ddd"}, bson.M{"$set": bson.M{"phone": "12345678"}})
// 	err = collection.Update(bson.M{"name": "aaa"}, bson.M{"phone": "1245", "name": "bbb"})

// 	//******删除数据************
// 	_, err = collection.RemoveAll(bson.M{"name": "Ale"})
// }
