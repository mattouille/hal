package hal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
)

// Handler constants
const (
	HEAR    = "HEAR"
	RESPOND = "RESPOND"
	TOPIC   = "TOPIC"
	ENTER   = "ENTER"
	LEAVE   = "LEAVE"
)

var (
	// Config is a global config
	Config = newConfig()
	// Logger is a global logger
	Logger = newLogger()
	// Router is a global HTTP muxer
	Router = newRouter()
	// HealthStatus is a global detailed status struct
	HealthStatus = newHealthStatus()
)

// New returns a Robot instance.
func New() (*Robot, error) {
	return NewRobot()
}

// Hear a message
func Hear(pattern string, fn func(res *Response) error) handler {
	return &Handler{Method: HEAR, Pattern: pattern, Run: fn}
}

// Respond creates a new listener for Respond messages
func Respond(pattern string, fn func(res *Response) error) handler {
	return &Handler{Method: RESPOND, Pattern: pattern, Run: fn}
}

// Topic returns a new listener for Topic messages
func Topic(pattern string, fn func(res *Response) error) handler {
	return &Handler{Method: TOPIC, Run: fn}
}

// Enter returns a new listener for Enter messages
func Enter(fn func(res *Response) error) handler {
	return &Handler{Method: ENTER, Run: fn}
}

// Leave creates a new listener for Leave messages
func Leave(fn func(res *Response) error) handler {
	return &Handler{Method: LEAVE, Run: fn}
}

// Close shuts down the robot. Unused?
func Close() error {
	return nil
}

// Creates a Logrus logger
func newLogger() *logrus.Logger {
	level, err := logrus.ParseLevel(Config.LogLevel)
	if err != nil {
		panic(err)
	}
	logger := logrus.New()
	logger.Level = level

	return logger
}

// newRouter initializes a new http.ServeMux and sets up several default routes
func newRouter() *http.ServeMux {
	router := http.NewServeMux()

	// A more human oriented healthcheck
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		healthJSON, err := json.Marshal(&HealthStatus)
		if err != nil {
			Logger.Fatal("Health cannot be determined")
		}
		fmt.Fprintf(w, "%s", healthJSON)
	})

	// Kubernetes Liveness and Readiness compatible endpoint
	router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		healthJSON, err := json.Marshal(evaluateHealth())
		if err != nil {
			Logger.Fatal("Health cannot be determined")
		}
		fmt.Fprintf(w, "%s", healthJSON)
	})

	router.HandleFunc("/hal/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server time is: %s\n", time.Now().UTC())
	})

	if Config.DebugEndpoint == true {
		// Robot Config
		router.HandleFunc("/hal/debug/config/robot", func(w http.ResponseWriter, r *http.Request) {
			configJSON, err := json.Marshal(&Config)
			if err != nil {
				Logger.Fatal("Robot config cannot be loaded.")
			}
			fmt.Fprintf(w, "%s", configJSON)
		})
	}

	return router
}
