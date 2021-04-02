package bloomfilter

import "testing"

func TestBloom(t *testing.T) {
	f := NewBloom(100, 3)
	s1 := []byte("abc")
	s2 := []byte("def")
	s3 := []byte("ghi")
	s4 := []byte("jkl")

	f.Add(s1)
	f.Add(s3)

	// `s1` and `s3` may be there in the set.
	// `s2` and `s4` _definitely_ not in the set.
	checkS1 := f.Test(s1)
	checkS2 := f.Test(s2)
	checkS3 := f.Test(s3)
	checkS4 := f.Test(s4)

	if !checkS1 {
		t.Errorf("%v should be in the bloom filter.", s1)
	}
	if checkS2 {
		t.Errorf("%v should not be in the bloom filter.", s2)
	}
	if !checkS3 {
		t.Errorf("%v should be in the bloom filter.", s3)
	}
	if checkS4 {
		t.Errorf("%v should not be in the bloom filter.", s4)
	}
}

func TestCap(t *testing.T) {
	f := NewBloom(10000, 10)
	if f.Cap() != f.m {
		t.Errorf("not calculating the capacity of the bloom filter properly.")
	}
}

func TestK(t *testing.T) {
	f := NewBloom(10000, 10)
	if f.K() != f.k {
		t.Errorf("not calculating the no. of hash functions properly.")
	}
}

func TestMax(t *testing.T) {
	a, b := uint(3), uint(1)
	if max(a, b) != a {
		t.Errorf("max(%d, %d) should be %d.", a, b, a)
	}
}
