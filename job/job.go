package job

import (
	"fmt"
	"localhostTest/go-worker-base/worker"
)

func jobTest( param[] int){
	fmt.Println("this is job1 test!");
	fmt.Println("this is test param[0] = ",param[0]," param[1] = " ,param[1])
	var returnParam []worker.ReturnType
	returnParam = append(returnParam, param[0]+param[1])
	returnParam = append(returnParam,param[0]*param[1])
	worker.WorkTaskReturn <- returnParam
}
type ParamType interface {}
func Run( param[] worker.ParamType){
	var paramJob[] int
	for _,p := range param{
		switch v:=p.(type){
		case int:
			var s int
			s = v
			paramJob  = append(paramJob,s);
		}

	}
	jobTest(paramJob);

}