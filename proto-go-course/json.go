package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJson(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	bytes, err := option.Marshal(pb)
	// bytes, err := protojson.Marshal(pb)

	if err != nil {
		log.Fatalln("cannot marshal input message", err)
	}

	return string(bytes)
}

func fromJson(msg string, pb proto.Message) proto.Message {
	err := protojson.Unmarshal([]byte(msg), pb)
	if err != nil {
		log.Fatalln("cannot unmarshal input message", err)
	}

	return pb
}
