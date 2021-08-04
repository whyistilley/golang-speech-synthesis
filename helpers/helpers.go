package helpers

import "log"

func Log(err error) {
	if err != nil {
		log.Println(err)
	}
}
