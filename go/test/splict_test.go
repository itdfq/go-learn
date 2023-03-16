package test

import (
	"reflect"
	"testing"
)

//注意：文件的命名：必须以_test.go 结尾
func TestSplit(t *testing.T) { //测试函数必须以 Test[a-Z]开头，必须接受一个*testing.T类型参数
	got := Splict("a:b:c", ":")        // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

func TestMoreSplit(t *testing.T) {
	got := Splict("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}
