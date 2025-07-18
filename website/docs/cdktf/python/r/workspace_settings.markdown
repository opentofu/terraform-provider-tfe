---
layout: "tfe"
page_title: "Terraform Enterprise: tfe_workspace_setting"
description: |-
  Manages workspace settings.
---


<!-- Please do not edit this file, it is generated. -->
# tfe_workspace_settings

~> **NOTE:** Manages or reads execution mode and agent pool settings for a workspace. This also interacts with the organization's default values for several settings, which can be managed with [tfe_organization_default_settings](organization_default_settings.html). If other resources need to identify whether a setting is a default or an explicit value set for the workspace, you can refer to the read-only `overwrites` argument.

~> **NOTE:** This resource manages values that can alternatively be managed by the  `tfe_workspace` resource. You should not attempt to manage the same property on both resources which could cause a permanent drift. Example properties available on both resources: `description`, `tags`, `auto_apply`, etc.

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
from imports.tfe.organization import Organization
from imports.tfe.workspace import Workspace
from imports.tfe.workspace_settings import WorkspaceSettings
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        test_organization = Organization(self, "test-organization",
            email="admin@company.com",
            name="my-org-name"
        )
        test = Workspace(self, "test",
            name="my-workspace-name",
            organization=test_organization.name
        )
        WorkspaceSettings(self, "test-settings",
            execution_mode="local",
            workspace_id=test.id
        )
```

With `execution_mode` of `agent`:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.tfe.agent_pool import AgentPool
from imports.tfe.agent_pool_allowed_workspaces import AgentPoolAllowedWorkspaces
from imports.tfe.organization import Organization
from imports.tfe.workspace import Workspace
from imports.tfe.workspace_settings import WorkspaceSettings
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        test_organization = Organization(self, "test-organization",
            email="admin@company.com",
            name="my-org-name"
        )
        test = Workspace(self, "test",
            name="my-workspace-name",
            organization=test_organization.name
        )
        test_agent_pool = AgentPool(self, "test-agent-pool",
            name="my-agent-pool-name",
            organization=test_organization.name
        )
        tfe_agent_pool_allowed_workspaces_test = AgentPoolAllowedWorkspaces(self, "test_3",
            agent_pool_id=test_agent_pool.id,
            allowed_workspace_ids=[test.id]
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        tfe_agent_pool_allowed_workspaces_test.override_logical_id("test")
        WorkspaceSettings(self, "test-settings",
            agent_pool_id=Token.as_string(tfe_agent_pool_allowed_workspaces_test.agent_pool_id),
            execution_mode="agent",
            workspace_id=test.id
        )
```

Using `remote_state_consumer_ids`:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformIterator, Op, conditional, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.tfe.workspace import Workspace
from imports.tfe.workspace_settings import WorkspaceSettings
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        # In most cases loops should be handled in the programming language context and
        #     not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
        #     you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
        #     you need to keep this like it is.
        test_for_each_iterator = TerraformIterator.from_list(
            Token.as_any(Fn.toset(["qa", "production"])))
        test = Workspace(self, "test",
            name="${" + test_for_each_iterator.value + "}-test",
            for_each=test_for_each_iterator
        )
        # In most cases loops should be handled in the programming language context and
        #     not inside of the Terraform context. If you are looping over something external, e.g. a variable or a file input
        #     you should consider using a for loop. If you are looping over something only known to Terraform, e.g. a result of a data source
        #     you need to keep this like it is.
        test_settings_for_each_iterator = TerraformIterator.from_list(
            Token.as_any(Fn.toset(["qa", "production"])))
        WorkspaceSettings(self, "test-settings",
            global_remote_state=False,
            remote_state_consumer_ids=Token.as_list(
                Fn.toset(
                    Fn.compact(
                        Token.as_list([
                            conditional(
                                Op.eq(test_settings_for_each_iterator.value, "production"),
                                Fn.lookup_nested(test, ["\"qa\"", "id"]), "")
                        ])))),
            workspace_id=Token.as_string(
                Fn.lookup_nested(
                    Fn.lookup_nested(test, [test_settings_for_each_iterator.value]), ["id"])),
            for_each=test_settings_for_each_iterator
        )
