package routes

import (
	"errors"
	"roles/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func helperGetDB(c *gin.Context) (*gorm.DB, error) {
	obj, ok := c.Get("db")

	if !ok {
		return nil, errors.New("no database available in context")
	}
	provider := obj.(model.Provider)
	return provider.GetDB()
}
