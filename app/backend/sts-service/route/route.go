package route

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/dto"
	errSts "github.com/tentativafc/investing-broker/app/backend/sts-service/error"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/service"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				switch err.(type) {
				case *errSts.AuthError:
					error := err.(*errSts.AuthError)
					c.JSON(error.Code(), dto.Error{Error: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
				case *errSts.NotFoundError:
					error := err.(*errSts.NotFoundError)
					c.JSON(error.Code(), dto.Error{Error: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
				case *errSts.BadRequestError:
					error := err.(*errSts.BadRequestError)
					c.JSON(error.Code(), dto.Error{Error: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
				case *errSts.GenericError:
					error := err.(*errSts.GenericError)
					c.JSON(error.Code(), dto.Error{Error: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
				default:
					c.JSON(http.StatusInternalServerError, dto.Error{Error: err.(error).Error(), StackTrace: string(debug.Stack())})
				}
			}
		}()
		c.Next()
	}
}

func CreateRoutes(s service.StsService) {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(Recovery())

	r.POST("/client_credentials", func(c *gin.Context) {
		var req dto.ClientCredentialsRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(errSts.NewBadRequestError("Error to parse body.", err))
		}
		cc, err := s.GenerateClientCredentials(req)
		if err != nil {
			panic(err)
		} else {
			c.JSON(http.StatusOK, cc)
		}
	})

	r.POST("/token", func(c *gin.Context) {
		var req dto.TokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(errSts.NewBadRequestError("Error to parse body.", err))
		}
		token, err := s.GenerateToken(req)
		if err != nil {
			panic(err)
		} else {
			c.JSON(http.StatusOK, dto.TokenResponse{Token: token})
		}
	})

	r.POST("/validate_token", func(c *gin.Context) {
		var req dto.ValidateTokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(errSts.NewBadRequestError("Error to parse body.", err))
		}
		resp, err := s.ValidateToken(req)
		if err != nil {
			panic(err)
		} else {
			c.JSON(http.StatusOK, resp)
		}
	})

	r.Run(":8080")
}
