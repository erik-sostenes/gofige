package repository

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
)

const (
	remoteUri = "mongodb+srv://%s:%s@cluster0.o3tc8ee.mongodb.net/?retryWrites=true&w=majority"
	localUri  = "mongodb://127.0.0.1:27017"
)

func TestNewClientMongo_CorrectConnection(t *testing.T) {
	timeout, _ := strconv.Atoi(os.Getenv("NoSQL_TIMEOUT"))

	tsc := map[string]struct {
		configuration MongoDB
		// err expeted error
		expetedErr error
	}{
		"Given a correct authentication configuration, a new mongo remote client will be created": {
			configuration: MongoDB{
				Uri:            fmt.Sprintf(remoteUri, os.Getenv("NoSQL_USER"), os.Getenv("NoSQL_PASSWORD")),
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
		ts := ts
		t.Run(name, func(t *testing.T) {
			_, err := NewClientMongo(ts.configuration)
			if err != nil {
				t.Fatalf("expected error %T, got %T error", ts.expetedErr, err)
			}
		})
	}
}

func TestNewClientMongo_WrongConnection(t *testing.T) {
	timeout, _ := strconv.Atoi(os.Getenv("NoSQL_TIMEOUT"))

	tsc := map[string]struct {
		configuration MongoDB
		// err expeted error
		expetedErr topology.ConnectionError
	}{
	"Given incorrect authentication configuation, a new mongo remote client will not be created": {
			configuration: MongoDB{
				Uri:            fmt.Sprintf(remoteUri, os.Getenv("NoSQL_USER"), os.Getenv("NoSQL_PASSWORD")),
				DatabaseName:   os.Getenv("NoSQL_DATABASE"),
				ConnectTimeout: time.Duration(timeout),
			},
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			_, err := NewClientMongo(ts.configuration)
			if !errors.As(err, &ts.expetedErr) {
				t.Fatalf("expected error %T, got %T error", ts.expetedErr, err)
			}
		})
	}
}
