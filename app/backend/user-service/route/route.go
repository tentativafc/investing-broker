package route

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/tentativafc/investing-broker/app/backend/user-service/dto"
	errUs "github.com/tentativafc/investing-broker/app/backend/user-service/error"
	"github.com/tentativafc/investing-broker/app/backend/user-service/service"
)

var gus service.UserService

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case errUs.IWithMessageAndStatusCode:
					error := err.(errUs.IWithMessageAndStatusCode)
					c.JSON(error.Status(), dto.Error{Message: error.Error(), Cause: error.Cause(), StackTrace: error.StackTrace()})
				default:
					c.JSON(http.StatusInternalServerError, dto.Error{Message: err.(error).Error(), StackTrace: string(debug.Stack())})
				}
			}
		}()
		c.Next()
	}
}

func CreateUser(c *gin.Context) {
	var req dto.User
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errUs.NewBadRequestError("Error to parse body.", err))
	}
	resp, err := gus.CreateUser(req)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusCreated, &resp)
	}
}

func UpdateUser(c *gin.Context) {
	var req dto.UserUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errUs.NewBadRequestError("Error to parse body.", err))
	}
	uId := c.Params.ByName("id")
	if _, ok := c.Request.Header["Authorization"]; !ok {
		panic(errUs.NewAuthError("Authorization header not found.", nil))
	}
	authorization := c.Request.Header["Authorization"][0]

	req.ID = uId
	resp, err := gus.UpdateUser(req, authorization)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, &resp)
	}
}

func Login(c *gin.Context) {
	var req dto.LoginData
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errUs.NewBadRequestError("Error to parse body.", err))
	}
	resp, err := gus.Login(req)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, &resp)
	}
}

func RecoverLogin(c *gin.Context) {
	var req dto.RecoverLoginData
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errUs.NewBadRequestError("Error to parse body.", err))
	}

	resp, err := gus.RecoverLogin(req)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, &resp)
	}
}

func GetUserById(c *gin.Context) {
	uId := c.Params.ByName("id")

	if _, ok := c.Request.Header["Authorization"]; !ok {
		panic(errUs.NewAuthError("Authorization header not found.", nil))
	}
	authorization := c.Request.Header["Authorization"][0]
	resp, err := gus.GetuserById(authorization, uId)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, &resp)
	}
}

func CreateRoutes(us service.UserService) {
	gus = us

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(Recovery())

	r.POST("/api/users/login", Login)
	r.GET("/api/users/:id", GetUserById)
	r.POST("/api/users/recover", RecoverLogin)
	r.POST("/api/users", CreateUser)
	r.PUT("/api/users/:id", UpdateUser)

	r.Run(":8081")
}
