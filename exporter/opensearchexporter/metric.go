package opensearchexporter

import (
	"context"

	"github.com/opensearch-project/opensearch-go/v4"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

type metricExporter struct {
	client       *opensearch.Client
	Index        string
	bulkAction   string
	model        mappingModel
	httpSettings confighttp.ClientConfig
	telemetry    component.TelemetrySettings
}

func newMetricExporter(cfg *Config, set exporter.Settings) *metricExporter {
	model := &encodeModel{
		dedup:             cfg.Dedup,
		dedot:             cfg.Dedot,
		sso:               cfg.Mode == MappingSS4O.String(),
		flattenAttributes: cfg.Mode == MappingFlattenAttributes.String(),
		timestampField:    cfg.TimestampField,
		unixTime:          cfg.UnixTimestamp,
		dataset:           cfg.Dataset,
		namespace:         cfg.Namespace,
	}

	return &metricExporter{
		telemetry:    set.TelemetrySettings,
		Index:        cfg.MetricsIndex,
		bulkAction:   cfg.BulkAction,
		httpSettings: cfg.ClientConfig,
		model:        model,
	}
}

func (l *metricExporter) Start(ctx context.Context,
	host component.Host) error {
	httpClient, err := l.httpSettings.ToClient(ctx, host, l.telemetry)
	if err != nil {
		return err
	}
	client, err := newOpenSearchClient(l.httpSettings.Endpoint,
		httpClient, l.telemetry.Logger)
	if err != nil {
		return err
	}
	l.client = client
	return nil
}

func (l *metricExporter) pushMetricData(ctx context.Context, ld pmetric.Metrics) error {
	// TODO: implement the indexer
	return nil
}
