/*此协程池 旨在方便的创建一个协程池供业务调用 各种缓冲应该在业务层处理*/
package worker

import (
	"fmt"
	"time"
)

type ParamType interface {}
type ReturnType interface {}
type Job func([]ParamType)
type taskWork struct {
	Run 		Job
	startBool 	bool
	params 		[]ParamType
}
var WorkMaxTask int
var WorkTaskPool chan taskWork
var WorkTaskReturn chan []ReturnType

//循环启动协程池
func StartPool(maxTask int){
	WorkMaxTask = maxTask
	WorkTaskPool = make(chan taskWork,maxTask)
	WorkTaskReturn =  make(chan []ReturnType,maxTask)
	for i:=0;i<maxTask;i++{
		var t  = taskWork{}
		fmt.Println("start task:",i)
		t.start()
	}
}
//启动任务
func (t *taskWork) start(){
	go func() {
		for {
			select {
			case funcRun:=<-WorkTaskPool:
				if(funcRun.startBool == true ){
					funcRun.Run(funcRun.params)
				}else{
					return
				}
			case <-time.After(time.Millisecond*1000):
				fmt.Print("time out");
			default:
				time.Sleep(10);
			}
		}
	}()
}
//消费任务
func Dispatch(funcJob Job,params ...ParamType){
	var paramSlice []ParamType
	for _,param := range params{
		paramSlice = append(paramSlice,param)
	}
	WorkTaskPool <- taskWork{funcJob,true,paramSlice}
}
//停止协程池
func StopPool(){
	var funcJob Job
	var paramSlice []ParamType
	for i:=0;i<WorkMaxTask;i++{
		WorkTaskPool <- taskWork{funcJob,false,paramSlice}
	}
}