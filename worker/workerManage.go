package worker

import (
	"fmt"
	"time"
)

type WorkPool struct {
	taskPool  chan taskWork
	workNum   int
	stopTopic bool
	key       string
	//暂时没有用，考虑后期 作为冗余队列使用
	taskQue chan taskWork
}

var WorkMap = make(map[string]WorkPool)

//得到一个线程池并返回 句柄
func GetPool(workKey string) WorkPool {
	if _, ok := WorkMap[workKey]; !ok {
		WorkMap[workKey] = WorkPool{key: workKey}
	}

	return WorkMap[workKey]

}

//开始work
func (p *WorkPool) Start(num int) {
	//说明已经存在
	if p.workNum != 0 {
		return
	}
	//taskPool 任务池 协程个数的两倍
	var queNum = num * 2
	if queNum >= 500 {
		num = 200
		queNum = num * 2
	}
	p.taskPool = make(chan taskWork, queNum)
	//p.taskQue = make(chan taskWork,queLen)
	//记录协程个数
	p.workNum = num
	//设置标志位
	p.stopTopic = false
	for i := 0; i < num; i++ {
		p.workInit(i)
		fmt.Println("start pool task:", i)
	}
}

//初始化 work池
func (p *WorkPool) workInit(id int) {
	go func(idNum int) {
		var i int = 0
		for {
			select {
			case task := <-p.taskPool:
				if task.startBool == true && task.Run != nil {
					//fmt.Print(idNum, "---")
					task.Run(task.params)
				}
			case <-time.After(time.Millisecond * 1000):
				fmt.Println("time out init")
			default:
				if p.stopTopic == true && len(p.taskPool) == 0 {
					fmt.Println("topic=", p.stopTopic)
					//work数递减
					p.workNum--
					return
				}
				i++
				//fmt.Println("default init",i);
				//time.Sleep(time.Millisecond*1000);
			}

		}
	}(id)

}

//停止一个workPool
func (p *WorkPool) Stop() {
	p.stopTopic = true
}
func (p *WorkPool) Run(funcJob Job, params ...ParamType) {
	p.taskPool <- taskWork{funcJob, true, params}
}
func (p *WorkPool) Run2(funcJob Job, params ...ParamType) {
	p.taskPool <- taskWork{funcJob, true, params}
}
