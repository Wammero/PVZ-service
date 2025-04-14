package metrics

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	RequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Общее количество HTTP-запросов",
		},
		[]string{"method", "path", "status"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Длительность HTTP-запросов",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	// Бизнесовые метрики
	CreatedPVZ = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "created_pvz_total",
			Help: "Количество созданных ПВЗ",
		},
	)

	CreatedOrders = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "created_orders_total",
			Help: "Количество созданных приёмок заказов",
		},
	)

	AddedProducts = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "added_products_total",
			Help: "Количество добавленных товаров",
		},
	)
)

func init() {
	prometheus.MustRegister(RequestCount)
	prometheus.MustRegister(RequestDuration)
	prometheus.MustRegister(CreatedPVZ)
	prometheus.MustRegister(CreatedOrders)
	prometheus.MustRegister(AddedProducts)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		_ = http.ListenAndServe(":9000", nil)
	}()
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rr, r)

		duration := time.Since(start).Seconds()

		RequestCount.WithLabelValues(r.Method, r.URL.Path, http.StatusText(rr.statusCode)).Inc()
		RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}

func Handler() http.Handler {
	return promhttp.Handler()
}
