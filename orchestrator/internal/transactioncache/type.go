package transactioncache

import (
	"sync"
)

type (
	Service struct {
		mapTx sync.Map
	}
)
