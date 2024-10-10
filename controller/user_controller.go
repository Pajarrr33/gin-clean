package controller

import (
	"gin-db/model"
	"gin-db/response"
	"gin-db/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
}

type userController struct {
	userUsecase 		usecase.UserUsecase
	crendetialUsecase 	usecase.CredentialUsecase
}

func NewUserController(userUsecase usecase.UserUsecase,crendetialUsecase usecase.CredentialUsecase) UserController {
	return &userController{userUsecase: userUsecase,crendetialUsecase: crendetialUsecase}
}

func (uc *userController) CreateUser(ctx *gin.Context) {
	var user model.User
	
	err := ctx.ShouldBind(&user)

	// Validate if input is valid
	if err != nil {
		response := response.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Input not valid",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate if there are any value with zero value
	if user.Age == 0 || user.Gender == "" || user.Name == "" {
		response := response.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Name, age, and gender is required",
		}
		ctx.JSON(http.StatusBadRequest,response)
		return
	}

	// Get Credential By Email
	email,_,_ := ctx.Request.BasicAuth()

	user.Credential,err = uc.crendetialUsecase.GetCredentialByEmail(email,user.Credential)
	if err != nil {
		response := response.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Email not found",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	createdUser,err := uc.userUsecase.CreateUser(&user)
	if err != nil {
		response := response.ResponseError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	createdUser.Credential = user.Credential

	// Response
	response := response.ResponseUserSingle{
		Code: http.StatusCreated,
		Data: user,
	}
	ctx.JSON(http.StatusCreated,response)

}