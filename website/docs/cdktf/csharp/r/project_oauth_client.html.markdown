---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_project_oauth_client"
description: |-
    Add an oauth client to a project
---


<!-- Please do not edit this file, it is generated. -->
# tfe_project_oauth_client

Adds and removes oauth clients from a project

## Example Usage

Basic usage:

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
        var tfeOrganizationTest = new Organization.Organization(this, "test", new OrganizationConfig {
            Email = "admin@company.com",
            Name = "my-org-name"
        });
        var tfeProjectTest = new Project.Project(this, "test_1", new ProjectConfig {
            Name = "my-project-name",
            Organization = Token.AsString(tfeOrganizationTest.Name)
        });
        /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
        tfeProjectTest.OverrideLogicalId("test");
        var tfeOauthClientTest = new OauthClient.OauthClient(this, "test_2", new OauthClientConfig {
            ApiUrl = "https://api.github.com",
            HttpUrl = "https://github.com",
            OauthToken = "oauth_token_id",
            Organization = tfeOrganizationTest,
            ServiceProvider = "github"
        });
        /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
        tfeOauthClientTest.OverrideLogicalId("test");
        var tfeProjectOauthClientTest =
        new ProjectOauthClient.ProjectOauthClient(this, "test_3", new ProjectOauthClientConfig {
            OauthClientId = Token.AsString(tfeOauthClientTest.Id),
            ProjectId = Token.AsString(tfeProjectTest.Id)
        });
        /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
        tfeProjectOauthClientTest.OverrideLogicalId("test");
    }
}
```

## Argument Reference

The following arguments are supported:

* `OauthClientId` - (Required) ID of the oauth client.
* `ProjectId` - (Required) Project ID to add the oauth client to.

## Attributes Reference

* `Id` - The ID of the oauth client attachment. ID format: `<project-id>_<oauth-client-id>`

## Import

Project OAuth Clients can be imported; use `<ORGANIZATION>/<PROJECT ID>/<OAUTH CLIENT NAME>`. For example:

```shell
terraform import tfe_project_oauth_client.test 'my-org-name/project/oauth-client-name'
```

<!-- cache-key: cdktf-0.17.0-pre.15 input-74e716b96a2f11693107d16605160ebd3c1858710f2cf85569428398a04b4072 -->