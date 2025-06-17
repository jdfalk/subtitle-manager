package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
)

// searchCmd lists available subtitles from a provider.
var searchCmd = &cobra.Command{
	Use:   "search [media] [lang]",
	Short: "Search for subtitles across providers",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.GetLogger("search")
		media, lang := args[0], args[1]
		key := viper.GetString("opensubtitles.api_key")
		names := providers.All()
		var all []string
		for i, name := range names {
			p, err := providers.Get(name, key)
			if err != nil {
				continue
			}
			s, ok := p.(providers.Searcher)
			if !ok {
				continue
			}
			urls, err := s.Search(context.Background(), media, lang)
			if err == nil {
				all = append(all, urls...)
			}
			time.Sleep(time.Duration(i+1) * time.Second)
		}
		for _, u := range all {
			fmt.Println(u)
		}
		logger.Infof("found %d results", len(all))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
