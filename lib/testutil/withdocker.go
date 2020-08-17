package testutil

import (
	"context"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"sync"
)

func GetDockerEnv(ctx context.Context, wg *sync.WaitGroup) (conf config.Config, err error) {
	_, ip, err := MongoContainer(ctx, wg)
	if err != nil {
		return conf, err
	}

	port, _, err := RyaContainer(ctx, wg, ip)
	if err != nil {
		return conf, err
	}

	conf.RyaUrl = "http://localhost:" + port
	return
}
