/*此协程池 旨在方便的创建一个协程池供业务调用 各种缓冲应该在业务层处理*/
package worker

import (
	"fmt"
	"time"
)

type Job func([]interface{})
type taskWork struct {
	Run       Job
	startBool bool
	params    []interface{}
}

var WorkMaxTask int
var WorkTaskPool chan taskWork
var WorkTaskReturn chan []interface{}

//启动任务
func (t *taskWork) start() {
	go func() {
		for {
			select {
			case funcRun := <-WorkTaskPool:
				if funcRun.startBool == true {
					funcRun.Run(funcRun.params)
				} else {
					fmt.Println("task  stop!")
					return
				}
			case <-time.After(time.Millisecond * 1000):
				fmt.Print("time out")
			}
		}
	}()
}

func (t *taskWork) stop() {
	fmt.Println("t stop ")
	t.startBool = false
}
func createTask() taskWork {
	var funcJob Job
	var paramSlice []interface{}
	return taskWork{funcJob, true, paramSlice}
}

//循环启动协程池
func StartPool(maxTask int) {
	WorkMaxTask = maxTask
	WorkTaskPool = make(chan taskWork, maxTask)
	WorkTaskReturn = make(chan []interface{}, maxTask)

	for i := 0; i < maxTask; i++ {
		var t = createTask()
		fmt.Println("start task:", i)
		t.start()
	}
}

//消费任务
func Dispatch(funcJob Job, params ...interface{}) {
	WorkTaskPool <- taskWork{funcJob, true, params}
}

//停止协程池
func StopPool() {
	var funcJob Job
	var paramSlice []interface{}
	for i := 0; i < WorkMaxTask; i++ {
		WorkTaskPool <- taskWork{funcJob, false, paramSlice}
	}
}
