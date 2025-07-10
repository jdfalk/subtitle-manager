package cmd

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/cache"
	"github.com/jdfalk/subtitle-manager/pkg/logging"
	"github.com/jdfalk/subtitle-manager/pkg/providers"
	web "github.com/jdfalk/subtitle-manager/pkg/webserver"
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

		mgr, err := cache.NewManagerFromViper()
		if err != nil {
			logger.Warnf("cache disabled: %v", err)
			mgr = nil
		}

		req := web.SearchRequest{
			Providers: names,
			MediaPath: media,
			Language:  lang,
		}
		data, _ := json.Marshal(req)
		cacheKey := fmt.Sprintf("%x", sha1.Sum(data))

		ctx := context.Background()
		if mgr != nil {
			if cached, err := mgr.GetProviderSearchResults(ctx, cacheKey); err == nil && cached != nil {
				var urls []string
				if err := json.Unmarshal(cached, &urls); err == nil {
					for _, u := range urls {
						fmt.Println(u)
					}
					logger.Infof("found %d results (cached)", len(urls))
					return nil
				}
			}
		}

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
			urls, err := s.Search(ctx, media, lang)
			if err == nil {
				all = append(all, urls...)
			}
			time.Sleep(time.Duration(i+1) * time.Second)
		}

		if mgr != nil && len(all) > 0 {
			if encoded, err := json.Marshal(all); err == nil {
				mgr.SetProviderSearchResults(ctx, cacheKey, encoded)
			}
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
