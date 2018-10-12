package internal

import (
	"github.com/golang/protobuf/proto"
	"github.com/lisuiheng/fsm"
)

//go:generate protoc -I . --go_out=. internal.proto

// MarshalUser encodes a source to binary protobuf format.
func MarshalUser(s *fsm.User) ([]byte, error) {
	return proto.Marshal(&User{
		ID:       s.ID,
		Name:     s.Name,
		Username: s.Username,
		Password: string(s.Password),
		Role:     int64(s.Role),
	})
}

// UnmarshalUser decodes a source from binary protobuf data.
func UnmarshalUser(data []byte, s *fsm.User) error {
	var pb User
	if err := proto.Unmarshal(data, &pb); err != nil {
		return err
	}

	s.ID = pb.ID
	s.Name = pb.Name
	s.Username = pb.Username
	s.Password = fsm.Sensitivity(pb.Password)
	s.Role = int(pb.Role)
	return nil
}
