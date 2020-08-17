package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

	s.logger.Info("Starting api server...")

	return http.ListenAndServe(s.config.Bindaddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

// routes
func (s *APIServer) main() http.HandlerFunc {

	//type request struct {
	//	name string
	//}

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "main")
	}
}

// router
func (s *APIServer) configureRouter() {
	s.router.HandleFunc(string("/main"), s.main())
}
