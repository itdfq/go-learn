package main

import (
	"fmt"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
//
//
//
// 示例 1：
//
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//
// 示例 2：
//
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//
//
//
// 提示：
//
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成
//
//
// Related Topics 字典树 字符串 👍 2724 👎 0

func main() {
	strs := []string{
		"ab", "a"}
	fmt.Println("开始执行")
	result := longestCommonPrefix(strs)
	fmt.Println(result)
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	min := len(strs[0])
	for _, s := range strs {
		if len(s) < min {
			min = len(s)
		}
	}
	i := 0

a:
	for ; i < min; i++ {
		ch := strs[0][i]
		for j := 1; j < len(strs); j++ {
			//如果第j个字符串的第i个字符和上一个不一致
			if ch != strs[j][i] {
				break a
			}
			//如果前面的都相同
			if ch == strs[j][i] && j == len(strs)-1 {
				continue a
			}
		}
	}
	return strs[0][:i]

}

//leetcode submit region end(Prohibit modification and deletion)
