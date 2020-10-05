package question

/**
请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。

给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。
*/
func Reversess(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	str := []rune(s) // 能处理中文字符；切片可以交换位置
	for i := 0; i < l/2; i++ {
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}
	return string(str)
}
