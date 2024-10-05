package main

import (
	"log"
	"reflect"

	pb "github.com/nithinkm8055/proto-go-course/proto"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A name",
		SimpleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {

	return &pb.Complex{
		OneDummy: &pb.Dummy{
			Id:   1,
			Name: "one dummy",
		},
		MultipleDummies: []*pb.Dummy{
			{
				Id:   2,
				Name: "two dummy",
			},
			{
				Id:   3,
				Name: "three dummy",
			},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		Color: pb.ColorType_COLOR_TYPE_BLUE,
	}
}

func doOneOfs(message any) {
	switch x := message.(type) {
	case *pb.Result_Id:
		log.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		log.Println(message.(*pb.Result_Message).Message)
	default:
		log.Printf("message has unexpected type: %v", x)

	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"id1": {Id: 1},
			"id2": {Id: 2},
			"id3": {Id: 3},
			"id4": {Id: 4},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"
	writeToFile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)
}

func doJson(p proto.Message, t reflect.Type) {

	x := toJson(p)
	log.Println("JSON representation is", x)

	message := reflect.New(t).Interface().(proto.Message)
	log.Println(fromJson(x, message))

}

func main() {
	log.Println(doSimple())
	log.Println(doComplex())
	log.Println(doEnum())
	doOneOfs(&pb.Result_Id{
		Id: 1,
	})
	doOneOfs(&pb.Result_Message{
		Message: "this is a one of message",
	})

	log.Println(doMap())

	doFile(doSimple())

	doJson(doSimple(), reflect.TypeOf(pb.Simple{}))
}
