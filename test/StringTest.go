package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "INSERT INTO `trade_online`.`aliexpress_full_manage_jit_stock` (`id`, `category_id`, `jit_stock`, `modify_time`, `modifier_id`, `modifier`, `create_time`, `creator`, `creator_id`) VALUES (3, 201371403, 0, NULL, NULL, NULL, '2023-07-27 18:50:03', 'system', -1);" +
		"\nINSERT INTO `trade_online`.`aliexpress_full_manage_jit_stock` (`id`, `category_id`, `jit_stock`, `modify_time`, `modifier_id`, `modifier`, `create_time`, `creator`, `creator_id`) VALUES (4, 201380001, 0, NULL, NULL, NULL, '2023-07-27 18:50:03', 'system', -1);\n"
	str := replace(s)
	fmt.Println("替换后的结果为：", str)
}

/**
字符串替换
自动匹配字符串，去除括号里面的id 以及trade_online
*/
func replace(str string) string {
	str = strings.Trim(str, "\n")
	a1 := strings.ReplaceAll(str, "`trade_online`.", "")
	a2 := strings.ReplaceAll(a1, "`id`,", "")
	arrayStr := strings.Split(a2, ";")
	var rsult strings.Builder
	for _, as := range arrayStr {
		len := len(as)
		if len < 5 {
			continue
		}
		//println("字符串总长度为", len)
		//获取最后一个（
		lastkuohao := strings.LastIndex(as, "(")
		newStr := as[lastkuohao:]
		//println("新字符串为：", newStr)
		//print(" (的指针为：", lastkuohao)
		end := strings.Index(newStr, ",")
		//println("逗号位置", end)
		end = end + lastkuohao + 2
		//print(" ,号的指针为：", end)

		//println("当前字符串为：", as)

		qian := as[:lastkuohao]
		//print("\n前半部分字符串为：", qian)
		hou := as[end:]
		//print("\n;后半部分字符串：", hou)
		qian = qian + "(" + hou + ";"

		rsult.WriteString(qian)
		//print("\n结果:", qian)
		//print("----------------")
		//print("\n")
		//ce := as[start+1 : end ]
		//print(ce)

	}

	return rsult.String()

}
