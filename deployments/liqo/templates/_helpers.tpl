{{/*
Expand the name of the chart.
*/}}
{{- define "liqo.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "liqo.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "liqo.chart" -}}
{{- printf "%s-%s" .Chart.Name (include "liqo.version" .) | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create version used to select the liqo version to be installed .
*/}}
{{- define "liqo.version" -}}
{{- if .Values.tag }}
{{- .Values.tag }}
{{- else if .Chart.AppVersion }}
{{- .Chart.AppVersion }}
{{- else }}
{{- fail "At least one between .Values.version and .Chart.AppVersion should be set" }}
{{- end }}
{{- end }}

{{/*
Create version used to select the liqo version to be installed.
*/}}
{{- define "liqo.suffix" -}}
{{- if or (eq .Values.tag "") (eq .Chart.AppVersion .Values.tag) }}
{{- print "" }}
{{- else }}
{{- print "-ci" }}
{{- end }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "liqo.labels" -}}
{{ include "liqo.selectorLabels" . }}
helm.sh/chart: {{ quote (include "liqo.chart" .) }}
app.kubernetes.io/version: {{ quote (include "liqo.version" .) }}
app.kubernetes.io/managed-by: {{ quote .Release.Service }}
{{- end }}

{{/*
Selector labels, it accepts a dict which contains fields "name" and "module"
*/}}
{{- define "liqo.selectorLabels" -}}
app.kubernetes.io/name: {{ quote .name }}
app.kubernetes.io/instance: {{ quote (printf "%s-%s" .Release.Name .name) }}
app.kubernetes.io/component: {{ quote .module }}
app.kubernetes.io/part-of: {{ quote (include "liqo.name" .) }}
{{- end }}

{{/*
Create a name prefixed with the chart name, it accepts a dict which contains the field "name".
*/}}
{{- define "liqo.prefixedName" -}}
{{- printf "%s-%s" (include "liqo.name" .) .name }}
{{- end }}

{{/*
Create the file name of a role starting from a prefix, it accepts a dict which contains the field "prefix".
*/}}
{{- define "liqo.role-filename" -}}
{{- printf "files/%s-%s" .prefix "Role.yaml" }}
{{- end }}

{{/*
Create the file name of a cluster role starting from a prefix, it accepts a dict which contains the field "prefix".
*/}}
{{- define "liqo.cluster-role-filename" -}}
{{- printf "files/%s-%s" .prefix "ClusterRole.yaml" }}
{{- end }}

{{/*
Gateway pod labels
*/}}
{{- define "liqo.gatewayPodLabels" -}}
net.liqo.io/gatewayPod: "true"
{{- end }}

{{/*
Auth pod labels
*/}}
{{- define "liqo.authServiceLabels" -}}
run: auth-service
{{- end }}

{{/*
Webhook pod labels
*/}}
{{- define "liqo.webhookServiceLabels" -}}
webhook.liqo.io/backend: "liqo-webhook"
{{- end }}

{{/*
Gateway service labels
*/}}
{{- define "liqo.gatewayServiceLabels" -}}
net.liqo.io/gateway: "true"
{{- end }}


