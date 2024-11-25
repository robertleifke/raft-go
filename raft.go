package raft

import (
	// "bytes"
	// "os"
	// "log"
	// "errors"
	// "math/rand"
	"sync"
	// "time"
	// "fmt"
)

type NodeState int 

const (
	leader 		NodeState = 0
	candidate 	NodeState = 1
	follower 	NodeState = 2

	heartbeatInterval int = 100
	electionTimeoutMin int = 150
	electionTimeoutMax int = 300
)

type Node struct {
	id			uint64 				// id of the node
	mu        	sync.Mutex          // synchronize access to shared state
	peerList    Set 				// list of peers
	rpc     	*ClientEnd 			// RPC end points of all peers
	persister 	*Persister          // Object to hold this peer's persisted state
	database 	*Database			// database
	self        int             // this peer's index into peers[]
	state       NodeState 			// state of the node

	// // state the follower need to maintain
	// currentTerm int        			// latest term server has seen
	// votedFor    int         		// candidateID
	// log         []LogEntry 			// log entries
	// commitIndex int         		// index of highest log entry known to be committed
	// lastApplied int         		// index of highest log entry applied to state machine
	// seed     rand.Source 			// the source used to generate random timeout duration
	// timer    *time.Timer 			// the timer used for timeout
	// voteRecv int         			// the number of votes that have been received

	//state the Leader need to maintain	
	nextIndex  []int // index of the next log entry
	matchIndex []int // index of the highest log entry known to be replicated on each server

	// apply        chan ApplyMsg // channel to send ApplyMsg
	stateUpdate  chan struct{} // channel to receive signal indicating server state has changed
	commitUpdate chan struct{} // channel to receive signal indicating commitIndex has changed
}

// func NewNode(id uint64, peerList Set, rpc []*labrpc.ClientEnd, persister *Persister, database *Database, ready <-chan interface{}, commitChan chan CommitEntry) *Node {
// 	node := &Node{
// 		id:                 id,                      // id is the id of the Raft node
// 		peerList:           peerList,                // List of peers of this Raft node
// 		rpc:             server,                  // Server is the server of the Raft node. Issue RPCs to the peers
// 		database:                 db,                      // Database is the storage of the Raft node
// 		commitChan:         commitChan,              // CommitChan is the channel the channel where this Raft Node is going to report committed log entries
// 		newCommitReady:     make(chan struct{}, 16), // NewCommitReady is an internal notification channel used to notify that new log entries may be sent on commitChan.
// 		trigger:            make(chan struct{}, 1),  // Trigger is the channel used to trigger the Raft node to send a AppendEntries RPC to the peers when some relevant event occurs
// 		currentTerm:        0,                       // CurrentTerm is the current term of the Raft node
// 		votedFor:           -1,                      // VotedFor is the candidate id that received a vote in the current term
// 		log:                make([]LogEntry, 0),     // Log is the log of the Raft node
// 		commitIndex:        0,                       // CommitIndex is the index of the last committed log entry
// 		lastApplied:        0,                       // LastApplied is the index of the last applied log entry
// 		state:              Follower,                // State is the state of the Raft node
// 		electionResetEvent: time.Now(),              // ElectionResetEvent is the time at which the Raft node had last reset its election timer
// 		nextIndex:          make(map[uint64]uint64), // NextIndex is the index of the next log entry to send to each peer
// 		matchIndex:         make(map[uint64]uint64), // MatchIndex is the index of the highest log entry known to be replicated on the leader's peers
// 		rpc:                rpc,
// 		persister:          persister,
// 		database:           database,
// 	}

// 	if node.db.HasData() {
// 		// If the database has data, load  the
// 		// currentTerm, votedFor, and log from
// 		// the database before crashing.

// 		node.restoreFromStorage()
// 	}

// 	// Start the Raft node
// 	go func() {
// 		<-ready                              // Wait for the peers to be initialized
// 		node.mu.Lock()                       // Lock the Raft node
// 		node.electionResetEvent = time.Now() // Reset the election timer
// 		node.mu.Unlock()                     // Unlock the Raft node
// 		node.runElectionTimer()              // Start the election timer
// 	}()

// 	go node.sendCommit() // Start the commit channel as a goroutine
// 	return node
}

