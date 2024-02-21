package main

import (
	"fmt"
	pb "go-proto/pb"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	products := &pb.Products{
		Pagination: &pb.Pagination{
			PerPage:     10,
			Total:       2,
			CurrentPage: 1,
			LastPage:    1,
		},
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Product 1",
				Price: 100,
				Stock: 1,
				Category: &pb.Category{
					Id:   1,
					Name: "Makanan",
				},
			},
			{
				Id:    2,
				Name:  "Product 2",
				Price: 101,
				Stock: 100,
				Category: &pb.Category{
					Id:   2,
					Name: "Minuman",
				},
			},
		},
	}

	data, err := proto.Marshal(products)
	if err != nil {
		log.Fatal("marshal error", err)
	}
	//compact binary wire format
	fmt.Printf("data: %+v\n", data)
	fmt.Println("--------------------------")

	decode := &pb.Products{}
	if err = proto.Unmarshal(data, decode); err != nil {
		log.Fatal("Unmarshal error", err)
	}
	fmt.Printf("decode: %+v", decode)
}
