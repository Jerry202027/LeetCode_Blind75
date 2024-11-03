package main

// https://leetcode.com/problems/reverse-bits/?envType=problem-list-v2&envId=oizxjoit

// test case
// num := uint32(0b00000010100101000001111010011100)
// fmt.Println(reverseBits(num))

// my concept (kind of brute-force solution)
// time O(32), space O(1)
// func reverseBits(num uint32) uint32 {
// 	numStr := fmt.Sprintf("%032b", num)
// 	reverseNum := 0

// 	for i, e := range numStr {
// 		if e == '1' {
// 			reverseNum += 1 << i
// 		}
// 	}

// 	return uint32(reverseNum)
// }

////////////////////////////////////////////////////////////////////////////////////////////////

// bits operation(the efficacy is better than string operations)
// time O(32), space O(1)
// func reverseBits(num uint32) uint32 {
// 	var reverseNum uint32 = 0
// 	for i := 0; i < 32; i++ {
// 		reverseNum = (reverseNum << 1) | (num & 1)
// 		num >>= 1
// 	}

// 	return reverseNum
// }

// //////////////////////////////////////////////////////////////////////////////////////////////
// lookup Table Method
// time O(1), space O(256)
var lookupTable [256]uint32

func initLookupTable() {
	for i := 0; i < 256; i++ {
		lookupTable[i] = reverseByte(uint8(i))
	}
}

func reverseByte(b uint8) uint32 {
	var result uint32 = 0
	for i := 0; i < 8; i++ {
		result = (result << 1) | uint32(b&1)
		b >>= 1
	}

	return result
}

func reverseBits(num uint32) uint32 {
	initLookupTable()
	return (lookupTable[num&0xFF] << 24) | (lookupTable[(num>>8)&0xFF] << 16) | (lookupTable[(num>>16)&0xFF] << 8) | (lookupTable[(num>>24)&0xFF])
}
