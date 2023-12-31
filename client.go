package main

import (
	"context"
	"fmt"

	"log"

	constants "github.com/Menaha-Chandrasekar/netxd_constants"
	pb "github.com/Menaha-Chandrasekar/netxd_customer/customer_proto"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(constants.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)
	

	response, err := client.CreateCustomer(context.Background(), &pb.CustomerRequest{
		CustomerId: 400,
		FirstName:  "Niranjana",
		LastName:   "P C",
		BankId:     7895,
		Balance:    50000,
		
	})
	if err != nil {
		log.Fatalf("Failed to call: %v", err)
	}

	fmt.Printf("CustomerID: %d\nCreatedTime:%v\n", response.CustomerId,response.CreatedAt)
}
