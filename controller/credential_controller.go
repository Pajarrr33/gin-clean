package controller

import (
	"fmt"
	"gin-db/model"
	"gin-db/response"
	"gin-db/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CredentialController interface{
	Register(ctx *gin.Context)
}

type credentialController struct {
	usecase usecase.CredentialUsecase
}

func NewCredentialController (usecase usecase.CredentialUsecase) CredentialController {
	return &credentialController{usecase: usecase}
}

func (cc *credentialController) Register(ctx *gin.Context) {
	var credential model.Credential

	err := ctx.ShouldBind(&credential)

	// Validate if input is valid
	if err != nil {
		response := response.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Input not valid",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// Validate if email or password empty 
	if credential.Email == "" || credential.Password == "" {
		response := response.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Email and password is required",
		}
		ctx.JSON(http.StatusBadRequest,response)
		return
	}

	// Check if email is exist 
	isEmailExist,err := cc.usecase.IsEmailExist(credential.Email)

	// Check if error appear
	if err != nil {
		response := response.ResponseError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError,response)
		return
	}

	fmt.Println(isEmailExist)
	if isEmailExist {
		response := response.ResponseError{
			Code:    http.StatusBadRequest,
			Message: "Email is already exist",
		}
		ctx.JSON(http.StatusBadRequest,response)
		return
	}

	createdCredential,err := cc.usecase.Register(&credential)
	if err != nil {
		response := response.ResponseError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError,response)
		return
	}

	response := response.ResponseCredentialSingle{
		Code: http.StatusCreated,
		Data: *createdCredential,
	}
	ctx.JSON(http.StatusCreated,response)
}