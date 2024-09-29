package publisher

import "github.com/toolsascode/helm-s3-publisher/internal/publish"

func pubCmd() publisherInterface {
	return &publish.Commands{}
}

func Publisher() {
	pubCmd().Run()
}
