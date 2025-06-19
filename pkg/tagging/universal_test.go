// file: pkg/tagging/universal_test.go
package tagging

import (
	"strconv"
	"testing"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/testutil"
)

// MockEntity implements TaggedEntity for testing.
type MockEntity struct {
	ID   string
	Type string
	Tags []string
}

func (m *MockEntity) GetEntityType() string {
	return m.Type
}

func (m *MockEntity) GetEntityID() string {
	return m.ID
}

func (m *MockEntity) GetTags() []string {
	return m.Tags
}

func (m *MockEntity) SetTags(tags []string) {
	m.Tags = tags
}

func TestTagManager_CreateTag(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	tm := NewTagManager(db)

	// Test creating a basic tag
	tag, err := tm.CreateTag("test-tag", "user", "media", "#FF0000", "Test tag description")
	if err != nil {
		t.Fatalf("Failed to create tag: %v", err)
	}

	if tag.Name != "test-tag" {
		t.Errorf("Expected tag name 'test-tag', got '%s'", tag.Name)
	}
	if tag.Type != "user" {
		t.Errorf("Expected tag type 'user', got '%s'", tag.Type)
	}
	if tag.EntityType != "media" {
		t.Errorf("Expected entity type 'media', got '%s'", tag.EntityType)
	}
	if tag.Color != "#FF0000" {
		t.Errorf("Expected color '#FF0000', got '%s'", tag.Color)
	}
	if tag.Description != "Test tag description" {
		t.Errorf("Expected description 'Test tag description', got '%s'", tag.Description)
	}

	// Test creating tag with defaults
	tag2, err := tm.CreateTag("default-tag", "", "", "", "")
	if err != nil {
		t.Fatalf("Failed to create default tag: %v", err)
	}

	if tag2.Type != "user" {
		t.Errorf("Expected default type 'user', got '%s'", tag2.Type)
	}
	if tag2.EntityType != "all" {
		t.Errorf("Expected default entity type 'all', got '%s'", tag2.EntityType)
	}
}

func TestTagManager_EntityTagging(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	tm := NewTagManager(db)

	// Create test tag
	tag, err := tm.CreateTag("entity-tag", "user", "all", "", "")
	if err != nil {
		t.Fatalf("Failed to create tag: %v", err)
	}

	tagID, err := strconv.ParseInt(tag.ID, 10, 64)
	if err != nil {
		t.Fatalf("Failed to parse tag ID: %v", err)
	}

	// Create mock entity
	entity := &MockEntity{
		ID:   "test-entity-1",
		Type: "provider",
		Tags: []string{},
	}

	// Test tagging entity
	err = tm.TagEntity(tagID, entity)
	if err != nil {
		t.Fatalf("Failed to tag entity: %v", err)
	}

	// Verify tag association
	tags, err := tm.GetEntityTags(entity)
	if err != nil {
		t.Fatalf("Failed to get entity tags: %v", err)
	}

	if len(tags) != 1 {
		t.Fatalf("Expected 1 tag, got %d", len(tags))
	}

	if tags[0].Name != "entity-tag" {
		t.Errorf("Expected tag name 'entity-tag', got '%s'", tags[0].Name)
	}

	// Test untagging entity
	err = tm.UntagEntity(tagID, entity)
	if err != nil {
		t.Fatalf("Failed to untag entity: %v", err)
	}

	// Verify tag removal
	tags, err = tm.GetEntityTags(entity)
	if err != nil {
		t.Fatalf("Failed to get entity tags after removal: %v", err)
	}

	if len(tags) != 0 {
		t.Errorf("Expected 0 tags after removal, got %d", len(tags))
	}
}

