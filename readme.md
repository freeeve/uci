# uci 

[![Build Status](https://travis-ci.org/wfreeman/uci.png?branch=master)](https://travis-ci.org/wfreeman/uci)
[![Coverage Status](https://coveralls.io/repos/wfreeman/uci/badge.png)](https://coveralls.io/r/wfreeman/uci)

A golang API to interact with UCI chess engines. (should be considered experimental for the time being) [A description of how UCI works is available here.](http://wbec-ridderkerk.nl/html/UCIProtocol.html)

Many chess engines support UCI (Universal Chess Interface). This library is designed for use with Stockfish, but should work with other UCI engines.

## minimum viable snippet
```Go
package main

import (
	"fmt"
	"log"
	"github.com/wfreeman/uci"
)

func main() {
	eng, err := uci.NewEngine("/path/to/stockfish")
	if err != nil {
		log.Fatal(err)
	}
	
	eng.SetOptions(uci.Options{
		Hash:128,
		Ponder:false,
		OwnBook:true,
		MultiPV:4,
	})
	eng.SetFEN("rnb4r/ppp1k1pp/3bp3/1N3p2/1P2n3/P3BN2/2P1PPPP/R3KB1R b KQ - 4 11")
	results := eng.GoDepth(10)
	fmt.Println(results)
}
```

produces this output:

```

```
