// file: pkg/bazarr/mapper.go
package bazarr

// MapSettings converts Bazarr settings to Subtitle Manager configuration keys.
// The returned map can be written into Viper.
func MapSettings(s Settings) map[string]any {
	out := make(map[string]any)
	if gen, ok := s["general"].(map[string]any); ok {
		if host, ok := gen["bind_address"].(string); ok {
			out["web.host"] = host
		}
		if p, ok := gen["port"]; ok {
			switch v := p.(type) {
			case float64:
				out["web.port"] = int(v)
			case int:
				out["web.port"] = v
			}
		}
		if base, ok := gen["url_base"].(string); ok {
			out["web.base_url"] = base
		}
	}
	if auth, ok := s["auth"].(map[string]any); ok {
		if user, ok := auth["user"].(string); ok {
			out["auth.user"] = user
		}
		if pass, ok := auth["password"].(string); ok {
			out["auth.password"] = pass
		}
		if key, ok := auth["api_key"].(string); ok {
			out["auth.api_key"] = key
		}
	}
	if prov, ok := s["providers"].(map[string]any); ok {
		enabled := map[string]bool{}
		for name, v := range prov {
			if p, ok := v.(map[string]any); ok {
				if en, ok := p["enabled"].(bool); ok && en {
					enabled[name] = true
				}
			}
		}
		if len(enabled) > 0 {
			out["providers"] = enabled
		}
	}
	if langs, ok := s["languages"].([]any); ok {
		list := []string{}
		for _, v := range langs {
			if m, ok := v.(map[string]any); ok {
				if en, ok := m["enabled"].(bool); ok && en {
					if code, ok := m["code"].(string); ok {
						list = append(list, code)
					}
				}
			}
		}
		if len(list) > 0 {
			out["languages"] = list
		}
	}
	return out
}
