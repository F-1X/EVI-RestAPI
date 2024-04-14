package postgresRepository_test

import (
	"os"
	"testing"
)

func TestCreateAd(t *testing.T) {
	urlStr := os.Getenv("POSTGRES_URL")
	if urlStr == "" {
		urlStr = "postgresql://postgres:postgres@localhost/postgres"
		const format = "env POSTGRES_URL is empty, used default value: %s"
		t.Logf(format, urlStr)
	}
}
