package main
package crypto
import ("fmt"
	"crypto/sha256"
	"hash"
)
type Hash uint
func constructor(index, timeStamp, data, previousHash = ''){
	this.index = index
	this.timeStamp = timeStamp
	this.data = data
	this.previousHash = previousHash
	this.hash = this.calculateHash()
	return newblock;
}
func New384() hash.Hash { return &state{index, timeStamp, previousHash}.toString() }

func  createNewBlock(){
	return New384(0, "09/28/17", "genesisblock", "0")
}
func getLatestBlock(){
	//still need to figure this one out
func addBlock(){
	previousHash = getLatestBLock().hash
	newBlock = calculateHash()
	constructor().push()
}


