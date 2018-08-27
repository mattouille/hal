package hal

// HealthStatus struct
type healthStatus struct {
	AdapterStatus string
	Adapter       *string
	StoreStatus   string
	Store         *string
}

// HealthCheck struct
type healthCheck struct {
	Liveness  bool `json:"liveness"`
	Readiness bool `json:"readiness"`
	Shutdown  bool `json:"shutdown"`
}

// Creates a new instance of healthStatus
func newHealthStatus() *healthStatus {
	h := &healthStatus{
		AdapterStatus: "disconnected",
		Adapter:       &Config.AdapterName,
		StoreStatus:   "disconnected",
		Store:         &Config.StoreName,
	}

	return h
}

// evaluateHealth evaluates information and determines lifecycle effects
func evaluateHealth() *healthCheck {
	// Assume nothing, fail by default
	// Readiness is a bit irrelevant as we don't use a Service
	h := &healthCheck{
		Liveness:  false,
		Readiness: false,
		Shutdown:  false,
	}

	// If both critical components are connected, give it the thumbs up
	if HealthStatus.AdapterStatus == "connected" && HealthStatus.StoreStatus == "connected" {
		h.Liveness = true
		h.Readiness = true
	}
	// As long as the store is connected we can prevent the container from being killed
	if HealthStatus.StoreStatus == "connected" {
		h.Liveness = true
	}

	return h
}
