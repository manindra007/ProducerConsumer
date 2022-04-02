package writer

import (
	msg "ProducerConsumer/queue"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func Writer(wg *sync.WaitGroup,pr,prod int) {
	defer wg.Done()
	for i := 0; i < prod;{
		if err:=msg.Push(msg.Product{ Message : "msg- "+strconv.Itoa(i)});err!=nil{
			fmt.Println("error occured")
		}else {
			i++
		}
		time.Sleep(time.Duration(pr)*time.Second)
	}
}
