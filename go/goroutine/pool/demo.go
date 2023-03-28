package main

import (
	"fmt"
	"math/rand"
)

type Job struct {
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

//计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
//随机生成数字进行计算
func main() {
	//创建管道
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)
	//创建工作池
	createPool(64, jobChan, resultChan)
	//开启打印的协程
	go func(resultChan chan *Result) {
		//遍历循环管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randomnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	for {
		id++
		//生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job

	}

}

func createPool(num int, jobChan chan *Job, resutChan chan *Result) {
	//根据协程个数去执行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			//执行运算
			//遍历管道内所有数据，进行相加
			for job := range jobChan {
				r_num := job.RandNum
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				r := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- r
			}
		}(jobChan, resutChan)
	}
}
