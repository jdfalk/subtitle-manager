// file: pkg/tagging/universal.go
package tagging

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/jdfalk/subtitle-manager/pkg/database"
)

// TagManager provides universal tagging operations across all entity types.
// Implements a consistent interface for tagging media, users, providers, and other entities.
type TagManager struct {
	db *sql.DB
}

// NewTagManager creates a new universal tag manager.
func NewTagManager(db *sql.DB) *TagManager {
	return &TagManager{db: db}
}

// CreateTag creates a new tag with metadata support.
// tagType can be 'system', 'user', or 'custom'
// entityType can be 'all', 'media', 'user', 'provider', etc.
func (tm *TagManager) CreateTag(name, tagType, entityType, color, description string) (*database.Tag, error) {
	if tagType == "" {
		tagType = "user"
	}
	if entityType == "" {
		entityType = "all"
	}

	_, err := tm.db.Exec(`
		INSERT INTO tags (name, type, entity_type, color, description, created_at)
		VALUES (?, ?, ?, ?, ?, ?)`,
		name, tagType, entityType, color, description, time.Now())
	if err != nil {
		return nil, err
	}

	// Retrieve the created tag
	row := tm.db.QueryRow(`
		SELECT id, name, type, entity_type, color, description, created_at
		FROM tags WHERE name = ? ORDER BY id DESC LIMIT 1`, name)

	var tag database.Tag
	var id int64
	var nullColor, nullDescription sql.NullString
	err = row.Scan(&id, &tag.Name, &tag.Type, &tag.EntityType, &nullColor, &nullDescription, &tag.CreatedAt)
	if err != nil {
		return nil, err
	}

	tag.ID = strconv.FormatInt(id, 10)
	if nullColor.Valid {
		tag.Color = nullColor.String
	}
	if nullDescription.Valid {
		tag.Description = nullDescription.String
	}

	return &tag, nil
}

// TagEntity associates a tag with any entity type.
func (tm *TagManager) TagEntity(tagID int64, entity database.TaggedEntity) error {
	return database.AssignTagToEntity(tm.db, tagID, entity.GetEntityType(), entity.GetEntityID())
}

// UntagEntity removes a tag association from any entity type.
func (tm *TagManager) UntagEntity(tagID int64, entity database.TaggedEntity) error {
	return database.RemoveTagFromEntity(tm.db, tagID, entity.GetEntityType(), entity.GetEntityID())
}

// GetEntityTags retrieves all tags for a specific entity.
func (tm *TagManager) GetEntityTags(entity database.TaggedEntity) ([]database.Tag, error) {
	return database.ListTagsForEntity(tm.db, entity.GetEntityType(), entity.GetEntityID())
}

// GetTagByName retrieves a tag by its name.
func (tm *TagManager) GetTagByName(name string) (*database.Tag, error) {
	row := tm.db.QueryRow(`
		SELECT id, name, type, entity_type, color, description, created_at
		FROM tags WHERE name = ? LIMIT 1`, name)

	var tag database.Tag
	var id int64
	var nullColor, nullDescription sql.NullString
	err := row.Scan(&id, &tag.Name, &tag.Type, &tag.EntityType, &nullColor, &nullDescription, &tag.CreatedAt)
	if err != nil {
		return nil, err
	}

	tag.ID = strconv.FormatInt(id, 10)
	if nullColor.Valid {
		tag.Color = nullColor.String
	}
	if nullDescription.Valid {
		tag.Description = nullDescription.String
	}

	return &tag, nil
}

// GetTagsByType retrieves all tags of a specific type.
func (tm *TagManager) GetTagsByType(tagType string) ([]database.Tag, error) {
	rows, err := tm.db.Query(`
		SELECT id, name, type, entity_type, color, description, created_at
		FROM tags WHERE type = ? ORDER BY name`, tagType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []database.Tag
	for rows.Next() {
		var tag database.Tag
		var id int64
		var nullColor, nullDescription sql.NullString
		if err := rows.Scan(&id, &tag.Name, &tag.Type, &tag.EntityType, &nullColor, &nullDescription, &tag.CreatedAt); err != nil {
			return nil, err
		}
		tag.ID = strconv.FormatInt(id, 10)
		if nullColor.Valid {
			tag.Color = nullColor.String
		}
		if nullDescription.Valid {
			tag.Description = nullDescription.String
		}
		tags = append(tags, tag)
	}
	return tags, rows.Err()
}

// GetEntitiesWithTag retrieves all entity IDs of a specific type that have the given tag.
func (tm *TagManager) GetEntitiesWithTag(tagID int64, entityType string) ([]string, error) {
	rows, err := tm.db.Query(`
		SELECT entity_id FROM tag_associations
		WHERE tag_id = ? AND entity_type = ?
		ORDER BY entity_id`, tagID, entityType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []string
	for rows.Next() {
		var entityID string
		if err := rows.Scan(&entityID); err != nil {
			return nil, err
		}
		entities = append(entities, entityID)
	}
	return entities, rows.Err()
}

// BulkTagEntities assigns multiple tags to multiple entities efficiently.
func (tm *TagManager) BulkTagEntities(tagIDs []int64, entities []database.TaggedEntity) error {
	tx, err := tm.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	stmt, err := tx.Prepare(`
		INSERT OR IGNORE INTO tag_associations (tag_id, entity_type, entity_id, created_at)
		VALUES (?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now()
	for _, tagID := range tagIDs {
		for _, entity := range entities {
			if _, execErr := stmt.Exec(tagID, entity.GetEntityType(), entity.GetEntityID(), now); execErr != nil {
				return execErr
			}
		}
	}

	return tx.Commit()
}

// FilterByTags returns entity IDs that match all specified tags (AND logic).
func (tm *TagManager) FilterByTags(entityType string, tagNames []string) ([]string, error) {
	if len(tagNames) == 0 {
		return []string{}, nil
	}

	// Build query with multiple JOINs for AND logic
	query := `
		SELECT DISTINCT ta1.entity_id
		FROM tag_associations ta1
		JOIN tags t1 ON ta1.tag_id = t1.id
		WHERE ta1.entity_type = ? AND t1.name = ?`

	args := []interface{}{entityType, tagNames[0]}

	for i := 1; i < len(tagNames); i++ {
		query += `
		AND ta1.entity_id IN (
			SELECT ta` + strconv.Itoa(i+1) + `.entity_id
			FROM tag_associations ta` + strconv.Itoa(i+1) + `
			JOIN tags t` + strconv.Itoa(i+1) + ` ON ta` + strconv.Itoa(i+1) + `.tag_id = t` + strconv.Itoa(i+1) + `.id
			WHERE ta` + strconv.Itoa(i+1) + `.entity_type = ? AND t` + strconv.Itoa(i+1) + `.name = ?
		)`
		args = append(args, entityType, tagNames[i])
	}

	query += ` ORDER BY ta1.entity_id`

	rows, err := tm.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entities []string
	for rows.Next() {
		var entityID string
		if err := rows.Scan(&entityID); err != nil {
			return nil, err
		}
		entities = append(entities, entityID)
	}
	return entities, rows.Err()
}
