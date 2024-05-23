package test

import (
	"os"
	"testing"
)

func TestDeleteProject(t *testing.T) {
	err := os.RemoveAll("consumer\\file\\MoauntEverest\\")
	if err != nil {
		t.Fatalf("Error create function: %v", err)
	}
}