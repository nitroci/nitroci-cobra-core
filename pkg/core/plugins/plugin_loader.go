/*
Copyright 2021 The NitroCI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package plugins

import (
	pkgCPlugins "github.com/nitroci/nitroci-core/pkg/core/plugins"

	"github.com/spf13/cobra"
)

func LoadMapFromFlags(cmd *cobra.Command, flags []pkgCPlugins.Flags) map[string]interface{} {
	flagsMap := map[string]interface{}{}
	if len(flags) > 0 {
		for _, flag := range flags {
			switch flag.Type {
			case "string":
				value, _ := cmd.Flags().GetString(flag.Name)
				flagsMap[flag.Name] = value
			case "bool":
				value, _ := cmd.Flags().GetBool(flag.Name)
				flagsMap[flag.Name] = value
			default:
			}
		}
	}
	return flagsMap
}

func LoadFlags(cmd *cobra.Command, flags []pkgCPlugins.Flags) {
	if len(flags) > 0 {
		for _, flag := range flags {
			switch flag.Type {
			case "string":
				flagValue := ""
				if flag.Value != nil {
					flagValue = flag.Value.(string)
				}
				if flag.Shorthand == nil {
					cmd.Flags().String(flag.Name, flagValue, flag.Usage)
				} else {
					cmd.Flags().StringP(flag.Name, flag.Shorthand.(string), flagValue, flag.Usage)
				}
			case "bool":
				flagValue := false
				if flag.Value != nil {
					flagValue = flag.Value.(bool)
				}
				if flag.Shorthand == nil {
					cmd.Flags().Bool(flag.Name, flagValue, flag.Usage)
				} else {
					cmd.Flags().BoolP(flag.Name, flag.Shorthand.(string), flagValue, flag.Usage)
				}
			default:
			}
		}
	}
}
