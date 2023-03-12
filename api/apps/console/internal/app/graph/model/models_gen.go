// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	v11 "github.com/kloudlite/operator/apis/crds/v1"
	"github.com/kloudlite/operator/pkg/operator"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AppSpec struct {
	Frozen         *bool                  `json:"frozen"`
	Interception   *AppSpecInterception   `json:"interception"`
	NodeSelector   map[string]interface{} `json:"nodeSelector"`
	Services       []*AppSpecServices     `json:"services"`
	Tolerations    []*AppSpecTolerations  `json:"tolerations"`
	Containers     []*AppSpecContainers   `json:"containers"`
	Hpa            *AppSpecHpa            `json:"hpa"`
	Region         string                 `json:"region"`
	Replicas       *int                   `json:"replicas"`
	ServiceAccount *string                `json:"serviceAccount"`
}

type AppSpecContainers struct {
	ResourceCPU     *AppSpecContainersResourceCPU    `json:"resourceCpu"`
	ResourceMemory  *AppSpecContainersResourceMemory `json:"resourceMemory"`
	Args            []*string                        `json:"args"`
	Command         []*string                        `json:"command"`
	Env             []*AppSpecContainersEnv          `json:"env"`
	Image           string                           `json:"image"`
	LivenessProbe   *AppSpecContainersLivenessProbe  `json:"livenessProbe"`
	EnvFrom         []*AppSpecContainersEnvFrom      `json:"envFrom"`
	ImagePullPolicy *string                          `json:"imagePullPolicy"`
	Name            string                           `json:"name"`
	ReadinessProbe  *AppSpecContainersReadinessProbe `json:"readinessProbe"`
	Volumes         []*AppSpecContainersVolumes      `json:"volumes"`
}

type AppSpecContainersEnv struct {
	Key     string  `json:"key"`
	RefKey  *string `json:"refKey"`
	RefName *string `json:"refName"`
	Type    *string `json:"type"`
	Value   *string `json:"value"`
}

type AppSpecContainersEnvFrom struct {
	RefName string `json:"refName"`
	Type    string `json:"type"`
}

type AppSpecContainersEnvFromIn struct {
	RefName string `json:"refName"`
	Type    string `json:"type"`
}

type AppSpecContainersEnvIn struct {
	Key     string  `json:"key"`
	RefKey  *string `json:"refKey"`
	RefName *string `json:"refName"`
	Type    *string `json:"type"`
	Value   *string `json:"value"`
}

type AppSpecContainersIn struct {
	ResourceCPU     *AppSpecContainersResourceCPUIn    `json:"resourceCpu"`
	ResourceMemory  *AppSpecContainersResourceMemoryIn `json:"resourceMemory"`
	Args            []*string                          `json:"args"`
	Command         []*string                          `json:"command"`
	Env             []*AppSpecContainersEnvIn          `json:"env"`
	Image           string                             `json:"image"`
	LivenessProbe   *AppSpecContainersLivenessProbeIn  `json:"livenessProbe"`
	EnvFrom         []*AppSpecContainersEnvFromIn      `json:"envFrom"`
	ImagePullPolicy *string                            `json:"imagePullPolicy"`
	Name            string                             `json:"name"`
	ReadinessProbe  *AppSpecContainersReadinessProbeIn `json:"readinessProbe"`
	Volumes         []*AppSpecContainersVolumesIn      `json:"volumes"`
}

type AppSpecContainersLivenessProbe struct {
	FailureThreshold *int                                   `json:"failureThreshold"`
	HTTPGet          *AppSpecContainersLivenessProbeHTTPGet `json:"httpGet"`
	InitialDelay     *int                                   `json:"initialDelay"`
	Interval         *int                                   `json:"interval"`
	Shell            *AppSpecContainersLivenessProbeShell   `json:"shell"`
	TCP              *AppSpecContainersLivenessProbeTCP     `json:"tcp"`
	Type             string                                 `json:"type"`
}

type AppSpecContainersLivenessProbeHTTPGet struct {
	HTTPHeaders map[string]interface{} `json:"httpHeaders"`
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
}

type AppSpecContainersLivenessProbeHTTPGetIn struct {
	HTTPHeaders map[string]interface{} `json:"httpHeaders"`
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
}

