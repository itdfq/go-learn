package reptile_common

import "log"

func HandleError(err error) {
	if err == nil {
		log.Println("handle error, msg:", err)
		return
	}
}
