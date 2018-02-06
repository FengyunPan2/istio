// Copyright 2017 Istio Authors
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

// THIS FILE IS AUTOMATICALLY GENERATED.

package adapter

import (
	circonus "istio.io/istio/mixer/adapter/circonus"
	denier "istio.io/istio/mixer/adapter/denier"
	fluentd "istio.io/istio/mixer/adapter/fluentd"
	kubernetesenv "istio.io/istio/mixer/adapter/kubernetesenv"
	list "istio.io/istio/mixer/adapter/list"
	memquota "istio.io/istio/mixer/adapter/memquota"
	noop "istio.io/istio/mixer/adapter/noop"
	opa "istio.io/istio/mixer/adapter/opa"
	prometheus "istio.io/istio/mixer/adapter/prometheus"
	rbac "istio.io/istio/mixer/adapter/rbac"
	servicecontrol "istio.io/istio/mixer/adapter/servicecontrol"
	solarwinds "istio.io/istio/mixer/adapter/solarwinds"
	stackdriver "istio.io/istio/mixer/adapter/stackdriver"
	statsd "istio.io/istio/mixer/adapter/statsd"
	stdio "istio.io/istio/mixer/adapter/stdio"
	adptr "istio.io/istio/mixer/pkg/adapter"
)

// Inventory returns the inventory of all available adapters.
func Inventory() []adptr.InfoFn {
	return []adptr.InfoFn{
		circonus.GetInfo,
		denier.GetInfo,
		fluentd.GetInfo,
		kubernetesenv.GetInfo,
		list.GetInfo,
		memquota.GetInfo,
		noop.GetInfo,
		opa.GetInfo,
		prometheus.GetInfo,
		rbac.GetInfo,
		servicecontrol.GetInfo,
		solarwinds.GetInfo,
		stackdriver.GetInfo,
		statsd.GetInfo,
		stdio.GetInfo,
	}
}
