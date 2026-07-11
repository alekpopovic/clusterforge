package generator

import "fmt"

func RenderTerraformCloudBackend(organization, workspace string) (string, error) {
	if organization == "" {
		return "", fmt.Errorf("Terraform Cloud organization is required")
	}
	if workspace == "" {
		return "", fmt.Errorf("Terraform Cloud workspace is required")
	}
	return fmt.Sprintf("%s\n\nterraform {\n  cloud {\n    organization = %q\n\n    workspaces {\n      name = %q\n    }\n  }\n}\n", header, organization, workspace), nil
}
