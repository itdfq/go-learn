package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Open("info.log")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	//输出目录下的所有文件
	dir := "D:\\go-work\\go-learn"
	listAll(dir, 0)
}

//列出所有的目录文件
func listAll(path string, curHier int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name(), "\\")
			listAll(path+"/"+info.Name(), curHier+1)
		} else {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}
			fmt.Println(info.Name())
		}
	}
}
