# Subtitle Quality Scoring System

The subtitle quality scoring system evaluates and automatically selects the best
subtitle matches based on multiple criteria.

## Features

- **Provider Reliability**: Trusted providers and user reputation scoring
- **Release Match Quality**: Perfect release group matches and source quality
  alignment
- **Format Preferences**: SRT preferred, with configurable format priorities
- **Metadata Quality**: Upload date, download popularity, and user ratings
- **User Preferences**: Hearing impaired (HI) and forced subtitle preferences
- **File Size Correlation**: Optimal subtitle-to-video file size ratios

## CLI Usage

### Basic Usage

```bash
# Use scoring to fetch the best subtitle
subtitle-manager fetch-scored movie.mkv en output.srt

# With custom scoring weights
subtitle-manager fetch-scored \
  --provider-weight 0.3 \
  --release-weight 0.4 \
  --format-weight 0.1 \
  --metadata-weight 0.2 \
  --min-score 75 \
  movie.mkv en output.srt

# Allow and prefer hearing impaired subtitles
subtitle-manager fetch-scored \
  --allow-hi \
  --prefer-hi \
  movie.mkv en output.srt
```

### Available Options

- `--min-score`: Minimum score threshold (0-100, default: 50)
- `--provider-weight`: Provider reliability weight (0.0-1.0, default: 0.25)
- `--release-weight`: Release match weight (0.0-1.0, default: 0.35)
- `--format-weight`: Format preference weight (0.0-1.0, default: 0.15)
- `--metadata-weight`: Metadata quality weight (0.0-1.0, default: 0.25)
- `--allow-hi`: Allow hearing impaired subtitles (default: true)
- `--prefer-hi`: Prefer hearing impaired subtitles (default: false)

## Configuration

### Config File Settings

Add scoring configuration to your `~/.subtitle-manager.yaml`:

```yaml
scoring:
  # Scoring weights (must sum to ~1.0)
  provider_weight: 0.25
  release_weight: 0.35
  format_weight: 0.15
  metadata_weight: 0.25

  # Format preferences (in order of preference)
  preferred_formats:
    - 'srt'
    - 'ass'
    - 'ssa'
    - 'vtt'

  # Subtitle preferences
  allow_hi: true
  prefer_hi: false
  allow_forced: true
  prefer_forced: false

  # Quality thresholds
  min_score: 50
  max_age_days: 365
```

### Environment Variables

You can also use environment variables with the `SM_` prefix:

```bash
export SM_SCORING_MIN_SCORE=75
export SM_SCORING_PROVIDER_WEIGHT=0.3
export SM_SCORING_RELEASE_WEIGHT=0.4
```

## Web API

### Get Current Configuration

```bash
curl -X GET http://localhost:8080/api/scoring/config
```

### Update Configuration

```bash
curl -X POST http://localhost:8080/api/scoring/config \
  -H "Content-Type: application/json" \
  -d '{
    "providerWeight": 0.3,
    "releaseWeight": 0.4,
    "formatWeight": 0.1,
    "metadataWeight": 0.2,
    "preferredFormats": ["srt", "ass"],
    "allowHI": true,
    "preferHI": false,
    "minScore": 75,
    "maxAge": "8760h"
  }'
```

### Test Scoring

```bash
curl -X POST http://localhost:8080/api/scoring/test \
  -H "Content-Type: application/json" \
  -d '{
    "mediaPath": "/movies/Movie.2023.1080p.BluRay.x264-GROUP.mkv",
    "language": "en"
  }'
```

### Get Default Profile

```bash
curl -X GET http://localhost:8080/api/scoring/defaults
```

## Scoring Algorithm

### Provider Score (25% weight by default)

- **Trusted Providers**: +30 points for verified/trusted providers
- **Provider Reputation**: OpenSubtitles (+20), Subscene/Addic7ed (+15), Others
  (+5-10)
- **Translation Quality**: -20 for machine translated, -10 for auto-translated

### Release Score (35% weight by default)

- **Perfect Match**: +40 for exact release group match
- **Source Quality**: +30 for matching source (BluRay, WEB-DL, etc.)
- **Resolution Match**: +15 for matching resolution
- **Codec Match**: +10 for matching codec
- **Quality Penalties**: -30 for CAM on BluRay source, -25 for TS on WEB-DL

### Format Score (15% weight by default)

- **Preferred Formats**: Higher scores for preferred formats (SRT: +25, ASS/SSA:
  +20)
- **Format Order**: Points decrease for later preferences in the list

### Metadata Score (25% weight by default)

- **Upload Date**: Newer uploads get higher scores (1 week: +20, 1 month: +15,
  etc.)
- **Download Popularity**: Logarithmic scaling based on download count
- **User Ratings**: Rating \* 5, weighted by vote count
- **HI Preferences**: +15 bonus if preferred, -25 penalty if not allowed
- **Forced Preferences**: +10 bonus if preferred, -15 penalty if not allowed
- **HD Quality**: +10 bonus for HD content
- **File Size**: +5 bonus for appropriate subtitle-to-video size ratios

## Integration Examples

### Custom Provider Integration

```go
package main

import (
    "context"
    "github.com/jdfalk/subtitle-manager/pkg/scoring"
    "github.com/jdfalk/subtitle-manager/pkg/providers/opensubtitles"
)

func main() {
    // Search for subtitles
    client := opensubtitles.New("your-api-key")
    results, err := client.SearchWithResults(context.Background(), "movie.mkv", "en")
    if err != nil {
        panic(err)
    }

    // Convert to scoring format
    subtitles := make([]scoring.Subtitle, len(results))
    for i, result := range results {
        subtitles[i] = scoring.FromOpenSubtitlesResult(result, "opensubtitles")
    }

    // Extract media info
    media := scoring.FromMediaPath("movie.mkv")

    // Load profile and score
    profile := scoring.LoadProfileFromConfig()
    best, score := scoring.SelectBest(subtitles, media, profile)

    fmt.Printf("Best subtitle score: %d\n", score.Total)
}
```

### Batch Processing

```go
// Score multiple subtitle candidates
scored := scoring.ScoreSubtitles(subtitles, media, profile)

// Process results
for _, result := range scored {
    fmt.Printf("Subtitle: %s, Score: %d\n",
        result.Subtitle.Release, result.Score.Total)
}
```

## Best Practices

1. **Weight Configuration**: Ensure scoring weights sum to approximately 1.0
2. **Threshold Setting**: Start with min_score=50, adjust based on your quality
   requirements
3. **Format Preferences**: Order formats by your player compatibility and
   preference
4. **Provider Weighting**: Increase provider_weight if you have strong provider
   preferences
5. **Release Matching**: Increase release_weight for better source quality
   matching

## Troubleshooting

### Common Issues

1. **No Subtitles Meet Threshold**: Lower `min_score` or adjust weights
2. **Poor Quality Selection**: Increase `release_weight` and ensure proper media
   path parsing
3. **Wrong Format Selected**: Check `preferred_formats` configuration order
4. **HI/Forced Issues**: Verify `allow_hi`/`prefer_hi` and
   `allow_forced`/`prefer_forced` settings

### Debug Information

The `fetch-scored` command provides detailed scoring breakdown:

```
selected subtitle with score 89 (provider: 100, release: 90, format: 75, metadata: 100)
```

Use this information to tune your scoring weights and preferences.
