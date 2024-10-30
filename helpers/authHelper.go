package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MatchUserIDs(c *gin.Context, uid string) (err error) {
	userID := c.GetString("user_id")

	if userID != uid {
		err = fmt.Errorf("UnAuthenticated to access this resource")
		return err
	}

	return nil
}
