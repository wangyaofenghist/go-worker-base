package main

import (
	"localhostTest/go-worker-base/worker"
	"localhostTest/go-worker-base/job"
	"fmt"
)
var poolOne worker.WorkPool
func init(){
	worker.StartPool(6)
	poolOne = worker.GetPool("one")
	poolOne.Start(500);
}
var JobCReturn chan int;
func main(){
	//调用协程池进行处理
	worker.Dispatch(job.Run,1,2)
	worker.Dispatch(job.RunA,3,4)
	//获取协程池结果
	fmt.Println(<-worker.WorkTaskReturn)
	worker.StopPool();

	JobCReturn = make(chan int,3)
	poolOne.Run(job.Run,5,6);
	//var runcReturn worker.ReturnType
	//利用map 传递地址的特性 来拿回结果

	var resultChan = make(chan worker.ReturnType,200)
	fmt.Println(resultChan)
	for i:=0;i<2000;i++{
		var paramMap = make(map[string] worker.ParamType)
		paramMap["a"] = 7+i;
		paramMap["b"] = 8;
		poolOne.Run(job.RunC,paramMap);
		//runcReturn =<-resultChan
		//fmt.Println(runcReturn.(int))
	}
	//<-resultChan
	poolOne.Stop();

}
