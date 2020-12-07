package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"usercenter/handler/roleHandler"
	"usercenter/handler/userHandler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", index)
	router.GET("ping", ping)

	userRouter := router.Group(`users`)
	{
		userRouter.GET(``, userHandler.UserList)
		userRouter.GET(`/`, userHandler.UserList)
		userRouter.GET("/:id", userHandler.GetById)
		userRouter.PUT("/:id", userHandler.UpdateUser)
		userRouter.POST("", userHandler.AddUser)
		userRouter.DELETE("/:id", userHandler.DeleteUser)
		userRouter.PUT("/:id/attr/password", userHandler.UpdatePassword)
		userRouter.GET("/:id/roles", userHandler.GetAllRoleWithCheckedByUserId)
		userRouter.PUT("/:id/roles", userHandler.UpdateUserRole)
	}

	roleRouter := router.Group("roles")
	{
		roleRouter.GET("", roleHandler.RoleList)
		roleRouter.GET("/", roleHandler.RoleList)
		roleRouter.POST("", roleHandler.AddRole)
		roleRouter.GET("/:id", roleHandler.GetById)
		roleRouter.PUT("/:id", roleHandler.UpdateRole)
		roleRouter.DELETE("/:id", roleHandler.DeleteById)
	}

	return router
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "hello gin")
}
