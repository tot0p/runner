package runner

import "github.com/docker/docker/client"

var cli *client.Client

// Init docker client
func Init() {
	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
}

func Close() error {
	return cli.Close()
}
