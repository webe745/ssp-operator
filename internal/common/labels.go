package common

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	ssp "kubevirt.io/ssp-operator/api/v1beta2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	AppKubernetesNameLabel      = "app.kubernetes.io/name"
	AppKubernetesPartOfLabel    = "app.kubernetes.io/part-of"
	AppKubernetesVersionLabel   = "app.kubernetes.io/version"
	AppKubernetesManagedByLabel = "app.kubernetes.io/managed-by"
	AppKubernetesComponentLabel = "app.kubernetes.io/component"

	AppComponentTektonPipelines       AppComponent = "tektonPipelines"
	AppComponentTektonTasks           AppComponent = "tektonTasks"
	AppKubernetesManagedByValue       string       = "ssp-operator"
	TektonAppKubernetesManagedByValue string       = "tekton-tasks-operator"
)

type AppComponent string

func (a AppComponent) String() string {
	return string(a)
}

const (
	AppComponentMonitoring AppComponent = "monitoring"
	AppComponentSchedule   AppComponent = "schedule"
	AppComponentTemplating AppComponent = "templating"
)

// AddAppLabels to the provided obj
// Name will translate into the AppKubernetesNameLabel
// Component will translate into the AppKubernetesComponentLabel
// Instance wide labels will be taken from the request if available
func AddAppLabels(requestInstance *ssp.SSP, name string, component AppComponent, obj client.Object) client.Object {
	labels := getOrCreateLabels(obj)
	addInstanceLabels(requestInstance, labels)

	labels[AppKubernetesNameLabel] = name
	labels[AppKubernetesComponentLabel] = component.String()
	labels[AppKubernetesManagedByLabel] = AppKubernetesManagedByValue

	return obj
}

func getOrCreateLabels(obj client.Object) map[string]string {
	labels := obj.GetLabels()
	if labels == nil {
		labels = map[string]string{}
		obj.SetLabels(labels)
	}
	return labels
}

func addInstanceLabels(requestInstance *ssp.SSP, to map[string]string) {
	if requestInstance.Labels == nil {
		return
	}

	copyLabel(requestInstance.Labels, to, AppKubernetesPartOfLabel)
	copyLabel(requestInstance.Labels, to, AppKubernetesVersionLabel)
}

func copyLabel(from, to map[string]string, key string) {
	to[key] = from[key]
}

func GetAppNameSelector(name string) (labels.Selector, error) {
	appNameRequirement, err := labels.NewRequirement(AppKubernetesNameLabel, selection.Equals, []string{name})
	if err != nil {
		return nil, err
	}
	return labels.NewSelector().Add(*appNameRequirement), nil
}
