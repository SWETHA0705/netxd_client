package main

import (
	"context"
	"fmt"
	"log"
	netxdcustomerconstants "github.com/SWETHA0705/netxd_server/netxd_customer_constants"

	netxdcustomer "github.com/SWETHA0705/netxd_customer/customer"

	"google.golang.org/grpc"
)

func main(){
	conn, err := grpc.Dial(netxdcustomerconstants.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := netxdcustomer.NewCustomerServiceClient(conn)
	response, err := client.CreateCustomer(context.Background(), &netxdcustomer.Customer{
		CustomerId: 0,
		FirstName:  "v",
		LastName:   "ss",
		BankId:     45678,
		Balance:    3456,
		
	})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %s\n", response.CreatedAt)
}


