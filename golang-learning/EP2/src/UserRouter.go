package src

import (
	"EP2/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")

	user.GET("/", service.FindAllUsers)
	user.GET("/:id", service.FindByUserId)
	user.POST("/", service.PostUser)

	user.DELETE("/:id", service.DeleteUser)
	user.PUT("/:id", service.PutUser)
}

// 用 git 紀錄
