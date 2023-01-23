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

package retransmission

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Thresholds Custom thresholds for high retransmission rate. If not set, automatic mode is used.
//
//	**All** of these conditions must be met to trigger an alert.
type Thresholds struct {
	RetransmissionRatePercentage        int32 `json:"retransmissionRatePercentage"`        // Retransmission rate is higher than *X*% in 3 out of 5 samples.
	RetransmittedPacketsNumberPerMinute int32 `json:"retransmittedPacketsNumberPerMinute"` // Number of retransmitted packets is higher than *X* packets per minute in 3 out of 5 samples.
}

func (me *Thresholds) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"retransmission_rate": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "Retransmission rate is higher than *X*% in 3 out of 5 samples",
		},
		"retransmitted_packets": {
			Type:        schema.TypeInt,
			Required:    true,
			Description: "Number of retransmitted packets is higher than *X* packets per minute in 3 out of 5 samples",
		},
	}
}

func (me *Thresholds) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"retransmission_rate":   me.RetransmissionRatePercentage,
		"retransmitted_packets": me.RetransmittedPacketsNumberPerMinute,
	})
}

func (me *Thresholds) UnmarshalHCL(decoder hcl.Decoder) error {
	if value, ok := decoder.GetOk("retransmission_rate"); ok {
		me.RetransmissionRatePercentage = int32(value.(int))
	}
	if value, ok := decoder.GetOk("retransmitted_packets"); ok {
		me.RetransmittedPacketsNumberPerMinute = int32(value.(int))
	}
	return nil
}
