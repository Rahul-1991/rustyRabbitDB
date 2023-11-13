# RustyRabbit 🐇 
### An in-memory key value store speaking the redis language

RustyRabbit is an attempt to create a lite version of Redis supporting basic operations. Its written in Go programming language and can be installed on local system to play with. You can connect to this server using the redis-cli command.

## Features

- connect using redis-cli command
- Basic Operations like GET, SET, LPUSH, RPUSH, DEL, INCR, DECR, SAVE
- Provision to set TTL and expire key after TTL is passed (EX, PX, EXAT, PXAT). For usage please refer redis docs
- Handle concurrent connections
- Save backup of in-memory data in .rdb file using compressed format
- Load from the backup file on startup if backup file exists

Working on this project has its own perks. You understand how to handle concurrent tcp connections and learn the language of Redis (RESP). You will also understand how to serialize and deserialize in-memory data and save and load it to and from a compressed backup file.

## Setting It Up
To run rustyrabbit locally you need Go programming language installed.
```
git clone https://github.com/Rahul-1991/rustyRabbitDB
cd rustyRabbitDB
go run main.go
```

## Benchmarks

Using the redis benchmark utility rustyrabbit was tested by simulating multiple clients sending commands at the same time to test concurrency. The results as tested on macbook pro M1 are given below
```
> redis-benchmark -p 6380 -t SET,GET -q
SET: 139082.06 requests per second, p50=0.183 msec
GET: 132802.12 requests per second, p50=0.183 msec
```

## Contributions

Would love to have people go through my code and help in adding more features to rustyrabbit. You can always create a pull request by forking this repository and adding any enhancement in your branch.
