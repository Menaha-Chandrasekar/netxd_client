package main

import (
	"context"
	"fmt"

	"log"

	pb "module/netxd_customer/customer_proto"

	"google.golang.org/grpc"
	
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)
	

	response, err := client.CreateCustomer(context.Background(), &pb.CustomerRequest{CustomerId: 109})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %d  %v\n", response.CustomerId,response.CreatedAt)
}