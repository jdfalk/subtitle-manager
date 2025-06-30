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
		name      string
		profile   *LanguageProfile
		wantError bool
	}{
		{
			name: "valid profile",
			profile: &LanguageProfile{
				ID:          "test-1",
				Name:        "Test Profile",
				Languages:   []LanguageConfig{{Language: "en", Priority: 1, Forced: false, HI: false}},
				CutoffScore: 75,
				IsDefault:   false,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			wantError: false,
		},
		{
			name: "empty name",
			profile: &LanguageProfile{
				ID:          "test-2",
				Name:        "",
				Languages:   []LanguageConfig{{Language: "en", Priority: 1, Forced: false, HI: false}},
				CutoffScore: 75,
			},
			wantError: true,
		},
		{
			name: "no languages",
			profile: &LanguageProfile{
				ID:          "test-3",
				Name:        "Test Profile",
				Languages:   []LanguageConfig{},
				CutoffScore: 75,
			},
			wantError: true,
		},
		{
			name: "invalid cutoff score",
			profile: &LanguageProfile{
				ID:          "test-4",
				Name:        "Test Profile",
				Languages:   []LanguageConfig{{Language: "en", Priority: 1, Forced: false, HI: false}},
				CutoffScore: 150,
			},
			wantError: true,
		},
		{
			name: "duplicate priority",
			profile: &LanguageProfile{
				ID:   "test-5",
				Name: "Test Profile",
				Languages: []LanguageConfig{
					{Language: "en", Priority: 1, Forced: false, HI: false},
					{Language: "es", Priority: 1, Forced: false, HI: false},
				},
				CutoffScore: 75,
			},
			wantError: true,
		},
		{
			name: "empty language code",
			profile: &LanguageProfile{
				ID:          "test-6",
				Name:        "Test Profile",
				Languages:   []LanguageConfig{{Language: "", Priority: 1, Forced: false, HI: false}},
				CutoffScore: 75,
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.profile.Validate()
			if (err != nil) != tt.wantError {
				t.Errorf("LanguageProfile.Validate() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}

func TestLanguageProfileGetPrimaryLanguage(t *testing.T) {
	profile := &LanguageProfile{
		ID:   "test",
		Name: "Test Profile",
		Languages: []LanguageConfig{
			{Language: "es", Priority: 2, Forced: false, HI: false},
			{Language: "en", Priority: 1, Forced: false, HI: false},
			{Language: "fr", Priority: 3, Forced: false, HI: false},
		},
		CutoffScore: 75,
	}

	primary := profile.GetPrimaryLanguage()
	if primary == nil {
		t.Fatal("GetPrimaryLanguage() returned nil")
	}

	if primary.Language != "en" {
		t.Errorf("GetPrimaryLanguage() returned %q, want %q", primary.Language, "en")
	}

	if primary.Priority != 1 {
		t.Errorf("GetPrimaryLanguage() returned priority %d, want %d", primary.Priority, 1)
	}
}

func TestLanguageProfileMarshalUnmarshal(t *testing.T) {
	original := &LanguageProfile{
		ID:   "test",
		Name: "Test Profile",
		Languages: []LanguageConfig{
			{Language: "en", Priority: 1, Forced: false, HI: false},
			{Language: "es", Priority: 2, Forced: true, HI: true},
		},
		CutoffScore: 80,
		IsDefault:   true,
		CreatedAt:   time.Now().Truncate(time.Second),
		UpdatedAt:   time.Now().Truncate(time.Second),
	}

	// Test marshal
	data, err := original.MarshalConfig()
	if err != nil {
		t.Fatalf("MarshalConfig() failed: %v", err)
	}

	// Test unmarshal
	restored := &LanguageProfile{}
	if err := restored.UnmarshalConfig(data); err != nil {
		t.Fatalf("UnmarshalConfig() failed: %v", err)
	}

	// Verify languages are preserved
	if len(restored.Languages) != len(original.Languages) {
		t.Errorf("Language count mismatch: got %d, want %d", len(restored.Languages), len(original.Languages))
	}

	for i, lang := range restored.Languages {
		origLang := original.Languages[i]
		if lang.Language != origLang.Language ||
			lang.Priority != origLang.Priority ||
			lang.Forced != origLang.Forced ||
			lang.HI != origLang.HI {
			t.Errorf("Language config mismatch at index %d: got %+v, want %+v", i, lang, origLang)
		}
	}
}

func TestDefaultLanguageProfile(t *testing.T) {
	profile := DefaultLanguageProfile()

	if profile == nil {
		t.Fatal("DefaultLanguageProfile() returned nil")
	}

	if profile.ID != "default" {
		t.Errorf("DefaultLanguageProfile() ID = %q, want %q", profile.ID, "default")
	}

	if profile.Name != "Default English" {
		t.Errorf("DefaultLanguageProfile() Name = %q, want %q", profile.Name, "Default English")
	}

	if !profile.IsDefault {
		t.Error("DefaultLanguageProfile() IsDefault = false, want true")
	}

	if len(profile.Languages) != 1 {
		t.Errorf("DefaultLanguageProfile() Languages count = %d, want 1", len(profile.Languages))
	}

	if profile.Languages[0].Language != "en" {
		t.Errorf("DefaultLanguageProfile() Languages[0].Language = %q, want %q", profile.Languages[0].Language, "en")
	}

	if profile.CutoffScore != 75 {
		t.Errorf("DefaultLanguageProfile() CutoffScore = %d, want 75", profile.CutoffScore)
	}

	// Test validation
	if err := profile.Validate(); err != nil {
		t.Errorf("DefaultLanguageProfile() failed validation: %v", err)
	}
}

func TestValidationError(t *testing.T) {
	err := &ValidationError{
		Field:   "test_field",
		Message: "test message",
		Index:   5,
	}

	expected := "validation error in test_field[5]: test message"
	if err.Error() != expected {
		t.Errorf("ValidationError.Error() = %q, want %q", err.Error(), expected)
	}

	// Test without index
	err.Index = -1
	expected = "validation error in test_field: test message"
	if err.Error() != expected {
		t.Errorf("ValidationError.Error() = %q, want %q", err.Error(), expected)
	}
}