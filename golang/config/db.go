package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	uri := "mongodb+srv://Byambasuren11:80664525Bn$@cluster0.veh63.mongodb.net/Userss?retryWrites=true&w=majority&appName=Cluster0"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("❌ Mongo client үүсгэхэд алдаа:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("❌ Mongo холболт амжилтгүй:", err)
	}

		DB = client.Database("Userss")
		fmt.Println("✅ MongoDB холбогдлоо")
	}
// package config

// import (
// 	"fmt"
// 	"log"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDB() {
// 	dsn := "host=localhost user=postgres password=mypassword dbname=postgres-db port=5432 sslmode=disable"
// 	var err error
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("❌ Database connection error: ", err)
// 	}

// 	fmt.Println("✅ Database connected")
// }
