package middleware

import (
	"fmt"
	"gin-db/config"
	"gin-db/model"
	"gin-db/response"
	"net/http"
	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		email, password, status := ctx.Request.BasicAuth()

		// If BasicAuth status is false, return 401 Unauthorized and stop further processing
		fmt.Println(status)
		if !status {
			ctx.JSON(http.StatusUnauthorized, response.ResponseError{
                Code: http.StatusUnauthorized,
                Message: "Unauthorized",
            })
            ctx.Abort() // Prevent further handlers from being executed
			return
		}

		// Get Credential by email 
		credential := model.Credential{}

		db := config.ConnectDb()

		defer db.Close()

		query := "SELECT credential_id,email,password FROM credential WHERE email = $1"

		err := db.QueryRow(query,email).Scan(&credential.Credential_id,&credential.Email,&credential.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.ResponseError{
                Code: http.StatusInternalServerError,
                Message: "Account Not Found",
            })
            ctx.Abort() // Prevent further handlers from being executed
			return
		}

		if password != credential.Password {
			// Password is incorrect, return 400 Bad Request
			ctx.JSON(http.StatusBadRequest, response.ResponseError{
				Code: http.StatusBadRequest,
				Message: "Wrong Password",
			})
			ctx.Abort()  // Prevent further handlers from being executed
			return
		}

		ctx.Next()
	}
}