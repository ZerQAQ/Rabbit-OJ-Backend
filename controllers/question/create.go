package question

import (
	"Rabbit-OJ-Backend/controllers/auth"
	"Rabbit-OJ-Backend/models/forms"
	QuestionService "Rabbit-OJ-Backend/services/question"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	questionForm := &forms.QuestionEditForm{}
	authObject, err := auth.GetAuthObjRequireAdmin(c)
	if err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
		})

		return
	}

	if err := c.BindJSON(&questionForm); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	tid, err := QuestionService.Create(
		authObject.Uid,
		questionForm.Subject,
		questionForm.Content,
		questionForm.Difficulty,
		questionForm.TimeLimit,
		questionForm.SpaceLimit,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": tid,
	})
}
