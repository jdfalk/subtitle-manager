// file: pkg/bazarr/mapper.go
package bazarr

import (
	"fmt"
)

// MappingInfo describes where a Bazarr setting will be mapped in subtitle-manager
type MappingInfo struct {
	Key         string `json:"key"`
	Section     string `json:"section"`
	Description string `json:"description"`
	Value       any    `json:"value"`
	Original    any    `json:"original"`
}

// MapSettingsWithInfo converts a Bazarr `Settings` map into Subtitle Manager
// configuration keys and returns both the resulting map and a slice of
// `MappingInfo` describing each mapping.
//
// Parameters:
//   - s: Bazarr settings as returned by `FetchSettings`.
//
// Returns:
//   - map[string]any: configuration keys compatible with viper.
//   - []MappingInfo: list of details about how each setting was mapped.
func MapSettingsWithInfo(s Settings) (map[string]any, []MappingInfo) {
	out := make(map[string]any)
	var mappings []MappingInfo

	addMapping := func(key, section, description string, value, original any) {
		out[key] = value
		mappings = append(mappings, MappingInfo{
			Key:         key,
			Section:     section,
			Description: description,
			Value:       value,
			Original:    original,
		})
	}

	// General settings
	if gen, ok := s["general"].(map[string]any); ok {
		if host, ok := gen["ip"].(string); ok && host != "" {
			addMapping("web.host", "Web Server", "Server bind address", host, host)
		}
		if p, ok := gen["port"]; ok {
			var port int
			switch v := p.(type) {
			case float64:
				port = int(v)
			case int:
				port = v
			}
			if port > 0 {
				addMapping("web.port", "Web Server", "Server port", port, p)
			}
		}
		if base, ok := gen["base_url"].(string); ok && base != "" {
			addMapping("web.base_url", "Web Server", "Base URL path", base, base)
		}
		if theme, ok := gen["theme"].(string); ok && theme != "" {
			addMapping("ui.theme", "User Interface", "UI theme preference", theme, theme)
		}
		if debug, ok := gen["debug"].(bool); ok {
			addMapping("log.debug", "Logging", "Enable debug logging", debug, debug)
		}
		if minScore, ok := gen["minimum_score"]; ok {
			addMapping("matching.minimum_score", "Subtitle Matching", "Minimum score for series subtitles", minScore, minScore)
		}
		if minScoreMovie, ok := gen["minimum_score_movie"]; ok {
			addMapping("matching.minimum_score_movie", "Subtitle Matching", "Minimum score for movie subtitles", minScoreMovie, minScoreMovie)
		}
		if pageSize, ok := gen["page_size"]; ok {
			addMapping("ui.page_size", "User Interface", "Number of items per page", pageSize, pageSize)
		}
		if subfolder, ok := gen["subfolder"].(string); ok && subfolder != "" {
			addMapping("subtitles.subfolder", "Subtitle Storage", "Subtitle folder organization", subfolder, subfolder)
		}
		if enabled, ok := gen["enabled_providers"].([]any); ok && len(enabled) > 0 {
			providers := make([]string, 0, len(enabled))
			for _, p := range enabled {
				if providerName, ok := p.(string); ok {
					providers = append(providers, providerName)
				}
			}
			if len(providers) > 0 {
				addMapping("providers.enabled", "Subtitle Providers", "Enabled subtitle providers", providers, enabled)
			}
		}
		if upgradeSubs, ok := gen["upgrade_subs"].(bool); ok {
			addMapping("subtitles.upgrade_enabled", "Subtitle Management", "Enable subtitle upgrades", upgradeSubs, upgradeSubs)
		}
		if upgradeFreq, ok := gen["upgrade_frequency"]; ok {
			addMapping("subtitles.upgrade_frequency", "Subtitle Management", "Upgrade check frequency (hours)", upgradeFreq, upgradeFreq)
		}
		if wantedFreq, ok := gen["wanted_search_frequency"]; ok {
			addMapping("search.wanted_frequency", "Search Settings", "Wanted search frequency (hours)", wantedFreq, wantedFreq)
		}
		if useEmbedded, ok := gen["use_embedded_subs"].(bool); ok {
			addMapping("subtitles.use_embedded", "Subtitle Sources", "Use embedded subtitles", useEmbedded, useEmbedded)
		}
		if usePlex, ok := gen["use_plex"].(bool); ok {
			addMapping("integrations.plex.enabled", "Integrations", "Enable Plex integration", usePlex, usePlex)
		}
		if useRadarr, ok := gen["use_radarr"].(bool); ok {
			addMapping("integrations.radarr.enabled", "Integrations", "Enable Radarr integration", useRadarr, useRadarr)
		}
		if useSonarr, ok := gen["use_sonarr"].(bool); ok {
			addMapping("integrations.sonarr.enabled", "Integrations", "Enable Sonarr integration", useSonarr, useSonarr)
		}
		if pathMappings, ok := gen["path_mappings"].([]any); ok && len(pathMappings) > 0 {
			addMapping("path_mappings.series", "Path Mappings", "Series path mappings", pathMappings, pathMappings)
		}
		if pathMappingsMovie, ok := gen["path_mappings_movie"].([]any); ok && len(pathMappingsMovie) > 0 {
			addMapping("path_mappings.movies", "Path Mappings", "Movie path mappings", pathMappingsMovie, pathMappingsMovie)
		}
	}

	// Authentication settings
	if auth, ok := s["auth"].(map[string]any); ok {
		if key, ok := auth["apikey"].(string); ok && key != "" {
			addMapping("auth.api_key", "Authentication", "API key for external access", key, key)
		}
		if user, ok := auth["username"].(string); ok && user != "" {
			addMapping("auth.username", "Authentication", "Web UI username", user, user)
		}
		if pass, ok := auth["password"].(string); ok && pass != "" {
			addMapping("auth.password", "Authentication", "Web UI password", pass, pass)
		}
	}

	// Plex integration
	if plex, ok := s["plex"].(map[string]any); ok {
		if ip, ok := plex["ip"].(string); ok && ip != "" {
			addMapping("integrations.plex.host", "Plex Integration", "Plex server IP address", ip, ip)
		}
		if port, ok := plex["port"]; ok {
			addMapping("integrations.plex.port", "Plex Integration", "Plex server port", port, port)
		}
		if apikey, ok := plex["apikey"].(string); ok && apikey != "" {
			addMapping("integrations.plex.token", "Plex Integration", "Plex authentication token", apikey, apikey)
		}
		if ssl, ok := plex["ssl"].(bool); ok {
			addMapping("integrations.plex.ssl", "Plex Integration", "Use SSL for Plex connection", ssl, ssl)
		}
		if movieLib, ok := plex["movie_library"].(string); ok && movieLib != "" {
			addMapping("integrations.plex.movie_library", "Plex Integration", "Plex movie library name", movieLib, movieLib)
		}
		if seriesLib, ok := plex["series_library"].(string); ok && seriesLib != "" {
			addMapping("integrations.plex.series_library", "Plex Integration", "Plex series library name", seriesLib, seriesLib)
		}
	}

	// Radarr integration
	if radarr, ok := s["radarr"].(map[string]any); ok {
		if ip, ok := radarr["ip"].(string); ok && ip != "" {
			addMapping("integrations.radarr.host", "Radarr Integration", "Radarr server IP address", ip, ip)
		}
		if port, ok := radarr["port"]; ok {
			addMapping("integrations.radarr.port", "Radarr Integration", "Radarr server port", port, port)
		}
		if apikey, ok := radarr["apikey"].(string); ok && apikey != "" {
			addMapping("integrations.radarr.api_key", "Radarr Integration", "Radarr API key", apikey, apikey)
		}
		if ssl, ok := radarr["ssl"].(bool); ok {
			addMapping("integrations.radarr.ssl", "Radarr Integration", "Use SSL for Radarr connection", ssl, ssl)
		}
		if baseUrl, ok := radarr["base_url"].(string); ok && baseUrl != "" {
			addMapping("integrations.radarr.base_url", "Radarr Integration", "Radarr base URL path", baseUrl, baseUrl)
		}
		if timeout, ok := radarr["http_timeout"]; ok {
			addMapping("integrations.radarr.timeout", "Radarr Integration", "HTTP timeout (seconds)", timeout, timeout)
		}
		if syncInterval, ok := radarr["movies_sync"]; ok {
			addMapping("integrations.radarr.sync_interval", "Radarr Integration", "Movie sync interval (minutes)", syncInterval, syncInterval)
		}
	}

	// Sonarr integration
	if sonarr, ok := s["sonarr"].(map[string]any); ok {
		if ip, ok := sonarr["ip"].(string); ok && ip != "" {
			addMapping("integrations.sonarr.host", "Sonarr Integration", "Sonarr server IP address", ip, ip)
		}
		if port, ok := sonarr["port"]; ok {
			addMapping("integrations.sonarr.port", "Sonarr Integration", "Sonarr server port", port, port)
		}
		if apikey, ok := sonarr["apikey"].(string); ok && apikey != "" {
			addMapping("integrations.sonarr.api_key", "Sonarr Integration", "Sonarr API key", apikey, apikey)
		}
		if ssl, ok := sonarr["ssl"].(bool); ok {
			addMapping("integrations.sonarr.ssl", "Sonarr Integration", "Use SSL for Sonarr connection", ssl, ssl)
		}
		if baseUrl, ok := sonarr["base_url"].(string); ok && baseUrl != "" {
			addMapping("integrations.sonarr.base_url", "Sonarr Integration", "Sonarr base URL path", baseUrl, baseUrl)
		}
		if timeout, ok := sonarr["http_timeout"]; ok {
			addMapping("integrations.sonarr.timeout", "Sonarr Integration", "HTTP timeout (seconds)", timeout, timeout)
		}
		if episodeSync, ok := sonarr["episodes_sync"]; ok {
			addMapping("integrations.sonarr.episode_sync_interval", "Sonarr Integration", "Episode sync interval (minutes)", episodeSync, episodeSync)
		}
		if seriesSync, ok := sonarr["series_sync"]; ok {
			addMapping("integrations.sonarr.series_sync_interval", "Sonarr Integration", "Series sync interval (minutes)", seriesSync, seriesSync)
		}
	}

	// Notification providers
	if notifications, ok := s["notifications"].(map[string]any); ok {
		if providers, ok := notifications["providers"].([]any); ok {
			enabledNotifs := make([]map[string]any, 0)
			for _, p := range providers {
				if provider, ok := p.(map[string]any); ok {
					if enabled, ok := provider["enabled"].(bool); ok && enabled {
						if name, ok := provider["name"].(string); ok {
							enabledNotifs = append(enabledNotifs, map[string]any{
								"name":    name,
								"url":     provider["url"],
								"enabled": true,
							})
						}
					}
				}
			}
			if len(enabledNotifs) > 0 {
				addMapping("notifications.providers", "Notifications", "Enabled notification providers", enabledNotifs, providers)
			}
		}
	}

	// Database settings
	if postgres, ok := s["postgresql"].(map[string]any); ok {
		if enabled, ok := postgres["enabled"].(bool); ok && enabled {
			if host, ok := postgres["host"].(string); ok && host != "" {
				addMapping("database.postgres.host", "Database", "PostgreSQL host", host, host)
			}
			if port, ok := postgres["port"]; ok {
				addMapping("database.postgres.port", "Database", "PostgreSQL port", port, port)
			}
			if db, ok := postgres["database"].(string); ok && db != "" {
				addMapping("database.postgres.database", "Database", "PostgreSQL database name", db, db)
			}
			if user, ok := postgres["username"].(string); ok && user != "" {
				addMapping("database.postgres.username", "Database", "PostgreSQL username", user, user)
			}
			if pass, ok := postgres["password"].(string); ok && pass != "" {
				addMapping("database.postgres.password", "Database", "PostgreSQL password", pass, pass)
			}
		}
	}

	// Backup settings
	if backup, ok := s["backup"].(map[string]any); ok {
		if folder, ok := backup["folder"].(string); ok && folder != "" {
			addMapping("backup.folder", "Backup", "Backup folder path", folder, folder)
		}
		if freq, ok := backup["frequency"].(string); ok && freq != "" {
			addMapping("backup.frequency", "Backup", "Backup frequency", freq, freq)
		}
		if retention, ok := backup["retention"]; ok {
			addMapping("backup.retention_days", "Backup", "Backup retention (days)", retention, retention)
		}
	}

	// Provider-specific settings
	mapProviderSettings(s, addMapping)

	// Scoring settings
	if seriesScores, ok := s["series_scores"].(map[string]any); ok {
		addMapping("scoring.series", "Subtitle Scoring", "Series subtitle scoring weights", seriesScores, seriesScores)
	}
	if movieScores, ok := s["movie_scores"].(map[string]any); ok {
		addMapping("scoring.movies", "Subtitle Scoring", "Movie subtitle scoring weights", movieScores, movieScores)
	}

	// Embedded subtitles settings
	if embedded, ok := s["embeddedsubtitles"].(map[string]any); ok {
		if includeSrt, ok := embedded["include_srt"].(bool); ok {
			addMapping("subtitles.embedded.include_srt", "Embedded Subtitles", "Include SRT embedded subtitles", includeSrt, includeSrt)
		}
		if includeAss, ok := embedded["include_ass"].(bool); ok {
			addMapping("subtitles.embedded.include_ass", "Embedded Subtitles", "Include ASS embedded subtitles", includeAss, includeAss)
		}
		if timeout, ok := embedded["timeout"]; ok {
			addMapping("subtitles.embedded.timeout", "Embedded Subtitles", "Extraction timeout (seconds)", timeout, timeout)
		}
	}

	// Proxy settings
	if proxy, ok := s["proxy"].(map[string]any); ok {
		if proxyType, ok := proxy["type"].(string); ok && proxyType != "" {
			addMapping("network.proxy.type", "Network", "Proxy type", proxyType, proxyType)
		}
		if url, ok := proxy["url"].(string); ok && url != "" {
			addMapping("network.proxy.url", "Network", "Proxy URL", url, url)
		}
		if port, ok := proxy["port"].(string); ok && port != "" {
			addMapping("network.proxy.port", "Network", "Proxy port", port, port)
		}
		if user, ok := proxy["username"].(string); ok && user != "" {
			addMapping("network.proxy.username", "Network", "Proxy username", user, user)
		}
		if pass, ok := proxy["password"].(string); ok && pass != "" {
			addMapping("network.proxy.password", "Network", "Proxy password", pass, pass)
		}
	}

	return out, mappings
}

