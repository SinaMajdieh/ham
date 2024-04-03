package ham

import (
	"strconv"

	. "github.com/SinaMajdieh/ham/common"
	file "github.com/SinaMajdieh/ham/file"
)

// Constructing the message in to blocks of BlockSize
func blockalize(message string) []Block {

	i := 0
	j := 1
	blocks := []Block{}
	block := NewBlock()

	// going over the message
	for i < len(message) {
		// block is complete
		if j >= BlockSize {
			blocks = append(blocks, block)
			block = NewBlock()
			j = 1
		}
		// index is a parity check skipping it
		if IsParityBit(j) {
			j++
			continue
		}

		//storing the data in block
		bit, _ := strconv.Atoi(string(message[i]))
		block[j] = bit
		i++
		j++
	}
	// if the last block is not zero add it to blocks
	if block.IsZero() == false {
		blocks = append(blocks, block)
	}

	return blocks
}

// Arranging the blocks of the message so that if a sequence of bits were flipped the message remain error robust
func arrangeBlocksToString(blocks []Block) string {
	arranged_message := ""
	// setting the remaining storage to be zero
	if len(blocks) != BlocksInStorage {
		FillStorage(&blocks)
	}
	// Arranging the first bit of every block and then the second bit of every block and ...
	for i := 0; i < BlockSize; i++ {
		for _, b := range blocks {
			arranged_message += strconv.Itoa(b[i])
		}
	}

	return arranged_message
}

// Convert the message into an error robust message
// with the ext of .ham
// saving it into a file
func HamIt(input string, inputIsFileName bool, saveOutputAsFile bool) string {
	var message string

	// input is a file name
	if inputIsFileName {
		// Reading the contents of the file
		message = file.ReadFile(input)
	} else {
		message = input
	}

	// converting them into the Bitstring (binary representation)
	bitstring := Byteslice(message).AsBitString()

	// constructing the message into blocks
	blocks := blockalize(string(bitstring))

	// setting the parity bits for all blocks
	for i := range blocks {
		blocks[i].SetParities()
	}

	// Arranging the blocks so they are error robust to a sequence of altered bits
	ham_message := arrangeBlocksToString(blocks)

	//write output into a file
	if saveOutputAsFile {
		// writing to the file
		file.WriteFile(input, ham_message)
		return ""
	}

	return ham_message
}
