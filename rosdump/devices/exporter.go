package devices

import (
	"context"
	"fmt"
	"io"

	"git.ecadlabs.com/ecad/rostools/rosdump/config"
	"github.com/sirupsen/logrus"
)

type Metadata map[string]interface{}

type Exporter interface {
	Export(context.Context) (io.ReadCloser, Metadata, error)
	Metadata() Metadata // For logging purposes
}

type NewExporterFunc func(config.Options, *logrus.Logger) (Exporter, error)

var registry = make(map[string]NewExporterFunc)

func registerExporter(name string, fn NewExporterFunc) {
	registry[name] = fn
}

func NewExporter(name string, options config.Options, logger *logrus.Logger) (Exporter, error) {
	if fn, ok := registry[name]; ok {
		return fn(options, logger)
	}

	return nil, fmt.Errorf("Unknown exporter driver: `%s'", name)
}
