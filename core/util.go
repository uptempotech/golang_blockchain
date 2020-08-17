package core

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dgraph-io/badger"
)

func retry(dir string, originalOpts badger.Options) (*badger.DB, error) {
	lockPath := filepath.Join(dir, "LOCK")
	if err := os.Remove(lockPath); err != nil {
		return nil, fmt.Errorf(`removing "LOCK": %s`, err)
	}
	retryOpts := originalOpts
	retryOpts.Truncate = true
	db, err := badger.Open(retryOpts)
	return db, err
}

func openDB(dir string, opts badger.Options) (*badger.DB, error) {
	var db *badger.DB
	var err error

	if db, err = badger.Open(opts); err != nil {
		if strings.Contains(err.Error(), "LOCK") {
			if db, err := retry(dir, opts); err == nil {
				log.Println("database unlocked, value log truncated")
				return db, nil
			}
			log.Println("could not unlock database:", err)
		}
		return nil, err
	}
	return db, nil
}

func intToBytes(num int) []byte {
	s := strconv.Itoa(num)

	return []byte(s)
}

func bytesToInt(data []byte) int {
	i, err := strconv.Atoi(string(data))
	Handle(err)

	return i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
