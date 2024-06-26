---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_run_trigger"
description: |-
  Manages run triggers
---


<!-- Please do not edit this file, it is generated. -->
# tfe_run_trigger

HCP Terraform provides a way to connect your workspace to one or more workspaces within your organization, 
known as "source workspaces". These connections, called run triggers, allow runs to queue automatically in 
your workspace on successful apply of runs in any of the source workspaces. You can connect your workspace 
to up to 20 source workspaces.

## Example Usage

Basic usage:

```go
import constructs "github.com/aws/constructs-go/constructs"
import "github.com/hashicorp/terraform-cdk-go/cdktf"
/*Provider bindings are generated by running cdktf get.
See https://cdk.tf/provider-generation for more details.*/
import "github.com/aws-samples/dummy/gen/providers/tfe/organization"
import "github.com/aws-samples/dummy/gen/providers/tfe/workspace"
import "github.com/aws-samples/dummy/gen/providers/tfe/runTrigger"
type myConvertedCode struct {
	terraformStack
}

func newMyConvertedCode(scope construct, name *string) *myConvertedCode {
	this := &myConvertedCode{}
	cdktf.NewTerraformStack_Override(this, scope, name)
	tfeOrganizationTestOrganization := organization.NewOrganization(this, jsii.String("test-organization"), &organizationConfig{
		email: jsii.String("admin@company.com"),
		name: jsii.String("my-org-name"),
	})
	tfeWorkspaceTestSourceable := workspace.NewWorkspace(this, jsii.String("test-sourceable"), &workspaceConfig{
		name: jsii.String("my-sourceable-workspace-name"),
		organization: cdktf.Token_AsString(tfeOrganizationTestOrganization.id),
	})
	tfeWorkspaceTestWorkspace := workspace.NewWorkspace(this, jsii.String("test-workspace"), &workspaceConfig{
		name: jsii.String("my-workspace-name"),
		organization: cdktf.Token_*AsString(tfeOrganizationTestOrganization.id),
	})
	runTrigger.NewRunTrigger(this, jsii.String("test"), &runTriggerConfig{
		sourceableId: cdktf.Token_*AsString(tfeWorkspaceTestSourceable.id),
		workspaceId: cdktf.Token_*AsString(tfeWorkspaceTestWorkspace.id),
	})
	return this
}
```

## Argument Reference

The following arguments are supported:

* `WorkspaceId` - (Required) The id of the workspace that owns the run trigger. This is the 
  workspace where runs will be triggered.
* `SourceableId` - (Required) The id of the sourceable. The sourceable must be a workspace.

## Attributes Reference

* `Id` - The ID of the run trigger.

## Import

Run triggers can be imported; use `<RUN TRIGGER ID>` as the import ID. For example:

```shell
terraform import tfe_run_trigger.test rt-qV9JnKRkmtMa4zcA
```

<!-- cache-key: cdktf-0.17.0-pre.15 input-4cc4ec907e7f9bb7aeb8b049d62ceae16783fcb2b4d804b22d75b4f91ed5eaeb -->