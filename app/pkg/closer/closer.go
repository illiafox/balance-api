package closer

import (
	"fmt"
	"io"
	"sync"

	"balance-service/app/pkg/logger"
	"go.uber.org/zap"
)

type Closers struct {
	mutex   sync.Mutex
	closers []io.Closer
}

func (c *Closers) Add(closer io.Closer) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.closers = append(c.closers, closer)
}

func (c *Closers) Close(logger logger.Logger) {
	if c.closers == nil {
		panic("closers are empty")
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i := len(c.closers) - 1; i >= 0; i-- {
		closer := c.closers[i]
		//
		if err := closer.Close(); err != nil {
			logger.Error(fmt.Sprintf("close %T", closer), zap.Error(err))
		}
	}

	c.closers = nil
}
