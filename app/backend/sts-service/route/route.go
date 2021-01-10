package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/dto"
	"github.com/tentativafc/investing-broker/app/backend/sts-service/service"
)

func CreateRoutes(s service.StsService) {

	r := gin.Default()
	r.POST("/client_credentials", func(c *gin.Context) {

		var req dto.ClientCredentialsRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		cc, err := s.CreateClientCredentials(req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		c.JSON(http.StatusOK, cc)

	})

	r.Run()
}
