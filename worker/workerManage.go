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
	if queNum >= 5000 {
		num = 2000
		queNum = 5000
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
		//var i int = 0
		for {
			select {
			case task := <-p.taskPool:
				if task.startBool == true && task.Run != nil {
					//fmt.Print(idNum, "---")
					task.Run(task.params)
				}
				//防止从channal 中读取数据超时
			case <-time.After(time.Millisecond * 1000):
				fmt.Println("time out init")
				if p.stopTopic == true && len(p.taskPool) == 0 {
					fmt.Println("topic=", p.stopTopic)
					//work数递减
					p.workNum--
					return
				}

				/*default:
				i++
				//fmt.Println("default init",i);
				//不添加 会导致 cpu 急剧上升
				fmt.Println(p.taskPool, 111111111111111)
				if len(p.taskPool) > 1000 {
					time.Sleep(time.Millisecond * 100)
				} else {

					time.Sleep(time.Millisecond * 1000)
				}*/

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

//用select 去做
func (p *WorkPool) Run2(funcJob Job, params ...ParamType) {
	task := taskWork{funcJob, true, params}
	select {
	//正常写入
	case p.taskPool <- task:
		//写入超时 说明队列满了 写入备用队列
	case <-time.After(time.Millisecond * 1000):
		if p.taskQue == nil {
			p.taskQue = make(chan taskWork, p.workNum*2)
			go p.queToPool()
		}
		//说明需要扩充进程
		if p.workNum < 1000 {
			p.workNum++
			p.workInit(p.workNum)
		}
		p.taskQue <- task

	}
}

//
func (p *WorkPool) queToPool() {
	for {
		task := <-p.taskQue
		select {
		case p.taskPool <- task:

		case <-time.After(time.Millisecond * 1000):
			//说明 1s l 还写入不进去 就要抛弃任务了 否则就危险了
		}
	}

}
