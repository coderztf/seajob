package task

var(
	Task chan string
)

func init(){
	//初始化任务队列
	Task = make(chan string,5)
}


