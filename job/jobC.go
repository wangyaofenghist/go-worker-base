package job

import (
	"fmt"
)

func RunC(param []interface{}) {
	var a int
	var b int
	var paramMap map[string]interface{}
	var resultChan chan interface{}
	for _, val := range param {
		switch v := val.(type) {
		case map[string]interface{}:
			paramMap = v
		case chan interface{}:
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
