package contest

import (
	ContestService "Rabbit-OJ-Backend/services/contest"
	"github.com/gin-gonic/gin"
)

func Question(c *gin.Context) {
	cid := c.Param("cid")
	questions, err := ContestService.QuestionExtended(cid)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": questions,
	})
}