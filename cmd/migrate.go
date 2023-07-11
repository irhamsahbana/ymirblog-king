// Package cmd is the command surface of ymirblog cli tool provided by kubuskotak.
// # This manifest was generated by ymir. DO NOT EDIT.
package cmd

import (
	"entgo.io/ent/dialect"
	"github.com/spf13/cobra"
)

type migrateOptions struct {
	Name    string
	Dialect string
	DSN     string
}

func newMigrateCmd() *cobra.Command {
	m := &migrateOptions{}
	cmd := &cobra.Command{
		Use:   `migrate`,
		Short: "Print version info",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Run(cmd, args)
		},
	}
	cmd.Flags().StringVarP(&m.Name, "filename", "n", "init", "migrate -n init")
	cmd.Flags().StringVarP(&m.Dialect, "dialect", "q", "sqlite3", "migrate -q mysql")
	cmd.Flags().StringVarP(&m.DSN, "data source name", "s", "sqlite://.data/blogs.migration.db?_fk=1", "migrate -s 'sqlite://.data/blogs.migration.db?_fk=1'")

	return cmd
}

// Run is ent func to migrate.
func (m *migrateOptions) Run(cmd *cobra.Command, _ []string) error {
	switch m.Dialect {
	case dialect.SQLite, dialect.MySQL:
		return nil
	default:
		return cmd.Usage()
	}
}
