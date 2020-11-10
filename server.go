package main

import (
	"context"
	"fmt"
	"net"

	"grpc-gorm-mysql/mysql"
	pb "grpc-gorm-mysql/proto"

	"google.golang.org/grpc"
)

const port = ":8080"

type myServer struct {
}

func (m *myServer) Insert(ctx context.Context, in *pb.InsDelUpdRequest) (*pb.Reply, error) {
	mysql.InsDelUpd("insert", in.GetId(), in.GetName(), in.GetPrice(), in.GetTypeId(), in.GetCreateTime())
	return &pb.Reply{Result: "Insert completed."}, nil
}
func (m *myServer) Delete(ctx context.Context, in *pb.InsDelUpdRequest) (*pb.Reply, error) {
	mysql.InsDelUpd("delete", in.GetId(), in.GetPrice(), in.GetTypeId(), in.GetCreateTime())
	return &pb.Reply{Result: "Delete Completed"}, nil
}
func (m *myServer) Update(ctx context.Context, in *pb.InsDelUpdRequest) (*pb.Reply, error) {
	mysql.InsDelUpd("update", in.GetId(), in.GetName(), in.GetPrice(), in.GetTypeId(), in.GetCreateTime())
	return &pb.Reply{Result: "Update completed."}, nil
}
func (m *myServer) Select(ctx context.Context, in *pb.SelectRequest) (*pb.Reply, error) {
	result := mysql.Select(in.GetTable(), in.GetColumns(), in.GetCondition())
	return &pb.Reply{Result: result}, nil
}

func (m *myServer) ExecSql(ctx context.Context, in *pb.SqlRequest) (*pb.Reply, error) {
	mysql.ExecSql(in.GetSql())
	return &pb.Reply{Result: "Execution completed."}, nil
}
func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}

	server := grpc.NewServer()
	pb.RegisterOperationServer(server, &myServer{})

	fmt.Println("grpc service starts...")
	server.Serve(list)
}
