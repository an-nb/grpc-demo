package server

import (
	"context"
	"database/sql"
	"fmt"
	"grpc-demo/invoicer"
	"log"
)

type MyInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
	DB *sql.DB
}

func (s *MyInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	// insert the data into the database
	result, err := s.DB.Exec("INSERT INTO invoice (name, price, sender, receiver) VALUES (?, ?, ?, ?)", req.Product.Name, req.Product.Price, req.Sender, req.Receiver)
	if err != nil {
		log.Printf("Failed to insert data into database: %v", err)
		return nil, err
	}

	// get the ID of the inserted row
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Failed to get the ID of the inserted row: %v", err)
		return nil, err
	}

	message := fmt.Sprintf("product %s created with ID %d, sender: %s, receiver: %s", req.Product.Name, id, req.Sender, req.Receiver)

	// return the response
	return &invoicer.CreateResponse{
		Message: message,
	}, nil
}

func (s MyInvoicerServer) Get(ctx context.Context, req *invoicer.GetRequest) (*invoicer.GetResponse, error) {
	var (
		name     string
		price    int64
		sender   string
		receiver string
	)

	// get the data from the database
	err := s.DB.QueryRow("SELECT name, price, sender, receiver FROM invoice WHERE id=?", req.Id).Scan(&name, &price, &sender, &receiver)
	if err != nil {
		return nil, fmt.Errorf("error fetching data from db: %s", err)
	}

	// return the response
	return &invoicer.GetResponse{
		Product: &invoicer.Product{
			Name:  name,
			Price: price,
		},
		Sender:   sender,
		Receiver: receiver,
	}, nil
}
