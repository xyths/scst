package swap

import (
	"github.com/sero-cash/go-sero/common"
	"github.com/sero-cash/go-sero/seroclient"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"testing"
)

const address = "5bYB1DSe18ad7oFiJSnt861ett8fT5GFSva3BBrdYvyU8SK6CJoyDrMyijZWiGmkT1ZDaQt2AMZr5yxu23n62YnQ"

func getSwapExchange(t *testing.T) *SwapExchange {
	url := os.Getenv("URL")
	conn, err := seroclient.Dial(url) // like "http://127.0.0.1:8545"
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	addr := common.Base58ToAddress(address)
	ex, err := NewSwapExchange(addr, conn)
	if err != nil {
		t.Fatal(err)
	}
	return ex
}
func TestSwapExchangeCaller_GetTokens(t *testing.T) {
	ex := getSwapExchange(t)
	var byte32t [32]byte
	token := os.Getenv("TOKEN")
	copy(byte32t[:], token)
	tokens, err := ex.GetTokens(nil, byte32t)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Base Token %s", token)
	for i := 0; i < len(tokens); i++ {
		t.Logf("\ttoken[%d] %s", i, string(tokens[i][:]))
	}
}

// there is bugs, no more than one
func TestSwapExchangeCaller_PairList(t *testing.T) {
	ex := getSwapExchange(t)
	start_ := os.Getenv("START")
	end_ := os.Getenv("END")
	start, ok := big.NewInt(0).SetString(start_, 0)
	if !ok {
		t.Logf("bad start")
	}
	end, ok := big.NewInt(0).SetString(end_, 0)
	if !ok {
		t.Logf("bad end")
	}
	pairs, err := ex.PairList(nil, start, end)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(pairs); i++ {
		p := pairs[i]
		t.Logf("pair[%d] %s-%s, Reserve: %s %s, TotalShares: %s, MyShare: %s, ShareReward: %s, Mining: %v",
			i, string(p.TokenA[:]), string(p.TokenB[:]),
			p.ReserveA, p.ReserveB, p.TotalShares,
			p.MyShare, p.ShareRreward, p.Mining,
		)
	}
}

var tests = []struct {
	A, B string
}{
	{"SERO", "USDT"}, // true
	{"sero", "USDT"}, // false
	{"USDT", "SERO"}, // true
	{"PFID", "USDT"},
}

func TestSwapExchangeCaller_HashKey(t *testing.T) {
	ex := getSwapExchange(t)
	for _, tt := range tests {
		var a, b [32]byte
		copy(a[:], []byte(tt.A))
		copy(b[:], []byte(tt.B))
		key, err := ex.HashKey(nil, a, b)
		if err != nil {
			t.Fail()
		}
		t.Logf("%s-%s: %s", tt.A, tt.B, string(key[:]))
	}
}

func TestSwapExchangeCaller_HasPair(t *testing.T) {
	ex := getSwapExchange(t)
	for _, tt := range tests {
		var a, b [32]byte
		copy(a[:], []byte(tt.A))
		copy(b[:], []byte(tt.B))
		key, err := ex.HashKey(nil, a, b)
		if err != nil {
			t.Fail()
		}
		has, err := ex.HasPair(nil, key)
		t.Logf("%s-%s: key %s, hasPair: %v", tt.A, tt.B, string(key[:]), has)
	}
}

func TestSwapExchangeCaller_PairInfo(t *testing.T) {
	ex := getSwapExchange(t)
	for _, tt := range tests {
		var a, b [32]byte
		copy(a[:], []byte(tt.A))
		copy(b[:], []byte(tt.B))
		key, err := ex.HashKey(nil, a, b)
		if err != nil {
			t.Fail()
		}
		has, err := ex.HasPair(nil, key)
		if has {
			p, err1 := ex.PairInfo(nil, key)
			if err1 != nil {
				t.Fatal(err1)
			}

			t.Logf("%s-%s A[%s]-B[%s], Reserve: %s %s, TotalShares: %s, MyShare: %s, ShareReward: %s, Mining: %v",
				tt.A, tt.B, string(p.TokenA[:]), string(p.TokenB[:]),
				p.ReserveA, p.ReserveB, p.TotalShares,
				p.MyShare, p.ShareRreward, p.Mining,
			)
		} else {
			t.Logf("no %s-%s pair", tt.A, tt.B)
		}

	}
}

func TestSwapExchangeCaller_EstimateSwap(t *testing.T) {
	ex := getSwapExchange(t)
	as := strings.ToUpper(os.Getenv("A"))
	bs := strings.ToUpper(os.Getenv("B"))
	ns := os.Getenv("N")
	n1, err := strconv.Atoi(ns)
	if err != nil {
		t.Fatal(err)
	}
	n := big.NewInt(int64(n1 * 1e18))
	var a, b [32]byte
	copy(a[:], []byte(as))
	copy(b[:], []byte(bs))
	key, err := ex.HashKey(nil, a, b)
	if err != nil {
		t.Fatal(err)
	}
	has, err := ex.HasPair(nil, key)
	if err != nil {
		t.Fatal(err)
	}
	if !has {
		t.Logf("no %s-%s pair", as, bs)
		return
	}
	amount_, err := ex.EstimateSwap(nil, key, a, n)
	amount := float64(amount_.Int64()) / 1e18
	t.Logf("%s-%s %s %s = %f %s", as, bs, ns, as, amount, bs)
}

func TestSwapExchangeCaller_EstimateSwapBuy(t *testing.T) {
	ex := getSwapExchange(t)
	as := strings.ToUpper(os.Getenv("A"))
	bs := strings.ToUpper(os.Getenv("B"))
	ns := os.Getenv("N")
	n1, err := strconv.Atoi(ns)
	if err != nil {
		t.Fatal(err)
	}
	n := big.NewInt(int64(n1 * 1e18))
	var a, b [32]byte
	copy(a[:], []byte(as))
	copy(b[:], []byte(bs))
	key, err := ex.HashKey(nil, a, b)
	if err != nil {
		t.Fatal(err)
	}
	has, err := ex.HasPair(nil, key)
	if err != nil {
		t.Fatal(err)
	}
	if !has {
		t.Logf("no %s-%s pair", as, bs)
		return
	}
	amount_, err := ex.EstimateSwapBuy(nil, key, a, n)
	amount := float64(amount_.Int64()) / 1e18
	t.Logf("%s-%s %s %s = %f %s", as, bs, ns, as, amount, bs)
}
