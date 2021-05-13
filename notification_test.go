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

package main_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/dtcookie/dynatrace/api/config/notifications"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/config"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/testbase"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type NotificationTest struct {
	resourceKey string
}

func NewNotificationTest() ResourceTest {
	return &NotificationTest{resourceKey: "dynatrace_notification"}
}

func (test *NotificationTest) Anonymize(m map[string]interface{}) {
	delete(m, "id")
	delete(m, "name")
	delete(m, "instanceName")
	delete(m, "displayName")
	delete(m, "metadata")
}

func (test *NotificationTest) ResourceKey() string {
	return test.resourceKey
}

func (test *NotificationTest) CreateTestCase(file string, localJSONFile string, t *testing.T) (*resource.TestCase, error) {
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
		PreCheck:          func() { testAccPreCheck(t) },
		IDRefreshName:     resourceName,
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      test.CheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					test.CheckExists(resourceName, t),
					compareLocalRemote(test, resourceName, localJSONFile, t),
				),
			},
		},
	}, nil
}

func TestAccNotifications(t *testing.T) {
	if disabled, ok := testbase.DisabledTests["notifications"]; ok && disabled {
		t.Skip()
	}
	test := NewNotificationTest()
	var err error
	var testCase *resource.TestCase
	if testCase, err = test.CreateTestCase(
		"test_data/notifications/example_a.tf",
		"test_data/notifications/example_a.json",
		t,
	); err != nil {
		t.Fatal(err)
		return
	}
	resource.Test(t, *testCase)
}

func (test *NotificationTest) URL(id string) string {
	envURL := testAccProvider.Meta().(*config.ProviderConfiguration).DTenvURL
	reqPath := "%s/notifications/%s"
	return fmt.Sprintf(reqPath, envURL, id)
}

func (test *NotificationTest) CheckDestroy(s *terraform.State) error {
	providerConf := testAccProvider.Meta().(*config.ProviderConfiguration)
	restClient := notifications.NewService(providerConf.DTenvURL, providerConf.APIToken)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dynatrace_notification" {
			continue
		}

		id := rs.Primary.ID

		if _, err := restClient.Get(id); err != nil {
			// HTTP Response "404 Not Found" signals a success
			if strings.Contains(err.Error(), `"code": 404`) {
				return nil
			}
			// any other error should fail the test
			return err
		}
		return fmt.Errorf("Custom Service still exists: %s", rs.Primary.ID)
	}

	return nil
}

func (test *NotificationTest) CheckExists(n string, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		providerConf := testAccProvider.Meta().(*config.ProviderConfiguration)
		restClient := notifications.NewService(providerConf.DTenvURL, providerConf.APIToken)

		if rs, ok := s.RootModule().Resources[n]; ok {
			if _, err := restClient.Get(rs.Primary.ID); err != nil {
				return err
			}
			return nil
		}

		return fmt.Errorf("Not found: %s", n)
	}
}