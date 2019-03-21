package comsume

import (
	"fmt"
	"testing"
	"time"
)

func TestMain1(t *testing.T)  {
	t1 := time.Now()
	ch := make(chan int, 30)

	go product(ch)

	go consume(ch, "111")
	go consume(ch, "222")
	go consume(ch, "333")

	t2 := t1.Sub(time.Now())

	fmt.Println(t2)

	time.Sleep(time.Minute)
}
