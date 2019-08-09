package blockchain

import (
  "bytes"
  "encoding/gob"
)


type Block struct {
  Hash []byte
  Data []byte
  PrevHash []byte
  Nonce int
}

func CreateBlock(data string, prevHash []byte) *Block {
  block := &Block{[]byte{}, []byte(data), prevHash, 0}
  pow := NewProof(block)
  nonce, hash := pow.Run()

  block.Hash = hash[:]
  block.Nonce = nonce

  return block
}

func Genesis() *Block {
  return CreateBlock("Genesis", []byte{})
}

//BadgerDB Only accepts arrays of bytes and slices of bytes so we must create a serialized and deserialized functions
func (b *Block) Serialize() []byte {
  var res bytes.Buffer
  encoder := gob.NewEncoder(&res)

  err := encoder.Encode(b) //error handeling
  Handle(err)
  return res.Bytes() //Byte representation of our block
}

func DeSerialize(data []byte) *Block {
  var block Block

  decoder := gob.NewDecoder(bytes.NewReader(data))
  err = decoder.Decode(&block)
  Handle(err)
  return &block
}

func Handle(err error) {
  if err != nil {
    log.Panic(err)
  }
}
