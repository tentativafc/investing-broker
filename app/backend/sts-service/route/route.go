package route

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/dto"
	errSts "github.com/tentativafc/investing-broker/app/backend/sts-service/error"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/service"
)

var gss service.StsService

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case errSts.IWithMessageAndStatusCode:
					error := err.(errSts.IWithMessageAndStatusCode)
					c.JSON(error.Status(), dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
				default:
					c.JSON(http.StatusInternalServerError, dto.Error{Message: err.(error).Error(), StackTrace: string(debug.Stack())})
				}
			}
		}()
		c.Next()
	}
}

func GenerateClientCredentials(c *gin.Context) {
	var req dto.ClientCredentialsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errSts.NewBadRequestError("Error to parse body.", err))
	}
	cc, err := gss.GenerateClientCredentials(req)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, cc)
	}
}

func GenerateToken(c *gin.Context) {
	var req dto.TokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errSts.NewBadRequestError("Error to parse body.", err))
	}
	token, err := gss.GenerateToken(req)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, dto.TokenResponse{Token: token})
	}
}

func ValidateToken(c *gin.Context) {
	var req dto.ValidateTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errSts.NewBadRequestError("Error to parse body.", err))
	}
	resp, err := gss.ValidateToken(req)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, resp)
	}
}

func CreateRoutes(ss service.StsService) {

	gss = ss

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(Recovery())

	r.POST("/client_credentials", GenerateClientCredentials)
	r.POST("/token", GenerateToken)
	r.POST("/validate_token", ValidateToken)

	r.Run(":8080")
}
