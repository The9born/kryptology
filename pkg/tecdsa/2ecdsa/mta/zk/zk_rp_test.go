package zk

import (
	"math/big"
	"testing"

	tt "github.com/coinbase/kryptology/internal"
	"github.com/coinbase/kryptology/pkg/paillier"
	"github.com/stretchr/testify/require"
)

func makeNewPaillierPublicKey(t *testing.T, n *big.Int) *paillier.PublicKey {
	t.Helper()
	publicKey, err := paillier.NewPubkey(n)
	require.NoError(t, err)
	return publicKey
}

func TestZKRP(t *testing.T) {
	q := tt.B10("138339111123263837359132451917124015517170754257582557304041157577680753093647565539079571019759527352525282230182764675776568911651258485531289666832959634563223004878122780813127971174409241627423824739817920620625033514978109452240689980096834418477667585185553068234098818928032774937863800840934496419899")
	p := tt.B10("165498465971525536497859961269214938631289964308823560526920537236787050377699904896554622379770774622567664583533323254169290844053351296829514419428489585830394868303448384771151376037064711115810339324861594209655768995895643373763166292366557525131878080032169065959558884224551806641003919879441772258023")

	sk, _ := paillier.NewSecretKey(p, q)
	pk := makeNewPaillierPublicKey(t, sk.N)

	st := &ModStatement{}
	st.N = pk.N
	ws := &ModWitness{}
	ws.p = p
	ws.q = q

	proof := rpprove(st, ws)
	res := rpverify(st, proof)

	require.Equal(t, true, res)
}
