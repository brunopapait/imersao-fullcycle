package main

import (
	"github.com/brunopapait/imersao-fullcycle/codepix/application/grpc"
	"github.com/brunopapait/imersao-fullcycle/codepix/infrastructure/db"
	"github.com/jinzhu/gorm"
	"os"
)

var database *gorm.DB

func main(){
	database = db.ConnectDB(os.Getenv("env"))

	grpc.StartGrpcServer(database, 50051)
}
