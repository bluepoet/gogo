package main

import (
	"github.com/bluepoet/gogo/mail/sendmail"
	"time"
	"fmt"
)

func main()  {
	startTime := time.Now()

	sendmail.ParallelSendmail("from@daum.net","to@daum.net","subject", "go test body!!", 10)

	endTime := time.Now()
	fmt.Printf("diff time : %v", endTime.Sub(startTime))
}
