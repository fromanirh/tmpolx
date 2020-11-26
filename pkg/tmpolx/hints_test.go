/*
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
 *
 * Copyright 2020 Red Hat, Inc.
 */

package tmpolx

import (
	"fmt"
	"testing"

	"k8s.io/kubernetes/pkg/kubelet/cm/topologymanager"
)

func TestParseHints(t *testing.T) {
	rawHints := []string{
		// cpu:[{01 true} {10 true} {11 false}]
		`{"R":"cpu", "H":[{"M":"01","P":true},{"M":"10","P":true},{"M":"11","P":false}]}`,
		// nvidia.com/gpu:[{01 true} {11 false}]
		`{"R":"nvidia.com/gpu", "H":[{"M":"01","P":true},{"M":"11","P":false}]}`,
		// openshift.io/intelsriov:[{10 true} {11 false}]
		`{"R":"openshift.io/intelsriov", "H":[{"M":"10","P":true},{"M":"11","P":false}]}`,
	}

	tmpx := &TMPolx{
		hints: make(map[string][]topologymanager.TopologyHint),
	}
	err := tmpx.ParseHints(rawHints)
	if err != nil {
		t.Errorf("failed to parse hints: %v", err)
	}
}