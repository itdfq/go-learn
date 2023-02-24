package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	a := "gppher"
	b := "hello world"
	// Compare 函数，用于比较两个字符串的大小，如果两个字符串相等，返回为 0。如果 a 小于 b ，返回 -1 ，反之返回 1 。不推荐使用这个函数，直接使用 == != > < >= <= 等一系列运算符更加直观。
	fmt.Println(strings.Compare(a, b))            //-1
	fmt.Println(strings.ContainsAny("team", "i")) //false
	// 并列必须用&连接
	fmt.Println(strings.ContainsAny("failure", "u & i")) //true
	//空格是或
	fmt.Println(strings.ContainsAny("in failure", "s g")) //true

	fmt.Println(strings.ContainsAny("foo", "")) //false
	fmt.Println(strings.ContainsAny("", ""))    //false

	//子字符串出现次数
	fmt.Println(strings.Count("cheese", "e"))    //3
	fmt.Println(len("谷歌中国"))                     //12
	fmt.Println(strings.Count("谷歌中国", ""))       //5
	fmt.Println(strings.Count("fivevev", "vev")) //1

	//字符串分隔
	fmt.Printf("Fields are：%q\n", strings.Fields(" foo bar baz"))
	//strings.FieldsFunc("Fields,are,qwe", IsStr)
	s := strings.Split("Fields,are,qwe", ",")
	for _, v := range s {
		fmt.Println("相应的结果", v)
	}
	//判断某个字符串是否有前缀或后缀
	fmt.Println(strings.HasPrefix("Gopher", "Go")) //true
	fmt.Println(strings.HasPrefix("Gopher", "C"))  //false
	fmt.Println(strings.HasPrefix("Gopher", ""))   //true
	fmt.Println(strings.HasSuffix("Amigo", "go"))  //true
	fmt.Println(strings.HasSuffix("Amigo", "Ami")) //false
	fmt.Println(strings.HasSuffix("Amigo", ""))    //true
	//统计字符串出现的位置
	fmt.Println("统计字符串出现的位置")
	//第一次出现的位置
	fmt.Println(strings.Index("abcde", "c")) //2
	//最后一次出现的位置
	fmt.Println(strings.LastIndex("abcdec", "c")) //5
	//字符串数组连接起来
	sa := []string{"Fields", "are", "qwe"}
	fmt.Println(strings.Join(sa, ",")) //Fields,are,qwe
	//将字符串重复多少次
	fmt.Println("ba:" + strings.Repeat("na", 2)) //nana
	//字符串替换
	fmt.Println(strings.Map(mapping, "Hello你#￥%……\\n（'World\\n,好Hello^(&(*界gopher...")) //结果：
	//hello
	//nworldn
	//hello
	//gopher

	//子字符串替换 n 换前n个
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))      //oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1)) //moo moo moo
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))  //moo moo moo
	//大小写转换
	fmt.Println(strings.ToLower("HELLO WORLD")) //hello world
	fmt.Println(strings.ToLower("Ā Á Ǎ À"))     //ā á ǎ à
	fmt.Println(strings.ToLower("Önnek İş"))    //önnek iş
	fmt.Println(strings.ToUpper("hello world")) //HELLO WORLD
	fmt.Println(strings.ToUpper("ā á ǎ à"))     //Ā Á Ǎ À
	fmt.Println(strings.ToUpper("örnek iş"))    //ÖRNEK IŞ
	//标题处理
	fmt.Println(strings.Title("hElLo wOrLd"))                                      // 首字母大写  HElLo WOrLd
	fmt.Println(strings.ToTitle("hElLo wOrLd"))                                    // 全部大写 HELLO WORLD
	fmt.Println(strings.Title("dünyanın ilk borsa yapısı Aizonai kabul edilir"))   //Dünyanın Ilk Borsa Yapısı Aizonai Kabul Edilir
	fmt.Println(strings.ToTitle("dünyanın ilk borsa yapısı Aizonai kabul edilir")) //DÜNYANIN ILK BORSA YAPISI AIZONAI KABUL EDILIR
	//修剪，去除
	x := "!!!@@@你好,!@#$ Gophers###$$$"
	fmt.Println(strings.Trim(x, "@#$!%^&*()_+=-"))                  // 去除左边和右边符合条件的(中间不动)        你好,!@#$ Gophers
	fmt.Println(strings.TrimLeft(x, "@#$!%^&*()_+=-"))              //去除左边   你好,!@#$ Gophers###$$$
	fmt.Println(strings.TrimRight(x, "@#$!%^&*()_+=-"))             //去除右边   !!!@@@你好,!@#$ Gophers
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n")) //去除间隔符 Hello, Gophers
	fmt.Println(strings.TrimPrefix(x, "!"))                         //去除前缀一次  !!@@@你好,!@#$ Gophers###$$$
	fmt.Println(strings.TrimSuffix(x, "$"))                         //去除后缀一次    !!!@@@你好,!@#$ Gophers###$$
}

func mapping(r rune) rune {
	switch {
	case r >= 'A' && r <= 'Z': // 大写字母转小写
		return r + 32
	case r >= 'a' && r <= 'z': // 小写字母不处理
		return r
	case unicode.Is(unicode.Han, r): // 汉字换行
		return '\n'
	}
	return -1 // 过滤所有非字母、汉字的字符

}
