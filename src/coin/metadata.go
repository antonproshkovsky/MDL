package coin

/*
	Output more detailed record of transactions and outputs in each block
*/

// Transaction metadata
type TransactionSummary struct {
    TransactionId   uint64   // transaction ID, redundant
    TransactionHash SHA256   // outer hash
    Time            uint64   // time of transaction
    BlockId         uint64   // block id, redundant
    Inputs          []uint64 // integer id for inputs, redudant
    OutTxIdx        uint64   // starting index for transaction output, redundant
    OutTxCnt        uint16   // number of outputs, redundant
    Fee             uint64
}

// Transaction output metadata
type TxOutMeta struct {
    Hash  SHA256 // hash of transactiontion hash + the struct
    Coins uint64 // number of coins
    Hours uint64 // coin hours
    Nonce [32]byte

    TxId  uint64 // id of transaction that created block
    Time  uint64 // time created
    Block uint64 // block id that created TxOut
}

// Block header metadata
type BlockMeta struct {
    TxSeq0 uint64 // start, transaction seq
    UxSeq0 uint64 // start, unspent output seq

    TxSeq1 uint64 // end, transaction seq
    UxSeq1 uint64 // end, unspent output seq

    UxXor0 SHA256 // xor of hashes of unspent before block
    UxXor1 SHA256 // xor of hash of unspent after block execution
}
