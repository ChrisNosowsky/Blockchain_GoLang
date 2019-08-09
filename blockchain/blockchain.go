package blockchain


import (
  "github.com/dgraph-io/badger"
)

const (
  dbPath = "./tmp/blocks"
)

type BlockChain struct {
  LastHash []byte //Last hash of the last block in the chain
  Database *badger.DB //Pointer ot our badger DB. Lets create a path to that
}



func InitBlockChain() *BlockChain {
  var lastHash []byte

  opts := badger.DefaultOptions
  opts.Dir = dbPath //Where the database will store our keys and metadata
  opts.ValueDir = dbPath //Store all the values

  db, err := badger.Open(opts) //tuple with a pointer to the database that opens the database
  Handle(err)

  err := db.Update(func(txn *badger.Txn) error { //accessing the database. Allows us to do read and write. View = read-only
    //check if blockchain has already been stored. If it has, create a new blockchain instance in memory
    //get last hash and push it into this instance in memory
    //if no blockchain, store genesis and store it as a last block hash
    if _, err:= txn.Get([]byte("lh")); err == badger.ErrKeyNotFound { //Means no transactions or keys found in db so we don't have a blockchain
      fmt.Println("No existing blockchain found")
      genesis := Genesis()
      fmt.Println("Genesis proved")
      err = txn.Set(genesis.Hash, genesis.Serialize())

      err = txn.Set([]byte("lh"), genesis.Hash)

      lastHash = genesis.Hash //put it into our memory storage
      return err
    } else { //if we already have a db and a blockchain
      item, err := txn.Get([]byte("lh")) //Error 1. Get last hash and handle the error.
      Handle(err)
      lastHash, err = item.Value() //get item struct
      return err
    }
  })
  Handle(err)

  blockchain := BlockChain{lastHash, db}
  return &blockchain
}



func (chain *BlockChain) AddBlock(data string) {
  var lastHash []byte

  err := chain.Database.View(func(txn *badger.Txn) error {
    item, err := txn.Get([]byte("lh"))
    Handle(err)
    lastHash, err = item.Value()

    return err
  })
  Handle(err)

  newBlock := CreateBlock(data, lastHash)
  err = chain.Database.Update(func(txn *badger.Txn) error {
    err := txn.Set(newBlock.Hash, newBlock.Serialize())
    Handle(err)
    err = txn.Set([]byte("lh"), newBlock.Hash)
  })
}
