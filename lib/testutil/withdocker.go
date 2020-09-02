package testutil

import (
	"context"
	"github.com/SENERGY-Platform/semantic-repository/lib/config"
	"net"
	"sync"
)

func GetDockerEnv(ctx context.Context, wg *sync.WaitGroup, conf *config.Config) (err error) {
	_, ip, err := MongoContainer(ctx, wg)
	if err != nil {
		return err
	}

	port, _, err := RyaContainer(ctx, wg, ip)
	if err != nil {
		return err
	}

	conf.RyaUrl = "http://localhost:" + port
	return
}

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}
