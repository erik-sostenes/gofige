package repository

import (
	"os"
	"strconv"
	"time"
)

// MongoDB contains the environment variables to configure the mongodb connection
type MongoDB struct {
	Uri,
	DatabaseName string
	ConnectTimeout time.Duration
}

var (
	timeout, _ = strconv.Atoi(os.Getenv("NoSQL_TIMEOUT"))
	// Config settings to a mongodb connection
	Config = MongoDB {
		Uri:            "mongodb://127.0.0.1:27017",
		DatabaseName:   "students",
		ConnectTimeout: time.Duration(timeout),
	}
)
