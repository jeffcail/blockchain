package BLC

import "fmt"

type Wallets struct {
	Wallets map[string]*Wallet
}

// NewWallets
func NewWallets() *Wallets {
	wallets := &Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	return wallets
}

// CreateNewWallet
func (w *Wallets) CreateNewWallet() {
	wallet := NewWallet()
	fmt.Printf("address: %s\n", wallet.GetAddress())
	w.Wallets[string(wallet.GetAddress())] = wallet
}
