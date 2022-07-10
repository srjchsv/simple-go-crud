package handlers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/srjchsv/simple-go-crud/internal/db"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func Home(ctx *gin.Context) {
	db := db.SqlConnect()
	var users []User

	db.Order("created_at asc").Find(&users)
	defer db.Close()

	ctx.HTML(200, "index.html", gin.H{
		"users": users,
	})
}

func CreateUser(ctx *gin.Context) {
	db := db.SqlConnect()
	defer db.Close()

	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	log.Printf("created user %v with %s email", name, email)
	db.Create(&User{Name: name, Email: email})

	ctx.Redirect(302, "/")
}

func DeleteUser(ctx *gin.Context) {
	db := db.SqlConnect()
	defer db.Close()

	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		ctx.String(400, "id is not a number")
		return
	}
	var user User
	db.First(&user, id)
	db.Delete(&user)

	ctx.Redirect(302, "/")
}
