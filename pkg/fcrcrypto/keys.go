package fcrcrypto

/*
 * Copyright 2020 ConsenSys Software Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */


import (
	"crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha512"
    "crypto/x509"
    "encoding/pem"
    "hash"
    "fmt"
)

// GenKeyPairV1 generates a key pair
func GenKeyPairV1() (*ecdsa.PrivateKey, error) {
    return GenKeyPair(elliptic.P256())
}

// GenKeyPair generates a key pair
func GenKeyPair(curve elliptic.Curve) (*ecdsa.PrivateKey, error) {
    return ecdsa.GenerateKey(curve, rand.Reader)
}

// PEMEncodePrivateKey converts a private key to a string
func PEMEncodePrivateKey(privateKey *ecdsa.PrivateKey) string {
    x509Encoded, err := x509.MarshalECPrivateKey(privateKey)
    if err != nil {
        panic(err)
    }
    pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})
    return string(pemEncoded)
}


// PEMEncodePublicKey converts a public key to a string
func PEMEncodePublicKey(publicKey *ecdsa.PublicKey) string {
    x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
    pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
    return string(pemEncodedPub)
}


// PEMDecodePrivateKey converts a string to a private key
func PEMDecodePrivateKey(pemEncoded string) *ecdsa.PrivateKey {
    block, _ := pem.Decode([]byte(pemEncoded))
    x509Encoded := block.Bytes
    privateKey, _ := x509.ParseECPrivateKey(x509Encoded)
    return privateKey
}

// PEMDecodePublicKey converts a string to a public key
func PEMDecodePublicKey(pemEncodedPub string) *ecdsa.PublicKey {
    blockPub, _ := pem.Decode([]byte(pemEncodedPub))
    x509EncodedPub := blockPub.Bytes
    genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
    publicKey := genericPublicKey.(*ecdsa.PublicKey)
    return publicKey
}

// HashPublicKeyV1 is the initial version of public key to hash algorithm.
// New algorithms may be used with later versions of the protocol
func HashPublicKeyV1(publicKey *ecdsa.PublicKey) []byte {
    return HashPublicKey(publicKey, sha512.New512_256())
}


// HashPublicKey generates a message digest that matches the public key.
func HashPublicKey(publicKey *ecdsa.PublicKey, hasher hash.Hash) []byte {
    xBytes := publicKey.X.Bytes()
    yBytes := publicKey.Y.Bytes()


    if _, err := hasher.Write(xBytes); err != nil {
        // The message digest algorithm should never cause an error.
		panic(fmt.Sprintf("Message digest algorithm is unable to process hashes: %v", err))
    }
    if _, err := hasher.Write(yBytes); err != nil {
        // The message digest algorithm should never cause an error.
		panic(fmt.Sprintf("Message digest algorithm is unable to process hashes: %v", err))
    }
    return hasher.Sum(nil)
}