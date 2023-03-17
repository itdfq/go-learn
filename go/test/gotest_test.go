package test

import "testing"

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil { //try a unit test on function
		t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") //记录一些你期望记录的信息
	}
}

func Test_Division_2(t *testing.T) {
	if _, e := Division(6, 0); e == nil { //try a unit test on function
		t.Error("Division did not work as expected.") // 如果不是如预期的那么就报错
	} else {
		t.Log("one test passed.", e) //记录一些你期望记录的信息
	}
}

//压力测试 压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母
//如果要执行压力测试需要带上参数-test.bench，语法:-test.bench="test_name_regex",例如go test -test.bench=".*"表

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		Division(4, 5)
	}
}

//goos: windows
//goarch: amd64
//pkg: go-learn/go/test
//cpu: 11th Gen Intel(R) Core(TM) i5-11400 @ 2.60GHz
//Benchmark_Division
//Benchmark_Division-12           1000000000 (执行测试)              0.2434 ns/op  （平均每次0.2434纳秒）
//PASS

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
