module kubevirt.io/ssp-operator/api

go 1.19

require (
	github.com/openshift/api v0.0.0-20230406152840-ce21e3fe5da2 // release-4.13
	k8s.io/apimachinery v0.26.4
	kubevirt.io/containerized-data-importer-api v1.57.0-rc1
	kubevirt.io/controller-lifecycle-operator-sdk/api v0.2.4
	sigs.k8s.io/controller-runtime v0.14.5
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/openshift/custom-resource-status v1.1.2 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/api v0.26.2 // indirect
	k8s.io/klog/v2 v2.80.1 // indirect
	k8s.io/utils v0.0.0-20221128185143-99ec85e7a448 // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
)

replace golang.org/x/net => golang.org/x/net v0.13.0
