# Scan API

## Env
>recommend
```text
go 1.10+
mongo 3.6.3
```

## Project structure
```text
├── api: api interface
│   ├── handlers: router handler
│   └── routers:  the http router
├── chart: chart data processor
│   ├── address: address chart processor
│   ├── block: block count and reward chart processor
│   ├── blockdifficulty: block difficulty chart processor
│   ├── blocktime: block average time chart processor
│   ├── hashrate: hashrate chart processor
│   ├── topminers: topminers chart processor
│   ├── tx_history: transaction hsitory chart processor
├── cmd: app entrance
|   ├── chart_service: chart service entrance
|   ├── node_service: node service entrance
│   └── scan_server:  http service entrance
├── database: mongodb database
├── log: third logger warpper
├── node: node service
├── rpc:  json rpc
├── rpc:  scan server
├── vendor: third dependencies

```

## Start
```text
cd scan-api
# generate the executable file
make
# start scan_server
./build/server/scan_server -c server.json
# start chart_service
./build/chart/chart_service -c server.json
# start node_service
./build/node/node_service -c server.json
```

## Config
```text
# scan_server

"GinMode":"debug",
# gin framewrok run mode debug 

    "Addr": ":8888", 
    "LimitConnection": 0,
    "DefaultHammerTime": 30,
    "RpcURL": "127.0.0.1:55028",
    "WriteLog": true,
    "LogLevel": "debug",
    "LogFile": "scan-api.log",
    "DisableConsoleColor": true,
    "LimitConnections": 0,
    "MaxHeaderBytes": 20, 
    "ReadTimeout":300,
    "IdleTimeout": 0,
    "WriteTimeout":120,
    "SyncSwitch":true,
    "BlockCacheLimit":1024,
    "TransCacheLimit":1024,
    "DataBaseConnUrl":"127.0.0.1:27017",
    "DataBaseName":"seele",
    "Interval":30
```