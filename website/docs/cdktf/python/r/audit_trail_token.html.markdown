---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_audit_trail_token"
description: |-
  Generates a new audit trail token in organization, replacing any existing token.
---


<!-- Please do not edit this file, it is generated. -->
# tfe_audit_trail_token

Generates a new audit trail token in organization, replacing any existing token.

Note that only organizations that have the [audit-logging entitlement](https://developer.hashicorp.com/terraform/cloud-docs/api-docs#audit-logging) may create audit trail tokens.

## Example Usage

Basic usage:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.tfe.audit_trail_token import AuditTrailToken
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        AuditTrailToken(self, "test",
            organization="my-org-name"
        )
```

## Argument Reference

The following arguments are supported:

* `organization` - (Optional) Name of the organization. If omitted, organization must be defined in the provider config.
* `force_regenerate` - (Optional) If set to `true`, a new token will be
  generated even if a token already exists. This will invalidate the existing
  token!
* `expired_at` - (Optional) The token's expiration date. The expiration date must be a date/time string in RFC3339
format (e.g., "2024-12-31T23:59:59Z"). If no expiration date is supplied, the expiration date will default to null and
never expire.

## Example Usage

When a token has an expiry:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.tfe.audit_trail_token import AuditTrailToken
from imports.time.rotating import Rotating
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        # The following providers are missing schema information and might need manual adjustments to synthesize correctly: time.
        #     For a more precise conversion please use the --provider flag in convert.
        example = Rotating(self, "example",
            rotation_days=30
        )
        AuditTrailToken(self, "test",
            expired_at=Token.as_string(example.rotation_rfc3339),
            organization=Token.as_string(org.name)
        )
```

## Attributes Reference

* `id` - The ID of the token.
* `token` - The generated token.

## Import

Audit trail tokens can be imported; use `<ORGANIZATION NAME>` as the import ID.
For example:

```shell
terraform import tfe_audit_trail_token.test my-org-name
```

<!-- cache-key: cdktf-0.20.8 input-2c87d1201e89a4243d53e9ef85752b50f8797786174b1c7744ecc61fd3ced2cd -->