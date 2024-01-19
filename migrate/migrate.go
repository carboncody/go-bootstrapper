package main

import (
	"fmt"
	"log"

	"github.com/carboncody/go-bootstrapper/initializers"
	"github.com/carboncody/go-bootstrapper/models"
)

func init() {
	config, err := initializers.LoadConfig()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&models.User{}, &models.Task{}, &models.TaskLabel{})
	fmt.Println("👍 Migration complete")
}
