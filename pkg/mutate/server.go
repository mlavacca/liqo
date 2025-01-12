package mutate

import (
	"fmt"
	"io/ioutil"
	"k8s.io/klog"
	"net/http"
	"time"
)

type MutationConfig struct {
	CertFile string
	KeyFile  string
}

type MutationServer struct {
	mux    *http.ServeMux
	server *http.Server

	config *MutationConfig
}

func NewMutationServer(c *MutationConfig) (*MutationServer, error) {
	s := &MutationServer{}
	s.config = c

	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/mutate", s.handleMutate)

	s.server = &http.Server{
		Addr:           ":8443",
		Handler:        s.mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1048576
	}

	return s, nil
}

func (s *MutationServer) handleMutate(w http.ResponseWriter, r *http.Request) {
	// read the body / request
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		s.sendError(err, w)
		return
	}

	// mutate the request
	mutated, err := s.Mutate(body)
	if err != nil {
		s.sendError(err, w)
		return
	}

	// and write it back
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(mutated)

	if err := r.Body.Close(); err != nil {
		klog.Error("error in body closing")
	}
}

func (s *MutationServer) sendError(err error, w http.ResponseWriter) {
	klog.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = fmt.Fprintf(w, "%s", err)
}

func (s *MutationServer) Serve() {
	klog.Fatal(s.server.ListenAndServeTLS(s.config.CertFile, s.config.KeyFile))
}
