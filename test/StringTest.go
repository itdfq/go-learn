package main

import (
	"fmt"
	"strings"
)

func main() {
	//s := "INSERT INTO `ep_trade_test`.`prod_listing_aliexpress_full_manage` (`id`, `store_acct_id`, `prod_p_id`, `aliexpress_template_id`, `store_p_sku`, `prod_p_sku`, `item_id`, `smt_category_id`, `title`, `main_images`, `market1_images`, `market2_images`, `product_type`, `desc_images`, `image_attr_id`, `sku_attr_ids`, `mobile_desc`, `pc_desc`, `adjust_price_status`, `adjust_price_type`, `adjust_price_country`, `configuration_data`, `video_file_url`, `poster_url`, `media_id`, `video_fail_reason`, `listing_time`, `list_timing`, `listing_status`, `listing_resp_code`, `listing_resp_msg`, `remark`, `status`, `create_time`, `creator_id`, `creator`, `modify_time`, `modifier_id`, `modifier`) VALUES (49, 35904, 74284, 46695, 'CETS1S41', 'CETS1S41', 1005005885057604, 200004218, 'White Memory Foam Replacement Ear pads sleeve Ear Tips Earbuds Cover Earplugs Cap For Apple Airpods Pro', 'https://ae01.alicdn.com/kf/Sf67ae2ce4e74479385195b05ef7e2d2eC.jpg,https://ae01.alicdn.com/kf/S8f2a62ad8a764a7fa7f0ae900dff67aby.jpg,https://ae01.alicdn.com/kf/S201254522ff4489e968a5b51b7fe189fb.jpg,https://ae01.alicdn.com/kf/S4e7c4a9efa35482193fd56c3bdbef536D.jpg,https://ae01.alicdn.com/kf/Sd9f829d009b245a7aaf64b869a4c0ac0E.jpg,https://ae01.alicdn.com/kf/Sf3d570ec0b0142de8797f5da02dfe84b2.jpg', 'https://ae01.alicdn.com/kf/S15adc0e69688465aacd3abc1fdb7e0fef.jpg', 'https://ae01.alicdn.com/kf/S032e046c31e74e7eba7a7503c2d5fe70g.jpg', 1, 'https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9638ff33402880c00bcd.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a4cf402880c00bd1.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a67c402880c00bdd.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a517402880c00bd4.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a61d402880c00bd9.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a52f402880c00bd6.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a63f402880c00bdb.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a536402880c00bd7.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a649402880c00bdc.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a4fb402880c00bd3.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a636402880c00bda.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a4bf402880c00bd0.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a4e0402880c00bd2.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a562402880c00bd8.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a4a0402880c00bcf.jpg,https://img.test.epean.cn/sub/ep/img/prodtpl/1/230727/9639a51c402880c00bd5.jpg', 14, '0,14,0', '{\\\"moduleList\\\":[{\\\"texts\\\":[{\\\"class\\\":\\\"title\\\",\\\"content\\\":\\\"\\\"},{\\\"class\\\":\\\"body\\\",\\\"content\\\":\\\"Material:Quality soft memory foam.\\\\r\\\\nColor:White\\\\r\\\\nSize: S / M / L\\\\r\\\\nCompatible with AirPods Pro\\\\r\\\\n1:1 designed for airpods pro\\\\r\\\\nTip: there is a mesh inside of earbuds\\\\r\\\\nNote: large size ear tips must be compressed before inserting into case.\\\\nFeature:\\\\r\\\\nBlocks out more sound improving noise cancellation.\\\\r\\\\nStronger hold of headphones in ears.\\\\r\\\\nMore comfortable fit.\\\\r\\\\nOur memory foam tips and all airpods pro memory foam ear tips currently available use stem adapters.\\\\r\\\\nOur stem adapters are very high quality!\\\\r\\\\nFits and locks precisely to the headphones and locking mechanism. \\\\r\\\\nEar tips have very firm hold to the stem adapters, will not easily slide off or stay in ear when removed. \\\\r\\\\nTips fit in charge case while on the headphones however the ends must be tucked into the bottom of the case before closing lid.\\\\r\\\\n\\\\r\\\\nPackage:\\\\r\\\\n1 pair memory foam Ear Tips or 3 pairs memory foam Ear Tips.\\\"}],\\\"type\\\":\\\"text\\\"},{\\\"images\\\":[{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/Sf67ae2ce4e74479385195b05ef7e2d2eC.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S15adc0e69688465aacd3abc1fdb7e0fef.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S201254522ff4489e968a5b51b7fe189fb.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S65a8aa0a25d04d1887b01744b25553dfp.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S4494e196aa2f427481ca6b4cee32b58c8.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/Sf3d570ec0b0142de8797f5da02dfe84b2.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S8f2a62ad8a764a7fa7f0ae900dff67aby.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/Sd9f829d009b245a7aaf64b869a4c0ac0E.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S6811134f0ea349c0896c2744f12cfacbx.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S84179875f78640bead17cc4d941ca7403.jpg\\\"}],\\\"type\\\":\\\"image\\\"},{\\\"images\\\":[{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S94a625c294e84def98c8d3681e77bb2aK.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/Sfe483df713c7404cb6610f0923a6f78cB.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S7ee64f412e58484f89c4f8365a2a0136b.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S4e7c4a9efa35482193fd56c3bdbef536D.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S9fbcbf84561e4312b923911c5bc30b27w.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S0af2422ddbc64bef8581713df316a712f.jpg\\\"}],\\\"type\\\":\\\"image\\\"}],\\\"version\\\":\\\"2.0.0\\\"}', '{\\\"moduleList\\\":[{\\\"texts\\\":[{\\\"class\\\":\\\"title\\\",\\\"content\\\":\\\"\\\"},{\\\"class\\\":\\\"body\\\",\\\"content\\\":\\\"Material:Quality soft memory foam.\\\\r\\\\nColor:White\\\\r\\\\nSize: S / M / L\\\\r\\\\nCompatible with AirPods Pro\\\\r\\\\n1:1 designed for airpods pro\\\\r\\\\nTip: there is a mesh inside of earbuds\\\\r\\\\nNote: large size ear tips must be compressed before inserting into case.\\\\nFeature:\\\\r\\\\nBlocks out more sound improving noise cancellation.\\\\r\\\\nStronger hold of headphones in ears.\\\\r\\\\nMore comfortable fit.\\\\r\\\\nOur memory foam tips and all airpods pro memory foam ear tips currently available use stem adapters.\\\\r\\\\nOur stem adapters are very high quality!\\\\r\\\\nFits and locks precisely to the headphones and locking mechanism. \\\\r\\\\nEar tips have very firm hold to the stem adapters, will not easily slide off or stay in ear when removed. \\\\r\\\\nTips fit in charge case while on the headphones however the ends must be tucked into the bottom of the case before closing lid.\\\\r\\\\n\\\\r\\\\nPackage:\\\\r\\\\n1 pair memory foam Ear Tips or 3 pairs memory foam Ear Tips.\\\"}],\\\"type\\\":\\\"text\\\"},{\\\"images\\\":[{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/Sf67ae2ce4e74479385195b05ef7e2d2eC.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S15adc0e69688465aacd3abc1fdb7e0fef.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S201254522ff4489e968a5b51b7fe189fb.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S65a8aa0a25d04d1887b01744b25553dfp.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S4494e196aa2f427481ca6b4cee32b58c8.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/Sf3d570ec0b0142de8797f5da02dfe84b2.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S8f2a62ad8a764a7fa7f0ae900dff67aby.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/Sd9f829d009b245a7aaf64b869a4c0ac0E.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S6811134f0ea349c0896c2744f12cfacbx.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S84179875f78640bead17cc4d941ca7403.jpg\\\"}],\\\"type\\\":\\\"image\\\"},{\\\"images\\\":[{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S94a625c294e84def98c8d3681e77bb2aK.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/Sfe483df713c7404cb6610f0923a6f78cB.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S7ee64f412e58484f89c4f8365a2a0136b.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S4e7c4a9efa35482193fd56c3bdbef536D.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S9fbcbf84561e4312b923911c5bc30b27w.jpg\\\"},{\\\"style\\\":{},\\\"url\\\":\\\"https://ae01.alicdn.com/kf/S0af2422ddbc64bef8581713df316a712f.jpg\\\"}],\\\"type\\\":\\\"image\\\"}],\\\"version\\\":\\\"2.0.0\\\"}', 0, NULL, NULL, '', NULL, NULL, NULL, NULL, '2023-07-27 17:29:23', NULL, 1, NULL, '刊登成功', '刊登成功', 1, '2023-07-27 16:37:49', 1560, '段方钦', '2023-07-27 17:29:23', -1, 'system');\n"
	s := "INSERT INTO `trade_online`.`shopee_account_auth_info` (`id`, `store_acct_id`, `shop_id`, `merchant_id`, `access_token_shop`, `access_token_merchant`, `merchant_name`, `shop_is_cnsc`, `shop_is_affiliate`, `partner_id`, `secret_key`, `access_token_refresh_time_shop`, `access_token_expire_in_shop`, `refresh_token_shop`, `refresh_token_refresh_time_shop`, `refresh_token_expire_in_shop`, `auth_time_shop`, `auth_expire_time_shop`, `refresh_result_shop`, `access_token_refresh_time_merchant`, `access_token_expire_in_merchant`, `refresh_token_merchant`, `refresh_token_refresh_time_merchant`, `refresh_token_expire_in_merchant`, `auth_time_merchant`, `auth_expire_time_merchant`, `refresh_result_merchant`, `create_time`, `modify_time`, `creator`, `modifier`, `creator_id`, `modifier_id`) VALUES (5410, 35252, 183962177, 1000079, '6e416352666d5a68747370614d524d72', '6f56654d6a5172436b59794c6f747659', '上海亿品贸易有限公司', 1, 1, 2002462, 'e994c49d70c47b6cd07fa6889d06eb2f38fb4222d0d45de46f8ed76b7ad74339', '2023-08-29 10:27:34', '2023-08-29 14:27:34', '4d766c4262766b485048684973506647', '2023-08-29 10:27:34', '2023-09-28 10:27:34', '2023-08-18 08:55:59', '2024-08-18 08:55:59', 'ok', '2023-08-18 08:55:59', '2023-08-29 14:06:07', '686e5273476c526c4a4d534145515257', '2023-08-29 10:06:07', '2023-09-28 10:06:07', '2023-08-18 08:55:59', '2024-08-18 08:55:59', 'ok', '2022-03-01 11:35:43', '2022-03-01 11:35:43', 'system', 'system', -1, -1);\n"
	str := replace(s)
	str2 := replace2(s)
	fmt.Printf("替换后的结果为：%v\n", str)
	fmt.Printf("替换后的结果2：%v\n", str2)
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

func replace2(str string) string {
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, "`trade_online`.", "")
	str = strings.ReplaceAll(str, "`id`,", "")
	arrayStr := strings.Split(str, ";")
	var result []string
	for _, as := range arrayStr {
		if len(as) < 5 {
			continue
		}
		lastkuohao := strings.LastIndex(as, "(")
		newStr := as[lastkuohao:]
		end := strings.Index(newStr, ",") + lastkuohao + 2

		qian := as[:lastkuohao]
		hou := as[end:]
		qian = qian + "(" + hou + ";"

		result = append(result, qian)
	}

	return strings.Join(result, "")
}
