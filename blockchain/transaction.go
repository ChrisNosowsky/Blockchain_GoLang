package blockchain


type Transaction struct {
  ID []byte
  Inputs []TxInput
  Outputs []TxOutput
}

type TxOutput struct { //These outputs are indivisible. Being referenced by the input
  Value int //Value in tokens
  PubKey string //Value needed to unlock the tokens which are inside the value field
}

type TxInput struct {
  ID []byte //Transaction the output is inside out
  Out int //Index where the transaction is in
  Sig string //users account
}

func (tx *Transaction) SetID() {
  var encoded bytes.Buffer
  var hash [32]byte

  encode := gob.NewEncoder(&encoded)
  err := encode.Encode(tx)
  Handle(err)

  hash = sha256.Sum256(encoded.Bytes())
  tx.ID = hash[:]
}

func CoinbaseTx(to, data string) *Transaction {
  if data == "" {
    data = fmt.Sprintf("Coins to %s", to)
  }
  txin := TxInput{[]byte{}, -1, data}
  txout := TxOutput{100, to}

  tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
  tx.SetID()

  return &tx
}

func (tx *Transaction) IsCoinbase() bool {
  return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs
}
