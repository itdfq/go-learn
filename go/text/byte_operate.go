package main

import (
	"bytes"
	"fmt"
)

func main() {
	//是否存在于某个字切片
	b := []byte{0, 1, 0, 1, 0, 1}
	a := []byte{0, 1}
	fmt.Println(bytes.Compare(b, a)) //1
	//出现次数
	fmt.Println(bytes.Count(b, a)) //3

	//转换为rune[] 可以适用于汉子和多种字节字符
	x := []byte("你好世界")
	for k, v := range x {
		fmt.Printf("%d:%s->%b |", k, string(v), v)
	}
	r := bytes.Runes(x)
	fmt.Println()
	for k, v := range r {
		fmt.Printf("%d:%s|", k, string(v))
	}
	//还是先了 Reader 类型
	r1 := bytes.NewReader(x)
	d1 := make([]byte, len(x))
	n, _ := r1.Read(d1)
	fmt.Println(n, string(d1))

	r2 := bytes.Reader{}
	r2.Reset(x)
	d2 := make([]byte, len(x))
	n, _ = r2.Read(d2)
	fmt.Println(n, string(d2))
	fmt.Println("-------------------------")
	//读取一个字符
	ch, size, _ := r1.ReadRune()
	fmt.Println(size, string(ch))
	//进度下标指向前一个字符，如果 r.i <= 0 返回错误，且只能在每次 ReadRune 方法后使用一次，否则返回错误。
	_ = r1.UnreadRune()
	ch, size, _ = r1.ReadRune()
	fmt.Println(size, string(ch))
	_ = r1.UnreadRune()

	by, _ := r1.ReadByte()
	fmt.Println(by)
	_ = r1.UnreadByte()
	by, _ = r1.ReadByte()
	fmt.Println(by)
	_ = r1.UnreadByte()
	//读取数据至 d13
	d13 := make([]byte, 6)
	n1, _ := r1.Read(d13)
	fmt.Println(n1, string(d13))

	//读取 r.s[off:] 的数据至b，该方法忽略进度下标 i，不使用也不修改。
	d23 := make([]byte, 6)
	n3, _ := r1.ReadAt(d23, 0)
	fmt.Println(n3, string(d2))

	w1 := &bytes.Buffer{}
	// 根据 whence 的值，修改并返回进度下标 i ，当 whence == 0 ，进度下标修改为 off，当 whence == 1 ，进度下标修改为 i+off，当 whence == 2 ，进度下标修改为 len[s]+off.
	// off 可以为负数，whence 的只能为 0，1，2，当 whence 为其他值或计算后的进度下标越界，则返回错误。
	_, _ = r1.Seek(0, 0)
	//写入到
	_, _ = r1.WriteTo(w1)
	fmt.Println(w1.String())
}
