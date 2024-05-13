---
layout: ""
page_title: dynatrace_log_debug_settings Resource - terraform-provider-dynatrace"
subcategory: "Log Monitoring"
description: |-
  The resource `dynatrace_log_debug_settings` covers configuration for log debug settings
---

# dynatrace_log_debug_settings (Resource)

-> This resource requires the API token scopes **Read settings** (`settings.read`) and **Write settings** (`settings.write`)

## Dynatrace Documentation

- Troubleshooting Log Monitoring (Logs Classic) - https://docs.dynatrace.com/docs/observe-and-explore/log-monitoring/lmc-troubleshooting

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:logmonitoring.log-debug-settings`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_log_debug_settings` downloads existing log debug settings

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_log_debug_settings" "#name#" {
  enabled     = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) This setting is enabled (`true`) or disabled (`false`)

### Read-Only

- `id` (String) The ID of this resource.
 