/**
* @license
* Copyright 2020 Dynatrace LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package alerting

import (
	"context"

	alertingapi "github.com/dtcookie/dynatrace/api/config/alerting"
	"github.com/dtcookie/hcl"
	"github.com/dtcookie/opt"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/config"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/hcl2sdk"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/logging"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Resource produces terraform resource definition for Management Zones
func Resource() *schema.Resource {
	return &schema.Resource{
		Schema:        hcl2sdk.Convert(new(alertingapi.Profile).Schema()),
		CreateContext: logging.Enable(Create),
		UpdateContext: logging.Enable(Update),
		ReadContext:   logging.Enable(Read),
		DeleteContext: logging.Enable(Delete),
		Importer:      &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext},
	}
}

func NewService(m interface{}) *alertingapi.Service {
	conf := m.(*config.ProviderConfiguration)
	apiService := alertingapi.NewService(conf.DTenvURL, conf.APIToken)
	return apiService
}

// Create expects the configuration within the given ResourceData and sends it to the Dynatrace Server in order to create that resource
func Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := new(alertingapi.Profile)
	if err := config.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	config.ID = nil
	config.Metadata = nil
	abc := NewService(m)
	objStub, err := abc.Create(config)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(objStub.ID)
	return Read(ctx, d, m)
}

// Update expects the configuration within the given ResourceData and send them to the Dynatrace Server in order to update that resource
func Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config := new(alertingapi.Profile)
	if err := config.UnmarshalHCL(hcl.DecoderFrom(d)); err != nil {
		return diag.FromErr(err)
	}
	config.ID = opt.NewString(d.Id())
	config.Metadata = nil
	if err := NewService(m).Update(config); err != nil {
		return diag.FromErr(err)
	}
	return Read(ctx, d, m)
}

// func traverseSlice(elems []interface{}) []interface{} {
// 	return elems
// }

// func unique(v interface{}) int {
// 	return 0
// }

// func traverseMap(m map[string]interface{}) map[string]interface{} {
// 	if m == nil {
// 		return m
// 	}
// 	for k, v := range m {
// 		switch tv := v.(type) {
// 		case hcl.Set:
// 			m[k] = schema.NewSet(nil, []interface{}(tv))
// 		case []interface{}:
// 			m[k] = traverseSlice(tv)
// 		case map[string]interface{}:
// 			m[k] = traverseMap(tv)
// 		default:
// 			m[k] = tv
// 		}
// 	}
// 	return m
// }

// Read queries the Dynatrace Server for the configuration
func Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	config, err := NewService(m).Get(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	marshalled, err := config.MarshalHCL()
	if err != nil {
		return diag.FromErr(err)
	}
	for k, v := range marshalled {
		d.Set(k, v)
	}
	return diag.Diagnostics{}
}

// Delete the configuration
func Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if err := NewService(m).Delete(d.Id()); err != nil {
		return diag.FromErr(err)
	}
	return diag.Diagnostics{}
}