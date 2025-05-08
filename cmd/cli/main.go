package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/nifle3/bfg/internal/core"
	"github.com/nifle3/bfg/internal/executor/interpreter"
	"github.com/nifle3/bfg/internal/reader/file"
	"github.com/nifle3/bfg/internal/reader/repl"

	"github.com/urfave/cli/v3"
)

const (
	argSourceFile = "input-file"
)

func startRepl(ctx context.Context, _ *cli.Command) error {
	repl := repl.New()
	executor := interpreter.New()
	core.Run(executor, repl)

	return nil
}

func startInterpreter(ctx context.Context, cmd *cli.Command) error {
	sourceFile := cmd.StringArg(argSourceFile)
	sourceFile = strings.TrimSpace(sourceFile)
	isCorrectExtension := strings.HasSuffix(sourceFile, ".bf")

	if !isCorrectExtension {
		return fmt.Errorf("%s has invalid extension", sourceFile)
	}

	reader := file.New(sourceFile)
	executor := interpreter.New()

	core.Run(executor, reader)

	return nil
}

func startCompile(ctx context.Context, _ *cli.Command) error {
	return nil
}

func main() {
	cmd := cli.Command{
		Commands: []*cli.Command{
			{
				Name:   "repl",
				Usage:  "Start repl",
				Action: startRepl,
			},
			{
				Name:   "interpreter",
				Usage:  "Start program from file with .bf",
				Action: startInterpreter,
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: argSourceFile,
					},
				},
			},
			{
				Name:   "compile",
				Usage:  "Compile program from file with .bf",
				Action: startCompile,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "target",
						Usage:   "wasm|llvm(default)",
						Value:   "llvm",
						Aliases: []string{"t"},
					},
					&cli.StringFlag{
						Name:    "output",
						Usage:   "File to output default a.exe or a",
						Value:   "a",
						Aliases: []string{"o"},
					},
				},
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: argSourceFile,
					},
				},
			},
		},
		DefaultCommand: "interpreter",
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
