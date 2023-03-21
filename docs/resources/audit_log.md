---
layout: ""
page_title: dynatrace_audit_log Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_audit_log` covers configuration for audit log
---

# dynatrace_audit_log (Resource)

## Dynatrace Documentation

- Manage audit logs - https://www.dynatrace.com/support/help/manage/data-privacy-and-security/configuration/audit-logs

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:audit-log`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_audit_log` downloads all existing audit log configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_audit_log" "#name#" {
  enabled = false
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) This setting is enabled (`true`) or disabled (`false`)

### Read-Only

- `id` (String) The ID of this resource.
 