```

This resource may be used as a data source when no optional arguments are defined:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Token, TerraformOutput, Op, Fn, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.tfe.data_tfe_workspace import DataTfeWorkspace
from imports.tfe.workspace_settings import WorkspaceSettings
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        test = DataTfeWorkspace(self, "test",
            name="my-workspace-name",
            organization="my-org-name"
        )
        tfe_workspace_settings_test = WorkspaceSettings(self, "test_1",
            workspace_id=Token.as_string(test.id)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        tfe_workspace_settings_test.override_logical_id("test")
        TerraformOutput(self, "workspace-explicit-local-execution",
            value=Fn.alltrue(
                Token.as_any([
                    Op.eq(tfe_workspace_settings_test.execution_mode, "local"),
                    Fn.lookup_nested(tfe_workspace_settings_test.overwrites, ["0", "\"execution_mode\""
                    ])
                ]))
        )
```

This resource can be used to self manage a workspace created from `terraform init` and a cloud block:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import Fn, Token, TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.tfe.data_tfe_workspace import DataTfeWorkspace
from imports.tfe.workspace_settings import WorkspaceSettings
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        self = DataTfeWorkspace(self, "self",
            name=Token.as_string(
                Fn.lookup_nested(Fn.split("/", tfc_workspace_slug.string_value), ["1"])),
            organization=Token.as_string(
                Fn.lookup_nested(Fn.split("/", tfc_workspace_slug.string_value), ["0"]))
        )
        tfe_workspace_settings_self = WorkspaceSettings(self, "self_1",
            assessments_enabled=True,
            tags={
                "prod": "true"
            },
            workspace_id=Token.as_string(self.id)
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        tfe_workspace_settings_self.override_logical_id("self")
```

## Argument Reference

The following arguments are supported:

* `workspace_id` - (Required) ID of the workspace.
* `agent_pool_id` - (Optional) The ID of an agent pool to assign to the workspace. Requires `execution_mode`
  to be set to `agent`. This value _must not_ be provided if `execution_mode` is set to any other value.
* `execution_mode` - (Optional) Which [execution mode](https://developer.hashicorp.com/terraform/cloud-docs/workspaces/settings#execution-mode)
  to use. Using HCP Terraform, valid values are `remote`, `local` or `agent`. When set to `local`, the workspace will be used for state storage only. **Important:** If you omit this attribute, the resource configures the workspace to use your organization's default execution mode (which in turn defaults to `remote`), removing any explicit value that might have previously been set for the workspace.
* `global_remote_state` - (Optional) Whether the workspace allows all workspaces in the organization to access its state data during runs. If false, then only specifically approved workspaces can access its state (`remote_state_consumer_ids`). By default, HashiCorp recommends you do not allow other workspaces to access their state. We recommend that you follow the principle of least privilege and only enable state access between workspaces that specifically need information from each other.
* `remote_state_consumer_ids` - (Optional) The set of workspace IDs set as explicit remote state consumers for the given workspace. To set this attribute, global_remote_state must be false.
* `auto_apply` - (Optional) Whether to automatically apply changes when a Terraform plan is successful. Defaults to `false`.
* `assessments_enabled` - (Optional) Whether to regularly run health assessments such as drift detection on the workspace. Defaults to `false`.
* `description` - (Optional) A description for the workspace.
* `tags` - (Optional) A map of key value tags for this workspace.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The workspace ID.
* `overwrites` - Can be used to check whether a setting is currently inheriting its value from another resource.
  - `execution_mode` - Set to `true` if the execution mode of the workspace is being determined by the setting on the workspace itself. It will be `false` if the execution mode is inherited from another resource (e.g. the organization's default execution mode)
  - `agent_pool` - Set to `true` if the agent pool of the workspace is being determined by the setting on the workspace itself. It will be `false` if the agent pool is inherited from another resource (e.g. the organization's default agent pool)
* `effective_tags` - A map of key value tags for this workspace, including any tags inherited from the parent project.

## Import

Workspaces can be imported; use `<WORKSPACE ID>` or `<ORGANIZATION NAME>/<WORKSPACE NAME>` as the
import ID. For example:

```shell
terraform import tfe_workspace_settings.test ws-CH5in3chf8RJjrVd
```

```shell
terraform import tfe_workspace_settings.test my-org-name/my-wkspace-name
```

<!-- cache-key: cdktf-0.20.8 input-ae37db8425cb90bf9d23460da7f257de504c63f7e3ecf246d845b2d4ccfe3df1 -->