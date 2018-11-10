/*
Copyright 2018 Iguazio Systems Ltd.

Licensed under the Apache License, Version 2.0 (the "License") with
an addition restriction as set forth herein. You may not use this
file except in compliance with the License. You may obtain a copy of
the License at http://www.apache.org/licenses/LICENSE-2.0.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
implied. See the License for the specific language governing
permissions and limitations under the License.

In addition, you may not use the software for any purposes that are
illegal under applicable law, and the grant of the foregoing license
under the Apache 2.0 license is conditioned upon your compliance with
such restriction.
*/

package tsdbctl

import (
	"fmt"
	"github.com/spf13/cobra"
)

type versionCommandeer struct {
	cmd            *cobra.Command
	rootCommandeer *RootCommandeer
}

func newVersionCommandeer(rootCommandeer *RootCommandeer) *versionCommandeer {
	commandeer := &versionCommandeer{
		rootCommandeer: rootCommandeer,
	}

	cmd := &cobra.Command{
		Aliases: []string{"ver"},
		Use:     "version",
		Hidden:  false,
		Short:   "Displays version",
		Long:    "Displays version information",
		Example: "- tsdbctl version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("tsdbctl build information:\n  OS: %s\n  Architecture: %s\n  Version: %s\n  SHA: %s\n  Branch: %s\n",
				rootCommandeer.BuildInfo.Os,
				rootCommandeer.BuildInfo.Architecture,
				rootCommandeer.BuildInfo.Version,
				rootCommandeer.BuildInfo.Revision,
				rootCommandeer.BuildInfo.Branch)
			return nil
		},
	}

	commandeer.cmd = cmd

	return commandeer
}
