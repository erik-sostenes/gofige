package repository

import (
	"errors"
	"os"
	"strconv"
	"testing"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
)

const (
	remoteUri = "mongodb+srv://%s:%s@cluster0.o3tc8ee.mongodb.net/?retryWrites=true&w=majority"
	localUri = "mongodb://127.0.0.1:27017"  
)

func TestNewMongoClient(t *testing.T) {
	timeout, _ := strconv.Atoi(os.Getenv("NoSQL_TIMEOUT"))

	tsc := map[string]struct {
		configuration MongoDB
		// expected is the type of error to expect
		expected topology.ConnectionError
	}{
		"Given a correct authentication configuration, a new mongo remote client will be created": {
			configuration: MongoDB{
				Uri:            fmt.Sprintf(remoteUri, os.Getenv("NoSQL_USER"), os.Getenv("NoSQL_PASSWORD")),
				DatabaseName:   os.Getenv("NoSQL_DATABASE"),
				ConnectTimeout: time.Duration(timeout),
			},
		},
		"Given incorrect authentication configuation, a new mongo remote client will not be created": {
			configuration: MongoDB{
				Uri:            fmt.Sprintf(remoteUri, os.Getenv("NoSQL_USER"), os.Getenv("NoSQL_PASSWORD")),
				User:           "some_user",
				Password:       "some_user",
				DatabaseName:   os.Getenv("NoSQL_DATABASE"),
				ConnectTimeout: time.Duration(timeout),
			},
		},
	"Given a correct authentication configuation, a new mongo local client will be created": {
			configuration: MongoDB{
				Uri:            localUri,
				DatabaseName:   "students",
				ConnectTimeout: time.Duration(timeout),
			},
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			_, err := NewMongoClient(ts.configuration)
			if err != nil {
				if !errors.As(err, &ts.expected) {
					t.Fatalf("expected error %T, got %T error", ts.expected, err)
				}
			}
		})
	}
}
