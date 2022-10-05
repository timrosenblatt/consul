package proxycfgglue

import (
	"context"

	"github.com/hashicorp/consul/agent/cache"
	cachetype "github.com/hashicorp/consul/agent/cache-types"
	"github.com/hashicorp/consul/agent/proxycfg"
)

// CacheHTTPChecks satisifies the proxycfg.HTTPChecks interface by sourcing
// data from the agent cache.
func CacheHTTPChecks(c *cache.Cache) proxycfg.HTTPChecks {
	return &cacheProxyDataSource[*cachetype.ServiceHTTPChecksRequest]{c, cachetype.ServiceHTTPChecksName}
}

func ServerHTTPChecks(deps ServerDataSourceDeps) proxycfg.HTTPChecks {
	return serverHTTPChecks{deps}
}

type serverHTTPChecks struct {
	deps ServerDataSourceDeps
}

func (c serverHTTPChecks) Notify(ctx context.Context, req *cachetype.ServiceHTTPChecksRequest, correlationID string, ch chan<- proxycfg.UpdateEvent) error {
	c.deps.Logger.Debug("no-op")
	return nil
}