type AppSpecContainersLivenessProbeIn struct {
	FailureThreshold *int                                     `json:"failureThreshold"`
	HTTPGet          *AppSpecContainersLivenessProbeHTTPGetIn `json:"httpGet"`
	InitialDelay     *int                                     `json:"initialDelay"`
	Interval         *int                                     `json:"interval"`
	Shell            *AppSpecContainersLivenessProbeShellIn   `json:"shell"`
	TCP              *AppSpecContainersLivenessProbeTCPIn     `json:"tcp"`
	Type             string                                   `json:"type"`
}

type AppSpecContainersLivenessProbeShell struct {
	Command []*string `json:"command"`
}

type AppSpecContainersLivenessProbeShellIn struct {
	Command []*string `json:"command"`
}

type AppSpecContainersLivenessProbeTCP struct {
	Port int `json:"port"`
}

type AppSpecContainersLivenessProbeTCPIn struct {
	Port int `json:"port"`
}

type AppSpecContainersReadinessProbe struct {
	InitialDelay     *int                                    `json:"initialDelay"`
	Interval         *int                                    `json:"interval"`
	Shell            *AppSpecContainersReadinessProbeShell   `json:"shell"`
	TCP              *AppSpecContainersReadinessProbeTCP     `json:"tcp"`
	Type             string                                  `json:"type"`
	FailureThreshold *int                                    `json:"failureThreshold"`
	HTTPGet          *AppSpecContainersReadinessProbeHTTPGet `json:"httpGet"`
}

type AppSpecContainersReadinessProbeHTTPGet struct {
	HTTPHeaders map[string]interface{} `json:"httpHeaders"`
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
}

type AppSpecContainersReadinessProbeHTTPGetIn struct {
	HTTPHeaders map[string]interface{} `json:"httpHeaders"`
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
}

type AppSpecContainersReadinessProbeIn struct {
	InitialDelay     *int                                      `json:"initialDelay"`
	Interval         *int                                      `json:"interval"`
	Shell            *AppSpecContainersReadinessProbeShellIn   `json:"shell"`
	TCP              *AppSpecContainersReadinessProbeTCPIn     `json:"tcp"`
	Type             string                                    `json:"type"`
	FailureThreshold *int                                      `json:"failureThreshold"`
	HTTPGet          *AppSpecContainersReadinessProbeHTTPGetIn `json:"httpGet"`
}

type AppSpecContainersReadinessProbeShell struct {
	Command []*string `json:"command"`
}

type AppSpecContainersReadinessProbeShellIn struct {
	Command []*string `json:"command"`
}

type AppSpecContainersReadinessProbeTCP struct {
	Port int `json:"port"`
}

type AppSpecContainersReadinessProbeTCPIn struct {
	Port int `json:"port"`
}

type AppSpecContainersResourceCPU struct {
	Max *string `json:"max"`
	Min *string `json:"min"`
}

type AppSpecContainersResourceCPUIn struct {
	Max *string `json:"max"`
	Min *string `json:"min"`
}

type AppSpecContainersResourceMemory struct {
	Max *string `json:"max"`
	Min *string `json:"min"`
}

type AppSpecContainersResourceMemoryIn struct {
	Max *string `json:"max"`
	Min *string `json:"min"`
}

type AppSpecContainersVolumes struct {
	Items     []*AppSpecContainersVolumesItems `json:"items"`
	MountPath string                           `json:"mountPath"`
	RefName   string                           `json:"refName"`
	Type      string                           `json:"type"`
}

type AppSpecContainersVolumesIn struct {
	Items     []*AppSpecContainersVolumesItemsIn `json:"items"`
	MountPath string                             `json:"mountPath"`
	RefName   string                             `json:"refName"`
	Type      string                             `json:"type"`
}

type AppSpecContainersVolumesItems struct {
	Key      string  `json:"key"`
	FileName *string `json:"fileName"`
}

type AppSpecContainersVolumesItemsIn struct {
	Key      string  `json:"key"`
	FileName *string `json:"fileName"`
}

type AppSpecHpa struct {
	Enabled         *bool `json:"enabled"`
	MaxReplicas     *int  `json:"maxReplicas"`
	MinReplicas     *int  `json:"minReplicas"`
	ThresholdCPU    *int  `json:"thresholdCpu"`
	ThresholdMemory *int  `json:"thresholdMemory"`
}

