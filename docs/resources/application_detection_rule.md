---
layout: ""
page_title: "dynatrace_application_detection_rule Resource - terraform-provider-dynatrace"
description: |-
  The resource `dynatrace_application_detection_rule` covers configuration for application detection rules
---

# dynatrace_application_detection_rule (Resource)

## Dynatrace Documentation

- Check application detection rules - https://www.dynatrace.com/support/help/how-to-use-dynatrace/real-user-monitoring/setup-and-configuration/web-applications/additional-configuration/application-detection-rules

- Applications detection rules API - https://www.dynatrace.com/support/help/dynatrace-api/configuration-api/rum/application-detection-configuration

## Export Example Usage

- `terraform-provider-dynatrace export dynatrace_application_detection_rule` downloads all existing application detection rule configuration

The full documentation of the export feature is available [here](https://registry.terraform.io/providers/dynatrace-oss/dynatrace/latest/docs#exporting-existing-configuration-from-a-dynatrace-environment).

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `application_identifier` (String) The Dynatrace entity ID of the application, for example APPLICATION-4A3B43
- `filter_config` (Block List, Min: 1, Max: 1) The condition of an application detection rule (see [below for nested schema](#nestedblock--filter_config))

### Optional

- `name` (String) The unique name of the Application detection rule
- `order` (String) The order of the rule in the rules list

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--filter_config"></a>
### Nested Schema for `filter_config`

Required:

- `application_match_target` (String) Where to look for the pattern value, possible values are `DOMAIN` or `URL`
- `application_match_type` (String) The operator used for matching the application detection rule, possible values are `BEGINS_WITH`, `CONTAINS`, `ENDS_WITH`, `EQUALS`, `MATCHES`
- `pattern` (String) The value to look for with the application detection rule
 