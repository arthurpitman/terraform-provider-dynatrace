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

package advanceddetectionrule

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type ProcessGroupExtraction struct {
	Delimiter      *Delimiter `json:"delimiter"`      // Optionally delimit this property between *From* and *To*.
	Property       string     `json:"property"`       // Possible values: `DOTNET_COMMAND`, `DOTNET_COMMAND_PATH`, `ASP_NET_CORE_APPLICATION_PATH`, `AWS_ECR_ACCOUNT_ID`, `AWS_ECR_REGION`, `AWS_ECS_CLUSTER`, `AWS_ECS_CONTAINERNAME`, `AWS_ECS_FAMILY`, `AWS_ECS_REVISION`, `AWS_LAMBDA_FUNCTION_NAME`, `AWS_REGION`, `APACHE_SPARK_MASTER_IP_ADDRESS`, `APACHE_CONFIG_PATH`, `CATALINA_BASE`, `CATALINA_HOME`, `CLOUD_FOUNDRY_APP_NAME`, `CLOUD_FOUNDRY_APPLICATION_ID`, `CLOUD_FOUNDRY_INSTANCE_INDEX`, `CLOUD_FOUNDRY_SPACE_NAME`, `CLOUD_FOUNDRY_SPACE_ID`, `COLDFUSION_JVM_CONFIG_FILE`, `SERVICE_NAME`, `COMMAND_LINE_ARGS`, `CONTAINER_ID`, `CONTAINER_IMAGE_VERSION`, `CONTAINER_NAME`, `DECLARATIVE_ID`, `CONTAINER_IMAGE_NAME`, `RUXIT_CLUSTER_ID`, `RUXIT_NODE_ID`, `EXE_NAME`, `EXE_PATH`, `ELASTIC_SEARCH_CLUSTER_NAME`, `ELASTIC_SEARCH_NODE_NAME`, `EQUINOX_CONFIG_PATH`, `GLASSFISH_DOMAIN_NAME`, `GLASSFISH_INSTANCE_NAME`, `PG_ID_CALC_INPUT_KEY_LINKAGE`, `GAE_INSTANCE`, `GAE_SERVICE`, `GOOGLE_CLOUD_PROJECT`, `HYBRIS_BIN_DIR`, `HYBRIS_CONFIG_DIR`, `HYBRIS_DATA_DIR`, `IBM_CICS_REGION`, `IBM_CICS_IMS_APPLID`, `IBM_CICS_IMS_JOBNAME`, `IBM_CTG_NAME`, `IBM_IMS_CONNECT`, `IBM_IMS_CONTROL`, `IBM_IMS_MPR`, `IBM_IMS_SOAP_GW_NAME`, `IIB_BROKER_NAME`, `IIB_EXECUTION_GROUP_NAME`, `IIS_APP_POOL`, `IIS_ROLE_NAME`, `JBOSS_HOME`, `JBOSS_MODE`, `JBOSS_SERVER_NAME`, `JAVA_JAR_FILE`, `JAVA_JAR_PATH`, `JAVA_MAIN_CLASS`, `KUBERNETES_BASEPODNAME`, `KUBERNETES_CONTAINERNAME`, `KUBERNETES_FULLPODNAME`, `KUBERNETES_NAMESPACE`, `KUBERNETES_PODUID`, `MSSQL_INSTANCE_NAME`, `NODEJS_APP_NAME`, `NODEJS_APP_BASE_DIR`, `NODEJS_SCRIPT_NAME`, `ORACLE_SID`, `PHP_CLI_SCRIPT_PATH`, `PHP_CLI_WORKING_DIR`, `SOFTWAREAG_INSTALL_ROOT`, `SOFTWAREAG_PRODUCTPROPNAME`, `SPRINGBOOT_APP_NAME`, `SPRINGBOOT_PROFILE_NAME`, `SPRINGBOOT_STARTUP_CLASS`, `TIBCO_BUSINESSWORKS_CE_APP_NAME`, `TIBCO_BUSINESSWORKS_CE_VERSION`, `TIBCO_BUSINESSWORKS_APP_NODE_NAME`, `TIBCO_BUSINESSWORKS_APP_SPACE_NAME`, `TIBCO_BUSINESSWORKS_DOMAIN_NAME`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE_PATH`, `TIBCO_BUSINESSWORKS_HOME`, `VARNISH_INSTANCE_NAME`, `WEBLOGIC_NAME`, `WEBLOGIC_CLUSTER_NAME`, `WEBLOGIC_DOMAIN_NAME`, `WEBLOGIC_HOME`, `WEBSPHERE_LIBERTY_SERVER_NAME`, `WEBSPHERE_CELL_NAME`, `WEBSPHERE_CLUSTER_NAME`, `WEBSPHERE_NODE_NAME`, `WEBSPHERE_SERVER_NAME`
	StandaloneRule bool       `json:"standaloneRule"` // If this option is selected, the default Dynatrace behavior is disabled for these detected processes. Only this rule is used to separate the process group.\n\nIf this option is not selected, this rule contributes to the default Dynatrace process group detection. \n\n[See our help page for examples.](https://dt-url.net/1722wrz)
}

