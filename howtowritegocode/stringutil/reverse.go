// 包 stringutil 主要包括处理字符串的工具函数
package stringutil

// 将字符串从头到尾调换顺序
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