type AppSpecHpaIn struct {
	Enabled         *bool `json:"enabled"`
	MaxReplicas     *int  `json:"maxReplicas"`
	MinReplicas     *int  `json:"minReplicas"`
	ThresholdCPU    *int  `json:"thresholdCpu"`
	ThresholdMemory *int  `json:"thresholdMemory"`
}

type AppSpecIn struct {
	Frozen         *bool                   `json:"frozen"`
	Interception   *AppSpecInterceptionIn  `json:"interception"`
	NodeSelector   map[string]interface{}  `json:"nodeSelector"`
	Services       []*AppSpecServicesIn    `json:"services"`
	Tolerations    []*AppSpecTolerationsIn `json:"tolerations"`
	Containers     []*AppSpecContainersIn  `json:"containers"`
	Hpa            *AppSpecHpaIn           `json:"hpa"`
	Region         string                  `json:"region"`
	Replicas       *int                    `json:"replicas"`
	ServiceAccount *string                 `json:"serviceAccount"`
}

type AppSpecInterception struct {
	Enabled   *bool  `json:"enabled"`
	ForDevice string `json:"forDevice"`
}

type AppSpecInterceptionIn struct {
	Enabled   *bool  `json:"enabled"`
	ForDevice string `json:"forDevice"`
}

type AppSpecServices struct {
	Name       *string `json:"name"`
	Port       int     `json:"port"`
	TargetPort *int    `json:"targetPort"`
	Type       *string `json:"type"`
}

type AppSpecServicesIn struct {
	Name       *string `json:"name"`
	Port       int     `json:"port"`
	TargetPort *int    `json:"targetPort"`
	Type       *string `json:"type"`
}

type AppSpecTolerations struct {
	Effect            *string `json:"effect"`
	Key               *string `json:"key"`
	Operator          *string `json:"operator"`
	TolerationSeconds *int    `json:"tolerationSeconds"`
	Value             *string `json:"value"`
}

type AppSpecTolerationsIn struct {
	Effect            *string `json:"effect"`
	Key               *string `json:"key"`
	Operator          *string `json:"operator"`
	TolerationSeconds *int    `json:"tolerationSeconds"`
	Value             *string `json:"value"`
}

type Check struct {
	Status     *bool   `json:"status"`
	Message    *string `json:"message"`
	Generation *int    `json:"generation"`
}

type ManagedResources struct {
	Enabled    *bool                 `json:"enabled"`
	Kind       *string               `json:"kind"`
	ObjectMeta *v1.ObjectMeta        `json:"metadata"`
	Overrides  *v11.JsonPatch        `json:"overrides"`
	Spec       *ManagedResourcesSpec `json:"spec"`
	Status     *operator.Status      `json:"status"`
	APIVersion *string               `json:"apiVersion"`
}

type ManagedResourcesIn struct {
	Enabled    *bool                   `json:"enabled"`
	Kind       *string                 `json:"kind"`
	ObjectMeta *v1.ObjectMeta          `json:"metadata"`
	Overrides  *v11.JsonPatch          `json:"overrides"`
	Spec       *ManagedResourcesSpecIn `json:"spec"`
	APIVersion *string                 `json:"apiVersion"`
}

type ManagedResourcesSpec struct {
	Inputs   map[string]interface{}        `json:"inputs"`
	MresKind *ManagedResourcesSpecMresKind `json:"mresKind"`
	MsvcRef  *ManagedResourcesSpecMsvcRef  `json:"msvcRef"`
}

type ManagedResourcesSpecIn struct {
	Inputs   map[string]interface{}          `json:"inputs"`
	MresKind *ManagedResourcesSpecMresKindIn `json:"mresKind"`
	MsvcRef  *ManagedResourcesSpecMsvcRefIn  `json:"msvcRef"`
}

type ManagedResourcesSpecMresKind struct {
	Kind string `json:"kind"`
}

type ManagedResourcesSpecMresKindIn struct {
	Kind string `json:"kind"`
}

type ManagedResourcesSpecMsvcRef struct {
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind"`
	Name       string  `json:"name"`
}

type ManagedResourcesSpecMsvcRefIn struct {
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind"`
	Name       string  `json:"name"`
}

type ManagedServiceSpec struct {
	Tolerations  []*ManagedServiceSpecTolerations `json:"tolerations"`
	Inputs       map[string]interface{}           `json:"inputs"`
	MsvcKind     *ManagedServiceSpecMsvcKind      `json:"msvcKind"`
	NodeSelector map[string]interface{}           `json:"nodeSelector"`
	Region       string                           `json:"region"`
}

