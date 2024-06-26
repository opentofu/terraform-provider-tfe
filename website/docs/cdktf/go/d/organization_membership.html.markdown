---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_organization_membership"
description: |-
  Get information on an organization membership.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: tfe_organization_membership

Use this data source to get information about an organization membership.

~> **NOTE:** This data source requires using the provider with HCP Terraform or
an instance of Terraform Enterprise at least as recent as v202004-1.

~> **NOTE:** If a user updates their email address, configurations using the email address should
be updated manually.

## Example Usage

### Fetch by email

```go
import constructs "github.com/aws/constructs-go/constructs"
import cdktf "github.com/hashicorp/terraform-cdk-go/cdktf"
/*Provider bindings are generated by running cdktf get.
See https://cdk.tf/provider-generation for more details.*/
import "github.com/aws-samples/dummy/gen/providers/tfe/dataTfeOrganizationMembership"
type myConvertedCode struct {
	terraformStack
}

func newMyConvertedCode(scope construct, name *string) *myConvertedCode {
	this := &myConvertedCode{}
	cdktf.NewTerraformStack_Override(this, scope, name)
	dataTfeOrganizationMembership.NewDataTfeOrganizationMembership(this, jsii.String("test"), &dataTfeOrganizationMembershipConfig{
		email: jsii.String("user@company.com"),
		organization: jsii.String("my-org-name"),
	})
	return this
}
```

### Fetch by username

```go
import constructs "github.com/aws/constructs-go/constructs"
import cdktf "github.com/hashicorp/terraform-cdk-go/cdktf"
/*Provider bindings are generated by running cdktf get.
See https://cdk.tf/provider-generation for more details.*/
import "github.com/aws-samples/dummy/gen/providers/tfe/dataTfeOrganizationMembership"
type myConvertedCode struct {
	terraformStack
}

func newMyConvertedCode(scope construct, name *string) *myConvertedCode {
	this := &myConvertedCode{}
	cdktf.NewTerraformStack_Override(this, scope, name)
	dataTfeOrganizationMembership.NewDataTfeOrganizationMembership(this, jsii.String("test"), &dataTfeOrganizationMembershipConfig{
		organization: jsii.String("my-org-name"),
		username: jsii.String("my-username"),
	})
	return this
}
```

### Fetch by organization membership ID

```go
import constructs "github.com/aws/constructs-go/constructs"
import cdktf "github.com/hashicorp/terraform-cdk-go/cdktf"
/*Provider bindings are generated by running cdktf get.
See https://cdk.tf/provider-generation for more details.*/
import "github.com/aws-samples/dummy/gen/providers/tfe/dataTfeOrganizationMembership"
type myConvertedCode struct {
	terraformStack
}

func newMyConvertedCode(scope construct, name *string) *myConvertedCode {
	this := &myConvertedCode{}
	cdktf.NewTerraformStack_Override(this, scope, name)
	dataTfeOrganizationMembership.NewDataTfeOrganizationMembership(this, jsii.String("test"), &dataTfeOrganizationMembershipConfig{
		organization: jsii.String("my-org-name"),
		organizationMembershipId: jsii.String("ou-xxxxxxxxxxx"),
	})
	return this
}
```

## Argument Reference

The following arguments are supported:

* `Organization` - (Optional) Name of the organization.
* `Email` - (Optional) Email of the user.
* `Username` - (Optional) The username of the user.
* `OrganizationMembershipId` - (Optional) ID belonging to the organziation membership.

~> **NOTE:** While `Email` and `Username` are optional arguments, one or the other is required if `OrganizationMembershipId` argument is not provided.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `Id` - The organization membership ID.
* `UserId` - The ID of the user associated with the organization membership.
* `Username` - The username of the user associated with the organization membership.

<!-- cache-key: cdktf-0.17.0-pre.15 input-1069b52dde4b0bd38635c64ace6287da1da5393d94556544ac6f7883ba353f34 -->