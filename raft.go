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

	// "github.com/dianchengwangCHN/raft-key-value-store/labgob"
	// "github.com/dianchengwangCHN/raft-key-value-store/labrpc"
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
	id			uint64 // id of Raft node
	mu        	sync.Mutex          // Lock to protect shared access to this peer's state
	peers     []*labrpc.ClientEnd // RPC end points of all peers
	persister *Persister          // Object to hold this peer's persisted state
	me        int                 // this peer's index into peers[]

	// Your data here (2A, 2B, 2C).
	// Look at the paper's Figure 2 for a description of what
	// state a Raft server must maintain.
	state       ServerState // state of server
	currentTerm int         // latest term server has seen
	votedFor    int         // candidateID
	log         []LogEntry  // log entries
	commitIndex int         // index of highest log entry known to be committed
	lastApplied int         // index of highest log entry applied to state machine

	seed     rand.Source // the source used to generate random timeout duration
	timer    *time.Timer // the timer used for timeout
	voteRecv int         // the number of votes that have been received

	//state the Leader need to maintain
	nextIndex  []int // index of the next log entry
	matchIndex []int // index of the highest log entry known to be replicated on each server

	applyCh        chan ApplyMsg // channel to send ApplyMsg
	stateUpdateCh  chan struct{} // channel to receive signal indicating server state has changed
	commitUpdateCh chan struct{} // channel to receive signal indicating commitIndex has changed
}

func (rf *Raft) startLogCommitterDaemon() {
	for {
		<-rf.commitUpdateCh

		for i, commitIndex := rf.lastApplied+1, rf.commitIndex; i <= commitIndex; i++ {
			// DPrintf("server %d LogCommitter tried to get the Lock\n", rf.me)
			rf.mu.Lock()
			// DPrintf("server %d LogCommitter got the Lock\n", rf.me)
			if offset := rf.log[0].EntryIndex; i > offset {
				rf.applyCh <- ApplyMsg{
					CommandValid: true,
					Command:      rf.log[i-offset].Command,
					CommandIndex: rf.log[i-offset].EntryIndex,
				}
				rf.lastApplied = i
			} else {
				i = offset
			}
			rf.mu.Unlock()
			// DPrintf("server %d applies Entries %d\n", rf.me, i)
		}
	}
}