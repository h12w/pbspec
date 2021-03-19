package proto2json

import (
	"io/ioutil"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/descriptorpb"
)

type (
	TypeSet struct {
		msgTypes  map[string]*MessageType
		enumTypes map[string]*EnumType
	}
	MessageType struct {
		*descriptorpb.DescriptorProto
		File *descriptorpb.FileDescriptorProto
	}
	EnumType struct {
		*descriptorpb.EnumDescriptorProto
	}
)

func LoadTypeSet(filename string) (*TypeSet, error) {
	var fileSet descriptorpb.FileDescriptorSet
	jsonBuf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := protojson.Unmarshal(jsonBuf, &fileSet); err != nil {
		return nil, err
	}
	msgTypes := make(map[string]*MessageType)
	enumTypes := make(map[string]*EnumType)
	for _, file := range fileSet.File {
		for _, msgDesc := range file.MessageType {
			typeName := file.GetPackage() + "." + msgDesc.GetName()
			msgTypes[typeName] = &MessageType{
				DescriptorProto: msgDesc,
				File:            file,
			}
		}
		for _, enumType := range file.EnumType {
			typeName := file.GetPackage() + "." + enumType.GetName()
			enumTypes[typeName] = &EnumType{
				EnumDescriptorProto: enumType,
			}
		}
	}
	return &TypeSet{
		msgTypes:  msgTypes,
		enumTypes: enumTypes,
	}, nil
}

func (s *TypeSet) GetMsgType(longName string) (*MessageType, bool) {
	longName = strings.TrimPrefix(longName, ".")
	item, ok := s.msgTypes[longName]
	return item, ok
}

func (s *TypeSet) GetEnumType(longName string) (*EnumType, bool) {
	longName = strings.TrimPrefix(longName, ".")
	item, ok := s.enumTypes[longName]
	return item, ok
}

func (t *MessageType) GoPackage() string {
	dirs := strings.Split(t.File.GetOptions().GetGoPackage(), "/")
	return dirs[len(dirs)-1]
}
