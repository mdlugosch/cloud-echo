package main
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)
// Response definiert das Datenformat unserer JSON-Antwort
type Response struct {
	Message   string    `json:"message"`
	Hostname  string    `json:"hostname"`
	Timestamp time.Time `json:"timestamp"`
	Status    string    `json:"status"`
}
func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Hostname des Kubernetes-Pods ermitteln
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unbekannt"
	}
	data := Response{
		Message:   "Tekton CI/CD Pipeline erfolgreich durchlaufen!",
		Hostname:  hostname,
		Timestamp: time.Now().UTC(),
		Status:    "Healthy",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
func main() {
	http.HandleFunc("/", echoHandler)
	fmt.Println("Server lauscht auf Port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {		log.Fatalf("Fehler beim Starten des Servers: %s\n", err)
	}
}

