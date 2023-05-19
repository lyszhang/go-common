/**
 * @Author: lyszhang
 * @Email: ericlyszhang@gmail.com
 * @Description:
 * @File:
 * @Version: 1.0.0
 * @Data:
 */

package main

import (
	log "github.com/inconshreveable/log15"
	"github.com/lyszhang/go-lib/logmgr"
	"os"
)

func main() {
	// flexible configuration
	log.Root().SetHandler(log.MultiHandler(
		log.StreamHandler(os.Stderr, log.LogfmtFormat()),
		log.LvlFilterHandler(
			log.LvlDebug,
			log.StreamHandler(logmgr.NewRotate(".", "test.log"), log.JsonFormat()))))

	log.Info("Hello")
	srvlog := log.New("module", "app/server")

	// all log messages can have key/value context
	srvlog.Warn("abnormal conn rate", "rate", 100, "low", 20, "high", 30)

	log.Info("world")
}
