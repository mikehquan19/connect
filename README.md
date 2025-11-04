# Content for the API Workshop

## How to run: 

Setup the Mongo database on Atlas, and get the connection string for the cluster.

After that, put the connection string to the ```.env``` file

Run:
```
go build -o connect server.go && chmod +x connect
./connect

OR 

go run server.go
```
