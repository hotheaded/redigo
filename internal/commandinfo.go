2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed unde an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the Liceedigo/internal"

import (
	"strings"
)

const (
	WatchState = 1 << iota
	MultiState
	SubscribeState
	MonitorState
)

type CommandInf

func init() {
	for n, ci := range commandInfos {
		commandInfos[strings.ToLower(n)] = ci
	}
}

func LookupCommandInfo(commandName string) CommandInfo {
	if ci, ok := commandInfos[commandName]; ok {
		return ci
	}
	return commandInfos[strings.ToUpper(commandName)]
}
