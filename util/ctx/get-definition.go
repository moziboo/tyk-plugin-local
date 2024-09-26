package ctx

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Definition represents the full API definition including configuration data.
type APIDefinition struct {
	Name       string                 `json:"name"`
	APIID      string                 `json:"api_id"`
	OrgID      string                 `json:"org_id"`
	ConfigData map[string]interface{} `json:"config_data"`
	Proxy      ProxyConfig            `json:"proxy"`
}

// Proxy details for how the API should proxy requests.
type ProxyConfig struct {
	ListenPath      string `json:"listen_path"`
	TargetURL       string `json:"target_url"`
	StripListenPath bool   `json:"strip_listen_path"`
}

// Logger for simple logging
var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

// readConfig reads configuration data from a JSON file and unmarshals into Definition.
func readConfig() (*APIDefinition, error) {
	var def APIDefinition
	data, err := os.ReadFile("apps/decrypt-test.json")
	if err != nil {
		logger.Println("Error reading config:", err)
		return nil, err
	}
	err = json.Unmarshal(data, &def)
	if err != nil {
		logger.Println("Error parsing config:", err)
		return nil, err
	}
	return &def, nil
}

// GetDefinition returns a Definition object containing configuration data.
func GetDefinition(r *http.Request) *APIDefinition {
	logger.Println("Request received:", r.URL.Path)
	def, err := readConfig()
	if err != nil {
		logger.Println("Error reading config:", err)
		return nil
	}
	return def
}
