package graph

import (
	"strings"
	"testing"
)

func TestLogicalGraph(t *testing.T) {
	graph, err := Build("dev", "")
	if err != nil {
		t.Fatal(err)
	}
	if len(graph.Nodes) != 4 || len(graph.Edges) != 3 || !strings.Contains(graph.DOT(), `"network" -> "cluster"`) {
		t.Fatalf("graph = %#v\n%s", graph, graph.DOT())
	}
}

func TestStackGraphIncludesDependencies(t *testing.T) {
	graph, err := Build("dev", "platform")
	if err != nil {
		t.Fatal(err)
	}
	if got := strings.Join(graph.Nodes, ","); got != "network,cluster,platform" {
		t.Fatalf("nodes = %s", got)
	}
}

func TestUnknownStackFails(t *testing.T) {
	if _, err := Build("dev", "database"); err == nil {
		t.Fatal("expected unknown stack error")
	}
}
