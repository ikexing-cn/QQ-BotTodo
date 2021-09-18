package tools

import (
	"strconv"

	"github.com/robfig/cron/v3"
)

var Cron = cron.New(cron.WithSeconds())

func CronInit() {
	Cron.Start()
	sqlCrons := DBQuery("select * from cron order by entry_id")
	for _, cron_ := range sqlCrons {
		CronAddFunc(cron_.Expression, cron_.Message, strconv.Itoa(cron_.UserID), cron_.Count, false)
	}
}

func CronRemove(entryID string) {
	atoi, err := strconv.Atoi(entryID)
	CheckErr(err)
	Cron.Remove(cron.EntryID(atoi))
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
				EntryID:    int(entryID),
				Expression: exps,
				Message:    message})
			SendPrivateMsg("数据已添加", userID)
		}
		return int(entryID)
	}
	return -1
}
