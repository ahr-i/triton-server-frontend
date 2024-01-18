package errController

import (
	"log"
	"os"
)

func ErrorCheck(err error, msg string, filePath string) {
	if err != nil {
		log.Println("===== ERROR CATCH =====")
		log.Println("Error:", err)
		log.Println("Msg:", msg)
		log.Println("In", filePath)
		log.Println("=======================")

		os.Exit(1)
	}
}
