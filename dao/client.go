package dao

import (
	"context"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/go-redis/redis"
	"github.com/lisuiheng/fsm"
	boltStore "github.com/lisuiheng/fsm/dao/bolt"
	"os"
	"time"
)

const (
	// ErrUnableToOpen means we had an issue establishing a connection (or creating the database)
	ErrUnableToOpen = "Unable to open boltdb; is there a chronograf already running?  %v"
	// ErrUnableToBackup means we couldn't copy the db file into ./backup
	ErrUnableToBackup = "Unable to backup your database prior to migrations:  %v"
	// ErrUnableToInitialize means we couldn't create missing Buckets (maybe a timeout)
	ErrUnableToInitialize = "Unable to boot boltdb:  %v"
	// ErrUnableToMigrate means we had an issue changing the db schema
	ErrUnableToMigrate = "Unable to migrate boltdb:  %v"
)

// Client is a client for the boltDB data store.
type Client struct {
	Now func() time.Time

	BoltDB     *bolt.DB
	RedisDB    *redis.Client
	UsersStore UsersStore
	Logger     fsm.Logger
}

// NewClient initializes all stores
//func NewRedisClient(hostPort string, password string, database int) *Client {
//	c := &Client{Now: time.Now}
//	rClient := redis.NewClient(&redis.Options{
//		Addr:     hostPort,
//		Password: password,
//		DB:       database,
//	})
//	c.RedisDB = rClient
//	c.UsersStore = &redisStore.UsersStore{Client: c}
//	return c
//}

func NewboltClient(ctx context.Context, path string) *Client {
	c := &Client{Now: time.Now}

	if err := c.OpenBolt(ctx, path); err != nil {
		c.Logger.Error(err)
		os.Exit(1)
	}
	c.UsersStore = &boltStore.UsersStore{DB: c.BoltDB}
	return c
}

// Open / create boltDB file.
func (c *Client) OpenBolt(ctx context.Context, path string) (err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		//c.isNew = true
	} else if err != nil {
		return err
	}

	// Open database file.
	c.BoltDB, err = bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return fmt.Errorf(ErrUnableToOpen, err)
	}
	if err = c.initializeBolt(ctx); err != nil {
		return fmt.Errorf(ErrUnableToInitialize, err)
	}
	return nil
}

// initialize creates Buckets that are missing
func (c *Client) initializeBolt(ctx context.Context) error {
	if err := c.BoltDB.Update(func(tx *bolt.Tx) error {
		// Always create Servers bucket.
		// Always create Users bucket.
		if _, err := tx.CreateBucketIfNotExists(boltStore.UsersBucket); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

// Close the connection to the bolt database
func (c *Client) Close() error {
	if c.BoltDB != nil {
		return c.BoltDB.Close()
	}
	return nil
}
