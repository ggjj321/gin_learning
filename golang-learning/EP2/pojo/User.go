package pojo

import (
	"EP2/database"
	"fmt"
	"log"
)

type User struct {
	Id       int    `json:"UserId"`
	Name     string `json:"UserName"`
	Password string `json:"UserPassword"`
	Email    string `json:"UserEmail"`
}

func FindAllUsers() []User {
	var users []User
	database.DBConnect.Table("\"Demo\".users").Find(&users)
	fmt.Println("access")
	return users
}

func FindByUserId(userId string) User {
	var user User
	database.DBConnect.Table("\"Demo\".users").Where("id = ?", userId).First(&user)
	return user
}

func CreateUser(user User) User {
	database.DBConnect.Table("\"Demo\".users").Create(&user)
	return user
}

func DeleteUser(userId string) bool {
	user := User{}
	result := database.DBConnect.Table("\"Demo\".users").Where("id = ?", userId).Delete(&user)
	log.Println(result)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func UpdateUser(userId string, user User) User {
	// 比較 update updates
	database.DBConnect.Table("\"Demo\".users").Model(&user).Where("id = ?", userId).Updates(user)
	return user
}
