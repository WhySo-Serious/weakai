package rnntest

import (
	"testing"

	"github.com/unixpickle/weakai/neuralnet"
	"github.com/unixpickle/weakai/rnn"
)

func TestBidirectional(t *testing.T) {
	outNet := neuralnet.Network{
		&neuralnet.DenseLayer{
			InputCount:  4,
			OutputCount: 2,
		},
		&neuralnet.Sigmoid{},
	}
	outNet.Randomize()
	bd := &rnn.Bidirectional{
		Forward: &rnn.RNNSeqFunc{
			Block: rnn.NewLSTM(3, 2),
		},
		Backward: &rnn.RNNSeqFunc{
			Block: rnn.NewLSTM(3, 2),
		},
		Output: &rnn.NetworkSeqFunc{
			Network: outNet,
		},
	}
	test := SeqFuncTest{
		S:        bd,
		Params:   bd.Parameters(),
		TestSeqs: rnnSeqFuncTests,
	}
	test.Run(t)
}
