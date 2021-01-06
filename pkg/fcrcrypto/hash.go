package fcrcrypto

import (
	"hash"
	"golang.org/x/crypto/blake2b"
)

// GetBlockchainHasher returns a message digest implementation that hashes according to the 
// algorithms used by the Filecoin blockchain.
func GetBlockchainHasher() (hash.Hash, error) {
	return blake2b.New256(nil)
}


// BlockchainHash message digests some data using the algorithm used by the Filecoin blockchain.
func BlockchainHash(data []byte) []byte {
	hash := blake2b.Sum256(data)
	return hash[:]
}


// GetRetrievalV1Hasher returns a message digest implementation that hashes according to the 
// algorithms used by version one of the Filecoin retrieval protocol.
func GetRetrievalV1Hasher() (hash.Hash, error) {
	return blake2b.New256(nil)
}

// RetrievalV1Hash message digests some data using the algorithm used by version one of the 
// Filecoin retrieval protocol.
func RetrievalV1Hash(data []byte) []byte {
	hash := blake2b.Sum256(data)
	return hash[:]
}
