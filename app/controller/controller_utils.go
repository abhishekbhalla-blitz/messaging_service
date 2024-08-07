package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func convertBodyFromJson(target interface{}, c *gin.Context) error {
	if err := c.BindJSON(target); err != nil {
		log.Error("Error converting body to json. Error: ", err)
		return err
	}
	return nil
}
