package cron

import (
	"strconv"

	"github.com/robfig/cron/v3"
	"todo-everyday/src/utils"
)

var Cron = cron.New(cron.WithSeconds())

func AddFunc(exps string, message string, userID string, count int) int {
	sql := "SELECT * FROM cron WHERE message = " + exps + " AND user_id = " + userID + " AND expression = " + exps
	query := utils.DBQuery(sql)
	if len(query) == 0 {
		entryID, _ := Cron.AddFunc(exps, func() {
			utils.SendPrivateMsg(message, userID)
		})
		intID, _ := strconv.Atoi(userID)
		utils.DBInsert(utils.SqlCron{
			UserID:     intID,
			Count:      count,
			EntityID:   int(entryID),
			Expression: exps,
			Message:    message})
		utils.SendPrivateMsg("数据已添加", userID)
		return int(entryID)
	}
	return -1
}
