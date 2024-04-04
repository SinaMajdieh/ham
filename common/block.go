package common

import (
	"fmt"
	"math"
)

const (
	// the size of a block
	BlockSize = 256
	// number of bits
	NumberOfBits = 8
	// total number of blocks to be saved
	BlocksInStorage = 3125
)

// the parity bits
var (
	parity_bit = [...]int{128, 64, 32, 16, 8, 4, 2, 1}
)

type Block [BlockSize]int

// Creates a new block setting all its values to 0
func NewBlock() Block {
	block := make([]int, BlockSize)
	for i := 0; i < BlockSize; i++ {
		block[i] = 0
	}
	return Block(block)
}

// If all the values of the block is 0 return true
func (block *Block) IsZero() bool {
	for _, v := range block {
		if v == 1 {
			return false
		}
	}
	return true
}

// Checking if the bit is a parity bit
func IsParityBit(index int) bool {
	for _, b := range parity_bit {
		if index == b {
			return true
		}
	}
	return false
}

// Stringer method for representing the block
func (block Block) String() string {
	result := ""
	side := int(math.Sqrt(float64(BlockSize)))
	for i, bit := range block {
		result += fmt.Sprint(bit, " ")
		if (i+1)%side == 0 {
			result += fmt.Sprintln()
		}
	}
	return result
}

// Check all the parities of the block
func (block Block) checkParities() []int {
	// count the number of ones in each parity group
	var parity_count [NumberOfBits]int

	// going over the block
	// checking to see that the bit belongs to which gropes
	// adding the count of that group
	for i := 1; i < BlockSize; i++ {
		if block[i] == 1 {
			binary_string := fmt.Sprintf("%08b", i)
			for j := NumberOfBits - 1; j >= 0; j-- {
				if string(binary_string[j]) == "1" {
					parity_count[j]++
				}
			}
		}
	}

	return parity_count[:]
}

// the number of ones in a block
func (block Block) numberOfOnes() int {
	number_of_ones := 0
	for _, bit := range block {
		if bit == 1 {
			number_of_ones++
		}
	}
	return number_of_ones
}

// Getting the indexes of all the ones in the block
func (block Block) getOnesPositions() []int {
	ones_positions := []int{}
	for i, bit := range block {
		if bit == 1 {
			ones_positions = append(ones_positions, i)
		}
	}
	return ones_positions
}

// Xor the indexes of all the ones in the block
func (block Block) XorOnes() int {
	//getting the ones position
	ones_positions := block.getOnesPositions()
	if len(ones_positions) == 0 {
		return 0
	}

	// xor the positions
	xor_result := ones_positions[0]
	for i := 1; i < len(ones_positions); i++ {
		xor_result = xor_result ^ ones_positions[i]
	}

	return xor_result
}

// correcting the noise if there were any and returning the index of the noise
func (block *Block) CorrectNoise() int {
	noise_index := block.XorOnes()
	if noise_index != 0 {
		block[noise_index] = block[noise_index] ^ 1
	}
	return noise_index
}

// Setting all the parity bits
func (block *Block) SetParities() {
	// getting the parity counts
	parity_count := block.checkParities()

	// Ensuring the number of ones in each parity group is even
	for i, bit := range parity_count {
		if bit%2 != 0 {
			block[parity_bit[i]] = 1
		}
	}

	// setting the 0 parity bit
	if block.numberOfOnes()%2 != 0 {
		block[0] = 1
	}
}

// filling all the remaining blocks
func FillStorage(blocks *[]Block) {
	for len(*blocks) != BlocksInStorage {
		*blocks = append(*blocks, NewBlock())
	}
}
