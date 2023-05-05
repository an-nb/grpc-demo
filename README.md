## 1. Clone the repository

## 2. Prerequisites
https://grpc.io/docs/languages/go/quickstart/

## 3. Create a new .env file in the project root directory
- Set the env as in ```.env.example```

## 4. Set up the database
- Create a new database as you set in .env
- Create a new ```invoice``` table
```
CREATE TABLE invoice (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    sender VARCHAR(255) NOT NULL,
    receiver VARCHAR(255) NOT NULL
);
```

## 5. Start the server
- ```make run_server```

## 6. Test using Postman
- New > gRPC Request > Import the .proto file
- Enter the Server URL: ```grpc://localhost:{port}``` or ```localhost:{port}``` 
- Choose "Use Example Message" to load the correct message format

## 7. General steps to build a gRPC server with Go
- First step: create a .proto file
  - Define messages
  - Define a service
    - With RPC methods
- Then you need to generate the Go code that will be used (with protoc)
  - To encode/decode messages with Protocal Buffers
  - To handle incoming gRPC requests
  - Generate the Go code by running ```make generate_grpc_code```
- Note:
  - You need to regenerate the code each time you change your .proto file
  - Do not touch the generated code as changes will be removed at the next generation