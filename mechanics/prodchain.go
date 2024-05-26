package mechanics

import (
	_ "embed"
	"encoding/json"
	"errors"
	"sync"

	"github.com/ruudiRatlos/s10s/openapi"
)

type prodChain struct {
	Export  string   `json:"export"`
	Imports []string `json:"imports"`
}

//go:embed production-chain.json
var chainRaw []byte
var chain map[openapi.TradeSymbol][]openapi.TradeSymbol
var chainOnce sync.Once
var chainErr error

func LoadChain() (map[openapi.TradeSymbol][]openapi.TradeSymbol, error) {
	chainOnce.Do(func() {
		list := []prodChain{}
		chainErr = json.Unmarshal(chainRaw, &list)
		if chainErr != nil {
			return
		}
		chain = map[openapi.TradeSymbol][]openapi.TradeSymbol{}
		for _, pc := range list {
			sym, chainErr := openapi.NewTradeSymbolFromValue(pc.Export)
			if chainErr != nil {
				return
			}
			chain[*sym] = []openapi.TradeSymbol{}
			for _, imp := range pc.Imports {
				impS, chainErr := openapi.NewTradeSymbolFromValue(imp)
				if chainErr != nil {
					return
				}
				chain[*sym] = append(chain[*sym], *impS)
			}
		}
	})
	return chain, chainErr
}

var ErrTradeSymbolNotFound error = errors.New("trade symbol not found")

func FindMaterials(good openapi.TradeSymbol) ([]openapi.TradeSymbol, error) {
	chain, err := LoadChain()
	if err != nil {
		return nil, err
	}
	r, e := chain[good]
	if !e {
		return nil, ErrTradeSymbolNotFound
	}

	return r, nil
}
