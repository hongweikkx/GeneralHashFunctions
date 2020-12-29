package GeneralHashFunctions

import "math/bits"

// RSHash from Robert Sedgwicks Algorithms
func RSHash(str string) uint {
	strLen := len(str)
	var b uint = 378551
	var a uint = 63689
	var hash uint = 0
	for i := 0; i < strLen; i++ {
		hash = hash*a + uint(str[i])
		a = a * b
	}
	return hash
}

// JSHash is a bitwise hash function written by Justin Sobel
func JSHash(str string) uint {
	strLen := len(str)
	var hash uint = 1315423911
	for i := 0; i < strLen; i++ {
		hash ^= (hash << 5) + uint(str[i]) + (hash >> 2)
	}
	return hash
}

// PJWHash: This hash algorithm is based on work by Peter J. Weinberger of Renaissance Technologies.
// 			The book Compilers (Principles, Techniques and Tools) by Aho, Sethi and Ulman,
// 			recommends the use of hash functions that employ the hashing methodology found in this particular algorithm.
func PJWHash(str string) uint {
	var BitsInUnsignedInt uint = bits.UintSize
	var ThreeQuarters uint = (BitsInUnsignedInt * 3) / 4
	var OneEighth uint = BitsInUnsignedInt / 8
	var HighBits uint = 0xFFFFFFFF << (BitsInUnsignedInt - OneEighth)
	var hash uint = 0
	var test uint = 0
	for i := 0; i < len(str); i++ {
		hash = (hash << OneEighth) + uint(str[i])
		test = hash & HighBits
		if test != 0 {
			hash = (hash ^ (test >> ThreeQuarters)) & (^HighBits)
		}
	}
	return hash
}

// ELFHash: Similar to the PJW Hash function, but tweaked for 32-bit processors. It is a widley used hash function on UNIX based systems.
func ELFHash(str string) uint {
	var hash uint = 0
	var x uint = 0
	for i := 0; i < len(str); i++ {
		hash = (hash << 4) + uint(str[i])
		x = hash & 0xF0000000
		if x != 0 {
			hash ^= x >> 24
		}
		hash &= ^x
	}
	return hash
}

// BKDRHash: This hash function comes from Brian Kernighan and Dennis Ritchie's book "The C Programming Language".
//           It is a simple hash function using a strange set of possible seeds which all constitute a pattern of 31....31...31 etc,
//           it seems to be very similar to the DJB hash function.
func BKDRHash(str string) uint {
	var seed uint = 131 /* 31 131 1313 13131 131313 etc.. */
	var hash uint = 0
	for i := 0; i < len(str); i++ {
		hash = hash*seed + uint(str[i])
	}
	return hash
}

// SDBHash: This is the algorithm of choice which is used in the open source SDBM project.
//			The hash function seems to have a good over-all distribution for many different data sets.
//			It seems to work well in situations where there is a high variance in the MSBs of the elements in a data set.
func SDBHash(str string) uint {
	var hash uint = 0
	for i := 0; i < len(str); i++ {
		hash = uint(str[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}

// DJBHash: An algorithm produced by Professor Daniel J. Bernstein and shown first to the world on the usenet newsgroup comp.lang.c.
//			It is one of the most efficient hash functions ever published.
func DJBHash(str string) uint {
	var hash uint = 5381
	for i := 0; i < len(str); i++ {
		hash = (hash << 5) + hash + uint(str[i])
	}
	return hash
}

// DEKHash: An algorithm proposed by Donald E. Knuth in The Art Of Computer Programming Volume 3,
//			under the topic of sorting and search chapter 6.4.
func DEKHash(str string) uint {
	strLen := len(str)
	var hash uint = uint(strLen)
	for i := 0; i < strLen; i++ {
		hash = (hash << 5) ^ (hash >> 27) ^ uint(str[i])
	}
	return hash
}

func BPHash(str string) uint {
	var hash uint = 0
	for i := 0; i < len(str); i++ {
		hash = hash<<7 ^ uint(str[i])
	}
	return hash
}

func FNVHash(str string) uint {
	var fnvPrime uint = 0x811C9DC5
	var hash uint = 0
	for i := 0; i < len(str); i++ {
		hash *= fnvPrime
		hash ^= uint(str[i])
	}
	return hash
}

// APHash: An algorithm produced by me Arash Partow.
//		   I took ideas from all of the above hash functions making a hybrid rotative and additive hash function algorithm.
//         There isn't any real mathematical analysis explaining why one should use this hash function instead of the others described above
//         other than the fact that I tired to resemble the design as close as possible to a simple LFSR.
//         An empirical result which demonstrated the distributive abilities of the hash algorithm was obtained using a hash-table with 100003 buckets,
//         hashing The Project Gutenberg Etext of Webster's Unabridged Dictionary, the longest encountered chain length was 7,
//         the average chain length was 2, the number of empty buckets was 4579.
func APHash(str string) uint {
	var hash uint = 0xAAAAAAAA
	for i := 0; i < len(str); i++ {
		if i&1 == 0 {
			hash ^= (hash << 7) ^ uint(str[i])*(hash>>3)
		} else {
			hash ^= ^((hash << 11) + (uint(str[i]) ^ (hash >> 5)))
		}
	}
	return hash
}
