---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_slug"
description: |-
  Manages files.
---
# Data Source: tfe_slug

This data source is used to represent configuration files on a local filesystem
intended to be uploaded to Terraform Cloud/Enterprise, in lieu of those files being
sourced from a configured VCS provider.

A unique checksum is generated for the specified local directory, which allows
resources such as `tfePolicySet` track the files and upload a new gzip compressed
tar file containing configuration files (a Terraform "slug") when those files change.

## Example Usage

Tracking a local directory to upload the Sentinel configuration and policies:

```java
import software.constructs.*;
import com.hashicorp.cdktf.*;
/*Provider bindings are generated by running cdktf get.
See https://cdk.tf/provider-generation for more details.*/
import gen.providers.tfe.dataTfeSlug.*;
import gen.providers.tfe.policySet.*;
public class MyConvertedCode extends TerraformStack {
    public MyConvertedCode(Construct scope, String name) {
        super(scope, name);
        DataTfeSlug dataTfeSlugTest = new DataTfeSlug(this, "test", new DataTfeSlugConfig()
                .sourcePath("policies/my-policy-set")
                );
        PolicySet tfePolicySetTest = new PolicySet(this, "test_1", new PolicySetConfig()
                .name("my-policy-set")
                .organization("my-org-name")
                .slug(Token.asStringMap(dataTfeSlugTest))
                );
        /*This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.*/
        tfePolicySetTest.overrideLogicalId("test");
    }
}
```

## Argument Reference

The following arguments are supported:

* `sourcePath` - (Required) The path to the directory where the files are located.

<!-- cache-key: cdktf-0.17.0-pre.15 input-8546d8f7537661b4b2d4d594c4cb0f6ebd1b0e70ae9e1bcc6e145f33bd763e70 -->