package main

import (
	"errors"
	"fmt"
	"time"
)

//定义全局错误
var ErrNotFound = errors.New("Employee not found!")

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	employee, err := getInformation(100)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(*employee)
	}

	//也可以使用这种方式进行处理
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("NOT FOUND:%v\n", err)
	} else {
		fmt.Println(*employee)
	}
}
func getInformation(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}
	employee, err := apiCall(id)
	if err != nil {
		return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
	}

	return employee, err
}
func apiCall(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil

}

//加入尝试重试机制
func getInformation1(id int) (*Employee, error) {
	for tries := 0; tries < 3; tries++ {
		employee, err := apiCall(1000)
		if err == nil {
			return employee, nil
		}

		fmt.Println("Server is not responding, retrying ...")
		time.Sleep(time.Second * 2)
	}

	return nil, fmt.Errorf("server has failed to respond to get the employee information")
}
