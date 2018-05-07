package job

import (
	"fmt"
	"github.com/wangyaofenghist/go-worker-base/worker"
)

func RunC(param []worker.ParamType) {
	var a int
	var b int
	var paramMap map[string]worker.ParamType
	var resultChan chan worker.ReturnType
	for _, val := range param {
		switch v := val.(type) {
		case map[string]worker.ParamType:
			paramMap = v
		case chan worker.ReturnType:
			resultChan = v
		}
	}
	a = paramMap["a"].(int)
	b = paramMap["b"].(int)
	c := cTest(a, b)
	if resultChan != nil {
		resultChan <- c
	} else {
		fmt.Println(c)
	}

	//time.Sleep(time.Millisecond*10);
}
func cTest(a int, b int) int {
	return a + b
}
