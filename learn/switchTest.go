package main

import "fmt"

func main() {
	//定义局部变量
	var grade = "B"
	var marks = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	switch {
	case grade == "A":
		println("优秀")
	case grade == "B":
		println("良好")
	case grade == "C":
		println("一般")
	default:
		println("差")
	}
	fmt.Printf("你的等级是 %s\n\n", grade)

	println("----------------------------------")

	//使用fallthrough 会强制执行后面的case语句
	/*
		switch {
		case false:
			fmt.Println("1、case 条件语句为 false")
			fallthrough
		case true:
			fmt.Println("2、case 条件语句为 true")
			fallthrough
		case false:
			fmt.Println("3、case 条件语句为 false")
			fallthrough
		case true:
			fmt.Println("4、case 条件语句为 true")
		case false:
			fmt.Println("5、case 条件语句为 false")
			fallthrough
		default:
			fmt.Println("6、默认 case")
		}

	*/
}
