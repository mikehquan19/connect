# GRAPH QL USING GOLANG 

## GraphQL: 

- Users get exactly what they want so that they can avoid overfetching and underfetching => smaller JSON payloads and faster network transfers.

- Single endpoint that allows you to do everything. 

- Support of nested data


Run
```
go run github.com/99designs/gqlgen init  
``` 
Then it will do some code generation to all the stuff 

Edit the schema 

Run
```
go get github.com/99designs/gqlgen@v0.17.81 
go run github.com/99designs/gqlgen generate          
```

We got the resolver, then implement the resolver