package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/srjchsv/simple-go-crud/internal/db"
	"github.com/srjchsv/simple-go-crud/internal/handlers"
	"github.com/srjchsv/simple-go-crud/internal/router"
)

func main() {
	db := db.SqlConnect()
	db.AutoMigrate(&handlers.User{})
	defer db.Close()

	r := router.Register()
	r.Run()
}
