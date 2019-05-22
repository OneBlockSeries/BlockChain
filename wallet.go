package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	//"bytes"
	"golang.org/x/crypto/ripemd160"
	"log"
)

const version = byte(0x01)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubKey
}
func CreateWallet() *Wallet {
	pri, pub := NewPair()
	wallet := Wallet{pri, pub}
	return &wallet
}
func (w Wallet) CreateAddress() []byte {
	//地址=版本号+两次公钥哈希 PIPEMD160(SHA256(PubKey))+checksum(两次哈希公钥哈希结果sha256(sha256（pubkeyhash）)
	//比特币地址=base58Encode(地址)

	pubkeyhash := w.PubkeyTwicehash()
	temp := append([]byte{version}, pubkeyhash...)

	checksumtemp1 := sha256.Sum256(temp)
	checksumtemp2 := sha256.Sum256(checksumtemp1[:])
	checksum := checksumtemp2[:4] //取前四个字节
	address := append(temp, checksum...)
	bitaddress := Base58Encode(address)
	return bitaddress
}
func (w Wallet) PubkeyTwicehash() []byte {

	publicSHA256 := sha256.Sum256(w.PublicKey)

	RIPEMD160Hasher := ripemd160.New()
	_, err := RIPEMD160Hasher.Write(publicSHA256[:])
	if err != nil {
		log.Panic(err)
	}
	publicRIPEMD160 := RIPEMD160Hasher.Sum(nil)

	return publicRIPEMD160
}
