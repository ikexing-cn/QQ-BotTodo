package main

import (
	"strings"
)

type TodoType int32

const (
	Monday    TodoType = 0
	Tuesday   TodoType = 1
	Wednesday TodoType = 2
	Thursday  TodoType = 3
	Friday    TodoType = 4
	Saturday  TodoType = 5
	Sunday    TodoType = 6
	EveryDay  TodoType = 7
	Error     TodoType = -1
)

func getType(t string) TodoType {
	switch strings.ToLower(t) {
	case "everyday":
		return EveryDay
	case "monday":
		return Monday
	case "tuesday":
		return Tuesday
	case "wednesday":
		return Wednesday
	case "thursday":
		return Thursday
	case "friday":
		return Friday
	case "saturday":
		return Saturday
	case "sunday":
		return Sunday
	default:
		return Error
	}
}

func CheckCommand(str string) []string {
	split := strings.Split(str, " ")

	if split[0] == "todo" {
		if getType(split[1]) == Error {
			return []string{"不存在此日期"}
		}
		if !strings.Contains(split[2], ":") {
			return []string{"时间不正确"}
		}
	} else {
		return []string{"错误的指令"}
	}
	return []string{split[1], split[2], split[3]}
}
