package main

import (
	"errors"
	"time"

	"github.com/go-redsync/redsync/v4"
)

func lockTrxRedisOnly(key string, expiration time.Duration) error {
	val, err := incr(ctx, key, expiration)
	if err != nil {
		return err
	} else if val != 1 {
		return errors.New("There are other processes in progress, please try again later")
	}
	return err
}

func lockTrx(key string, expiration time.Duration) (*redsync.Mutex, error) {

	// Create a new mutex
	mutex := rs.NewMutex(key, redsync.WithExpiry(expiration))
	err := mutex.Lock()
	if err != nil {
		return nil, err
	}
	return mutex, nil
}

func unlockTrx(mutex redsync.Mutex) error {
	_, err := mutex.Unlock()
	return err
}
