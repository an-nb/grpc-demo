package main

import (
	"database/sql"
	"fmt"
	"grpc-demo/invoicer"
	"log"
	"net"
	"os"

	server "grpc-demo/server"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// get environment variables
	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// construct DSN string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	// establish a connection to the MySQL server
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL server: %v", err)
	}

	defer db.Close()

	// create the gRPC server
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Cannot create listener: %s", err)
	}

	fmt.Printf("Starting server on %s", port)

	serverRegistrar := grpc.NewServer()
	service := &server.MyInvoicerServer{DB: db}
	invoicer.RegisterInvoicerServer(serverRegistrar, service)

	// serve the gRPC server
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("Impossible to serve: %s", err)
	}
}
