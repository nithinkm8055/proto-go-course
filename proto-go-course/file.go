package main

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func writeToFile(fname string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("cant serialize to bytes", err)
	}

	if err = os.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("cant write to file", err)
	}

	fmt.Println("data has been written")
}

func readFromFile(fname string, pb proto.Message) {
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("error reading from file", err)
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("cannot unmarshal", err)
	}
	log.Println("read to simple message ", pb)
}
