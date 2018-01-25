package main

import (
	"flag"
	"os"

	cmn "github.com/tendermint/tmlibs/common"
	"github.com/tendermint/tmlibs/log"

	priv_val "github.com/tendermint/tendermint/types/priv_validator"
)

func main() {
	var (
		chainID     = flag.String("chain-id", "mychain", "chain id")
		privValPath = flag.String("priv", "", "priv val file path")
		socketAddr  = flag.String("socket.addr", ":46659", "socket bind addr")

		logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "priv_val")
	)
	flag.Parse()

	logger.Info("args", "chainID", *chainID, "privPath", *privValPath)

	privVal := priv_val.LoadPrivValidatorJSON(*privValPath)

	pvss := priv_val.NewPrivValidatorSocketServer(
		logger,
		*socketAddr,
		*chainID,
		privVal,
	)
	// pvss.Start()

	cmn.TrapSignal(func() {
		pvss.Stop()
	})
}
