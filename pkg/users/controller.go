package users

import (
	golog "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"file-system/pkg/helpers"
	jwtToken "file-system/pkg/jwt"
	"file-system/pkg/response"
)

type UserController struct {
	service UserService
}

func NewUserController(service UserService) UserController {
	return UserController{
		service: service,
	}
}

func (controller *UserController) Create(ctx *gin.Context) {
	log.Info().Msg("create user")
	createUserRequest := CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helpers.ErrorHelper(err)

	result := controller.service.Create(createUserRequest)

	userClaims := jwtToken.UserClaims{
		ID:        result.ID.String(),
		Username:  result.UserName,
		Role:      result.Role,
		CreatedAt: result.CreatedAt.String(),
	}

	tokenString, err := jwtToken.CreateToken(userClaims)

	if err != nil {
		helpers.ErrorHelper(err)
	}
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	Response := response.Response{
		Code:   http.StatusNoContent,
		Status: "success",
		Data:   map[string]interface{}{"token": tokenString},
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, Response)
}

// Delete implements UserService.
func (controller *UserController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete user")
	userId := ctx.Param("userId")

	controller.service.Delete(userId)
	Response := response.Response{
		Code:   http.StatusNoContent,
		Status: "success",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusNoContent, Response)
}

// List implements UserService.
func (controller *UserController) List(ctx *gin.Context) {
	log.Info().Msg("list user")
	users := controller.service.List()

	Response := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   users,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, Response)
}

// Retrieve implements UserService.
func (controller *UserController) Retrieve(ctx *gin.Context) {
	log.Info().Msg("list user")
	userId := ctx.Param("userId")

	user := controller.service.Retrieve(userId)
	Response := response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   user,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, Response)
}

// Update implements UserService.
func (controller *UserController) Update(ctx *gin.Context) {
	log.Info().Msg("update user")
	userId := ctx.Param("userId")
	updateUserRequest := UpdateUserRequest{}

	err := ctx.ShouldBindJSON(&updateUserRequest)
	helpers.ErrorHelper(err)

	golog.Println(updateUserRequest)

	controller.service.Update(userId, updateUserRequest)
	Response := response.Response{
		Code:   http.StatusNoContent,
		Status: "success",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusNoContent, Response)

}
