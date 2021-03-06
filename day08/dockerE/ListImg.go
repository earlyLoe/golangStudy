package dockerE

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func GetImageList() interface{} {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithHost("http://yyy:2376"),
		client.WithVersion("1.36"))
	if nil != err {
		panic(err)
	}

	images, err := cli.ImageList(ctx, types.ImageListOptions{})

	if nil != err {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image)
	}

	return images
}
