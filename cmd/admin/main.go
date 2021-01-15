package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bo-er/todo-simpler/admin/app"
	"github.com/urfave/cli/v2"
)

func main() {
	ctx := NewTagContext(context.Background(), "__main__")

	app := cli.NewApp()
	app.Name = "todo-demo"
	app.Usage = "This is a SWS Todo Demo server."
	app.Commands = []*cli.Command{
		newWebCmd(ctx),
	}
	err := app.Run(os.Args)
	fmt.Println("运行参数是:", os.Args)
	if err != nil {
		log.Fatalf("服务运行失败，错误原因是:%s", err.Error())
	}
}

func newWebCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "todosimpler",
		Usage: "运行todo simpler admin服务",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "配置文件(.json,.yaml,.toml)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return app.Run(ctx,
				app.SetConfigFile(c.String("conf")))
		},
	}
}

type tagKey struct{}

// NewTagContext ...
func NewTagContext(ctx context.Context, tag string) context.Context {
	return context.WithValue(ctx, tagKey{}, tag)
}