func TestTagManager_BulkTagging(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	tm := NewTagManager(db)

	// Create test tags
	tag1, err := tm.CreateTag("bulk-tag-1", "user", "all", "", "")
	if err != nil {
		t.Fatalf("Failed to create tag 1: %v", err)
	}

	tag2, err := tm.CreateTag("bulk-tag-2", "user", "all", "", "")
	if err != nil {
		t.Fatalf("Failed to create tag 2: %v", err)
	}

	tagID1, _ := strconv.ParseInt(tag1.ID, 10, 64)
	tagID2, _ := strconv.ParseInt(tag2.ID, 10, 64)

	// Create mock entities
	entities := []database.TaggedEntity{
		&MockEntity{ID: "bulk-entity-1", Type: "provider"},
		&MockEntity{ID: "bulk-entity-2", Type: "provider"},
		&MockEntity{ID: "bulk-entity-3", Type: "media"},
	}

	// Test bulk tagging
	err = tm.BulkTagEntities([]int64{tagID1, tagID2}, entities)
	if err != nil {
		t.Fatalf("Failed to bulk tag entities: %v", err)
	}

	// Verify each entity has both tags
	for i, entity := range entities {
		tags, err := tm.GetEntityTags(entity)
		if err != nil {
			t.Fatalf("Failed to get tags for entity %d: %v", i, err)
		}

		if len(tags) != 2 {
			t.Errorf("Expected 2 tags for entity %d, got %d", i, len(tags))
		}

		// Check tag names
		tagNames := make(map[string]bool)
		for _, tag := range tags {
			tagNames[tag.Name] = true
		}

		if !tagNames["bulk-tag-1"] || !tagNames["bulk-tag-2"] {
			t.Errorf("Entity %d missing expected tags", i)
		}
	}
}

func TestTagManager_FilterByTags(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	tm := NewTagManager(db)

	// Create test tags
	tag1, _ := tm.CreateTag("filter-tag-1", "user", "all", "", "")
	tag2, _ := tm.CreateTag("filter-tag-2", "user", "all", "", "")
	tag3, _ := tm.CreateTag("filter-tag-3", "user", "all", "", "")

	tagID1, _ := strconv.ParseInt(tag1.ID, 10, 64)
	tagID2, _ := strconv.ParseInt(tag2.ID, 10, 64)
	tagID3, _ := strconv.ParseInt(tag3.ID, 10, 64)

	// Create entities with different tag combinations
	entity1 := &MockEntity{ID: "filter-entity-1", Type: "media"}
	entity2 := &MockEntity{ID: "filter-entity-2", Type: "media"}
	entity3 := &MockEntity{ID: "filter-entity-3", Type: "media"}

	// Entity 1: tag1, tag2
	_ = tm.TagEntity(tagID1, entity1)
	_ = tm.TagEntity(tagID2, entity1)

	// Entity 2: tag1, tag3
	_ = tm.TagEntity(tagID1, entity2)
	_ = tm.TagEntity(tagID3, entity2)

	// Entity 3: tag2, tag3
	_ = tm.TagEntity(tagID2, entity3)
	_ = tm.TagEntity(tagID3, entity3)

	// Test filtering by single tag
	entities, err := tm.FilterByTags("media", []string{"filter-tag-1"})
	if err != nil {
		t.Fatalf("Failed to filter by single tag: %v", err)
	}

	expectedSingle := map[string]bool{"filter-entity-1": true, "filter-entity-2": true}
	if len(entities) != 2 {
		t.Errorf("Expected 2 entities for single tag filter, got %d", len(entities))
	}
	for _, entityID := range entities {
		if !expectedSingle[entityID] {
			t.Errorf("Unexpected entity in single tag filter: %s", entityID)
		}
	}

	// Test filtering by multiple tags (AND logic)
	entities, err = tm.FilterByTags("media", []string{"filter-tag-1", "filter-tag-2"})
	if err != nil {
		t.Fatalf("Failed to filter by multiple tags: %v", err)
	}

	if len(entities) != 1 || entities[0] != "filter-entity-1" {
		t.Errorf("Expected only filter-entity-1 for AND filter, got %v", entities)
	}
}

func TestTagManager_GetTagsByType(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	tm := NewTagManager(db)

	// Create tags of different types
	_, err := tm.CreateTag("system-tag", "system", "all", "", "")
	if err != nil {
		t.Fatalf("Failed to create system tag: %v", err)
	}

	_, err = tm.CreateTag("user-tag", "user", "media", "", "")
	if err != nil {
		t.Fatalf("Failed to create user tag: %v", err)
	}

	_, err = tm.CreateTag("custom-tag", "custom", "provider", "", "")
	if err != nil {
		t.Fatalf("Failed to create custom tag: %v", err)
	}

	// Test getting system tags
	systemTags, err := tm.GetTagsByType("system")
	if err != nil {
		t.Fatalf("Failed to get system tags: %v", err)
	}

	if len(systemTags) != 1 || systemTags[0].Name != "system-tag" {
		t.Errorf("Expected 1 system tag named 'system-tag', got %v", systemTags)
	}

	// Test getting user tags
	userTags, err := tm.GetTagsByType("user")
	if err != nil {
		t.Fatalf("Failed to get user tags: %v", err)
	}

	if len(userTags) != 1 || userTags[0].Name != "user-tag" {
		t.Errorf("Expected 1 user tag named 'user-tag', got %v", userTags)
	}

	// Test getting custom tags
	customTags, err := tm.GetTagsByType("custom")
	if err != nil {
		t.Fatalf("Failed to get custom tags: %v", err)
	}

	if len(customTags) != 1 || customTags[0].Name != "custom-tag" {
		t.Errorf("Expected 1 custom tag named 'custom-tag', got %v", customTags)
	}
}

