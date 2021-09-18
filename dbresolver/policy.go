package dbresolver

import (
	"math/rand"

	"github.com/Hive-Intelligence/gorm"
)

type Policy interface {
	Resolve([]gorm.ConnPool) gorm.ConnPool
}

type RandomPolicy struct {
}

func (RandomPolicy) Resolve(connPools []gorm.ConnPool) gorm.ConnPool {
	return connPools[rand.Intn(len(connPools))]
}
