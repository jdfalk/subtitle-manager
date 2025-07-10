package cmd

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/cache"
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

		type cacheReq struct {
			Providers []string `json:"providers"`
			MediaPath string   `json:"mediaPath"`
			Language  string   `json:"language"`
		}

		req := cacheReq{Providers: names, MediaPath: media, Language: lang}
		reqData, _ := json.Marshal(req)
		sum := sha1.Sum(reqData)
		cacheKey := fmt.Sprintf("%x", sum)

		var mgr *cache.Manager
		if c, err := cache.NewManagerFromViper(); err == nil {
			mgr = c
			defer mgr.Close()
			if data, err := mgr.GetProviderSearchResults(cmd.Context(), cacheKey); err == nil && data != nil {
				var cached []string
				if err := json.Unmarshal(data, &cached); err == nil {
					for _, u := range cached {
						fmt.Println(u)
					}
					logger.Infof("found %d results (cached)", len(cached))
					return nil
				}
			}
		} else {
			logger.Warnf("cache disabled: %v", err)
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
			urls, err := s.Search(cmd.Context(), media, lang)
			if err == nil {
				all = append(all, urls...)
			}
			time.Sleep(time.Duration(i+1) * time.Second)
		}

		if mgr != nil && len(all) > 0 {
			if data, err := json.Marshal(all); err == nil {
				mgr.SetProviderSearchResults(cmd.Context(), cacheKey, data)
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
