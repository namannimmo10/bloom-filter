package bloomfilter

import (
	"github.com/spaolacci/murmur3"
	"github.com/willf/bitset"
)

// A `BloomFilter` is a space-efficient probabilistic data structure
// that checks whether an element is a member of a set. False positives are
// possible, but false negatives are not.
type BloomFilter struct {
	m uint
	k uint
	b *bitset.BitSet
}

func max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

// `NewBloom` will create a new bloom filter with 'm' bits and 'k' hashing functions.
func NewBloom(m, k uint) *BloomFilter {
	return &BloomFilter{max(1, m), max(1, k), bitset.New(m)}
}

func (f *BloomFilter) Add(a []byte) *BloomFilter {
	h := baseHashes(a)
	for i := uint(0); i < f.k; i++ {
		f.b.Set(f.location(h, i))
	}
	return f
}

// `baseHashes` returns the four hash values of data that are used to create k
// hashes.
func baseHashes(data []byte) [4]uint64 {
	a1 := []byte{1} // to grab another bit of data
	hasher := murmur3.New128()
	hasher.Write(data)
	v1, v2 := hasher.Sum128()
	hasher.Write(a1)
	v3, v4 := hasher.Sum128()
	return [4]uint64{
		v1, v2, v3, v4,
	}
}

// `location` returns the ith hashed location using the four base hash values
func location(h [4]uint64, i uint) uint64 {
	ii := uint64(i)
	return h[ii%2] + ii*h[2+(((ii+(ii%2))%4)/2)]
}

// `location` returns the ith hashed location using the four base hash values
func (f *BloomFilter) location(h [4]uint64, i uint) uint {
	return uint(location(h, i) % uint64(f.m))
}

// `Test` will find whether an element is a member of a given set. It returns false
// when an element is definitely not in the set, and returns true if it might be
// present; hence it is probabilistic.
func (f *BloomFilter) Test(a []byte) bool {
	h := baseHashes(a)
	for i := uint(0); i < f.k; i++ {
		if !f.b.Test(f.location(h, i)) {
			return false
		}
	}
	return true
}

// `Cap()` gives the capacity of the bloom filter.
func (f *BloomFilter) Cap() uint {
	return f.m
}

// `K()` gives the number of hash functions in the bloom filter.
func (f *BloomFilter) K() uint {
	return f.k
}
