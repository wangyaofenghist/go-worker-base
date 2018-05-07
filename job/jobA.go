package job

import (
	"fmt"
	"github.com/wangyaofenghist/go-worker-base/worker"
)

func RunA(param []worker.ParamType) {
	fmt.Println(param)
}
