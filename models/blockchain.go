package models

import "time"

type Blockchain struct {
	GenesisBlock Block
	Chain        []Block
	Difficulty   int
}

func NewBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		PreviousHash: "0",
		TimeStamp:    time.Now(),
	}
	genesisBlock.Hash = genesisBlock.CalculateHash()
	return Blockchain{
		GenesisBlock: genesisBlock,
		Chain:        []Block{genesisBlock},
		Difficulty:   difficulty,
	}
}

// AddBlock añade un nuevo bloque a la cadena
func (b *Blockchain) AddBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.Chain[len(b.Chain)-1]
	newBlock := Block{
		Data:         blockData,
		PreviousHash: lastBlock.Hash,
		TimeStamp:    time.Now(),
	}
	newBlock.Mine(b.Difficulty)
	b.Chain = append(b.Chain, newBlock)
}

// IsValid verifica la validez de toda la cadena de bloques
func (b Blockchain) IsValid() bool {
	for i := range b.Chain[1:] {
		previousBlock := b.Chain[i]
		currentBlock := b.Chain[i+1]
		if currentBlock.Hash != currentBlock.CalculateHash() || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}
