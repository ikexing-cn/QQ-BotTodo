package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

const address = "http://127.0.0.1:5700"
const howtouse = "格式: 秒/分钟/小时/天/月份/星期\r\n" +
	"例子: 20/10/*/*/*/6\r\n" +
	"#表示：星期六的每个小时的第10分钟的第20秒提醒" +
	"\r\n" + "\r\n" +
	"关于字符的解释请输入/character 或者 /char"

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
		Reply(context, howtouse)
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

func SendPrivateMsg(message string, userID string) string {
	return getRequest(map[string]string{"message": message, "user_id": userID}, "/send_private_msg")
}

func errorCommand(context *gin.Context) []string {
	Reply(context, "错误的指令")
	return []string{"skip"}
}

func getRequest(params map[string]string, suffix string) string {
	paramsTemp := url.Values{}
	Url, _ := url.Parse(address + suffix)
	for k, v := range params {
		paramsTemp.Set(k, v)
	}

	Url.RawQuery = paramsTemp.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", Url.String(), strings.NewReader(""))
	if err != nil {
		log.Println(err)
	}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
