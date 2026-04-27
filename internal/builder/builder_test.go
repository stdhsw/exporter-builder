// Package builder
package builder

import (
	"testing"

	"github.com/stdhsw/exporter-builder/internal/config"
)

func TestGenerateExporter(t *testing.T) {
	cfg := config.NewConfig()
	if err := GenerateExporter(cfg); err != nil {
		t.Errorf("GenerateExporter() err = %v", err)
	}
}
