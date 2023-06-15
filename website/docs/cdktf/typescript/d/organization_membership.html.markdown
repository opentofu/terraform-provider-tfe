---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_organization_membership"
description: |-
  Get information on an organization membership.
---

# Data Source: tfe_organization_membership

Use this data source to get information about an organization membership.

~> **NOTE:** This data source requires using the provider with Terraform Cloud or
an instance of Terraform Enterprise at least as recent as v202004-1.

~> **NOTE:** If a user updates their email address, configurations using the email address should
be updated manually.

## Example Usage

### Fetch by email

```typescript
import * as constructs from "constructs";
import * as cdktf from "cdktf";
/*Provider bindings are generated by running cdktf get.
See https://cdk.tf/provider-generation for more details.*/
import * as tfe from "./.gen/providers/tfe";
class MyConvertedCode extends cdktf.TerraformStack {
  constructor(scope: constructs.Construct, name: string) {
    super(scope, name);
    new tfe.dataTfeOrganizationMembership.DataTfeOrganizationMembership(
      this,
      "test",
      {
        email: "user@company.com",
        organization: "my-org-name",
      }
    );
  }
}

```

### Fetch by username

```typescript
import * as constructs from "constructs";
import * as cdktf from "cdktf";
/*Provider bindings are generated by running cdktf get.
See https://cdk.tf/provider-generation for more details.*/
import * as tfe from "./.gen/providers/tfe";
class MyConvertedCode extends cdktf.TerraformStack {
  constructor(scope: constructs.Construct, name: string) {
    super(scope, name);
    new tfe.dataTfeOrganizationMembership.DataTfeOrganizationMembership(
      this,
      "test",
      {
        organization: "my-org-name",
        username: "my-username",
      }
    );
  }
}

```

## Argument Reference

The following arguments are supported:

* `organization` - (Required) Name of the organization.
* `email` - (Optional) Email of the user.
* `username` - (Optional) The username of the user.

~> **NOTE:** While `email` and `username` are optional arguments, one or the other is required.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The organization membership ID.
* `userId` - The ID of the user associated with the organization membership.
* `username` - The username of the user associated with the organization membership.

<!-- cache-key: cdktf-0.17.0-pre.15 input-4a14b3c66d2d0654b6af772796f3c1bdf8871760782092fd770887fd1a21bb84 -->