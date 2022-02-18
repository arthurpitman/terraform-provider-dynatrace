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

package monitors_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	sloapi "github.com/dtcookie/dynatrace/api/config/synthetic/monitors"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/config"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/testbase"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const HTTPResourceName = "dynatrace_http_monitor"

func (test *TestStruct) CreateHTTPTestCase(file string, localJSONFile string, t *testing.T) (*resource.TestCase, error) {
	var content []byte
	var err error
	if content, err = ioutil.ReadFile(file); err != nil {
		return nil, err
	}
	config := string(content)
	name := acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
	resourceName := test.ResourceKey() + "." + name
	config = strings.ReplaceAll(config, "#name#", name)
	return &resource.TestCase{
		PreCheck:          func() { testbase.TestAccPreCheck(t) },
		IDRefreshName:     resourceName,
		ProviderFactories: testbase.TestAccProviderFactories,
		// CheckDestroy:      test.CheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					CheckHTTPExists(resourceName, t),
					testbase.CompareLocalRemoteExt(test, resourceName, localJSONFile, t, true),
				),
			},
		},
	}, nil
}

func TestHTTPMonitorsA(t *testing.T) {
	test := &TestStruct{resourceKey: HTTPResourceName}
	var err error
	var testCase *resource.TestCase
	if testCase, err = test.CreateHTTPTestCase(
		TestDataFolder+"/example_http_a.tf",
		TestDataFolder+"/example_http_a.json",
		t,
	); err != nil {
		t.Fatal(err)
		return
	}
	resource.Test(t, *testCase)
}

func TestHTTPMonitorsB(t *testing.T) {
	test := &TestStruct{resourceKey: HTTPResourceName}
	var err error
	var testCase *resource.TestCase
	if testCase, err = test.CreateHTTPTestCase(
		TestDataFolder+"/example_http_b.tf",
		TestDataFolder+"/example_http_b.json",
		t,
	); err != nil {
		t.Fatal(err)
		return
	}
	resource.Test(t, *testCase)
}

func (test *TestStruct) CheckHTTPDestroy(s *terraform.State) error {
	providerConf := testbase.TestAccProvider.Meta().(*config.ProviderConfiguration)
	restClient := sloapi.NewService(providerConf.DTApiV2URL, providerConf.APIToken)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != ResourceName {
			continue
		}

		id := rs.Primary.ID

		if _, err := restClient.GetHTTP(id); err != nil {
			// HTTP Response "404 Not Found" signals a success
			if strings.Contains(err.Error(), `"code": 404`) {
				return nil
			}
			// any other error should fail the test
			return err
		}
		return fmt.Errorf("Configuration still exists: %s", rs.Primary.ID)
	}

	return nil
}

func CheckHTTPExists(n string, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providerConf := testbase.TestAccProvider.Meta().(*config.ProviderConfiguration)
		restClient := sloapi.NewService(providerConf.DTNonConfigEnvURL, providerConf.APIToken)

		var err error
		var cnt = 0

		for cnt < 20 {
			if rs, ok := s.RootModule().Resources[n]; ok {
				if _, err = restClient.GetHTTP(rs.Primary.ID); err != nil {
					cnt++
				}
				return nil
			} else {
				time.Sleep(5 * time.Second)
			}
		}
		return err
	}
}