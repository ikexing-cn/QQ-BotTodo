package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckCommand(str string, context *gin.Context) []string {
	split := strings.Split(str, " ")

	switch strings.ToLower(split[0]) {
	case "todo":
		if len(strings.Split(split[1], "/")) != 6 {
			return errorCommand(context)
		} else {
			return []string{split[0], strings.Replace(split[1], "/", " ", -1), split[2]}
		}
	case "howtouse", "htu":
		Reply(context, "格式: 秒/分钟/小时/天/月份/星期 \r\n例子: 20/10/*/*/*/6 \r\n"+
			"#表示：星期六的每个小时的第10分钟的第20秒提醒")
		return []string{"skip"}
	default:
		return errorCommand(context)
	}
}

func Reply(context *gin.Context, message string) {
	context.JSON(http.StatusOK, gin.H{
		"reply": message,
	})
}

func errorCommand(context *gin.Context) []string {
	Reply(context, "错误的指令")
	return []string{"skip"}
}