func (me *ProcessGroupExtraction) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"delimiter": {
			Type:        schema.TypeList,
			Description: "Optionally delimit this property between *From* and *To*.",
			Required:    true,
			Elem:        &schema.Resource{Schema: new(Delimiter).Schema()},
			MinItems:    1,
			MaxItems:    1,
		},
		"property": {
			Type:        schema.TypeString,
			Description: "Possible values: `DOTNET_COMMAND`, `DOTNET_COMMAND_PATH`, `ASP_NET_CORE_APPLICATION_PATH`, `AWS_ECR_ACCOUNT_ID`, `AWS_ECR_REGION`, `AWS_ECS_CLUSTER`, `AWS_ECS_CONTAINERNAME`, `AWS_ECS_FAMILY`, `AWS_ECS_REVISION`, `AWS_LAMBDA_FUNCTION_NAME`, `AWS_REGION`, `APACHE_SPARK_MASTER_IP_ADDRESS`, `APACHE_CONFIG_PATH`, `CATALINA_BASE`, `CATALINA_HOME`, `CLOUD_FOUNDRY_APP_NAME`, `CLOUD_FOUNDRY_APPLICATION_ID`, `CLOUD_FOUNDRY_INSTANCE_INDEX`, `CLOUD_FOUNDRY_SPACE_NAME`, `CLOUD_FOUNDRY_SPACE_ID`, `COLDFUSION_JVM_CONFIG_FILE`, `SERVICE_NAME`, `COMMAND_LINE_ARGS`, `CONTAINER_ID`, `CONTAINER_IMAGE_VERSION`, `CONTAINER_NAME`, `DECLARATIVE_ID`, `CONTAINER_IMAGE_NAME`, `RUXIT_CLUSTER_ID`, `RUXIT_NODE_ID`, `EXE_NAME`, `EXE_PATH`, `ELASTIC_SEARCH_CLUSTER_NAME`, `ELASTIC_SEARCH_NODE_NAME`, `EQUINOX_CONFIG_PATH`, `GLASSFISH_DOMAIN_NAME`, `GLASSFISH_INSTANCE_NAME`, `PG_ID_CALC_INPUT_KEY_LINKAGE`, `GAE_INSTANCE`, `GAE_SERVICE`, `GOOGLE_CLOUD_PROJECT`, `HYBRIS_BIN_DIR`, `HYBRIS_CONFIG_DIR`, `HYBRIS_DATA_DIR`, `IBM_CICS_REGION`, `IBM_CICS_IMS_APPLID`, `IBM_CICS_IMS_JOBNAME`, `IBM_CTG_NAME`, `IBM_IMS_CONNECT`, `IBM_IMS_CONTROL`, `IBM_IMS_MPR`, `IBM_IMS_SOAP_GW_NAME`, `IIB_BROKER_NAME`, `IIB_EXECUTION_GROUP_NAME`, `IIS_APP_POOL`, `IIS_ROLE_NAME`, `JBOSS_HOME`, `JBOSS_MODE`, `JBOSS_SERVER_NAME`, `JAVA_JAR_FILE`, `JAVA_JAR_PATH`, `JAVA_MAIN_CLASS`, `KUBERNETES_BASEPODNAME`, `KUBERNETES_CONTAINERNAME`, `KUBERNETES_FULLPODNAME`, `KUBERNETES_NAMESPACE`, `KUBERNETES_PODUID`, `MSSQL_INSTANCE_NAME`, `NODEJS_APP_NAME`, `NODEJS_APP_BASE_DIR`, `NODEJS_SCRIPT_NAME`, `ORACLE_SID`, `PHP_CLI_SCRIPT_PATH`, `PHP_CLI_WORKING_DIR`, `SOFTWAREAG_INSTALL_ROOT`, `SOFTWAREAG_PRODUCTPROPNAME`, `SPRINGBOOT_APP_NAME`, `SPRINGBOOT_PROFILE_NAME`, `SPRINGBOOT_STARTUP_CLASS`, `TIBCO_BUSINESSWORKS_CE_APP_NAME`, `TIBCO_BUSINESSWORKS_CE_VERSION`, `TIBCO_BUSINESSWORKS_APP_NODE_NAME`, `TIBCO_BUSINESSWORKS_APP_SPACE_NAME`, `TIBCO_BUSINESSWORKS_DOMAIN_NAME`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE`, `TIPCO_BUSINESSWORKS_PROPERTY_FILE_PATH`, `TIBCO_BUSINESSWORKS_HOME`, `VARNISH_INSTANCE_NAME`, `WEBLOGIC_NAME`, `WEBLOGIC_CLUSTER_NAME`, `WEBLOGIC_DOMAIN_NAME`, `WEBLOGIC_HOME`, `WEBSPHERE_LIBERTY_SERVER_NAME`, `WEBSPHERE_CELL_NAME`, `WEBSPHERE_CLUSTER_NAME`, `WEBSPHERE_NODE_NAME`, `WEBSPHERE_SERVER_NAME`",
			Required:    true,
		},
		"standalone_rule": {
			Type:        schema.TypeBool,
			Description: "If this option is selected, the default Dynatrace behavior is disabled for these detected processes. Only this rule is used to separate the process group.\n\nIf this option is not selected, this rule contributes to the default Dynatrace process group detection. \n\n[See our help page for examples.](https://dt-url.net/1722wrz)",
			Optional:    true,
		},
	}
}

func (me *ProcessGroupExtraction) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"delimiter":       me.Delimiter,
		"property":        me.Property,
		"standalone_rule": me.StandaloneRule,
	})
}

func (me *ProcessGroupExtraction) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"delimiter":       &me.Delimiter,
		"property":        &me.Property,
		"standalone_rule": &me.StandaloneRule,
	})
}
