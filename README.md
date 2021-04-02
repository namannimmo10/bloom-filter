A Bloom filter is a space-efficient probabilistic data structure that is used to check whether an element is a member of a set.
False positive matches are possible, but false negatives are not, therefore, each query may either return "possibly present" in the set or
"_definitely_ not in the set". Elements can be added to this set, but cannot be removed. In this implementation, `murmurhash` is used, which is a non-cryptographic hashing function.

Test coverage - 100%
```
$ go test -cover
PASS
coverage: 100.0% of statements
ok      github.com/namannimmo/bloomfilter       0.002s
```
