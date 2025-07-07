// file: cmd/tag_test.go
// version: 1.0.0
// guid: 1bcf09a8-2b3d-4a1d-9473-73f63af2dd14
package cmd

import (
	"bytes"
	"flag"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"

	"github.com/jdfalk/subtitle-manager/pkg/database"
	"github.com/jdfalk/subtitle-manager/pkg/testutil"
)

// TestTagCommands verifies tag add, list, and remove operations.
func TestTagCommands(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if err := testutil.CheckSQLiteSupport(); err != nil {
		t.Skipf("SQLite support not available: %v", err)
	}

	dir := t.TempDir()
	dbPath := dir + "/test.db"
	viper.Set("db_path", dbPath)
	defer viper.Reset()

	store, err := database.OpenStore(dbPath, "sqlite")
	if err != nil {
		t.Fatalf("open store: %v", err)
	}
	store.Close()

	// Add tag
	tagTypeFlag = "user"
	tagEntityType = "all"
	tagColor = "#ff0000"
	tagDesc = "test"
	if err := tagAddCmd.RunE(tagAddCmd, []string{"demo"}); err != nil {
		t.Fatalf("tag add: %v", err)
	}

	// List tags
	buf := &bytes.Buffer{}
	tagListCmd.SetOut(buf)
	if err := tagListCmd.RunE(tagListCmd, nil); err != nil {
		t.Fatalf("tag list: %v", err)
	}
	output := buf.String()
	if !strings.Contains(output, "demo") {
		t.Fatalf("expected tag in list")
	}

	// Remove tag
	if err := tagRemoveCmd.RunE(tagRemoveCmd, []string{"1"}); err != nil {
		t.Fatalf("tag remove: %v", err)
	}
}
