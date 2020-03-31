package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	collection := client.Database("test").Collection("stu")

	s1 := Student{"小红", 12}
	s2 := Student{"小兰", 10}
	s3 := Student{"小黄", 11}

	insertResult, err := collection.InsertOne(context.TODO(), s1)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)

	students := []interface{}{s2, s3}
	insertManyResult, err := collection.InsertMany(context.TODO(), students)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// updateone()方法允许你更新单个文档。它需要一个筛选器文档来匹配数据库中的文档，并需要一个更新文档来描述更新操作
	// 你可以使用bson.D类型来构建筛选文档和更新文档:
	filter := bson.D{{"name", "小兰"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// 创建一个Student变量用来接收查询的结果
	var result Student
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Found a single document: %+v\n", result)

	// 查询多个
	// 将选项传递给Find()
	findOptions := options.Find()
	findOptions.SetLimit(2)
	// 定义一个切片用来存储查询结果
	var results []*Student
	// 把bson.D{{}}作为一个filter来匹配所有文档
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem Student
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// 完成后关闭游标
	_ = cur.Close(context.TODO())
	log.Printf("Found multiple documents (array of pointers): %#v\n", results)

	// 删除名字是小黄的那个
	deleteResult1, err := collection.DeleteOne(context.TODO(), bson.D{{"name", "小黄"}})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted %v documents in the trainers collection\n", deleteResult1.DeletedCount)
	// 删除所有
	deleteResult2, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Deleted %v documents in the trainers collection\n", deleteResult2.DeletedCount)

	// 断开连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to MongoDB closed.")
}
