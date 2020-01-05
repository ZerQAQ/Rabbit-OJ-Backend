package contest

import (
	"Rabbit-OJ-Backend/models"
	"Rabbit-OJ-Backend/models/responses"
	"github.com/jinzhu/gorm"
)

func RegenerateUserScore(tx *gorm.DB, cid, uid string) error {
	contestInfo, err := Info(cid)
	if err != nil {
		return err
	}

	questionMapTidToId, contestQuestion, err := QuestionMapTidToId(cid)
	if err != nil {
		return err
	}

	progress := make([]responses.ScoreBoardProgress, len(questionMapTidToId))
	var contestSubmissionList []models.ContestSubmission

	if err := tx.Table("contest_submission").
		Where("cid = ? AND uid = ?", cid, uid).
		Find(&contestSubmissionList).
		Error; err != nil {
		return err
	}

	// calc Number
	for _, item := range contestSubmissionList {
		questionId := questionMapTidToId[item.Tid]

		if item.Status == StatusPending {
			continue
		} else if item.Status == StatusAC {
			progress[questionId].Status = StatusAC

			totalTime := &progress[questionId].TotalTime
			if *totalTime == 0 || *totalTime > item.TotalTime {
				*totalTime = item.TotalTime
			}
		} else if item.Status == StatusERR {
			progress[questionId].Bug++
		}
	}

	score, totalTime := uint32(0), int64(0)
	// calc penalty
	for i, item := range progress {
		if item.Status == StatusAC {
			score += contestQuestion[i].Score

			currentTime := item.TotalTime + contestInfo.Penalty*int64(item.Bug)
			if currentTime > totalTime {
				totalTime = currentTime
			}
		}
	}

	if err := tx.Table("contest_user").
		Where("cid = ? AND uid = ?", cid, uid).
		Updates(map[string]interface{}{
			"score":      score,
			"total_time": totalTime,
		}).Error; err != nil {
		return err
	}

	return nil
}
