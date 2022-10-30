package BLC

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

const walletFile = "Wallets.dat"

type Wallets struct {+

	WalletsMap map[string]*Wallet
}

// NewWallets
func NewWallets() (*Wallets, error) {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		wallets := &Wallets{}
		wallets.WalletsMap = make(map[string]*Wallet)
		return wallets, err
	}

	fileContent, err := os.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	var wallets *Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}

	return wallets, nil
}

// CreateNewWallet
func (w *Wallets) CreateNewWallet() {
	wallet := NewWallet()
	fmt.Printf("address: %s\n", wallet.GetAddress())
	w.WalletsMap[string(wallet.GetAddress())] = wallet
	w.SaveWallets()
}

// 保存钱包数据到本地
func (w *Wallets) SaveWallets() {
	// 1.对Wallets加密
	var buffer bytes.Buffer
	// 注册接口类型, 否则会出现“gob: type not registered for interface: elliptic.p256Curve”
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(&w)
	if err != nil {
		log.Panic(err)
	}
	// 2.把加密后数据写入文件中
	err = os.WriteFile(walletFile, buffer.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
