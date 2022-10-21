package BLC

type TXOutPut struct {
	Value        int64
	ScriptPubKey string
}

// UnLockScriptPubKeyWithAddress
func (txOutput *TXOutPut) UnLockScriptPubKeyWithAddress(address string) bool {
	return txOutput.ScriptPubKey == address
}
