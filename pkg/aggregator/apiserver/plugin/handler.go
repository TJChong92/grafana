package plugin

import (
	"context"
	"net/http"
	"path"

	"github.com/grafana/grafana-plugin-sdk-go/backend"

	aggregationv0alpha1 "github.com/grafana/grafana/pkg/aggregator/apis/aggregation/v0alpha1"
)

type PluginClient interface {
	backend.QueryDataHandler
	backend.StreamHandler
	backend.AdmissionHandler
	backend.CallResourceHandler
}

type PluginContextProvider interface {
	GetPluginContext(ctx context.Context, pluginID, uid string) (backend.PluginContext, error)
}

type PluginHandler struct {
	mux      *http.ServeMux
	delegate http.Handler

	client                PluginClient
	pluginContextProvider PluginContextProvider

	dataplaneService aggregationv0alpha1.DataPlaneService
}

func NewPluginHandler(
	client PluginClient,
	dataplaneService aggregationv0alpha1.DataPlaneService,
	pluginContextProvider PluginContextProvider,
	delegate http.Handler,
) *PluginHandler {
	h := &PluginHandler{
		mux:                   http.NewServeMux(),
		delegate:              delegate,
		client:                client,
		pluginContextProvider: pluginContextProvider,
		dataplaneService:      dataplaneService,
	}
	h.registerRoutes()
	return h
}

func (h *PluginHandler) registerRoutes() {
	proxyPath := proxyPathBuilder(h.dataplaneService.Spec.Group, h.dataplaneService.Spec.Version)

	for _, service := range h.dataplaneService.Spec.Services {
		switch service.Type {
		case aggregationv0alpha1.AdmissionControlServiceType:
			// TODO: implement in future PR
		case aggregationv0alpha1.ConversionServiceType:
			// TODO: implement in future PR
		case aggregationv0alpha1.DataSourceProxyServiceType:
			// TODO: implement in future PR
		case aggregationv0alpha1.QueryServiceType:
			h.mux.Handle(proxyPath("/namespaces/{namespace}/connections/{uid}/query"), h.QueryDataHandler())
		case aggregationv0alpha1.RouteServiceType:
			// TODO: implement in future PR
		case aggregationv0alpha1.StreamServiceType:
			// TODO: implement in future PR
		}
	}

	// fallback to the delegate
	h.mux.Handle("/", h.delegate)
}

func (h *PluginHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.mux.ServeHTTP(w, req)
}

func proxyPathBuilder(group, version string) func(string) string {
	return func(suffix string) string {
		return path.Join("/apis", group, version, suffix)
	}
}
