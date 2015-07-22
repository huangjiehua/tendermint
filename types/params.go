package types

import (
	"fmt"
	"github.com/tendermint/tendermint/merkle"
)

type ParamsKey uint

const (
	ParamsKeyNewAccountTx ParamsKey = iota

	NumParams = 1 // NOTE: adjust this too
)

func (pk ParamsKey) String() string {
	switch pk {
	case ParamsKeyNewAccountTx:
		return "NewAccountTx"
	}
	return ""
}

// The number of params should be kept
// small and updated infrequently
// Each param should have a permission for modifying it
type Params struct {
	NewAccountTx *NewAccountTxParams `json:"new_account_tx"`
}

type NewAccountTxParams struct {
	PoWTarget []byte `json:"pow_target"`
	Balance   int64  `json:"balance"`
}

func (p *Params) Hash() []byte {
	pkvs := make([]interface{}, NumParams)
	//
	pkvs[ParamsKeyNewAccountTx] = p.NewAccountTx
	//
	return merkle.SimpleHashFromBinaries(pkvs)
}

// All relevant elements must be explicitly copied
func (p *Params) Copy() *Params {
	p2 := new(Params)
	if p2.NewAccountTx != nil {
		copy(p2.NewAccountTx.PoWTarget, p.NewAccountTx.PoWTarget)
		p2.NewAccountTx.Balance = p.NewAccountTx.Balance
	}
	return p2
}

func (p *Params) Get(key ParamsKey) (interface{}, error) {
	switch key {
	case ParamsKeyNewAccountTx:
		return p.NewAccountTx, nil
	default:
		return nil, fmt.Errorf("Invalid params key %d", key)
	}
}
