// DBDeployer - The MySQL Sandbox
// Copyright © 2006-2018 Giuseppe Maxia
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sandbox

// Templates for group replication

var (
	checkPxcNodesTemplate string = `#!/bin/sh
{{.Copyright}}
# Generated by dbdeployer {{.AppVersion}} using {{.TemplateName}} on {{.DateTime}}
multi_sb={{.SandboxDir}}
[ -z "$SLEEP_TIME" ] && SLEEP_TIME=1

CHECK_NODE="SHOW GLOBAL STATUS WHERE variable_name IN"
CHECK_NODE="$CHECK_NODE ('wsrep_local_state','wsrep_local_state_comment',"
CHECK_NODE="$CHECK_NODE 'wsrep_local_commits','wsrep_received',"
CHECK_NODE="$CHECK_NODE 'wsrep_cluster_size','wsrep_cluster_status',"
CHECK_NODE="$CHECK_NODE 'wsrep_local_state_uuid','wsrep_incoming_addresses',"
CHECK_NODE="$CHECK_NODE 'wsrep_gcomm_uuid','wsrep_connected','wsrep_ready')"

{{ range .Nodes}}
	echo "# Node {{.Node}}"
	$multi_sb/{{.NodeLabel}}{{.Node}}/use -t -e "$CHECK_NODE"
	sleep $SLEEP_TIME
{{end}}
`
	PxcTemplates = TemplateCollection{
		"check_pxc_nodes_template": TemplateDesc{
			Description: "Checks the status of PXC replication",
			Notes:       "",
			Contents:    checkPxcNodesTemplate,
		},
	}
)