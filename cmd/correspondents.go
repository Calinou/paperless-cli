// Copyright © 2019 Steve Garf <stgarf@gmail.com>
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

package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var correspondentsCmd = &cobra.Command{
	Use:     "correspondents",
	Aliases: []string{"correspondent", "correspond", "corr", "cor"},
	Short:   "Manage correspondents of Paperless instance",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debugf("Called 'correspondents' with args %v", args)
		fmt.Println("correspondents called")
	},
}

func init() {
	rootCmd.AddCommand(correspondentsCmd)
}
