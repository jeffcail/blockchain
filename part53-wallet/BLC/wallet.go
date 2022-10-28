package BLC

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"github.com/c/public-chain.io/common/utils"

	"golang.org/x/crypto/ripemd160"
)

const (
	version            = byte(0x00)
	addressCheckSumLen = 4
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewWallet() *Wallet {
	private, public := newKeyPair()
	wallet := Wallet{
		PrivateKey: private,
		PublicKey:  public,
	}
	return &wallet
}
func (w Wallet) GetAddress() []byte {
	pubKeyHash := HashPubKey(w.PublicKey)
	versionPayload := append([]byte{version}, pubKeyHash...)
	checksum := checkSum(versionPayload)
	fullPayload := append(versionPayload, checksum...)
	address := utils.Base58Encoding(fullPayload)
	return address
}

func HashPubKey(pubKey []byte) []byte {
	publicSHA256 := sha256.Sum256(pubKey)

	RIPEMD160Hash := ripemd160.New()
	_, err := RIPEMD160Hash.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hash.Sum(nil)
	return publicRIPEMD160
}

func ValidateAddress(address []byte) bool {
	base58Decoding := Base58Decoding(address)
	b := base58Decoding[len(base58Decoding)-addressCheckSumLen:]
	checkSumBytes := base58Decoding[:len(base58Decoding)-addressCheckSumLen]
	checkBytes := checkSum(checkSumBytes)
	if bytes.Compare(b, checkBytes) == 0 {
		return true
	}
	return false
}

func checkSum(payload []byte) []byte {
	firstSHA := sha256.Sum256(payload)
	secondSHA := sha256.Sum256(firstSHA[:])
	return secondSHA[:addressCheckSumLen]
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}
