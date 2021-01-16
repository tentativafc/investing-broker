package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tentativafc/investing-broker/app/backend/user-service/dto"
	errUs "github.com/tentativafc/investing-broker/app/backend/user-service/error"
	"github.com/tentativafc/investing-broker/app/backend/user-service/repo"
	"github.com/tentativafc/investing-broker/app/backend/user-service/service"
)

var ur repo.UserRepository = repo.NewUserRepository()
var us service.UserService = service.NewUserService(ur)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case *errUs.AuthError:
					error := err.(*errUs.AuthError)
					c.JSON(error.Code(), gin.H{"error": error.Error()})
				case *errUs.NotFoundError:
					error := err.(*errUs.NotFoundError)
					c.JSON(error.Code(), gin.H{"error": error.Error()})
				case *errUs.BadRequestError:
					error := err.(*errUs.BadRequestError)
					c.JSON(error.Code(), gin.H{"error": error.Error()})
				case *errUs.GenericError:
					error := err.(*errUs.GenericError)
					c.JSON(error.Code(), gin.H{"error": error.Error()})
				}
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.(error).Error()})
			}
		}()
		c.Next()
	}
}

func CreateUser(c *gin.Context) {
	var req dto.User
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errUs.NewBadRequestError("Error to parse body."))
	}
	resp, err := us.CreateUser(req)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusCreated, &resp)
	}
}

func UpdateUser(c *gin.Context) {
	var req dto.UserUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errUs.NewBadRequestError("Error to parse body."))
	}
	uId := c.Params.ByName("id")
	authorization := c.Request.Header["Authorization"][0]
	req.ID = uId
	resp, err := us.UpdateUser(req, authorization)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, &resp)
	}
}

func Login(c *gin.Context) {
	var req dto.LoginData
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errUs.NewBadRequestError("Error to parse body."))
	}
	resp, err := us.Login(req)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, &resp)
	}
}

func RecoverLogin(c *gin.Context) {
	var req dto.RecoverLoginData
	if err := c.ShouldBindJSON(&req); err != nil {
		panic(errUs.NewBadRequestError("Error to parse body."))
	}

	resp, err := us.RecoverLogin(req)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, &resp)
	}
}

func GetUserById(c *gin.Context) {
	uId := c.Params.ByName("id")
	authorization := c.Request.Header["Authorization"][0]
	resp, err := us.GetuserById(authorization, uId)
	if err != nil {
		panic(err)
	} else {
		c.JSON(http.StatusOK, &resp)
	}
}

func HandleRequests() {

	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(Recovery())

	r.POST("/users/login", Login)
	r.GET("/users/{id}", GetUserById)
	r.POST("/users/recover", RecoverLogin)
	r.POST("/users", CreateUser)
	r.PUT("/users/{id}", UpdateUser)

	r.Run(":8080")
}
