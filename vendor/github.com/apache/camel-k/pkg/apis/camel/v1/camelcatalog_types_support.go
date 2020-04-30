/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewCamelCatalog --
func NewCamelCatalog(namespace string, name string) CamelCatalog {
	return CamelCatalog{
		TypeMeta: metav1.TypeMeta{
			APIVersion: SchemeGroupVersion.String(),
			Kind:       CamelCatalogKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
	}
}

// NewCamelCatalogWithSpecs --
func NewCamelCatalogWithSpecs(namespace string, name string, spec CamelCatalogSpec) CamelCatalog {
	return CamelCatalog{
		TypeMeta: metav1.TypeMeta{
			APIVersion: SchemeGroupVersion.String(),
			Kind:       CamelCatalogKind,
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: spec,
	}
}

// NewCamelCatalogList --
func NewCamelCatalogList() CamelCatalogList {
	return CamelCatalogList{
		TypeMeta: metav1.TypeMeta{
			APIVersion: SchemeGroupVersion.String(),
			Kind:       CamelCatalogKind,
		},
	}
}

// GetDependencyID returns a Camel K recognizable maven dependency for the artifact
func (in *CamelArtifact) GetDependencyID() string {
	switch {
	case in.GroupID == "org.apache.camel" && strings.HasPrefix(in.ArtifactID, "camel-"):
		return "camel:" + in.ArtifactID[6:]
	case in.GroupID == "org.apache.camel.quarkus" && strings.HasPrefix(in.ArtifactID, "camel-quarkus-"):
		return "camel-quarkus:" + in.ArtifactID[14:]
	case in.Version == "":
		return "mvn:" + in.GroupID + ":" + in.ArtifactID
	default:
		return "mvn:" + in.GroupID + ":" + in.ArtifactID + ":" + in.Version
	}
}
