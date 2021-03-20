package pbspec

import (
	"io/ioutil"
	"os"
	"os/exec"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Load the descriptors in the proto files with the same command line arguments as protoc, e.g. -IPATH1 -IPATH2 ... PROTO_FILES
func Load(protocArgs []string) (*descriptorpb.FileDescriptorSet, error) {
	binFile, err := ioutil.TempFile(".", "pbspec-")
	if err != nil {
		return nil, err
	}
	binFile.Close()
	defer os.Remove(binFile.Name())
	protoCmd := exec.Command("protoc", append(protocArgs,
		"--include_imports", "--include_source_info",
		"-o"+binFile.Name(),
	)...)
	protoCmd.Stderr = os.Stderr
	if err := protoCmd.Run(); err != nil {
		return nil, err
	}
	var set descriptorpb.FileDescriptorSet
	binBuf, err := ioutil.ReadFile(binFile.Name())
	if err != nil {
		return nil, err
	}
	if err := proto.Unmarshal(binBuf, &set); err != nil {
		return nil, err
	}
	return &set, nil
}
