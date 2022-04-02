package listener

import (
	msg "ProducerConsumer/queue"
	"fmt"
	"sync"
	"time"
)

func Listen(wg *sync.WaitGroup,pr,prod int){
	defer wg.Done()
	for i:=0;i<prod;{
		if res,err := msg.Pop();err!=nil{
			fmt.Println(err.Error())
			time.Sleep(time.Duration(pr+2)*time.Second)
			continue
		}else {
			fmt.Println(res)
			i++
		}
		time.Sleep(time.Duration(pr)*time.Second)
	}
}
