---
layout: ""
page_title:  Resource - terraform-provider-dynatrace"
subcategory: "Documents"
description: |-
  The resource `dynatrace_document` covers configuration of Documents (dashboards and notebooks) in Dynatrace.
---

# dynatrace_document (Resource)

-> This resource is excluded by default in the export utility. You can, of course, specify that resource explicitly in order to export it. In that case, don't forget to specify the environment variables `DYNATRACE_AUTOMATION_CLIENT_ID` and `DYNATRACE_AUTOMATION_CLIENT_SECRET` for authentication.

## Dynatrace Documentation

- Dynatrace Documents - https://########.apps.dynatrace.com/platform/swagger-ui/index.html?urls.primaryName=Document%20Service

## Prerequisites

Using this resource requires an OAuth client to be configured within your account settings.
The scopes of the OAuth Client need to include `Create and edit documents (document:documents:write)`, `View documents (document:documents:read)`, `Delete documents (document:documents:delete)`.

Finally the provider configuration requires the credentials for that OAuth Client.
The configuration section of your provider needs to look like this.

```terraform
provider "dynatrace" {
  dt_env_url   = "https://########.live.dynatrace.com/"
  dt_api_token = "######.########################.################################################################"

  # Usually not required. Terraform will deduct it if `dt_env_url` has been specified
  # automation_env_url = "https://########.apps.dynatrace.com/"
  automation_client_id = "######.########"
  automation_client_secret = "######.########.################################################################"
}
```

-> In order to handle credentials in a secure manner we recommend to use the environment variables `DYNATRACE_AUTOMATION_CLIENT_ID` and `DYNATRACE_AUTOMATION_CLIENT_SECRET` as an alternative.

## Resource Example Usage

```terraform
resource "dynatrace_document" "this" {
  type    = "dashboard"
  name    = "Example Dashboard"
  content = file(format("%s/example-dashboard.json", path.module))
}

data "dynatrace_documents" "all-dashboard-and-notebooks" {}
```


<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `content` (String) Document content as JSON
- `name` (String) The name/name of the document
- `type` (String) Type of the document. Possible Values are `dashboard` and `notebook`

### Optional

- `actor` (String) The user context the executions of the document will happen with
- `owner` (String) The ID of the owner of this document

### Read-Only

- `id` (String) The ID of this resource.
- `version` (Number) The version of the document
