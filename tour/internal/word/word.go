package word

import (
	"strings"
	"unicode"
)

// 将字符串转为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 将字符串转为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 下滑线转大写骆驼峰
func UnderscoreToUpperCameCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线转小写骆驼峰
func UnderscoreToLowerCameCase(s string) string {
	s = UnderscoreToUpperCameCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 骆驼峰转下划线
func CameCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
