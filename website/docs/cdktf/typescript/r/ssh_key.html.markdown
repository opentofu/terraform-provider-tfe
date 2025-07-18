---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_ssh_key"
description: |-
  Manages SSH keys.
---


<!-- Please do not edit this file, it is generated. -->
# tfe_ssh_key

This resource represents an SSH key which includes a name and the SSH private
key. An organization can have multiple SSH keys available.

## Example Usage

Basic usage:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { SshKey } from "./.gen/providers/tfe/ssh-key";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SshKey(this, "test", {
      key: "private-ssh-key",
      name: "my-ssh-key-name",
      organization: "my-org-name",
    });
  }
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name to identify the SSH key.
* `organization` - (Optional) Name of the organization. If omitted, organization must be defined in the provider config.
* `key` - (Optional) The text of the SSH private key. One of `key` or `keyWo`
  must be provided.
* `keyWo` - (Optional, [Write-Only](https://developer.hashicorp.com/terraform/language/v1.11.x/resources/ephemeral#write-only-arguments)) The text of the SSH private key, guaranteed not to be
  written to plan or state artifacts. One of `key` or `keyWo` must be provided.

## Attributes Reference

* `id` The ID of the SSH key.

## Import

Because the Terraform Enterprise API does not return the private SSH key
content, this resource cannot be imported.

-> **Note:** Write-Only argument `keyWo` is available to use in place of `key`. Write-Only arguments are supported in HashiCorp Terraform 1.11.0 and later. [Learn more](https://developer.hashicorp.com/terraform/language/v1.11.x/resources/ephemeral#write-only-arguments).

<!-- cache-key: cdktf-0.20.8 input-4145429fa09f70dcbd1e53bbbe9cacef09123cd081cf03a8b4a30f26615f87e5 -->