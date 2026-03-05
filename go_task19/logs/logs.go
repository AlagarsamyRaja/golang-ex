package logs

import (
	"log"
	"os"
)

func Logs(info string) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	//defer file.Close()

	log.SetPrefix("Information:")

	log.SetFlags(log.LstdFlags | log.Llongfile)

	log.Println(info)


}