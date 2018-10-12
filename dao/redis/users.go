package dao

//
//import (
//	"context"
//	"github.com/lisuiheng/fsm"
//	"github.com/lisuiheng/fsm/dao"
//	"github.com/lisuiheng/fsm/dao/internal"
//	"github.com/rs/xid"
//)
//
//// UsersBucket is the bolt bucket to store lists of users
//var UsersBucket = "Users"
//
//// UsersStore is the bolt implementation to store users in a store.
//// Used store users that are associated in some way with a source
//type UsersStore struct {
//	Client *dao.Client
//}
//
//// All returns all known sources
//func (s *UsersStore) All(context.Context) ([]fsm.User, error) {
//	//data, err := s.Client.Get(string(UsersBucket)).Result()
//	//if err == redis.Nil {
//	//	return nil, filer2.ErrNotFound
//	//}
//	//
//	//if err != nil {
//	//	return nil, fmt.Errorf("get %s : %v", entry.FullPath, err)
//	//}
//
//	return []fsm.User{
//		{
//			Name: "hell",
//		},
//	}, nil
//}
//
//func (s *UsersStore) Add(ctx context.Context, src *fsm.User) (err error) {
//	src.ID = xid.New().String()
//	if v, err := internal.MarshalUser(src); err == nil {
//		s.Client.RedisDB.RPush(UsersBucket, v)
//	}
//	return err
//}
