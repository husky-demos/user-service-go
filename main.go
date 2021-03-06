package main

import (
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"net"
	userServiceGo "user-service-go/pb/user-service-go"
	"user-service-go/service"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	db, err := sqlx.Open("mysql", "root:rootroot@tcp(localhost:3306)/users")
	if err != nil {
		glog.Fatalln(err)
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)

	lis, err := net.Listen("tcp", "0.0.0.0:9002")
	if err != nil {
		glog.Fatalln(err)
	}
	server := grpc.NewServer()
	userServiceGo.RegisterUserServiceServer(server, service.NewUserService(db))
	if err = server.Serve(lis); err != nil {
		glog.Fatalln(err)
	}
}
