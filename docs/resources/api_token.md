---
layout: ""
page_title: dynatrace_api_token Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_api_token` covers configuration for API tokens
---

# dynatrace_api_token (Resource)

The token value can be retrieved with `dynatrace_api_token.<#name#>.token` after apply.

WARNING: The usage of `dynatrace_api_token` will introduce sensitive data within your Terraform state. The `token` property is flagged as `sensitive`, but the field will be stored as plain-text. More information can be found [here](https://developer.hashicorp.com/terraform/language/state/sensitive-data).

## Dynatrace Documentation

- Dynatrace API Tokens and authentication - https://www.dynatrace.com/support/help/dynatrace-api/basics/dynatrace-api-authentication

- Tokens API v2 - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/tokens-v2

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_api_token` downloads all existing API token configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_api_token" "#name#" {
  name                    = "#name#"
  enabled                 = false
  # personal_access_token = false
  scopes                  = [ "geographicRegions.read" ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the token.
- `scopes` (Set of String) A list of the scopes to be assigned to the token.

### Optional

- `creation_date` (String) Token creation date in ISO 8601 format (yyyy-MM-dd'T'HH:mm:ss.SSS'Z')
- `enabled` (Boolean) The token is enabled (true) or disabled (false), default disabled (false).
- `expiration_date` (String) The expiration date of the token.
- `last_used_date` (String) Token last used date in ISO 8601 format (yyyy-MM-dd'T'HH:mm:ss.SSS'Z')
- `last_used_ip_address` (String) Token last used IP address.
- `modified_date` (String) Token last modified date in ISO 8601 format (yyyy-MM-dd'T'HH:mm:ss.SSS'Z').
- `owner` (String) The owner of the token
- `personal_access_token` (Boolean) The token is a personal access token (true) or an API token (false).

### Read-Only

- `id` (String) The ID of this resource.
- `token` (String, Sensitive) The secret of the token.
 