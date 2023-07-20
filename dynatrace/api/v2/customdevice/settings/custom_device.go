package customdevice

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type CustomDevice struct {
	EntityId       string            `json:"entityId,omitempty"`       // The ID of the custom device.
	DisplayName    *string           `json:"displayName,omitempty"`    // The name of the custom device, displayed in the UI.
	CustomDeviceID string            `json:"customDeviceId,omitempty"` // A unique name that can be provided or generated by the provider
	Group          *string           `json:"group,omitempty"`          // User defined group of entity. Changing the group requires a new custom device to be created.
	IPAddresses    []string          `json:"ipAddresses,omitempty"`    // The list of IP addresses that belong to the custom device.
	ListenPorts    []int             `json:"listenPorts,omitempty"`    // The list of ports the custom devices listens to.
	Type           *string           `json:"type,omitempty"`           // The technology type definition of the custom device.
	FaviconUrl     *string           `json:"faviconUrl,omitempty"`     // The icon to be displayed for your custom component within Smartscape. Provide the full URL of the icon file.
	ConfigUrl      *string           `json:"configUrl,omitempty"`      // The URL of a configuration web page for the custom device, such as a login page for a firewall or router.
	DNSNames       []string          `json:"dnsNames,omitempty"`       // The list of DNS names related to the custom device.
	Properties     map[string]string `json:"properties,omitempty"`     // The list of key-value pair properties that will be shown beneath the infographics of your custom device.
}

type CustomDeviceGetResponse struct {
	Entities []*CustomDeviceResponse `json:"entities,omitempty"` // An unordered list of custom devices
}

type CustomDeviceResponse struct {
	EntityId          string                     `json:"entityId,omitempty"`
	DisplayName       *string                    `json:"displayName,omitempty"`
	Type              *string                    `json:"type,omitempty"`
	Properties        *PropertiesResponse        `json:"properties,omitempty"`
	FromRelationships *FromRelationshipsResponse `json:"fromRelationships,omitempty"`
}

type PropertiesResponse struct {
	DetectedName     string              `json:"detectedName,omitempty"`
	IPAddress        []string            `json:"ipAddress,omitempty"`
	ListenPorts      []int               `json:"listenPorts,omitempty"`
	CustomFavicon    *string             `json:"customFavicon,omitempty"`
	DNSNames         []string            `json:"dnsNames,omitempty"`
	CustomProperties []*CustomProperties `json:"customProperties,omitempty"`
}

type CustomProperties struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type FromRelationshipsResponse struct {
	IsInstanceOf []*IsInstanceOf `json:"isInstanceOf,omitempty"`
}

type IsInstanceOf struct {
	ID   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

func (me *CustomDevice) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"entity_id": {
			Type:        schema.TypeString,
			Description: "The Dynatrace EntityID of this resource. If you need to refer to this custom device within other resources you want to use this ID",
			Computed:    true,
		},
		"display_name": {
			Type:        schema.TypeString,
			Description: "The name of the custom device, displayed in the UI.",
			Required:    true,
		},
		"custom_device_id": {
			Type:        schema.TypeString,
			Description: "The unique name of the custom device. This Id can either be provided in the resource or generated by Terraform when the resource is created. If you use the ID of an existing device, the respective parameters will be updated",
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
		},
		"group": {
			Type:        schema.TypeString,
			Description: "User defined group of entity. Changing the group requires a new custom device to be created.",
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
		},
		"ip_addresses": {
			Type:        schema.TypeSet,
			Description: "The list of IP addresses that belong to the custom device.",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"listen_ports": {
			Type:        schema.TypeSet,
			Description: "The list of ports the custom devices listens to.",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeInt},
		},
		"type": {
			Type:        schema.TypeString,
			Description: "The technology type definition of the custom device.",
			Optional:    true,
			Computed:    true,
		},
		"favicon_url": {
			Type:        schema.TypeString,
			Description: "The icon to be displayed for your custom component within Smartscape. Provide the full URL of the icon file.",
			Optional:    true,
		},
		"config_url": {
			Type:        schema.TypeString,
			Description: "The URL of a configuration web page for the custom device, such as a login page for a firewall or router.",
			Optional:    true,
		},
		"dns_names": {
			Type:        schema.TypeSet,
			Description: "The list of DNS names related to the custom device.",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
		"properties": {
			Type:        schema.TypeMap,
			Description: "The list of key-value pair properties that will be shown beneath the infographics of your custom device.",
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
		},
	}
}

func (me *CustomDeviceGetResponse) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"entities": {
			Type:        schema.TypeString,
			Description: "The list of entities returned by the GET call.",
			Optional:    true,
			Computed:    true,
		},
	}
}

func (me *CustomDeviceGetResponse) MarshalHCL(properties hcl.Properties) error {
	if err := properties.EncodeAll(map[string]any{
		"entities": me.Entities,
	}); err != nil {
		return err
	}
	return nil
}

func (me *CustomDeviceGetResponse) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"entities": &me.Entities,
	})
}

func (me *CustomDevice) MarshalHCL(properties hcl.Properties) error {
	if err := properties.EncodeAll(map[string]any{
		"entity_id":        me.EntityId,
		"display_name":     me.DisplayName,
		"custom_device_id": me.CustomDeviceID,
		"group":            me.Group,
		"ip_addresses":     me.IPAddresses,
		"listen_ports":     me.ListenPorts,
		"type":             me.Type,
		"favicon_url":      me.FaviconUrl,
		"dns_names":        me.DNSNames,
		"properties":       me.Properties,
	}); err != nil {
		return err
	}
	return nil
}

func (me *CustomDevice) UnmarshalHCL(decoder hcl.Decoder) error {
	err := decoder.DecodeAll(map[string]any{
		"entity_id":        &me.EntityId,
		"display_name":     &me.DisplayName,
		"custom_device_id": &me.CustomDeviceID,
		"group":            &me.Group,
		"ip_addresses":     &me.IPAddresses,
		"listen_ports":     &me.ListenPorts,
		"type":             &me.Type,
		"favicon_url":      &me.FaviconUrl,
		"config_url":       &me.ConfigUrl,
		"dns_names":        &me.DNSNames,
		"properties":       &me.Properties,
	})

	return err
}

func (me *CustomDevice) Name() string {
	return *me.DisplayName
}
