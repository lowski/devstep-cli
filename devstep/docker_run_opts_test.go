package devstep_test

import (
	"github.com/fgrehm/devstep-cli/devstep"
	"github.com/fsouza/go-dockerclient"
	"testing"
)

func Test_ToCreateOptsPublishPorts(t *testing.T) {
	opts := &devstep.DockerRunOpts{
		Publish: []string{"80:80", "8080"},
	}
	config := opts.ToCreateOpts()

	expExposedPorts := map[docker.Port]struct{}{
		docker.Port("80"):   struct{}{},
		docker.Port("8080"): struct{}{},
	}

	exposedPorts := config.Config.ExposedPorts
	equals(t, expExposedPorts, exposedPorts)

	expPortBindings := map[docker.Port][]docker.PortBinding{
		docker.Port("80"):   []docker.PortBinding{docker.PortBinding{HostPort: "80"}},
		docker.Port("8080"): []docker.PortBinding{docker.PortBinding{}},
	}

	portBindings := config.HostConfig.PortBindings
	equals(t, expPortBindings, portBindings)
}
