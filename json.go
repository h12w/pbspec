package pbspec

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/descriptorpb"
)

func ToJSON(fileSet *descriptorpb.FileDescriptorSet) ([]byte, error) {
	return (protojson.MarshalOptions{
		Multiline: true,
		Indent:    "  ",
	}).Marshal(fileSet)
}

func FromJSON(jsonBytes []byte) (*descriptorpb.FileDescriptorSet, error) {
	var fileSet descriptorpb.FileDescriptorSet
	if err := protojson.Unmarshal(jsonBytes, &fileSet); err != nil {
		return nil, err
	}
	return &fileSet, nil
}
