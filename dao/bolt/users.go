package dao

import (
	"context"
	"github.com/boltdb/bolt"
	"github.com/lisuiheng/fsm"
	"github.com/lisuiheng/fsm/dao/internal"
	"github.com/rs/xid"
)

// UsersBucket is the bolt bucket to store lists of users
var UsersBucket = []byte("Users")

// UsersStore is the bolt implementation to store users in a store.
// Used store users that are associated in some way with a source
type UsersStore struct {
	DB *bolt.DB
}

// All returns all known sources
func (s *UsersStore) All(ctx context.Context) ([]fsm.User, error) {
	var srcs []fsm.User
	if err := s.DB.View(func(tx *bolt.Tx) error {
		var err error
		srcs, err = s.all(ctx, tx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return srcs, nil

}

// Add creates a new User in the UserStore.
func (s *UsersStore) Add(ctx context.Context, src *fsm.User) (err error) {
	if err = s.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(UsersBucket)

		id := xid.New()
		src.ID = id.String()
		if v, err := internal.MarshalUser(src); err != nil {
			return err
		} else if err := b.Put([]byte(src.ID), v); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// Delete removes the User from the UsersStore
func (s *UsersStore) Delete(ctx context.Context, src *fsm.User) error {
	if err := s.DB.Update(func(tx *bolt.Tx) error {
		if err := tx.Bucket(UsersBucket).Delete([]byte(src.ID)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

// Get returns a User if the id exists.
func (s *UsersStore) Get(ctx context.Context, id string) (src *fsm.User, err error) {
	src = &fsm.User{}
	if err = s.DB.View(func(tx *bolt.Tx) error {
		if v := tx.Bucket(UsersBucket).Get([]byte(id)); v == nil {
			return fsm.ErrUserNotFound
		} else if err = internal.UnmarshalUser(v, src); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return src, nil
}

// Update a User
func (s *UsersStore) Update(ctx context.Context, src *fsm.User) error {
	if err := s.DB.Update(func(tx *bolt.Tx) error {
		// Get an existing user with the same ID.
		b := tx.Bucket(UsersBucket)
		if v := b.Get([]byte(src.ID)); v == nil {
			return fsm.ErrUserNotFound
		}

		if v, err := internal.MarshalUser(src); err != nil {
			return err
		} else if err := b.Put([]byte(src.ID), v); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *UsersStore) all(ctx context.Context, tx *bolt.Tx) ([]fsm.User, error) {
	var srcs []fsm.User
	if err := tx.Bucket(UsersBucket).ForEach(func(k, v []byte) error {
		var src fsm.User
		if err := internal.UnmarshalUser(v, &src); err != nil {
			return err
		}
		srcs = append(srcs, src)
		return nil
	}); err != nil {
		return srcs, err
	}
	return srcs, nil
}
