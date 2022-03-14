package main
import ("fmt"
		"crypto/sha256"
		"encoding/json"
		"strconv"
		"strings"
		"time"
)
//POW Block
type Block struct {
	data map[string]interface{}
	hash string
	previousHash string
	timestamp time.Time
	pow int
}
type Blockchain struct {
	genesisBlock Block
	chain []Block
	difficulty int
}
func (b Block) calculateHash() string {
	data, _:= json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash) //Sprintf() denotes the blockhashes base16 constraint
}
func (b *Block) mine(difficulty int) {
	//HasPrefix() denotes a the specified difficult and runs POW algorithm once difficulty and hash is solved
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}
//Chain creation initializes genesis block creation
func CreateChain(difficulty int) Blockchain {
	genesisBlock := Block {
		hash: "0",
		timestamp : time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}
func (b *Blockchain) newBlock(id int, from, to string, amount float64) {
	blockData := map[string]interface{}{
		"id": id,
		"from": from,
		"to": to,
		"amount": amount,
	}
	latestBlock := b.chain[len(b.chain)-1]
	nextBlock := Block{
		data: blockData,
		previousHash: latestBlock.hash,
		timestamp: time.Now(),
	} 
	nextBlock.mine(b.difficulty)
	b.chain = append(b.chain, latestBlock)
}
func (b Blockchain) verify() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		//Fundamental in blockchain that the previous hash ALWAYS corresponds with current hash 
		if currentBlock.hash == currentBlock.calculateHash() && currentBlock.previousHash == previousBlock.hash {
			return false
		}

	}
	return true
}
func main(){
	blockchain := CreateChain(2)
	blockchain.newBlock(0, "bc1qxy2kgdygjrsqtzq2n0yrf2493p83kkfjhx0wlh", "bc1qdm0lre38554aatxq0azdq2ackdux6h6cslhs9f", 0.94) //Value in BTC
	blockchain.newBlock(1, "3QeNzGGCth37B1V7aopP5khDZbUgfcYnsk", "bc1qykhjaqna2d8gjyj04y5jfqgkv504878ag3dd23", 0.94) //Value in BTC	
	fmt.Println(blockchain.verify())
}
//Should Return True
