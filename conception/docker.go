package conception

import (
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"golang.org/x/net/context"
)

type Docker struct {
	client *client.Client
}

func NewDockerConceiver(host string) (Conceiver, error) {
	if host == "" {
		host = "unix:///var/run/docker.sock"
	}

	defaultHeaders := map[string]string{"User-Agent": "go-conception-1.0"}
	client, err := client.NewClient(host, "v1.21", nil, defaultHeaders)
	if err != nil {
		return nil, err
	}

	return &Docker{client: client}, nil
}

func (d *Docker) Run(opts Options) error {
	ctx := context.Background()

	config := &container.Config{
		Cmd:   opts.Command,
		Image: opts.Image,
	}
	container, err := d.client.ContainerCreate(ctx, config, nil, nil, "")
	if err != nil {
		return err
	}

	defer func() {
		removeOpts := types.ContainerRemoveOptions{
			// RemoveVolumes: true,
			// RemoveLinks:   true,
			Force: true,
		}
		d.client.ContainerRemove(ctx, container.ID, removeOpts)
	}()

	attachOpts := types.ContainerAttachOptions{
		Stream: true,
		Stdout: true,
		Stderr: true,
	}
	res, err := d.client.ContainerAttach(ctx, container.ID, attachOpts)
	if err != nil {
		return err
	}

	stdCopyErr := make(chan error, 1)
	go func() {
		defer res.Close()

		// Write to stdout and/or stderr if a write buffer is provided
		if opts.Stdout != nil || opts.Stderr != nil {
			// Demultiplex combined stdout, and stderr stream to separate streams
			_, err := stdcopy.StdCopy(opts.Stdout, opts.Stderr, res.Reader)
			stdCopyErr <- err
		}
	}()

	if err := d.client.ContainerStart(ctx, container.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	_, err = d.client.ContainerWait(ctx, container.ID)
	if err != nil {
		return err
	}

	select {
	case err := <-stdCopyErr:
		return err
	case <-ctx.Done():
	}

	return nil
}
