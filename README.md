# go-worker-base  #
##实现的通用协程池
job 目录是一个 要执行业务的demo

worker 目录是 业务协程池的封装

workerManage.go 为通用协程池实现

worker.go       为基础协程池数据结构

###实例：

```

poolOne = worker.GetPool("one")					//拿到一个协程池
poolOne.Start(50)								//定义协程数量
poolOne.Run(runFunc.Run, "test4", " aa ", " BB")//运行一个函数 runFunc.Run 为需要执行函数
```

