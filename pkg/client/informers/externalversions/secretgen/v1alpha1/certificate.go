// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	secretgenv1alpha1 "github.com/k14s/secretgen-controller/pkg/apis/secretgen/v1alpha1"
	versioned "github.com/k14s/secretgen-controller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/k14s/secretgen-controller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/k14s/secretgen-controller/pkg/client/listers/secretgen/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CertificateInformer provides access to a shared informer and lister for
// Certificates.
type CertificateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CertificateLister
}

type certificateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCertificateInformer constructs a new informer for Certificate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCertificateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCertificateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCertificateInformer constructs a new informer for Certificate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCertificateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecretgenV1alpha1().Certificates(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecretgenV1alpha1().Certificates(namespace).Watch(options)
			},
		},
		&secretgenv1alpha1.Certificate{},
		resyncPeriod,
		indexers,
	)
}

func (f *certificateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCertificateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *certificateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&secretgenv1alpha1.Certificate{}, f.defaultInformer)
}

func (f *certificateInformer) Lister() v1alpha1.CertificateLister {
	return v1alpha1.NewCertificateLister(f.Informer().GetIndexer())
}