package bootstrap

import (
	"fmt"
	"net/http"

	"github.com/miti997/api-gateway/internal/logging"
	"github.com/miti997/api-gateway/internal/routing"
)

type Bootstraper interface {
	Bootstrap()
}

type DefaultBootstraper struct {
	configFilePath       string
	routesFilePath       string
	loggerConfigFilePath string
	serverConfig         *ServerConfig
	routesConfig         *RoutesConfig
	loggerConfig         *LoggerConfig
}

func NewDefaultBootstraper(cfp string, rfp string, lcfp string) (*DefaultBootstraper, error) {
	serverConfig := &ServerConfig{}
	if err := serverConfig.Load(cfp); err != nil {
		return nil, fmt.Errorf("could not load server config: %v", err)
	}

	routesConfig := &RoutesConfig{}
	if err := routesConfig.Load(rfp); err != nil {
		return nil, fmt.Errorf("could not load routes config: %v", err)
	}

	loggerConfig := &LoggerConfig{}
	if err := loggerConfig.Load(lcfp); err != nil {
		return nil, fmt.Errorf("could not load logger config: %v", err)
	}

	return &DefaultBootstraper{
		configFilePath:       cfp,
		routesFilePath:       rfp,
		loggerConfigFilePath: lcfp,
		serverConfig:         serverConfig,
		routesConfig:         routesConfig,
		loggerConfig:         loggerConfig,
	}, nil
}

func (b *DefaultBootstraper) Bootstrap() error {
	sm := http.NewServeMux()

	l, e := logging.NewDefaultLogger(b.loggerConfig.FilePath, b.loggerConfig.FileName, b.loggerConfig.MaxSizeMB)

	if e != nil {
		return e
	}

	for _, route := range b.routesConfig.Routes {
		r, e := routing.NewRoute(route.Request, route.In, route.Out, l)
		if e != nil {
			return e
		}
		sm.HandleFunc(r.GetPattern(), r.HandleRequest)
	}

	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	s := &http.Server{
		Addr:    b.serverConfig.Address,
		Handler: sm,
	}

	fmt.Println("API Gateway is starting...")

	s.ListenAndServe()

	return nil
}
