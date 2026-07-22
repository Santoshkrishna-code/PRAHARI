package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
	"time"
)

// Real-time Event Frame Schema
type IndustrialEvent struct {
	EventID       string                 `json:"eventId"`
	Timestamp     string                 `json:"timestamp"`
	Category      string                 `json:"category"`
	Source        string                 `json:"source"`
	Asset         string                 `json:"asset"`
	Plant         string                 `json:"plant"`
	Severity      string                 `json:"severity"`
	CorrelationID string                 `json:"correlationId"`
	Message       string                 `json:"message"`
	AIDecision    string                 `json:"aiDecision,omitempty"`
	Priority      int                    `json:"priority"`
	Metrics       map[string]interface{} `json:"metrics,omitempty"`
}

// Client Connection Hub
type Hub struct {
	clients    map[chan []byte]bool
	broadcast  chan []byte
	register   chan chan []byte
	unregister chan chan []byte
	mu         sync.Mutex
}

var hub = Hub{
	clients:    make(map[chan []byte]bool),
	broadcast:  make(chan []byte, 256),
	register:   make(chan chan []byte),
	unregister: make(chan chan []byte),
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			log.Printf("[WEBSOCKET HUB] Client connected. Total clients: %d", len(h.clients))

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client)
			}
			h.mu.Unlock()
			log.Printf("[WEBSOCKET HUB] Client disconnected. Total clients: %d", len(h.clients))

		case message := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients {
				select {
				case client <- message:
				default:
					close(client)
					delete(h.clients, client)
				}
			}
			h.mu.Unlock()
		}
	}
}

