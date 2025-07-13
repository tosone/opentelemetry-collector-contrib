// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package opensearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opensearchexporter"

import (
	"time"
)

type dataStream struct {
	Dataset   string `json:"dataset,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Type      string `json:"type,omitempty"`
}

type ssoSpanEvent struct {
	Attributes             map[string]any `json:"attributes"`
	DroppedAttributesCount uint32         `json:"droppedAttributesCount"`
	Name                   string         `json:"name"`
	ObservedTimestamp      *time.Time     `json:"observedTimestamp,omitempty"`
	Timestamp              *time.Time     `json:"@timestamp,omitempty"`
}

type ssoSpanLinks struct {
	Attributes             map[string]any `json:"attributes,omitempty"`
	SpanID                 string         `json:"spanId,omitempty"`
	TraceID                string         `json:"traceId,omitempty"`
	TraceState             string         `json:"traceState,omitempty"`
	DroppedAttributesCount uint32         `json:"droppedAttributesCount,omitempty"`
}

type ssoSpanContext struct {
	ParentSpanID string `json:"parentSpanId"`
	SpanID       string `json:"spanId"`
	TraceID      string `json:"traceId"`
}

type ssoSpan struct {
	Context              ssoSpanContext `json:"context"`
	Attributes           map[string]any `json:"attributes,omitempty"`
	EndTime              time.Time      `json:"endTime"`
	Events               []ssoSpanEvent `json:"events,omitempty"`
	InstrumentationScope struct {
		Attributes map[string]any `json:"attributes,omitempty"`
		Name       string         `json:"name"`
		SchemaURL  string         `json:"schemaUrl"`
		Version    string         `json:"version"`
	} `json:"instrumentationScope,omitempty"`
	Kind      string            `json:"kind"`
	Links     []ssoSpanLinks    `json:"links,omitempty"`
	Name      string            `json:"name"`
	Resource  map[string]string `json:"resource,omitempty"`
	StartTime time.Time         `json:"startTime"`
	Status    struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Timestamp  time.Time `json:"@timestamp"`
	TraceState string    `json:"traceState"`
}

type ssoRecord struct {
	Attributes           map[string]any `json:"attributes,omitempty"`
	Body                 string         `json:"body"`
	InstrumentationScope struct {
		Attributes map[string]any `json:"attributes,omitempty"`
		Name       string         `json:"name,omitempty"`
		SchemaURL  string         `json:"schemaUrl,omitempty"`
		Version    string         `json:"version,omitempty"`
	} `json:"instrumentationScope,omitempty"`
	ObservedTimestamp *time.Time        `json:"observedTimestamp,omitempty"`
	Resource          map[string]string `json:"resource,omitempty"`
	SchemaURL         string            `json:"schemaUrl,omitempty"`
	Severity          struct {
		Text   string `json:"text,omitempty"`
		Number int64  `json:"number,omitempty"`
	} `json:"severity"`
	SpanID    string     `json:"spanId,omitempty"`
	Timestamp *time.Time `json:"@timestamp"`
	TraceID   string     `json:"traceId,omitempty"`
}
