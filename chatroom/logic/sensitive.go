package logic

import (
	"strings"

	"github.com/wushengyouya/chatroom/global"
)

// 违禁词替换
func FilterSensitive(content string) string {
	for _, word := range global.SensitiveWords {
		content = strings.ReplaceAll(content, word, "**")
	}
	return content
}
