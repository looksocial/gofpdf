package main

import (
    "os/exec"
    "strings"
    "testing"
)

// TestHelpFlag verifies that the help flag prints usage information.
func TestHelpFlag(t *testing.T) {
    // Ensure binary is built
    if err := exec.Command("go", "build").Run(); err != nil {
        t.Fatalf("build failed: %v", err)
    }

    out, err := exec.Command("./makefont", "--help").CombinedOutput()
    if err != nil {
        // help should not fail; but still validate output
        t.Logf("help returned error: %v", err)
    }
    got := string(out)
    if !strings.Contains(got, "Usage:") {
        t.Fatalf("expected help usage in output, got: %s", got)
    }
}

// TestNoArgs verifies that running without args prints an error/usage hint.
func TestNoArgs(t *testing.T) {
    // Ensure binary is built
    if err := exec.Command("go", "build").Run(); err != nil {
        t.Fatalf("build failed: %v", err)
    }

    out, err := exec.Command("./makefont").CombinedOutput()
    // even if error is non-nil, we only assert output contains guidance
    _ = err
    got := string(out)
    if !strings.Contains(got, "At least one Type1 or TrueType font must be specified") {
        t.Fatalf("expected missing-args message, got: %s", got)
    }
}


