package pbspec

import (
	"strings"

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
	Field struct {
		*descriptorpb.FieldDescriptorProto
	}
)

func NewTypeSet(fileSet *descriptorpb.FileDescriptorSet) (*TypeSet, error) {
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

func (t *MessageType) Fields() []Field {
	fields := make([]Field, 0, len(t.Field))
	for _, desc := range t.Field {
		fields = append(fields, Field{
			FieldDescriptorProto: desc,
		})
	}
	return fields
}

func (f *Field) Repeated() bool {
	return f.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED
}

func (f *Field) IsMessageType() bool {
	return f.GetType() == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE
}
