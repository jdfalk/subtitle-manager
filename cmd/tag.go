// file: cmd/tag.go
// version: 1.0.0
// guid: 3796780f-7c51-4ad6-b073-43e3d5e9e837

package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/tagging"
)

var (
	tagTypeFlag   string
	tagEntityType string
	tagColor      string
	tagDesc       string
)

var tagCmd = &cobra.Command{
	Use:   "tag",
	Short: "Manage tags",
}

var tagListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tags",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := database.OpenSQLStore(database.GetDatabasePath())
		if err != nil {
			return err
		}
		defer store.Close()
		tags, err := store.ListTags()
		if err != nil {
			return err
		}
		for _, t := range tags {
			cmd.Printf("%s\t%s\t%s\t%s\t%s\n", t.ID, t.Name, t.Type, t.Color, t.Description)
		}
		return nil
	},
}

var tagAddCmd = &cobra.Command{
	Use:   "add [name]",
	Args:  cobra.ExactArgs(1),
	Short: "Create a tag",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := database.OpenSQLStore(database.GetDatabasePath())
		if err != nil {
			return err
		}
		defer store.Close()
		tm := tagging.NewTagManager(store.DB())
		_, err = tm.CreateTag(args[0], tagTypeFlag, tagEntityType, tagColor, tagDesc)
		return err
	},
}

var tagRemoveCmd = &cobra.Command{
	Use:   "remove [id]",
	Args:  cobra.ExactArgs(1),
	Short: "Delete a tag",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := database.OpenSQLStore(database.GetDatabasePath())
		if err != nil {
			return err
		}
		defer store.Close()
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("invalid id: %w", err)
		}
		return store.DeleteTag(id)
	},
}

func init() {
	tagAddCmd.Flags().StringVar(&tagTypeFlag, "type", "user", "tag type")
	tagAddCmd.Flags().StringVar(&tagEntityType, "entity", "all", "entity type")
	tagAddCmd.Flags().StringVar(&tagColor, "color", "", "tag color")
	tagAddCmd.Flags().StringVar(&tagDesc, "desc", "", "tag description")

	tagCmd.AddCommand(tagListCmd)
	tagCmd.AddCommand(tagAddCmd)
	tagCmd.AddCommand(tagRemoveCmd)
	rootCmd.AddCommand(tagCmd)
}
