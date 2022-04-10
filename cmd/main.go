package cmd

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/ryatsuga/codepix-go/application/grpc"
	"github.com/ryatsuga/codepix-go/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
