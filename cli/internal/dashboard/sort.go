package dashboard

import "sort"

func sortExport(out *Export) {
	sort.Slice(out.Environments, func(i, j int) bool { return out.Environments[i].Name < out.Environments[j].Name })
	sort.Slice(out.Clusters, func(i, j int) bool { return out.Clusters[i].Name < out.Clusters[j].Name })
	sort.Slice(out.Apps, func(i, j int) bool { return out.Apps[i].Name < out.Apps[j].Name })
	sort.Slice(out.Runbooks, func(i, j int) bool { return out.Runbooks[i].Name < out.Runbooks[j].Name })
}
