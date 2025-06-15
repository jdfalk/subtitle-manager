package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
)

// searchCmd lists available subtitles from a provider.
var searchCmd = &cobra.Command{
	Use:   "search [provider] [media] [lang]",
	Short: "Search for subtitles without downloading",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("search")
		name, media, lang := args[0], args[1], args[2]
		key := viper.GetString("opensubtitles.api_key")
		p, err := providers.Get(name, key)
		if err != nil {
			return err
		}
		s, ok := p.(providers.Searcher)
		if !ok {
			return fmt.Errorf("provider %s does not support search", name)
		}
		urls, err := s.Search(context.Background(), media, lang)
		if err != nil {
			return err
		}
		for _, u := range urls {
			fmt.Println(u)
		}
		logger.Infof("found %d results", len(urls))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
