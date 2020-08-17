package core

// Handle is a generic error handler
func Handle(err error) {
	if err != nil {
		Error.Panic(err)
	}
}