type ManagedServiceSpecIn struct {
	Tolerations  []*ManagedServiceSpecTolerationsIn `json:"tolerations"`
	Inputs       map[string]interface{}             `json:"inputs"`
	MsvcKind     *ManagedServiceSpecMsvcKindIn      `json:"msvcKind"`
	NodeSelector map[string]interface{}             `json:"nodeSelector"`
	Region       string                             `json:"region"`
}

type ManagedServiceSpecMsvcKind struct {
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind"`
}

type ManagedServiceSpecMsvcKindIn struct {
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind"`
}

type ManagedServiceSpecTolerations struct {
	Effect            *string `json:"effect"`
	Key               *string `json:"key"`
	Operator          *string `json:"operator"`
	TolerationSeconds *int    `json:"tolerationSeconds"`
	Value             *string `json:"value"`
}

type ManagedServiceSpecTolerationsIn struct {
	Effect            *string `json:"effect"`
	Key               *string `json:"key"`
	Operator          *string `json:"operator"`
	TolerationSeconds *int    `json:"tolerationSeconds"`
	Value             *string `json:"value"`
}

type Patch struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type PatchIn struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type RouterSpec struct {
	RateLimit       *RouterSpecRateLimit `json:"rateLimit"`
	Region          *string              `json:"region"`
	Routes          []*RouterSpecRoutes  `json:"routes"`
	BasicAuth       *RouterSpecBasicAuth `json:"basicAuth"`
	Cors            *RouterSpecCors      `json:"cors"`
	Domains         []*string            `json:"domains"`
	HTTPS           *RouterSpecHTTPS     `json:"https"`
	MaxBodySizeInMb *int                 `json:"maxBodySizeInMB"`
}

type RouterSpecBasicAuth struct {
	Enabled    bool    `json:"enabled"`
	SecretName *string `json:"secretName"`
	Username   *string `json:"username"`
}

type RouterSpecBasicAuthIn struct {
	Enabled    bool    `json:"enabled"`
	SecretName *string `json:"secretName"`
	Username   *string `json:"username"`
}

type RouterSpecCors struct {
	Origins          []*string `json:"origins"`
	AllowCredentials *bool     `json:"allowCredentials"`
	Enabled          *bool     `json:"enabled"`
}

type RouterSpecCorsIn struct {
	Origins          []*string `json:"origins"`
	AllowCredentials *bool     `json:"allowCredentials"`
	Enabled          *bool     `json:"enabled"`
}

type RouterSpecHTTPS struct {
	Enabled       bool  `json:"enabled"`
	ForceRedirect *bool `json:"forceRedirect"`
}

type RouterSpecHTTPSIn struct {
	Enabled       bool  `json:"enabled"`
	ForceRedirect *bool `json:"forceRedirect"`
}

type RouterSpecIn struct {
	RateLimit       *RouterSpecRateLimitIn `json:"rateLimit"`
	Region          *string                `json:"region"`
	Routes          []*RouterSpecRoutesIn  `json:"routes"`
	BasicAuth       *RouterSpecBasicAuthIn `json:"basicAuth"`
	Cors            *RouterSpecCorsIn      `json:"cors"`
	Domains         []*string              `json:"domains"`
	HTTPS           *RouterSpecHTTPSIn     `json:"https"`
	MaxBodySizeInMb *int                   `json:"maxBodySizeInMB"`
}

type RouterSpecRateLimit struct {
	Enabled     *bool `json:"enabled"`
	Rpm         *int  `json:"rpm"`
	Rps         *int  `json:"rps"`
	Connections *int  `json:"connections"`
}

type RouterSpecRateLimitIn struct {
	Enabled     *bool `json:"enabled"`
	Rpm         *int  `json:"rpm"`
	Rps         *int  `json:"rps"`
	Connections *int  `json:"connections"`
}

type RouterSpecRoutes struct {
	Lambda  *string `json:"lambda"`
	Path    string  `json:"path"`
	Port    int     `json:"port"`
	Rewrite *bool   `json:"rewrite"`
	App     *string `json:"app"`
}

type RouterSpecRoutesIn struct {
	Lambda  *string `json:"lambda"`
	Path    string  `json:"path"`
	Port    int     `json:"port"`
	Rewrite *bool   `json:"rewrite"`
	App     *string `json:"app"`
}
