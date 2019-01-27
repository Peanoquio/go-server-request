package interfaces

// ServerCallInterface is the interface that needs to be implemented
type ServerCallInterface interface {
	Get(url string, jsonPayload map[string]interface{}) (map[string]interface{}, error)
	Post(url string, jsonPayload map[string]interface{}) (map[string]interface{}, error)
	Put(url string, jsonPayload map[string]interface{}) (map[string]interface{}, error)
	Patch(url string, jsonPayload map[string]interface{}) (map[string]interface{}, error)
	Delete(url string, jsonPayload map[string]interface{}) (map[string]interface{}, error)
}
