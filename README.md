# Distributed Cache System

The Distributed Cache System is a scalable, high-performance, and fault-tolerant caching solution built using Go and HashiCorp's Memberlist library. It allows data to be cached across multiple nodes in a distributed cluster, enabling robust data storage, efficient retrieval, and automatic failure detection using a gossip-based protocol.

This project extends the functionality of a basic cache by implementing distributed caching across multiple nodes, supporting cluster membership and dynamic peer discovery.

## Key Features

- **Distributed Cache:** Seamless caching across multiple nodes, enabling high availability and fault tolerance.
- **Cluster Membership:** Automatic peer discovery and cluster management using a gossip-based protocol, powered by HashiCorp’s Memberlist.
- **Failure Detection:** Real-time detection and handling of node failures, ensuring uninterrupted service.
- **HTTP API:** Simple REST API for interacting with the cache, supporting basic CRUD operations (Get, Put, Delete).
- **Configurable Gossip Parameters:** Fine-tune gossip and failure detection settings for optimized performance.

## Tech Stack

- **Language:** Go
- **Library:** HashiCorp Memberlist
- **Network Communication:** HTTP for client interaction, Gossip protocol for node-to-node communication

## Installation

Go 1.18 or later
Git


- ***Clone the Repository:***

   ```bash
    git clone https://github.com/notlelouch/distributed-cache.git
    cd distributed-cache
   ```
   
- ***Install Dependencies:***
  ```bash
  go mod tidy
  ```
  
## Usage

- ### Running a Distributed Cache Node
  Each instance of the cache runs on a specific port, and nodes can join a cluster by connecting to existing peers. To start a node(be in the root directory):
  - Start the first node(in one terminal):
   ```bash
    export PORT=8080
    export NODE_NAME=<node_name>
    make run
   ```
    - Join the cluster(in another terminal instance):
   ```bash
    export PORT=<port>
    export NODE_NAME=<node_name>
    export PEER=127.0.0.1:8080
    make run
   ```

     - Example:
     ```bash
    export PORT=8084
    export NODE_NAME=node4
    export PEER=127.0.0.1:8080
    make run
   ``` 
- ### Interacting with the Cache
  The cache can be accessed via simple HTTP requests. Each node in the cluster can handle HTTP requests to interact with the distributed cache.
  1. #### Get a Value:
      Retrieve a cached value by sending a `GET` request to `/cache/{key}`.
  
     ```bash
      curl -X GET http://localhost:8001/cache/myKey
     ```
     ***Response:*** Returns the cached value if found, or a 404 if the key is not in the cache.

  2. #### Put a Value:
      Store a value in the cache using a `PUT` request with `value` and `duration` as parameters. `duration` is the expiration time (in seconds) for the cached value
  
     ```bash
      curl -X PUT http://localhost:8001/cache/myKey -d "value=myValue&duration=60"
     ```
     ***Response:*** Returns the cached value if found, or a 404 if the key is not in the cache.
     ***Parameters:***
      - `value`: The value to store in the cache.
      - `duration`: How long (in seconds) the value should be stored.

  4. #### Delete a Value:
      Remove a cached value by sending a DELETE request to /cache/{key}.
     ```bash
      curl -X DELETE http://localhost:8001/cache/myKey
     ```
     ***Response:*** Deletes the key if found, no output on success.

## Project Structure

```
├── cmd/
│   └── server/
│       └── main.go               # Entry point, starts the cache and joins the cluster
├── pkg/
│   ├── cache/
│   │   ├── cache.go              # Core cache logic for managing data storage and expiration
│   │   └── cache_test.go         # Test file for cache.go
│   └── distributed/
│       ├── distributed.go        # Implementation of the distributed cache, cluster management, HTTP API handlers
│       └── distributed_test.go   # Test file for distributed.go
├── go.mod                        # Go module dependencies
├── go.sum                        # Go module versions
├── README.md                     # Project documentation
└── LICENSE                       # License file for the project
```
### Important Files

- **cache.go:** Contains the basic cache functionality (get, set, delete, expiration).
- **distributed.go:** Handles cluster membership, peer discovery, failure detection, and HTTP request handling.
- **main.go:** Entry point for running the distributed cache node.

## Configuration
The cache uses default settings for the gossip-based protocol and cluster management. However, you can customize the following parameters in `distributed.go` for tuning performance:
- **GossipInterval:** Time interval between gossip messages.
- **GossipNodes:** Number of nodes to gossip with in each interval.
- **ProbeInterval:** Frequency of checking for failed nodes.
- **ProbeTimeout:** Timeout for failure detection after missing heartbeats.

```bash
  config.GossipInterval = 300 * time.Millisecond
  config.GossipNodes = 3
  config.ProbeInterval = 1 * time.Second
  config.ProbeTimeout = 5 * time.Second
```

## How It Works
1. **Cluster Membership:** Each node maintains a list of active peers, using a gossip-based protocol to share state information about the current cluster.
2. **Failure Detection:** Nodes send periodic heartbeat messages to peers. If a node doesn’t respond within the timeout, it’s marked as suspect and eventually considered failed if it doesn’t recover.
3. **Cache Operations:** The HTTP API allows users to `GET`, `PUT`, and `DELETE` keys in the cache. Each node manages its local cache, and future updates may include cache synchronization across nodes.


## Contributing

Contributions are welcome! If you have ideas for improvements, new features, or bug fixes, please open an issue or submit a pull request.
