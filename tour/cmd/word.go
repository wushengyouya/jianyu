package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/wushengyouya/tour/internal/word"
)

const (
	ModeUppder                    = iota + 1 // 全部转大写
	ModeLower                                // 全部转小写
	ModeUnderscoreToUpperCameCase            // 下划线转大写骆驼峰
	ModeUnderscoreToLowerCase                // 下换线转小写骆驼峰
	ModeCameCaseToUnderscore                 // 驼峰转下划线
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1：全部转大写",
	"2：全部转小写",
	"3：下划线转大写驼峰",
	"4：下划线转小写驼峰",
	"5：驼峰转下划线",
}, "\n")
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUppder:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCameCase:
			content = word.UnderscoreToUpperCameCase(str)
		case ModeUnderscoreToLowerCase:
			content = word.UnderscoreToLowerCameCase(str)
		case ModeCameCaseToUnderscore:
			content = word.CameCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换格式，请执行 help word 查看帮助文档")
		}
		log.Printf("输出结果: %s", content)
	},
}
var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}
