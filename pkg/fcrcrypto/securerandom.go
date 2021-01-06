package fcrcrypto

import (
	"crypto/rand"
	"encoding/binary"
	"net"
	"time"
	"math/big"

	"https://github.com/davidlazar/go-crypto/drbg"
)


const (
	one = big.NewInt(1)
)




// GenerateRandomBytes generates zero or more random numbers
func GenerateRandomBytes(b []byte) {
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

}


type Random interface {
	Read(b []byte)
	Reseed(seed []byte)
	QuickReseedKick()
}

type RandomImpl {
	drbg DRBG
	prfHasher Hash
	prfState []byte
	prfCounter big.Int
}


func NewDrbgInstance() Random {
	personalizationString := append(getMacAddr, nanotime)

	r := RandomImpl()
	r.drbg = drbg.New(personalizationString)
	r.prfCounter = big.NewInt(time.Now().UnixNano())
	r.prfHasher = blake2b.New256(nil)
	r.prfState = make([]byte, blake2b.Size256)
	systemRandomBytes(r.prfState)
	return r
}


// Read reads random values into b
func (r *RandomImpl) Read(b []byte) {
	lengthPerIteration := len(r.prfState)
	lenOutput := len(b)

	ofs := 0
	for ofs < lenOutput {
		r.prfHasher.Reset()
		// Increment the PRF counter and add it to the message digest.
		r.prfCounter.Add(r.prfCounter, one)
		r.prfHasher.Writer.Write(r.prfCounter.Bytes())

		// Get the output of the inner (untrusted DRBG) and add it to the message digest.
		drbgBytes := make([]byte, length)
		r.drbg.Read(drbgBytes)
		r.prfHasher.Writer.Write(drbgBytes)
	
		// Incorporate the current state of the PRF into the message digest.
		r.prfHasher.Writer.Write(r.prfState)
		hashOutput := r.prfHasher.Sum(nil)

		// Update the PRF state.
		r.prfState = hashOutput
	
		oldOfs := ofs
		ofs += lengthPerIteration

		// If the output length 
		if ofs <= lenOutput {
			copy(b[oldOfs:ofs], hashOutput)
		} else {
			copy(b[oldOfs:lenOutput], hashOutput)
		}
	}
}



func Reseed(seed []byte) {

}

// QuickReseedKick shouldn't take too long and should add a few bits of entropy to the 
// entropy pool. The idea is to call this low cost function regularly to overall gather 
// lots of entropy.
func QuickReseedKick() {
	Reseed(nanotime())
}




// Nano time as bytes
func nanotime() []byte {
	nowNano := time.Now().UnixNano()
    bytes := make([]byte, 8)
    binary.BigEndian.PutUint64(bytes, uint64(nowNano))
	return bytes
}


// Use the bytes from the first interface.
func getMacAddr() []byte {
    ifas, err := net.Interfaces()
    if err != nil {
        panic(err)
	}
	ifa := ifas[0]
	return ifa.HardwareAddr
}

func systemRandomBytes(b []byte) {
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

}