// mapProviderSettings maps individual provider configurations
func mapProviderSettings(s Settings, addMapping func(string, string, string, any, any)) {
	providerMap := map[string]string{
		"opensubtitles":    "OpenSubtitles",
		"opensubtitlescom": "OpenSubtitles.com",
		"addic7ed":         "Addic7ed",
		"podnapisi":        "Podnapisi",
		"subscene":         "Subscene",
		"napiprojekt":      "NapiProjekt",
		"legendasdivx":     "LegendasDivx",
		"assrt":            "Assrt",
	}

	for provider, displayName := range providerMap {
		if config, ok := s[provider].(map[string]any); ok {
			sectionName := fmt.Sprintf("%s Provider", displayName)
			keyPrefix := fmt.Sprintf("providers.%s", provider)

			if username, ok := config["username"].(string); ok && username != "" {
				addMapping(fmt.Sprintf("%s.username", keyPrefix), sectionName, fmt.Sprintf("%s username", displayName), username, username)
			}
			if password, ok := config["password"].(string); ok && password != "" {
				addMapping(fmt.Sprintf("%s.password", keyPrefix), sectionName, fmt.Sprintf("%s password", displayName), password, password)
			}
			if apiKey, ok := config["api_key"].(string); ok && apiKey != "" {
				addMapping(fmt.Sprintf("%s.api_key", keyPrefix), sectionName, fmt.Sprintf("%s API key", displayName), apiKey, apiKey)
			}
			if token, ok := config["token"].(string); ok && token != "" {
				addMapping(fmt.Sprintf("%s.token", keyPrefix), sectionName, fmt.Sprintf("%s token", displayName), token, token)
			}
			if vip, ok := config["vip"].(bool); ok {
				addMapping(fmt.Sprintf("%s.vip", keyPrefix), sectionName, fmt.Sprintf("%s VIP status", displayName), vip, vip)
			}
			if ssl, ok := config["ssl"].(bool); ok {
				addMapping(fmt.Sprintf("%s.ssl", keyPrefix), sectionName, fmt.Sprintf("%s use SSL", displayName), ssl, ssl)
			}
		}
	}
}

// MapSettings converts Bazarr settings to Subtitle Manager configuration keys.
// This is the legacy function for backward compatibility.
func MapSettings(s Settings) map[string]any {
	config, _ := MapSettingsWithInfo(s)
	return config
}
