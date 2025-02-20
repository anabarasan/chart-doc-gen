/*
Copyright The Kubepack Authors.

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

package api

type DocInfo struct {
	Project       ProjectInfo    `json:"project"`
	Repository    RepositoryInfo `json:"repository"`
	Chart         ChartInfo      `json:"chart"`
	Prerequisites []string       `json:"prerequisites"`
	Release       ReleaseInfo    `json:"release"`
	Validation    []string       `json:"validation"`
}

type ProjectInfo struct {
	Name        string `json:"name"`
	ShortName   string `json:"shortName"`
	URL         string `json:"url"`
	Description string `json:"description"`
	App         string `json:"app"`
}

type RepositoryInfo struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

type ChartInfo struct {
	Name          string `json:"name"`
	Version       string `json:"version,omitempty"`
	Values        string `json:"values"`
	ValuesExample string `json:"valuesExample"`
}

type ReleaseInfo struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
