package tools

import (
	"strings"
	"unicode"

	pinyin "github.com/mozillazg/go-pinyin"
)

// TranslateToPinyin 将中文翻译成拼音
func TranslateToPinyin(str string) string {
	result := ""
	if IsChineseChar(str) { // 判断是否是中文
		// 默认
		a := pinyin.NewArgs()
		strArray := pinyin.Pinyin(str, a)

		for _, s := range strArray {
			for _, ss := range s {
				result += (ss + " ")
			}
		}
	} else {
		result = str
	}
	return strings.TrimSpace(result)
}

// IsChineseChar 判断是否是中文
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}
