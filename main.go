package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"net"
	"time"
	userServiceGo "user-service-go/pb/user-service-go"
	"user-service-go/service"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	db, err := sql.Open("mysql", "root:rootroot@tcp(localhost:3306)/users")
	if err != nil {
		glog.Fatalln(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
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
