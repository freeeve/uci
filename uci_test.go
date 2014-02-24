package uci

import (
	"fmt"
	"testing"
	. "launchpad.net/gocheck"
)

func Test(t *testing.T) { TestingT(t) }

type UCISuite struct{}

var _ = Suite(&UCISuite{})

func (s *UCISuite) TestUCIDepth4(c *C) {
	eng, err := NewEngine("./stockfish")
	c.Assert(err, IsNil)
	eng.SetFEN("rnb4r/ppp1k1pp/3bp3/1N3p2/1P2n3/P3BN2/2P1PPPP/R3KB1R b KQ - 4 11")
	res, err := eng.GoDepth(4)
	c.Assert(err, IsNil)
	c.Assert(res.BestMove, Equals, "c8d7")
	fmt.Println(res)
}

func (s *UCISuite) TestUCIDepth15(c *C) {
	eng, err := NewEngine("./stockfish")
	c.Assert(err, IsNil)
	eng.SetFEN("r1b1k1nr/ppq2pbp/2n1p1p1/1B2pN2/5P2/2N1B3/PPP3PP/R2QK2R w KQkq - 2 11")
	res, err := eng.GoDepth(15)
	c.Assert(err, IsNil)
	c.Assert(res.BestMove, Equals, "b5c6")
	fmt.Println(res)
}
