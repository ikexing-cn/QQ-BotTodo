package tools

import (
	"strconv"

	"github.com/robfig/cron/v3"
)

var Cron = cron.New(cron.WithSeconds())

func CronRemoveFunc(entityID int) {
	Cron.Remove(cron.EntryID(entityID))
}

func CronAddFunc_(exps string, message string, userID string, count int) int {
	return CronAddFunc(exps, message, userID, count, true)
}

func CronAddFunc(exps string, message string, userID string, count int, isFirst bool) int {
	sql := "SELECT * FROM cron WHERE message = '" + message + "' AND user_id = '" + userID + "' AND expression = '" + exps + "'"
	query := DBQuery(sql)
	if len(query) == 0 || !isFirst {
		entryID, _ := Cron.AddFunc(exps, func() {
			SendPrivateMsg(message, userID)
			DBUpdateCount(message, userID, exps)
		})
		if isFirst {
			intID, _ := strconv.Atoi(userID)
			DBInsert(SqlCron{
				UserID:     intID,
				Count:      count,
				EntityID:   int(entryID),
				Expression: exps,
				Message:    message})
			SendPrivateMsg("数据已添加", userID)
		}
		return int(entryID)
	}
	return -1
}
