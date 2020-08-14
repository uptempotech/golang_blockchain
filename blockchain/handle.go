package blockchain

import "log"

// Handle is a generic error handler
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
