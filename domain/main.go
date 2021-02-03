package main

import (
	"github.com/jinzhu/gorm"
	"github.com/pedrofelli/imersao/codepix-go/application/grpc"
	"github.com/pedrofelli/imersao/codepix-go/infrastructure/db"
	"os"
)

var database *gorm.DB

func main(){
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
