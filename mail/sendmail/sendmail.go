package sendmail

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
)

func Sendmail(from string, to string, subject string, body string)  {
	message := ""
	header := "subject: "+subject+"\n"
	header += "from: "+ from +"\n"
	message += header + "\n"+body

	servername := "mail server domain"

	for i := 0; i < 100; i++ {
		err := smtp.SendMail(servername + ":25", nil, from, []string{to}, []byte(message))
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Mail Send Success!!")
		}
	}
}

func ParallelSendmail(from string, to string, subject string, body string, n int) {
	message := ""
	header := "subject: "+subject+"\n"
	header += "from: "+ from +"\n"
	message += header + "\n"+body

	servername := "mail server domain"

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				err := smtp.SendMail(servername+":25", nil, from, []string{to}, []byte(message))
				if err != nil {
					log.Fatal(err)
				}else{
					fmt.Println("Mail Send Success!!")
				}
			}
		}()
	}
	wg.Wait()
}
