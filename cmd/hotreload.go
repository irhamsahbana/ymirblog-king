// Package cmd is the command surface of ymirblog cli tool provided by kubuskotak.
// # This manifest was generated by ymir. DO NOT EDIT.
package cmd

import (
	"context"
	"os"
	"path/filepath"

	"github.com/kubuskotak/asgard/hotreload"
	"github.com/kubuskotak/asgard/signal"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/abialemuel/ymirblog/pkg/infrastructure"
)

type hotReload struct {
	Path     string
	Filename string
}

func newHotReloadCmd() *cobra.Command {
	var (
		r   = &hotReload{}
		cmd = &cobra.Command{
			Use:   `reload`,
			Short: "hot reload",
			RunE: func(cmd *cobra.Command, args []string) error {
				return r.Run(cmd, args)
			},
		}
	)
	cmd.PersistentFlags().StringVarP(
		&r.Filename, "config", "c", "config.yaml", "config file name")
	cmd.PersistentFlags().StringVarP(
		&r.Path, "config-path", "d", "./", "config dir path")

	return cmd
}

// Run is start hot reload.
func (r *hotReload) Run(_ *cobra.Command, _ []string) error {
	var (
		wd, _   = os.Getwd()
		rest    = filepath.Join(wd, "pkg", "api", "rest")
		usecase = filepath.Join(wd, "pkg", "usecase")
		root    = filepath.Join(wd, "cmd", "root.go")
		errCh   = make(chan error)
		stopCh  = signal.SetupSignalHandler()
	)

	stop, err := hotreload.Start(
		hotreload.WithDirs([]string{root, rest, usecase}),
		hotreload.WithLogger(log.Logger),
	)
	if err != nil {
		return err
	}
	return signal.Graceful(infrastructure.Envs.Server.Timeout, stopCh, errCh, func(ctx context.Context) {
		_ = stop()
	})
}
