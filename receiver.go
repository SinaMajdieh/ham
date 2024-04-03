package ham

import (
	"strconv"

	. "github.com/SinaMajdieh/ham/common"
	file "github.com/SinaMajdieh/ham/file"
)

// Rearrange all the bits to their corresponding blocks, removing the noise if there are any and ignoring the leading zero blocks
func rearrangeBlocks(message string) []Block {
	blocks := []Block{}

	// going over all the blocks that can fit in storage
	for i := 0; i < BlocksInStorage; i++ {
		block := NewBlock()
		// Extracting the corresponding bits of the ith block
		for j := 0; j < BlockSize; j++ {
			block[j], _ = strconv.Atoi(string(message[(j*BlocksInStorage)+i]))
		}

		// Detecting the noise if there is any and correcting it
		_ = block.CorrectNoise()
		// if noise_position != 0 {
		// 	println("There was a noise in INDEX ", (i*BlockSize)+noise_position)
		// }

		// if the block is a leading zero breaking out of the loop
		if block.IsZero() {
			break
		}

		// adding the completed and error free block to the blocks
		blocks = append(blocks, block)
	}
	return blocks
}

// extract the data stored in non parity blocks
func extractData(blocks []Block) string {
	data := ""

	// going over blocks
	for _, block := range blocks {
		// going over the values stored in a block
		for i, bit := range block {
			// if the index of the value is a parity then skip
			// otherwise add it to data
			if IsParityBit(i) || i == 0 {
				continue
			}
			data += strconv.Itoa(bit)
		}
	}
	return data
}

func UnHamIt(input string, inputIsFileName bool) string {
	var message string

	//in put is a file name
	if inputIsFileName {
		//Read the contents of the file
		message = file.ReadFile(input)
	} else {
		message = input
	}

	// Rearrange the blocks
	blocks := rearrangeBlocks(message)

	// Extract the data that are not for parity check
	data := Bitstring(extractData(blocks))

	// Removing the leading zeros
	leading_zeros_index := len(data) - (len(data) % 8)
	data = data[:leading_zeros_index]

	// converting to bytes
	byte_message := data.AsByteSlice()

	return string(byte_message)
}
