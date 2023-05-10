---
layout: ""
page_title: dynatrace_cloud_foundry Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_cloud_foundry` covers configuration for Cloud Foundry monitoring
---

# dynatrace_cloud_foundry (Resource)

## Dynatrace Documentation

- Cloud Foundry monitoring - https://www.dynatrace.com/support/help/how-to-use-dynatrace/infrastructure-monitoring/container-platform-monitoring/cloud-foundry-monitoring

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:cloud.cloudfoundry`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_cloud_foundry` downloads all existing Cloud Foundry monitoring configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_cloud_foundry" "#name#" {
  enabled           = true
  active_gate_group = "Terraform"
  api_url           = "https://www.google.at/test/#name#"
  label             = "#name#"
  login_url         = "https://www.google.at/test/#name#"
  password          = "pass2"
  username          = "user"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `api_url` (String) Cloud Foundry API Target
- `enabled` (Boolean) This setting is enabled (`true`) or disabled (`false`)
- `label` (String) Name this connection
- `login_url` (String) Cloud Foundry Authentication Endpoint
- `password` (String, Sensitive) Cloud Foundry Password
- `username` (String) Cloud Foundry Username

### Optional

- `active_gate_group` (String) ActiveGate group

### Read-Only

- `id` (String) The ID of this resource.
 