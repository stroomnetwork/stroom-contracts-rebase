package stroomabi

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"testing"
)

const (
	tmpAbiDir       = "../build/tmp/stroomabi/"
	committedAbiDir = "./"
)

// List of files to test for ABI compatibility
var filesToTest = []string{
	"strbtc.go",
	"wstrbtc.go",
	"user_activator.go",
	"validator_registry.go",
	"proxy.go",
	"bitcoin_utils_wrapper.go",
}

// TestAbiCompatibility verifies that the committed Go bindings for the Solidity smart contracts are up-to-date.
func TestAbiCompatibility(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping contract tests in short mode because it requires ABI generation step")
	}

	for _, file := range filesToTest {
		t.Run(file, func(t *testing.T) {
			committedChecksum := readFileChecksum(t, committedAbiDir+file)
			tmpChecksum := readFileChecksum(t, tmpAbiDir+file)

			if committedChecksum != tmpChecksum {
				t.Errorf(
					"%s check failed. Looks like you changed `*.sol` files but forgot to generate ABI? "+
						"Run `make go-gen` and commit the updated files.",
					file,
				)
			}
		})
	}
}

// Helper function to calculate the checksum of a file
func readFileChecksum(t *testing.T, file string) string {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", file, err)
	}
	hash := sha256.Sum256(fileBytes)
	return hex.EncodeToString(hash[:])
}
