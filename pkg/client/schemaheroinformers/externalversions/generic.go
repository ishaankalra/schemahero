/*
Copyright 2021 The SchemaHero Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha4 "github.com/schemahero/schemahero/pkg/apis/databases/v1alpha4"
	schemasv1alpha4 "github.com/schemahero/schemahero/pkg/apis/schemas/v1alpha4"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=databases.schemahero.io, Version=v1alpha4
	case v1alpha4.SchemeGroupVersion.WithResource("databases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Databases().V1alpha4().Databases().Informer()}, nil

		// Group=schemas.schemahero.io, Version=v1alpha4
	case schemasv1alpha4.SchemeGroupVersion.WithResource("datatypes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Schemas().V1alpha4().DataTypes().Informer()}, nil
	case schemasv1alpha4.SchemeGroupVersion.WithResource("migrations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Schemas().V1alpha4().Migrations().Informer()}, nil
	case schemasv1alpha4.SchemeGroupVersion.WithResource("tables"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Schemas().V1alpha4().Tables().Informer()}, nil
	case schemasv1alpha4.SchemeGroupVersion.WithResource("views"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Schemas().V1alpha4().Views().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
