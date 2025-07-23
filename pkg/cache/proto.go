// file: pkg/cache/proto.go
// version: 1.0.0
// guid: 94a9a33b-5e6e-4b6c-9c34-24e2f5a1e3a1

package cache

import (
	commonpb "github.com/jdfalk/gcommon/pkg/common/proto"
	"google.golang.org/protobuf/types/known/durationpb"
)

// TTLConfigToProto converts a TTLConfig to a map of CachePolicy messages
// keyed by configuration name.
func TTLConfigToProto(cfg TTLConfig) map[string]*commonpb.CachePolicy {
	return map[string]*commonpb.CachePolicy{
		"provider_search_results": {DefaultTtl: durationpb.New(cfg.ProviderSearchResults)},
		"search_results":          {DefaultTtl: durationpb.New(cfg.SearchResults)},
		"tmdb_metadata":           {DefaultTtl: durationpb.New(cfg.TMDBMetadata)},
		"translation_results":     {DefaultTtl: durationpb.New(cfg.TranslationResults)},
		"user_sessions":           {DefaultTtl: durationpb.New(cfg.UserSessions)},
		"api_responses":           {DefaultTtl: durationpb.New(cfg.APIResponses)},
	}
}

// TTLConfigFromProto converts a map of CachePolicy messages to TTLConfig.
// Unknown keys are ignored.
func TTLConfigFromProto(policies map[string]*commonpb.CachePolicy) TTLConfig {
	var cfg TTLConfig
	if p, ok := policies["provider_search_results"]; ok && p.GetDefaultTtl() != nil {
		cfg.ProviderSearchResults = p.GetDefaultTtl().AsDuration()
	}
	if p, ok := policies["search_results"]; ok && p.GetDefaultTtl() != nil {
		cfg.SearchResults = p.GetDefaultTtl().AsDuration()
	}
	if p, ok := policies["tmdb_metadata"]; ok && p.GetDefaultTtl() != nil {
		cfg.TMDBMetadata = p.GetDefaultTtl().AsDuration()
	}
	if p, ok := policies["translation_results"]; ok && p.GetDefaultTtl() != nil {
		cfg.TranslationResults = p.GetDefaultTtl().AsDuration()
	}
	if p, ok := policies["user_sessions"]; ok && p.GetDefaultTtl() != nil {
		cfg.UserSessions = p.GetDefaultTtl().AsDuration()
	}
	if p, ok := policies["api_responses"]; ok && p.GetDefaultTtl() != nil {
		cfg.APIResponses = p.GetDefaultTtl().AsDuration()
	}
	return cfg
}
