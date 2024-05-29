package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"file-system/pkg/auth"
	"file-system/pkg/files"
	"file-system/pkg/users"
)

func NewRouter(userController users.UserController, fileController files.FileController) *gin.Engine {
	router := gin.Default()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	baseRouter := router.Group("/api/v1")
	authRouter := baseRouter.Group("/auth")
	protectedRouter := baseRouter.Group("/protected")
	userRouter := baseRouter.Group("/users")
	fileRouter := baseRouter.Group("/files")

	authRouter.POST("/login", auth.LoginHandler)
	authRouter.POST("/signup", userController.Create)

	protectedRouter.GET("", auth.ProtectedHandler, userController.List)
	userRouter.GET("/:userId", userController.Retrieve)
	userRouter.POST("", userController.Create)
	userRouter.PATCH("/:userId", userController.Update)
	userRouter.DELETE("/:userId", userController.Delete)

	fileRouter.GET("", fileController.List)
	fileRouter.GET("/:fileId", fileController.Retrieve)
	fileRouter.POST("", fileController.Create)
	fileRouter.PATCH("/:fileId", fileController.Update)
	fileRouter.DELETE("/:fileId", fileController.Delete)

	return router
}
