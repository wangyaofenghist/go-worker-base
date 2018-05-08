# go-worker-base  #
##实现的通用协程池
job 目录是一个 要执行业务的demo

worker 目录是 业务协程池的封装

workerManage.go 为通用协程池实现

worker.go       为基础协程池数据结构

###实例1：

```

poolOne = worker.GetPool("one")					//拿到一个协程池
poolOne.Start(50)								//定义协程数量
poolOne.Run(runFunc.Run, "test4", " aa ", " BB")//运行一个函数 runFunc.Run 为需要执行函数
```

### 实例2：

```
//初始化协程池
worker.StartPool(6)
//调用协程池进行处理
worker.Dispatch(job.Run, 1, 2)
worker.Dispatch(job.RunA, 3, 4)
//获取协程池结果
fmt.Println(<-worker.WorkTaskReturn)
worker.StopPool()
```

### 配合go-Call实现的简单demo

组合go-Call 通用回调方法，快速简单的实现业务并发，协程可控

go-Call 地址：https://github.com/wangyaofenghist/go-worker-base

```
package main

import (
   "fmt"
   "github.com/wangyaofenghist/go-Call/call"
   "github.com/wangyaofenghist/go-Call/test"
   "github.com/wangyaofenghist/go-worker-base/worker"
   "time"
)

//声明一号池子
var poolOne worker.WorkPool

//声明回调变量
var funcs call.CallMap

//以结构体方式调用
type runWorker struct{}

//初始化协程池 和回调参数
func init() {
   poolOne = worker.GetPool("one")
   poolOne.Start(50)
   funcs = call.CreateCall()

}

//通用回调
func (f *runWorker) Run(param []interface{}) {
   name := param[0].(string)
   //调用回调并拿回结果
   funcs.Call(name, param[1:]...)
}

//主函数
func main() {
   var resultChan = make(chan interface{})
   var runFunc runWorker = runWorker{}
   funcs.AddCall("test4", test.Test4)
   for i := 0; i < 10000; i++ {
      poolOne.Run(runFunc.Run, "test4", " aa ", " BB")
      poolOne.Run(runFunc.Run, "test4", " cc ", " dd")
      poolOne.Run(runFunc.Run, "test4", " ee ", " ff")
   }
   poolOne.Stop()
   <-resultChan
}
```