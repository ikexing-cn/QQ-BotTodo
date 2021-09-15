package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func main() {
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

func MessageExec(json string, context *gin.Context) {
	//userID := gjson.Get(json, "user_id").String()
	subType := gjson.Get(json, "sub_type").String()
	message := gjson.Get(json, "message").String()

	if subType != "friend" {
		return
	}

	if message[0:1] == "/" {
		command := CheckCommand(message[1:])
		context.JSON(http.StatusOK, gin.H{
			"reply": command[len(command)-1],
		})
	}
}