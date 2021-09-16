package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	address  = "http://127.0.0.1:5700"
	howtouse = "格式: 秒/分钟/小时/天/月份/星期\r\n" +
		"例子: 20/10/*/*/*/6\r\n" +
		"#表示：星期六的每个小时的第10分钟的第20秒提醒" +
		"\r\n" + "\r\n" +
		"关于字符的解释请输入/character 或者 /char"
	char = "星号 (*) : 将匹配允许值范围内的所有值.例如表达式是(* 5 * * * *) 那么代表每个小时的第五分钟内的每秒都会执行你指定的函数\r\n\r\n" +
		"斜杠 (/) : 斜杠代表允许值的增量.例如表达式是(3-59⁄15 * * * * *) 那么代表每分钟的第三秒和从第三秒开始往后的每15秒钟(直到第59秒结束)都会执行.\r\n" +
		"再举个例子,例如表达式为 (5⁄15 * * * * *)那么每一分钟的第5秒和以5秒开始往后的每15秒执行一次 (就是每分钟内的第5,20,35,50秒都会执行)\r\n\r\n" +
		"逗号 (,) : 分割列表中的项目. 例如表达式为 (* * * * * MON,WED,FRI) 则代表每个星期的星期一、星期三和星期五都会执行你指定的函数\r\n\r\n" +
		"横杠 (-) : 用于定义一个范围内(但必须在允许值范围内). 例如表达式为 (35-55 * * * * *) 那么在每分钟的第35到55秒内的每秒都会执行你指定的函数\r\n\r\n" +
		"问号 (?) : 可以使用?代替*代表月份中某一天或者某一个星期空白 (我自己不是太懂这个意思 XD)\r\n"
)

func CheckCommand(str string, context *gin.Context) []string {
	splits := strings.Split(str, " ")

	switch strings.ToLower(splits[0]) {
	case "todo":
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
	case "howtouse", "htu":
		return skipCommand(context, howtouse)
	case "char", "character":
		return skipCommand(context, char)
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

func skipCommand(context *gin.Context, text string) []string {
	Reply(context, text)
	return []string{"skip"}
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

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
