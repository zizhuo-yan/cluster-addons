/*

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AddonInstallerConfiguration struct {
	metav1.TypeMeta `json:",inline"`

	// DryRun indicates whether or not to actually install the listed addons
	DryRun bool `json:"dryRun"`
	// Addons is a list of addons to install
	Addons []Addon `json:"addons"`
}

// Addon names and references an addon to be installed.
// Only one of `KustomizeRef` or `ManifestRef` should be provided.
type Addon struct {
	// Name provides the name of the addon.
	// It is used for detecting upgrades and uninstalls as the AddonInstallerConfiguration changes.
	Name string `json:"name"`
	// KustomizeRef may be a folder-path or git URL /w optional subpath
	KustomizeRef string `json:"kustomizeRef"`
	// ManifestRef may be a file-path or HTTP/S served YAML file
	ManifestRef string `json:"manifestRef"`
}
