package safe

import (
	"testing"
)

func TestList(t *testing.T) {
	safe, err := createTestSafe()
	if err != nil {
		t.Error(err)
	}

	accounts := safe.List()
	if len(accounts) < 1 {
		t.Error(err)
	}
}
