---
layout: ""
page_title: dynatrace_service_http_failure Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_service_http_failure` covers configuration for service-level HTTP failure detection parameters
---

# dynatrace_service_http_failure (Resource)

## Dynatrace Documentation

- Configure service failure detection - https://www.dynatrace.com/support/help/platform-modules/applications-and-microservices/services/service-monitoring-settings/configure-service-failure-detection

- Settings API - https://www.dynatrace.com/support/help/dynatrace-api/environment-api/settings (schemaId: `builtin:failure-detection.service.http-parameters`)

## Export Example Usage

- `terraform-provider-dynatrace -export dynatrace_service_http_failure` downloads all existing service-level HTTP failure detection parameters

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs/guides/export-v2).

## Resource Example Usage

```terraform
resource "dynatrace_service_http_failure" "#name#" {
  enabled    = true
  service_id = "SERVICE-D892CFE7DFAB0D08"
  broken_links {
    broken_link_domains         = [ "www.google.com" ]
    http_404_not_found_failures = true
  }
  http_response_codes {
    client_side_errors                        = "401-599"
    fail_on_missing_response_code_client_side = true
    fail_on_missing_response_code_server_side = true
    server_side_errors                        = "501-599"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `enabled` (Boolean) This setting is enabled (`true`) or disabled (`false`)
- `service_id` (String) The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.

### Optional

- `broken_links` (Block List, Max: 1) HTTP 404 response codes are thrown when a web server can't find a certain page. 404s are classified as broken links on the client side and therefore aren't considered to be service failures. By enabling this setting, you can have 404s treated as server-side service failures. (see [below for nested schema](#nestedblock--broken_links))
- `http_response_codes` (Block List, Max: 1) HTTP response codes (see [below for nested schema](#nestedblock--http_response_codes))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--broken_links"></a>
### Nested Schema for `broken_links`

Required:

- `http_404_not_found_failures` (Boolean) Consider 404 HTTP response codes as failures

Optional:

- `broken_link_domains` (Set of String) If your application relies on other hosts at other domains, add the associated domain names here. Once configured, Dynatrace will consider 404s thrown by hosts at these domains to be service failures related to your application.


<a id="nestedblock--http_response_codes"></a>
### Nested Schema for `http_response_codes`

Required:

- `client_side_errors` (String) HTTP response codes which indicate client side errors
- `fail_on_missing_response_code_client_side` (Boolean) Treat missing HTTP response code as client side error
- `fail_on_missing_response_code_server_side` (Boolean) Treat missing HTTP response code as server side errors
- `server_side_errors` (String) HTTP response codes which indicate an error on the server side
 