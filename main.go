package main

import (
	"localhostTest/go-worker-base/worker"
	"localhostTest/go-worker-base/job"
	"fmt"
)
func init(){
	worker.StartPool(6)
}
func main(){
	//调用协程池进行处理
	worker.Dispatch(job.Run,1,2)
	worker.Dispatch(job.RunA,3,4)
	//获取协程池结果
	fmt.Println(<-worker.WorkTaskReturn)
	worker.StopPool();
}
