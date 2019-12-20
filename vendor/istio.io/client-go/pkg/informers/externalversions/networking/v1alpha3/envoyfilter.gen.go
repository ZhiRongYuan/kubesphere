// Code generated by informer-gen. DO NOT EDIT.

package v1alpha3

import (
	time "time"

	networkingv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	versioned "istio.io/client-go/pkg/clientset/versioned"
	internalinterfaces "istio.io/client-go/pkg/informers/externalversions/internalinterfaces"
	v1alpha3 "istio.io/client-go/pkg/listers/networking/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EnvoyFilterInformer provides access to a shared informer and lister for
// EnvoyFilters.
type EnvoyFilterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha3.EnvoyFilterLister
}

type envoyFilterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEnvoyFilterInformer constructs a new informer for EnvoyFilter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEnvoyFilterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEnvoyFilterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEnvoyFilterInformer constructs a new informer for EnvoyFilter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEnvoyFilterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha3().EnvoyFilters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha3().EnvoyFilters(namespace).Watch(options)
			},
		},
		&networkingv1alpha3.EnvoyFilter{},
		resyncPeriod,
		indexers,
	)
}

func (f *envoyFilterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEnvoyFilterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *envoyFilterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv1alpha3.EnvoyFilter{}, f.defaultInformer)
}

func (f *envoyFilterInformer) Lister() v1alpha3.EnvoyFilterLister {
	return v1alpha3.NewEnvoyFilterLister(f.Informer().GetIndexer())
}