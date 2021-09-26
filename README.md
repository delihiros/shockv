# ShockV

ShockV is a simple key-value store based on badgerDB with RESTful API.
It's best suited for experimental project which you need a lightweight data store.


## Usage

### Install

```
go install github.com/delihiros/shockv@latest
```

### Server

You have to start server first. The server will handle REST request from any kind of HTTP client.

To start ShockV server:

```
shockv server start
```

### Client

CRUD method is available through REST API. ShockV provides CLI client to ease interaction with ShockV.

To make database:

```
shockv client new --database hello --diskless false
curl "http://localhost:8080/new?database=hello&diskless=false"
```

if you set `--diskless` to `true` the database will be a diskless one and does not preserve data when server is down.

Setting data:

```
shockv client set --database hello --key 1 --value "xyz, abc"
curl -X POST -H "Content-Type: application/json" -d '{"key": "1", "value": "xyz, abc"}' http://localhost:8080/hello
```

Getting the value:

```
shockv client get --database hello --key 1
curl localhost:8080/hello/1
```

Deleting the data:

```
shockv client delete --database hello --key 1
curl -X DELETE http://localhost:8080/hello/1
```