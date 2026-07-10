package bindings

import (
	"strings"
	"testing"
)

func TestResolveDatabaseAndQueue(t *testing.T) {
	registry := Registry{Dependencies: map[string]Entry{
		"main": {Type: "rds-postgres", Module: "database"},
		"jobs": {Type: "sqs", Module: "queues", Key: "jobs"},
	}}
	result, err := Resolve(registry, map[string]Request{
		"database": {Type: "rds-postgres", Reference: "main", Env: map[string]string{"DATABASE_HOST": "endpoint"}},
		"queue":    {Type: "sqs", Reference: "jobs", Env: map[string]string{"QUEUE_URL": "queue_url"}},
	})
	if err != nil {
		t.Fatalf("resolve: %v", err)
	}
	if result.Environment["DATABASE_HOST"] != "module.database.endpoint" {
		t.Fatalf("database reference = %q", result.Environment["DATABASE_HOST"])
	}
	if result.Environment["QUEUE_URL"] != `module.queues.queue_urls["jobs"]` {
		t.Fatalf("queue reference = %q", result.Environment["QUEUE_URL"])
	}
}

func TestUnknownDependencyFails(t *testing.T) {
	_, err := Resolve(Registry{Dependencies: map[string]Entry{}}, map[string]Request{
		"database": {Type: "rds-postgres", Reference: "missing"},
	})
	if err == nil || !strings.Contains(err.Error(), "unknown registry entry") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestSecretBindingRendersOnlyComment(t *testing.T) {
	result, err := Resolve(Registry{Dependencies: map[string]Entry{"main": {Type: "rds-postgres"}}}, map[string]Request{
		"database": {Type: "rds-postgres", Reference: "main", Secrets: map[string]string{"DATABASE_PASSWORD": "password"}},
	})
	if err != nil {
		t.Fatalf("resolve: %v", err)
	}
	if len(result.Environment) != 0 || len(result.Comments) != 1 || strings.Contains(result.Comments[0], "plaintext") {
		t.Fatalf("result = %#v", result)
	}
}
