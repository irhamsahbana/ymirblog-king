// Package adapters are the glue between components and external sources.
// # This manifest was generated by ymir. DO NOT EDIT.
package adapters

import (
	"fmt"
	"net/http"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/go-resty/resty/v2"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type client interface {
	*sql.Driver | *http.Client | *resty.Client | *redis.Client | *mongo.Client
}

// Driver - interface adapter.
type Driver[T client] interface {
	Open() (T, error)
	Connect() error
	Disconnect() error
}

// Adapter components for external sources.
type Adapter struct {
}

// Option is Adapter type return func.
type Option func(adapter *Adapter)

// Sync - configure all adapters.
func (a *Adapter) Sync(opts ...Option) {
	for o := range opts {
		opt := opts[o]
		opt(a)
	}
}

// UnSync - release all adapter connection.
func (a *Adapter) UnSync() error {
	var errs []string
	if len(errs) > 0 {
		err := fmt.Errorf(strings.Join(errs, "\n"))
		log.Error().Err(err).Msg("UnSync adapter error")
		return err
	}
	return nil
}
