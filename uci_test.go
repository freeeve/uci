package uci

import (
	"testing"
	. "launchpad.net/gocheck"
)

func Test(t *testing.T) { TestingT(t) }

type UCISuite struct{}

var _ = Suite(&UCISuite{})

func (s *UCISuite) TestUCIDepth4(c *C) {
	eng, err := NewEngine("/Applications/stockfish-dd-64-modern")
	c.Assert(err, IsNil)
	eng.SetFEN("rnb4r/ppp1k1pp/3bp3/1N3p2/1P2n3/P3BN2/2P1PPPP/R3KB1R b KQ - 4 11")
	res, err := eng.GoDepth(4)
	c.Assert(err, IsNil)
	c.Assert(res.BestMove, Equals, "c8d7")
}
