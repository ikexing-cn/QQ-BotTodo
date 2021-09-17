package tools

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func list(context *gin.Context) []string {
	sqlCrons := DBQuery("select * from cron order by entry_id")
	reply := ""
	for i, cron := range sqlCrons {
		reply = reply + "ID：" + strconv.Itoa(cron.EntryID) + ", 发送者：" + strconv.Itoa(cron.UserID) +
			", 剩余次数：" + strconv.Itoa(cron.Count) + ", 表达式：" + cron.Expression + ", 发送消息：" + cron.Message
		if i != len(sqlCrons)-1 {
			reply = reply + "\r\n"
		}
	}
	return skipCommand(context, reply)
}

func todo(context *gin.Context, splits []string) []string {
	if len(splits) == 4 {
		if len(strings.Split(splits[1], "/")) == 6 {
			_, err := strconv.Atoi(splits[2])
			if err != nil {
				return skipCommand(context, "参数需为数字")
			}
			return []string{splits[0], strings.Replace(splits[1], "/", " ", -1), splits[2], splits[3]}
		}
	}
	return errorCommand(context)
}

func remove(context *gin.Context, splits []string) []string {
	if len(splits) == 2 {
		sqlCrons := DBQuery("select * from cron where entry_id = " + splits[1])
		if len(sqlCrons) != 0 {
			DBDelete(splits[1])
			return skipCommand(context, "执行成功")
		}
		return skipCommand(context, "ID不正确")
	}
	return errorCommand(context)
}
