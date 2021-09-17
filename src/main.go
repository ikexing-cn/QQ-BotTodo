package main

import (
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"todo-everyday/src/tools"
)

func main() {
	Init()
	r := gin.Default()
	r.POST("/", func(context *gin.Context) {
		dataRender := context.Request.Body
		rawData, _ := ioutil.ReadAll(dataRender)
		json := string(rawData)
		postType := gjson.Get(json, "post_type").String()
		if postType == "message" {
			MessageExec(json, context)
		}
	})
	_ = r.Run(":5701")
}

func Init() {
	tools.DBInit()
	tools.DBOrderByEntryID()
	tools.CronInit()
}

func MessageExec(json string, context *gin.Context) {
	userID := gjson.Get(json, "user_id").String()
	subType := gjson.Get(json, "sub_type").String()
	message := gjson.Get(json, "message").String()

	if subType != "friend" {
		return
	}

	if message[0:1] == "/" {
		result := tools.CheckCommand(message[1:], context)
		if result[0] == "skip" {
			return
		}
		if result[0] == "todo" {
			atoi, _ := strconv.Atoi(result[2])
			if tools.CronAddFunc_(result[1], result[3], userID, atoi) == -1 {
				tools.Reply(context, "此ToDo已经存在于数据库，请勿重复添加")
			}
		}
	}
}
