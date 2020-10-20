package token

import (
	"github.com/sero-cash/go-sero/common"
	"github.com/sero-cash/go-sero/seroclient"
	"log"
	"math/big"
	"os"
	"testing"
)

const address = "eX4rXhV94QMsMftenDiiJQCKckHvpvoLA8MoRtFXhmeMaLRb7ggM1UjG49BLviKHU1mjnTjQgJTLLc2dQQuiY1J"

func TestTokenCaller_GetBalance(t *testing.T) {
	url := os.Getenv("URL")
	//keyStr:=os.Getenv("KEYSTORE")
	//pass:=os.Getenv("PASS")
	conn, err := seroclient.Dial(url) // like "http://127.0.0.1:8545"
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	//auth, err := bind.NewTransactor(strings.NewReader(keyStr), pass, nil)
	//if err != nil {
	//	log.Fatalf("Failed to create authorized transactor: %v", err)
	//}
	addr := common.Base58ToAddress(address)
	token, err := NewToken(addr, conn)
	if err != nil {
		t.Fatal(err)
	}
	var balance struct {
		TokenList [][32]byte
		Balances  []*big.Int
	}
	balance, err = token.GetBalance(nil)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(balance.Balances); i++ {
		//t.Logf("token: %s, balance: %s", common.BytesToAddress(balance.TokenList[i]...), balance.Balances[i].String())
		t.Logf("token[%d] %s, balance: %s", i, string(balance.TokenList[i][:]), balance.Balances[i].String())
	}
}
