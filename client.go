package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "grpc-gorm-mysql/proto"
)

const address = "127.0.0.1.8080"

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	c := pb.NewOperationClient(conn)

	fmt.Printf("connected to the server")

loop:
	for {

		fmt.Printf("Please select the following operations:\n 1-select\t 2-insert\t 3-delete\t 4-update\t 5-sql\t 0-quit\n")
		var op string
		fmt.Scanln(&op)
		var result *pb.Reply
		switch op {
		case "1":
			fmt.Printf("enter the table name: ")
			var table string
			fmt.Scanln(&table)
			fmt.Println("enter the coloumn name: ")
			var columns string
			fmt.Scanln(&columns)
			fmt.Println("enter the condition: ")
			var con string
			fmt.Scanln(&con)
			result, err = c.Select(context.Background(), &pb.SelectRequest{Columns: columns, Table: table, Condition: con})

		case "2":
			fmt.Printf("Please enter the data: \n")
			var id int32
			var name string
			var price float32
			var typeId int32
			var createTime int64
			fmt.Scanf("%d %s %f %d %d", &id, &name, &price, &typeId, &createTime)
			result, err = c.Insert(context.Background(), &pb.InsDelUpdRequest{Id: id, Name: name, Price: price, TypeId: typeId, CreateTime: createTime})

		case "3":
			fmt.Printf("Please enter the data: \n")
			var id int32
			var name string
			var price float32
			var typeId int32
			var createTime int64
			fmt.Scanf("%d %s %f %d %d\n", &id, &name, &price, &typeId, &createTime)
			result, err = c.Delete(context.Background(), &pb.InsDelUpdRequest{Id: id, Name: name, Price: price, TypeId: typeId, CreateTime: createTime})

		case "4":
			fmt.Printf("Please enter the data: \n example: 3 wang 20  50 \n")
			var id int32
			var name string
			var price float32
			var typeId int32
			var createTime int64
			fmt.Scanf("%d %s %f %d %d\n", &id, &name, &price, &typeId, &createTime)
			result, err = c.Update(context.Background(), &pb.InsDelUpdRequest{Id: id, Name: name, Price: price, TypeId: typeId, CreateTime: createTime})

		case "5":
			fmt.Printf("Please enter the sql: \n")
			var sql string
			fmt.Scanln(&sql)
			result, err = c.ExecSql(context.Background(), &pb.SqlRequest{Sql: sql})

		case "0":
			break loop

		}
		if err != nil {
			fmt.Println("Failed to get reply.")
			return
		}
		fmt.Println(result)
	}

}
