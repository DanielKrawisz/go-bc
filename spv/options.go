package spv

import "github.com/libsv/go-bc"

// ClientOpts can be implemented to provided functional options for an spv.Client.
type ClientOpts func(*spvclient)

// WithBlockHeaderChain will inject the provided BlockHeaderChain into the spv.Client
func WithBlockHeaderChain(bhc bc.BlockHeaderChain) ClientOpts {
	return func(s *spvclient) {
		s.bhc = bhc
	}
}

// WithTxStore will inject the provided TxGetter into the spv.Client
func WithTxStore(txc TxStore) ClientOpts {
	return func(s *spvclient) {
		s.txc = txc
	}
}

// WithMerkleProofStore will inject the provided MerkleProofGetter into the spv.Client
func WithMerkleProofStore(mpc MerkleProofStore) ClientOpts {
	return func(s *spvclient) {
		s.mpc = mpc
	}
}
