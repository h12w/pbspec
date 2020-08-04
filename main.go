package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	binFile, err := ioutil.TempFile(".", "proto2json-")
	if err != nil {
		return err
	}
	binFile.Close()
	defer os.Remove(binFile.Name())
	protoCmd := exec.Command("protoc", append(os.Args[1:],
		"--include_imports", "--include_source_info",
		"-o"+binFile.Name(),
	)...)
	protoCmd.Stderr = os.Stderr
	if err := protoCmd.Run(); err != nil {
		return err
	}
	var set descriptorpb.FileDescriptorSet
	binBuf, err := ioutil.ReadFile(binFile.Name())
	if err != nil {
		return err
	}
	if err := proto.Unmarshal(binBuf, &set); err != nil {
		return err
	}
	outBuf, err := (protojson.MarshalOptions{
		Multiline: true,
		Indent:    "  ",
	}).Marshal(&set)
	if err != nil {
		return err
	}
	os.Stdout.Write(outBuf)
	return nil
}
