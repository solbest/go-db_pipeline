package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	env "github.com/caarlos0/env/v6"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// "github.com/prometheus/client_golang/prometheus/promauto"
)

type Config struct {
	REDIS RDSConfig `envPrefix:"REDIS_"`
}
type RDSConfig struct {
	DB_URL    string `env:"DB_URL"`
	DB_INDEX  int    `env:"DB_INDEX"`
	KEY_VALUE string `env:"KEY_VALUE"`
}

func randomName() string {
	// generate a random string of length 5
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 5)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func randomValue() float64 {
	// generate a random float between 0 and 100
	return rand.Float64() * 100
}

// type responseWriter struct {
// 	http.ResponseWriter
// 	statusCode int
// }

// func NewResponseWriter(w http.ResponseWriter) *responseWriter {
// 	return &responseWriter{w, http.StatusOK}
// }

// func (rw *responseWriter) WriteHeader(code int) {
// 	rw.statusCode = code
// 	rw.ResponseWriter.WriteHeader(code)
// }

// func prometheusMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		route := mux.CurrentRoute(r)
// 		path, _ := route.GetPathTemplate()

// 		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
// 		rw := NewResponseWriter(w)
// 		next.ServeHTTP(rw, r)

// 		statusCode := rw.statusCode

// 		responseStatus.WithLabelValues(strconv.Itoa(statusCode)).Inc()
// 		totalRequests.WithLabelValues(path).Inc()

// 		timer.ObserveDuration()
// 	})
// }
// var totalRequests = prometheus.NewCounterVec(
// 	prometheus.CounterOpts{
// 		Name: "http_requests_total",
// 		Help: "Number of get requests.",
// 	},
// 	[]string{"path"},
// )

// var responseStatus = prometheus.NewCounterVec(
// 	prometheus.CounterOpts{
// 		Name: "response_status",
// 		Help: "Status of HTTP response",
// 	},
// 	[]string{"status"},
// )

// var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
// 	Name: "http_response_time_seconds",
// 	Help: "Duration of HTTP requests.",
// }, []string{"path"})
var netBids = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "net",
	Subsystem: "prom",
	Name:      "bids",
	Help:      "The number of bids",
}, []string{"name"})
var netWins = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "net",
	Subsystem: "prom",
	Name:      "wins",
	Help:      "The number of wins",
}, []string{"name"})
var netSpends = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "net",
	Subsystem: "prom",
	Name:      "spends",
	Help:      "The amount spent",
}, []string{"name"})
var camBids = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "cam",
	Subsystem: "prom",
	Name:      "bids",
	Help:      "The number of bids",
}, []string{"name"})
var camWins = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "cam",
	Subsystem: "prom",
	Name:      "wins",
	Help:      "The number of wins",
}, []string{"name"})
var camSpends = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "cam",
	Subsystem: "prom",
	Name:      "spends",
	Help:      "The amount spent",
}, []string{"name"})

func init() {
	// prometheus.Register(totalRequests)
	// prometheus.Register(responseStatus)
	// prometheus.Register(httpDuration)
	prometheus.Register(netBids)
	prometheus.Register(netWins)
	prometheus.Register(netSpends)
	prometheus.Register(camBids)
	prometheus.Register(camWins)
	prometheus.Register(camSpends)
}

func bidsnetGen() {
	//Temp data area
	keys := []string{"advenue", "A2X", "Sonobi", "SmartyAds"}
	rand.Seed(time.Now().UnixNano())
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			index := rand.Intn(len(keys))
			name := keys[index]
			value := randomValue()
			netBids.With(prometheus.Labels{
				"name": name,
			}).Set(value)
			fmt.Printf("%s:%f \n", name, value)
		}
	}()
}
func winsnetGen() {
	//Temp data area
	keys := []string{"advenue", "A2X", "Sonobi", "SmartyAds"}
	rand.Seed(time.Now().UnixNano())
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			index := rand.Intn(len(keys))
			name := keys[index]
			value := randomValue()
			netWins.With(prometheus.Labels{
				"name": name,
			}).Set(value)
			fmt.Printf("%s:%f \n", name, value)
		}
	}()
}
func spendsnetGen() {
	//Temp data area
	keys := []string{"advenue", "A2X", "Sonobi", "SmartyAds"}
	rand.Seed(time.Now().UnixNano())
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			index := rand.Intn(len(keys))
			name := keys[index]
			value := randomValue()
			netSpends.With(prometheus.Labels{
				"name": name,
			}).Set(value)
			fmt.Printf("%s:%f \n", name, value)
		}
	}()
}
func bidscamGen() {
	//Temp data area
	keys := []string{"advenue", "A2X", "Sonobi", "SmartyAds"}
	rand.Seed(time.Now().UnixNano())
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			index := rand.Intn(len(keys))
			name := keys[index]
			value := randomValue()
			camBids.With(prometheus.Labels{
				"name": name,
			}).Set(value)
			fmt.Printf("%s:%f \n", name, value)
		}
	}()
}
func winscamGen() {
	//Temp data area
	keys := []string{"advenue", "A2X", "Sonobi", "SmartyAds"}
	rand.Seed(time.Now().UnixNano())
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			index := rand.Intn(len(keys))
			name := keys[index]
			value := randomValue()
			camWins.With(prometheus.Labels{
				"name": name,
			}).Set(value)
			fmt.Printf("%s:%f \n", name, value)
		}
	}()
}
func spendscamGen() {
	//Temp data area
	keys := []string{"advenue", "A2X", "Sonobi", "SmartyAds"}
	rand.Seed(time.Now().UnixNano())
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			index := rand.Intn(len(keys))
			name := keys[index]
			value := randomValue()
			camSpends.With(prometheus.Labels{
				"name": name,
			}).Set(value)
			fmt.Printf("%s:%f \n", name, value)
		}
	}()
}
func main() {
	//Load environment variable
	var config Config
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}
	err = env.Parse(&config)
	if err != nil {
		fmt.Println("Env load err:", err)
	}

	//Redis connection
	rds_db := redis.NewClient(&redis.Options{
		Addr:     config.REDIS.DB_URL,
		Password: "",
		DB:       0,
	})

	pong, err := rds_db.Ping().Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	rds_db.Do("SELECT", config.REDIS.DB_INDEX)
	fmt.Println("Connected to Redis:", pong)

	bidscamGen()
	bidsnetGen()
	winscamGen()
	winsnetGen()
	spendscamGen()
	spendsnetGen()

	// advenue, err := rds_db.HGet("advenue", "Bids_Network").Result()
	// if err != nil {
	// 	log.Fatalf("Fetch data error: %e", err)
	// }
	// value, err := strconv.ParseFloat(advenue, 64)
	// if err != nil {
	// 	fmt.Println("Error convert to float:", err)
	// 	return
	// }

	//Server
	router := mux.NewRouter()
	// router.Use(prometheusMiddleware)
	// Prometheus endpoint
	router.Path("/prom/net/bids").Handler(promhttp.Handler())
	router.Path("/prom/net/wins").Handler(promhttp.Handler())
	router.Path("/prom/net/spends").Handler(promhttp.Handler())
	router.Path("/prom/cam/bids").Handler(promhttp.Handler())
	router.Path("/prom/cam/wins").Handler(promhttp.Handler())
	router.Path("/prom/cam/spends").Handler(promhttp.Handler())

	// Serving static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Serving requests on port 8080")
	err = http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
