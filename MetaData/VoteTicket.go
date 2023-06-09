package MetaData

//go:generate msgp

type VoteTicket struct {
	VoteResult  []int    `msg:"VoteResult"`
	BlockHashes [][]byte `msg:"hashes"`
	Voter       string   `msg:"Voter"`
	Timestamp   float64  `msg:"Timestamp"`
	Sig         string   `msg:"Sig"`
}