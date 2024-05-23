package stubs

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/dynatrace/api"
)

func getStubsFolder() string {
	return strings.TrimSpace(os.Getenv("DYNATRACE_STUBS_FOLDER"))
}

func ShouldLoadStubs() bool {
	return getStubsFolder() != ""
}

func LoadStubs(name string) (api.Stubs, error) {
	filename := fmt.Sprintf("%s/%s.json", getStubsFolder(), name)
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open stubs file %s: %w", filename, err)
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read stubs file, error: %w", err)
	}

	stubs := api.Stubs{}
	if err := json.Unmarshal(jsonData, &stubs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal stubs file, error: %w", err)
	}

	return stubs, nil
}

func LoadStubsValue(name string, id string, v any) error {
	filename := fmt.Sprintf("%s/%s_%s.json", getStubsFolder(), name, id)
	jsonFile, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open stubs value file %s: %w", filename, err)
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return fmt.Errorf("failed to read stubs value file, error: %w", err)
	}

	if err := json.Unmarshal(jsonData, v); err != nil {
		return fmt.Errorf("failed to unmarshal stubs value file, error: %w", err)
	}

	return nil
}
