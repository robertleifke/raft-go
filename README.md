# Raft 

This repository implements Raft in Go.

## Overview 

In computer science there is a common problem whenever you deal with more than one computer, a `node`.  That is, to make sure all the nodes in a system can agree to what is happening - creating a shared state even when one node gets disconnected. This process of reaching agreement is called consensus. In the case of Raft, it's a simple algorithm for achieving this consensus through the use of an elected leader. The leader dictates what changes are accepted or rejected to reach consensus for which all the other nodes follow. These followers must implement the changes or get booted from the system. If the elected leader goes down, a new leader gets elected from potential candidates who democratically get elected by the followers.

## Usage

```bash
go run raft.go
```

## Dependencies

```bash
run go mod vendor to install dependencies
```

Requires the protoc compiler with the go output plugin installed as it uses Protocol Buffers.

Redis should be installed with default options.