func TestTagManager_GetEntitiesWithTag(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	tm := NewTagManager(db)

	// Create test tag
	tag, err := tm.CreateTag("entity-search-tag", "user", "all", "", "")
	if err != nil {
		t.Fatalf("Failed to create tag: %v", err)
	}

	tagID, _ := strconv.ParseInt(tag.ID, 10, 64)

	// Create entities and tag some of them
	providerEntities := []*MockEntity{
		{ID: "provider-1", Type: "provider"},
		{ID: "provider-2", Type: "provider"},
		{ID: "provider-3", Type: "provider"},
	}

	mediaEntities := []*MockEntity{
		{ID: "media-1", Type: "media"},
		{ID: "media-2", Type: "media"},
	}

	// Tag some providers and media
	_ = tm.TagEntity(tagID, providerEntities[0])
	_ = tm.TagEntity(tagID, providerEntities[2])
	_ = tm.TagEntity(tagID, mediaEntities[0])

	// Get providers with tag
	providerIDs, err := tm.GetEntitiesWithTag(tagID, "provider")
	if err != nil {
		t.Fatalf("Failed to get providers with tag: %v", err)
	}

	expectedProviders := []string{"provider-1", "provider-3"}
	if len(providerIDs) != 2 {
		t.Errorf("Expected 2 providers with tag, got %d", len(providerIDs))
	}

	for i, expectedID := range expectedProviders {
		if i >= len(providerIDs) || providerIDs[i] != expectedID {
			t.Errorf("Expected provider %s at index %d, got %s", expectedID, i, providerIDs[i])
		}
	}

	// Get media with tag
	mediaIDs, err := tm.GetEntitiesWithTag(tagID, "media")
	if err != nil {
		t.Fatalf("Failed to get media with tag: %v", err)
	}

	if len(mediaIDs) != 1 || mediaIDs[0] != "media-1" {
		t.Errorf("Expected media-1 with tag, got %v", mediaIDs)
	}
}

func TestTagManager_GetTagByName(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	tm := NewTagManager(db)

	// Create test tag
	originalTag, err := tm.CreateTag("named-tag", "custom", "media", "#00FF00", "Named tag for testing")
	if err != nil {
		t.Fatalf("Failed to create tag: %v", err)
	}

	// Retrieve tag by name
	retrievedTag, err := tm.GetTagByName("named-tag")
	if err != nil {
		t.Fatalf("Failed to get tag by name: %v", err)
	}

	// Verify all fields match
	if retrievedTag.ID != originalTag.ID {
		t.Errorf("Expected ID %s, got %s", originalTag.ID, retrievedTag.ID)
	}
	if retrievedTag.Name != originalTag.Name {
		t.Errorf("Expected name %s, got %s", originalTag.Name, retrievedTag.Name)
	}
	if retrievedTag.Type != originalTag.Type {
		t.Errorf("Expected type %s, got %s", originalTag.Type, retrievedTag.Type)
	}
	if retrievedTag.EntityType != originalTag.EntityType {
		t.Errorf("Expected entity type %s, got %s", originalTag.EntityType, retrievedTag.EntityType)
	}
	if retrievedTag.Color != originalTag.Color {
		t.Errorf("Expected color %s, got %s", originalTag.Color, retrievedTag.Color)
	}
	if retrievedTag.Description != originalTag.Description {
		t.Errorf("Expected description %s, got %s", originalTag.Description, retrievedTag.Description)
	}

	// Test non-existent tag
	_, err = tm.GetTagByName("non-existent-tag")
	if err == nil {
		t.Error("Expected error for non-existent tag, got nil")
	}
}
