---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_sentinel_policy"
description: |-
  Manages Sentinel policies.
---


<!-- Please do not edit this file, it is generated. -->
# tfe_sentinel_policy

Sentinel Policy as Code is an embedded policy as code framework integrated
with Terraform Enterprise.

Policies are configured on a per-organization level and are organized and
grouped into policy sets, which define the workspaces on which policies are
enforced during runs.

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
import { SentinelPolicy } from "./.gen/providers/tfe/sentinel-policy";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new SentinelPolicy(this, "test", {
      description: "This policy always passes",
      enforceMode: "hard-mandatory",
      name: "my-policy-name",
      organization: "my-org-name",
      policy: "main = rule { true }",
    });
  }
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the policy.
* `description` - (Optional) A description of the policy's purpose.
* `organization` - (Optional) Name of the organization. If omitted, organization must be defined in the provider config.
* `policy` - (Required) The actual policy itself.
* `enforceMode` - (Optional) The enforcement level of the policy. Valid
  values are `advisory`, `hard-mandatory` and `soft-mandatory`. Defaults
  to `soft-mandatory`.

## Attributes Reference

* `id` - The ID of the policy.

## Import

Sentinel policies can be imported; use `<ORGANIZATION NAME>/<POLICY ID>` as the
import ID. For example:

```shell
terraform import tfe_sentinel_policy.test my-org-name/pol-wAs3zYmWAhYK7peR
```

<!-- cache-key: cdktf-0.20.8 input-e8ffab24beacc256a0ab1c1b8664c815d96f0c934f344a8cf4867f273b7a4ab1 -->