package question

import (
	"Rabbit-OJ-Backend/db"
	"fmt"
	"github.com/jinzhu/gorm"
)

func UpdateAttemptCount(tid string) {
	fmt.Println("[ATTEMPT] INCREMENT COUNT " + tid)
	if err := db.DB.Table("question").
		Update("attempt", gorm.Expr("attempt + 1")).
		Where("tid = ?", tid).Error;
		err != nil {

		fmt.Println(err)
	}
}

func UpdateAcceptedCount(tid string) {
	if err := db.DB.Table("question").
		Update("accepted", gorm.Expr("accepted + 1")).
		Where("tid = ?", tid).Error;
		err != nil {

		fmt.Println(err)
	}
}
