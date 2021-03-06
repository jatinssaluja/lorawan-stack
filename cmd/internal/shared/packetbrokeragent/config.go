// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
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

package shared

import (
	"go.thethings.network/lorawan-stack/pkg/packetbrokeragent"
)

// DefaultPacketBrokerAgentConfig is the default configuration for the Packet Broker Agent.
var DefaultPacketBrokerAgentConfig = packetbrokeragent.Config{
	HomeNetwork: packetbrokeragent.HomeNetworkConfig{
		WorkerPool: packetbrokeragent.WorkerPoolConfig{
			Limit: 4096,
		},
		BlacklistForwarder: true,
	},
	Forwarder: packetbrokeragent.ForwarderConfig{
		WorkerPool: packetbrokeragent.WorkerPoolConfig{
			Limit: 1024,
		},
	},
}
