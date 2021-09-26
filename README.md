# ShockV

ShockV is a simple key-value store based on badgerDB with RESTful API.
It's best suited for experimental project which you need a data store.


## Usage

To start ShockV server:

```
shockv server start
```

To make database:

```
shockv client new --database hello --diskless false
```
if you set `--diskless` to `true` the database will be a diskless one and does not preserve data when server is down.

Setting data:

```
shockv client set --database hello --key 1 --value "xyz, abc"
```

Getting the value:

```
shockv client get --database hello --key 1
```