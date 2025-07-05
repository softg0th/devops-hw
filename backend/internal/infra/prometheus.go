package infra

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "requests_total",
			Help: "Total number of the spam messages",
		})
)

func StartPrometheus(
	prometheusPort string,
) <-chan error {
	errCh := make(chan error, 1)

	http.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(RequestsTotal)

	go func() {
		srvErr := http.ListenAndServe(prometheusPort, nil)
		if srvErr != nil {
			errCh <- srvErr
		}
		close(errCh)
	}()

	return errCh
}
