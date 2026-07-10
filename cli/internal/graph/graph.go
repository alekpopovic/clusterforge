package graph

import (
	"fmt"
	"strings"

	"github.com/textracta/clusterforge/cli/internal/config"
)

type Edge struct {
	From string
	To   string
}

type Logical struct {
	Environment string
	Nodes       []string
	Edges       []Edge
}

func Build(environment, selectedStack string) (Logical, error) {
	stacks := config.StackOrder()
	limit := len(stacks)
	if selectedStack != "" {
		limit = 0
		for index, stack := range stacks {
			if stack == selectedStack {
				limit = index + 1
				break
			}
		}
		if limit == 0 {
			return Logical{}, fmt.Errorf("unknown stack %q; expected network, cluster, platform, or apps", selectedStack)
		}
	}
	graph := Logical{Environment: environment, Nodes: append([]string{}, stacks[:limit]...)}
	for index := 1; index < len(graph.Nodes); index++ {
		graph.Edges = append(graph.Edges, Edge{From: graph.Nodes[index-1], To: graph.Nodes[index]})
	}
	return graph, nil
}

func (g Logical) DOT() string {
	var out strings.Builder
	fmt.Fprintf(&out, "digraph %q {\n", "clusterforge_"+g.Environment)
	out.WriteString("  rankdir=LR;\n")
	for _, node := range g.Nodes {
		fmt.Fprintf(&out, "  %q [label=%q];\n", node, node)
	}
	for _, edge := range g.Edges {
		fmt.Fprintf(&out, "  %q -> %q;\n", edge.From, edge.To)
	}
	out.WriteString("}\n")
	return out.String()
}

func (g Logical) Text() string {
	return fmt.Sprintf("environment %s: %s\n", g.Environment, strings.Join(g.Nodes, " -> "))
}