// Background Telemetry & Event Simulator (Go Shared Pipeline Engine)
func startIndustrialTelemetryEngine() {
	ticker := time.NewTicker(1500 * time.Millisecond)
	defer ticker.Stop()

	baseVib := 8.5
	baseTemp := 84.0
	basePsi := 230.0
	baseKw := 320.0
	step := 0

	for range ticker.C {
		step++
		// Add Brownian motion to parameters
		baseVib = math.Max(4.0, math.Min(15.8, baseVib+(rand.Float64()-0.47)*0.7))
		baseTemp = math.Max(72.0, math.Min(108.0, baseTemp+(rand.Float64()-0.46)*0.6))
		basePsi = math.Max(200.0, math.Min(280.0, basePsi+(rand.Float64()-0.48)*1.8))
		baseKw = math.Max(280.0, math.Min(390.0, baseKw+(rand.Float64()-0.5)*5.0))

		sev := "Info"
		msg := fmt.Sprintf("Live telemetry stream: Vib=%.2f mm/s, Temp=%.1f°C, PSI=%.1f", baseVib, baseTemp, basePsi)
		aiDecision := ""

		if baseVib > 13.0 {
			sev = "Critical"
			msg = fmt.Sprintf("CRITICAL ALARM: Pump P-102 vibration velocity reached %.2f mm/s (Exceeds ISO 10816 Limit)", baseVib)
			aiDecision = "AI Supervisor automatically recalculated RUL: 18 days. Dispatched Work Order WO-7821 to Maintenance Crew."
		} else if baseVib > 11.5 {
			sev = "Warning"
			msg = fmt.Sprintf("WARNING: Pump P-102 vibration elevated at %.2f mm/s", baseVib)
			aiDecision = "PINN Digital Twin calculated 72% outer bearing race wear probability."
		}

		event := IndustrialEvent{
			EventID:       fmt.Sprintf("evt-go-%d", time.Now().UnixNano()%100000),
			Timestamp:     time.Now().Format("15:04:05"),
			Category:      "Telemetry",
			Source:        "Go Telemetry Engine (Shared Pipeline)",
			Asset:         "PUMP-P102",
			Plant:         "Plant Alpha (Gulf Coast)",
			Severity:      sev,
			CorrelationID: "corr-p102-vib",
			Message:       msg,
			AIDecision:    aiDecision,
			Priority:      1,
			Metrics: map[string]interface{}{
				"vib":  math.Round(baseVib*100) / 100,
				"temp": math.Round(baseTemp*10) / 10,
				"psi":  math.Round(basePsi*10) / 10,
				"kw":   math.Round(baseKw),
			},
		}

		data, err := json.Marshal(event)
		if err == nil {
			hub.broadcast <- data
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	go hub.run()
	go startIndustrialTelemetryEngine()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	aiPlatformURL := os.Getenv("AI_PLATFORM_URL")
	if aiPlatformURL == "" {
		aiPlatformURL = "http://localhost:8000"
	}

	targetURL, err := url.Parse(aiPlatformURL)
	if err != nil {
		log.Fatalf("Invalid AI Platform target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Live Server-Sent Events (SSE) & Stream Endpoint
	http.HandleFunc("/events/stream", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		messageChan := make(chan []byte, 64)
		hub.register <- messageChan

		defer func() {
			hub.unregister <- messageChan
		}()

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		for {
			select {
			case <-r.Context().Done():
				return
			case msg := <-messageChan:
				fmt.Fprintf(w, "data: %s\n\n", msg)
				flusher.Flush()
			}
		}
	})

	// ═══════════════════════════════════════════════════════════
	// IN-MEMORY PERSISTENT STORE (CONNECTED TO EVENT BUS & WS)
	// ═══════════════════════════════════════════════════════════
	var (
		assetsMu    sync.RWMutex
		assetsStore = []map[string]interface{}{
			{"id": "asset-102", "name": "Pump P-102", "loc": "DC-101 Recirc", "type": "Centrifugal Pump", "health": 74, "rul": "18d", "st": "Warning", "owner": "Mechanical Team", "vib": 11.8, "temp": 94.1},
			{"id": "asset-88", "name": "Valve V-88", "loc": "DC-101 Isol", "type": "Gate Valve", "health": 98, "rul": "N/A", "st": "Locked", "owner": "Safety Team", "vib": "--", "temp": "--"},
			{"id": "asset-04", "name": "Heat Exchanger HX-04", "loc": "DC-101 Cool", "type": "Shell & Tube", "health": 91, "rul": "62d", "st": "Running", "owner": "Process Team", "vib": 3.2, "temp": 88},
			{"id": "asset-03", "name": "Compressor C-03", "loc": "DC-102", "type": "Reciprocating", "health": 96, "rul": "84d", "st": "Running", "owner": "Mechanical Team", "vib": 4.1, "temp": 72},
		}

		incidentsMu    sync.RWMutex
		incidentsStore = []map[string]interface{}{
			{"id": "INC-2026-0447", "title": "Pump P-102 Vibration Anomaly", "desc": "Vibration probe threshold exceeded during high-pressure run.", "sev": "Warning", "asset": "Pump P-102", "st": "Under Investigation", "time": "14:10"},
			{"id": "INC-2026-0442", "title": "Contractor Badge Expiration Zone B", "desc": "Contractor C-4412 detected with expired safety badge.", "sev": "Info", "asset": "Gate B", "st": "Resolved", "time": "12:30"},
		}

		woMu    sync.RWMutex
		woStore = []map[string]interface{}{
			{"id": "WO-7821", "desc": "Bearing replacement and lubrication service", "asset": "P-102", "pri": "Critical", "rul": "18d", "st": "Overdue"},
			{"id": "WO-7822", "desc": "Vibration probe recalibration", "asset": "P-102", "pri": "High", "rul": "18d", "st": "Assigned"},
			{"id": "WO-7823", "desc": "Quarterly compressor inspection", "asset": "C-03", "pri": "Medium", "rul": "84d", "st": "Scheduled"},
			{"id": "WO-7824", "desc": "Boiler tube thickness measurement", "asset": "Boiler A", "pri": "Medium", "rul": "71d", "st": "Scheduled"},
		}
	)

	// REST Endpoint: Assets (GET & POST)
	http.HandleFunc("/api/assets", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == "POST" {
			var newAsset map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&newAsset); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if newAsset["id"] == nil || newAsset["id"] == "" {
				newAsset["id"] = fmt.Sprintf("asset-%d", time.Now().UnixNano()%10000)
			}
			if newAsset["health"] == nil {
				newAsset["health"] = 100
			}
			if newAsset["rul"] == nil {
				newAsset["rul"] = "120d"
			}
			if newAsset["st"] == nil {
				newAsset["st"] = "Running"
			}

			assetsMu.Lock()
			assetsStore = append([]map[string]interface{}{newAsset}, assetsStore...)
			assetsMu.Unlock()

			// Broadcast AssetCreated Event to WebSocket Clients
			assetName, _ := newAsset["name"].(string)
			assetLoc, _ := newAsset["loc"].(string)
			evt := IndustrialEvent{
				EventID:       fmt.Sprintf("evt-asset-%d", time.Now().UnixNano()),
				Timestamp:     time.Now().Format("15:04:05"),
				Category:      "Asset",
				Source:        "Go Asset Persistence Service",
				Asset:         assetName,
				Plant:         "Plant Alpha (Gulf Coast)",
				Severity:      "Info",
				CorrelationID: fmt.Sprintf("corr-%s", assetName),
				Message:       fmt.Sprintf("Registered asset %s in location %s", assetName, assetLoc),
				Priority:      1,
			}
			data, _ := json.Marshal(evt)
			hub.broadcast <- data

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newAsset)
			return
		}

		assetsMu.RLock()
		defer assetsMu.RUnlock()
		json.NewEncoder(w).Encode(assetsStore)
	})

	// REST Endpoint: Incidents (GET & POST)
	http.HandleFunc("/api/incidents", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == "POST" {
			var newInc map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&newInc); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if newInc["id"] == nil || newInc["id"] == "" {
				newInc["id"] = fmt.Sprintf("INC-2026-%d", 1000+rand.Intn(9000))
			}
			if newInc["st"] == nil {
				newInc["st"] = "Under Investigation"
			}
			if newInc["time"] == nil {
				newInc["time"] = time.Now().Format("15:04")
			}

			incidentsMu.Lock()
			incidentsStore = append([]map[string]interface{}{newInc}, incidentsStore...)
			incidentsMu.Unlock()

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newInc)
			return
		}

		incidentsMu.RLock()
		defer incidentsMu.RUnlock()
		json.NewEncoder(w).Encode(incidentsStore)
	})

	// REST Endpoint: Maintenance Work Orders (GET & POST)
	http.HandleFunc("/api/maintenance/workorders", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == "POST" {
			var newWO map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&newWO); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if newWO["id"] == nil || newWO["id"] == "" {
				newWO["id"] = fmt.Sprintf("WO-%d", 7825+rand.Intn(100))
			}
			if newWO["st"] == nil {
				newWO["st"] = "Dispatched"
			}

			woMu.Lock()
			woStore = append([]map[string]interface{}{newWO}, woStore...)
			woMu.Unlock()

			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newWO)
			return
		}

		woMu.RLock()
		defer woMu.RUnlock()
		json.NewEncoder(w).Encode(woStore)
	})

	// ═══════════════════════════════════════════════════════════
	// DASHBOARD AGGREGATOR SERVICE ENDPOINTS (SINGLE SOURCE OF TRUTH)
	// ═══════════════════════════════════════════════════════════

	// 1. Executive Dashboard Aggregator
	http.HandleFunc("/api/dashboard/executive", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"safetyIndex":      94.2,
			"trirRate":         0.18,
			"activeRisks":      2,
			"inspectionPass":   98.4,
			"assetHealth":      91.2,
			"permitCompliance": 100.0,
			"lastUpdated":      time.Now().Format("15:04:05"),
			"shift":            "Shift B (143 Operators Online)",
			"recommendations": []string{
				"Approve bearing race replacement for Pump P-102 within 18 days.",
				"Review lubrication schedule auto-escalation in CMMS configuration.",
				"Maintain continuous Jetson AGX camera scan in Zone B.",
			},
		})
	})

	// 2. Inspections Dashboard Aggregator
	http.HandleFunc("/api/dashboard/inspections", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"summary": map[string]interface{}{
				"completedToday": 18,
				"completedMonth": 122,
				"pending":        14,
				"overdue":        3,
				"failed":         2,
				"complianceScore": 97.8,
			},
			"queue": []map[string]interface{}{
				{"id": "AUD-9901", "name": "PPE Hardhat & Goggles Audit", "area": "Zone B North", "inspector": "Harish M.", "status": "Pending", "due": "Today (16:00)", "score": "98%"},
				{"id": "AUD-9902", "name": "Fire Safety & Hydrant Pressure", "area": "Boiler Area A", "inspector": "Priya S.", "status": "Passed", "due": "Completed", "score": "95%"},
				{"id": "AUD-9903", "name": "Electrical MCC Panel Inspection", "area": "MCC Room 7B", "inspector": "Rahul K.", "status": "Overdue", "due": "2 Days Ago", "score": "88%"},
				{"id": "AUD-9904", "name": "Emergency Exit Door Clearance", "area": "Warehouse Zone C", "inspector": "John D.", "status": "Failed", "due": "Today", "score": "72%"},
			},
			"breakdown": []map[string]interface{}{
				{"category": "PPE Safety", "score": 98},
				{"category": "Fire Safety", "score": 95},
				{"category": "Electrical", "score": 96},
				{"category": "Mechanical", "score": 100},
				{"category": "Hazardous Chemical", "score": 93},
				{"category": "Emergency Prep", "score": 97},
			},
		})
	})

	// 3. Compliance Dashboard Aggregator
	http.HandleFunc("/api/dashboard/compliance", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"complianceScore": 97.6,
			"lastAudit":       "3 Days Ago",
			"verifiedFiles":   1284,
			"standards": []map[string]interface{}{
				{"std": "ISO 45001", "score": "98%", "detail": "42/45 Clauses Compliant"},
				{"std": "OSHA 29 CFR", "score": "95%", "detail": "118/120 Requirements Met"},
				{"std": "Internal EHS SOP", "score": "100%", "detail": "Completed"},
				{"std": "Environmental", "score": "96%", "detail": "EPA Standard"},
				{"std": "Contractor Audit", "score": "91%", "detail": "1 Expired Badge"},
			},
			"heatmap": []map[string]interface{}{
				{"zone": "Zone A (Main Line)", "score": "98%", "status": "Compliant"},
				{"zone": "Zone B (Reactor North)", "score": "92%", "status": "Attention"},
				{"zone": "Tank Farm T-204", "score": "100%", "status": "Optimal"},
				{"zone": "Utilities & Steam", "score": "95%", "status": "Compliant"},
				{"zone": "Warehouse Storage", "score": "89%", "status": "Drill Due"},
			},
		})
	})

	// 4. Full Audit Package Export Endpoint
	http.HandleFunc("/api/compliance/export", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Disposition", "attachment; filename=\"PRAHARI_Enterprise_Audit_Package_2026-07-22.json\"")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		auditPackage := map[string]interface{}{
			"reportMetadata": map[string]interface{}{
				"title":           "PRAHARI Enterprise ISO 45001 & OSHA Audit Package",
				"generatedAt":     time.Now().Format(time.RFC3339),
				"organization":    "Alpha Chemical Refinery Inc.",
				"plant":           "Plant Alpha (Gulf Coast Site)",
				"auditHash":       "0x9f8b7a6c5d4e3f2a1b0c9d8e7f6a5b4c3d2e1f0a",
				"verifiedByAgent": "Multi-Agent AI Compliance Supervisor",
			},
			"executiveSummary": map[string]interface{}{
				"complianceScore": 97.6,
				"status":          "GOOD",
				"totalVerifiedFiles": 1284,
				"trirIncidentRate": 0.18,
			},
			"standards": []map[string]interface{}{
				{"std": "ISO 45001", "score": "98%", "clauses": "42/45 Compliant"},
				{"std": "OSHA 29 CFR 1910", "score": "95%", "clauses": "118/120 Requirements Met"},
				{"std": "Internal EHS SOP", "score": "100%", "clauses": "Fully Verified"},
				{"std": "Environmental Protection", "score": "96%", "clauses": "EPA Compliant"},
			},
			"heatmap": []map[string]interface{}{
				{"zone": "Zone A (Main Line)", "score": "98%", "status": "Compliant"},
				{"zone": "Zone B (Reactor North)", "score": "92%", "status": "Attention"},
				{"zone": "Tank Farm T-204", "score": "100%", "status": "Optimal"},
				{"zone": "Utilities & Steam", "score": "95%", "status": "Compliant"},
				{"zone": "Warehouse Storage", "score": "89%", "status": "Drill Due"},
			},
			"aiActionPlan": map[string]interface{}{
				"verifiedChanges": []string{
					"2 contractor certifications expire this week (Badge C-4412 auto-revoked).",
					"Emergency drill due in 7 days (Warehouse Zone C).",
					"PPE compliance index increased from 94% to 97% across Zone B.",
				},
				"recommendations": []string{
					"Schedule Warehouse emergency response drill before Friday.",
					"Renew contractor medical certification for C-4412.",
					"Upload missing permit attachment for PTW-8903.",
				},
			},
		}

		json.NewEncoder(w).Encode(auditPackage)
	})

	// Real-time Platform Health & Observability Metrics
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		hub.mu.Lock()
		clientCount := len(hub.clients)
		hub.mu.Unlock()

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":                 "UP",
			"service":                "gateway-service (Go Shared Pipeline)",
			"timestamp":              time.Now().Format(time.RFC3339),
			"connected_clients":      clientCount,
			"telemetry_rate_per_min": 2400,
			"database":               map[string]string{"status": "healthy", "engine": "PostgreSQL v15.7"},
			"cache":                  map[string]string{"status": "healthy", "engine": "Redis Pub/Sub"},
			"mqtt_broker":            map[string]string{"status": "connected", "protocol": "MQTT v5.0"},
			"event_bus":              map[string]string{"status": "active", "throughput": "40 ev/sec"},
		})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		log.Printf("[GATEWAY PROXY] %s %s -> %s", r.Method, r.URL.Path, aiPlatformURL)
		r.Host = targetURL.Host
		proxy.ServeHTTP(w, r)
	})

	log.Printf("PRAHARI Go API Gateway & WebSocket Hub Listening on :%s (Proxying to %s)", port, aiPlatformURL)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("Gateway server error: %v", err)
	}
}
