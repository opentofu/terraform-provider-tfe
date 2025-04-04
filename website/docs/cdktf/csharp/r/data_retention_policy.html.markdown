---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_data_retention_policy"
description: |-
  Manages data retention policies for organizations and workspaces
---


<!-- Please do not edit this file, it is generated. -->
# tfe_data_retention_policy

Creates a data retention policy attached to either an organization or workspace. This resource is for Terraform Enterprise only.

## Example Usage

Creating a data retention policy for a workspace:

```csharp
using Constructs;
using HashiCorp.Cdktf;
/*Provider bindings are generated by running cdktf get.
See https://cdk.tf/provider-generation for more details.*/
using Gen.Providers.Tfe;
class MyConvertedCode : TerraformStack
{
    public MyConvertedCode(Construct scope, string name) : base(scope, name)
    {
        var tfeOrganizationTestOrganization = new Organization.Organization(this, "test-organization", new OrganizationConfig {
            Email = "admin@company.com",
            Name = "my-org-name"
        });
        var tfeWorkspaceTestWorkspace = new Workspace.Workspace(this, "test-workspace", new WorkspaceConfig {
            Name = "my-workspace-name",
            Organization = Token.AsString(tfeOrganizationTestOrganization.Name)
        });
        new DataRetentionPolicy.DataRetentionPolicy(this, "foobar", new DataRetentionPolicyConfig {
            DeleteOlderThan = new [] { new Struct {
                Days = 42
            } },
            WorkspaceId = Token.AsString(tfeWorkspaceTestWorkspace.Id)
        });
    }
}
```

Creating a data retention policy for an organization:

```csharp
using Constructs;
using HashiCorp.Cdktf;
/*Provider bindings are generated by running cdktf get.
See https://cdk.tf/provider-generation for more details.*/
using Gen.Providers.Tfe;
class MyConvertedCode : TerraformStack
{
    public MyConvertedCode(Construct scope, string name) : base(scope, name)
    {
        var tfeOrganizationTestOrganization = new Organization.Organization(this, "test-organization", new OrganizationConfig {
            Email = "admin@company.com",
            Name = "my-org-name"
        });
        new DataRetentionPolicy.DataRetentionPolicy(this, "foobar", new DataRetentionPolicyConfig {
            DeleteOlderThan = new [] { new Struct {
                Days = 1138
            } },
            Organization = Token.AsString(tfeOrganizationTestOrganization.Name)
        });
    }
}
```

Creating a data retention policy for an organization and exclude a single workspace from it:

```hcl
resource "tfe_organization" "test-organization" {
  name  = "my-org-name"
  email = "admin@company.com"
}

// create data retention policy the organization
resource "tfe_data_retention_policy" "foobar" {
  organization = tfe_organization.test-organization.name

  delete_older_than {
    days = 1138
  }
}

resource "tfe_workspace" "test-workspace" {
  name         = "my-workspace-name"
  organization = tfe_organization.test-organization.name
}

// create a policy that prevents automatic deletion of data in the test-workspace
resource "tfe_data_retention_policy" "foobar" {
  workspace_id = tfe_workspace.test-workspace.id

  dont_delete {}
}
```

## Argument Reference

The following arguments are supported:

* `Organization` - (Optional) The name of the organization you want the policy to apply to. Must not be set if `WorkspaceId` is set.
* `WorkspaceId` - (Optional) The ID of the workspace you want the policy to apply to. Must not be set if `Organization` is set.
* `DeleteOlderThan` - (Optional) If this block is set, the created policy will apply to any data older than the configured number of days. Must not be set if `DontDelete` is set.
* `DontDelete` - (Optional) If this block is set, the created policy will prevent other policies from deleting data from this workspace or organization. Must not be set if `DeleteOlderThan` is set.


## Import

A resource can be imported; use `<ORGANIZATION>/<WORKSPACE NAME>` or `<ORGANIZATION>` as the import ID. For example:

```shell
terraform import tfe_data_retention_policy.foobar my-org-name/my-workspace-name
```

<!-- cache-key: cdktf-0.17.0-pre.15 input-f96ec458c2bca8796e296f06011ad4fc674072a5029841651f6906f053e79d76 -->