// file: pkg/profiles/language_test.go
// version: 1.0.0
// guid: 9c8b7a6d-5e4f-0a9b-3c2d-6e5f7a8b9c0d

package profiles

import (
	"testing"
	"time"
)

func TestLanguageProfileValidation(t *testing.T) {
	tests := []struct {
		name    string
		profile *LanguageProfile
		wantErr bool
		wantMsg string
	}{
		{
			name: "valid profile",
			profile: &LanguageProfile{
				Name:        "Test Profile",
				Languages:   []LanguageConfig{{Language: "en", Priority: 1, Forced: false, HI: false}},
				CutoffScore: 75,
			},
			wantErr: false,
		},
		{
			name:    "empty_name",
			profile: &LanguageProfile{Name: "", Languages: []LanguageConfig{{Language: "en", Priority: 1}}, CutoffScore: 80},
			wantErr: true,
			wantMsg: "name[0]: Name cannot be empty",
		},
		{
			name: "empty language code",
			profile: &LanguageProfile{
				Name:        "Test Profile",
				Languages:   []LanguageConfig{{Language: "", Priority: 1, Forced: false, HI: false}},
				CutoffScore: 75,
			},
			wantErr: true,
			wantMsg: "languages[0]: Language code cannot be empty",
		},
		{
			name: "duplicate priority",
			profile: &LanguageProfile{
				Name:        "Test Profile",
				Languages:   []LanguageConfig{{Language: "en", Priority: 1, Forced: false, HI: false}, {Language: "es", Priority: 1, Forced: false, HI: false}},
				CutoffScore: 75,
			},
			wantErr: true,
			wantMsg: "languages[1]: Duplicate priority",
		},
		{
			name:    "cutoff_score_out_of_range",
			profile: &LanguageProfile{Name: "English", Languages: []LanguageConfig{{Language: "en", Priority: 1}}, CutoffScore: 101},
			wantErr: true,
			wantMsg: "cutoff_score[0]: Cutoff score must be between 0 and 100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.profile.ValidateProfile()
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				} else if err.Error() != tt.wantMsg {
					t.Errorf("unexpected error: got %q, want %q", err.Error(), tt.wantMsg)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestLanguageProfile_GetPrimaryLanguage(t *testing.T) {
	tests := []struct {
		name     string
		profile  *LanguageProfile
		expected *LanguageConfig
	}{
		{
			name: "single language",
			profile: &LanguageProfile{
				Languages: []LanguageConfig{{Language: "en", Priority: 1, Forced: false, HI: false}},
			},
			expected: &LanguageConfig{Language: "en", Priority: 1, Forced: false, HI: false},
		},
		{
			name: "multiple languages",
			profile: &LanguageProfile{
				Languages: []LanguageConfig{{Language: "en", Priority: 2, Forced: false, HI: false}, {Language: "es", Priority: 1, Forced: false, HI: false}, {Language: "fr", Priority: 3, Forced: false, HI: false}},
			},
			expected: &LanguageConfig{Language: "es", Priority: 1, Forced: false, HI: false},
		},
		{
			name:     "no languages",
			profile:  &LanguageProfile{Languages: []LanguageConfig{}},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.profile.GetPrimaryLanguage()
			if tt.expected == nil {
				if got != nil {
					t.Errorf("expected nil, got %v", got)
				}
			} else {
				if got == nil {
					t.Errorf("expected %v, got nil", tt.expected)
				} else if *got != *tt.expected {
					t.Errorf("expected %v, got %v", *tt.expected, *got)
				}
			}
		})
	}
}

func TestLanguageProfile_ToJSON_FromJSON(t *testing.T) {
	original := &LanguageProfile{
		ID:          "test-id",
		Name:        "Test Profile",
		CutoffScore: 85,
		IsDefault:   false,
		Languages:   []LanguageConfig{{Language: "en", Priority: 1, Forced: false, HI: false}, {Language: "es", Priority: 2, Forced: true, HI: true}},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	jsonStr, err := original.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON() error: %v", err)
	}

	restored := &LanguageProfile{}
	err = restored.FromJSON(jsonStr)
	if err != nil {
		t.Fatalf("FromJSON() error: %v", err)
	}

	if len(restored.Languages) != len(original.Languages) {
		t.Errorf("expected %d languages, got %d", len(original.Languages), len(restored.Languages))
	}
	for i, lang := range original.Languages {
		if i >= len(restored.Languages) {
			continue
		}
		restoredLang := restored.Languages[i]
		if lang.Language != restoredLang.Language || lang.Priority != restoredLang.Priority || lang.Forced != restoredLang.Forced || lang.HI != restoredLang.HI {
			t.Errorf("FromJSON() language[%d] = %v, want %v", i, restoredLang, lang)
		}
	}
}

func TestDefaultProfile(t *testing.T) {
	profile := DefaultProfile()
	if profile.Name != "Default" {
		t.Errorf("expected name 'Default', got %q", profile.Name)
	}
	if len(profile.Languages) != 1 || profile.Languages[0].Language != "en" {
		t.Errorf("expected default language 'en', got %+v", profile.Languages)
	}
	if !profile.IsDefault {
		t.Errorf("expected IsDefault true")
	}
	if err := profile.ValidateProfile(); err != nil {
		t.Errorf("expected valid default profile, got error: %v", err)
	}
}
