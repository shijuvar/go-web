package cloudendpoint

import (
	"log"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

// Register the API endpoints
func init() {
	taskService := &TaskService{}
	// Adds the TaskService to the server.
	api, err := endpoints.RegisterService(
		taskService,
		"tasks",
		"v1",
		"Tasks API",
		true,
	)
	if err != nil {
		log.Fatalf("Register service: %v", err)
	}

	// Get ServiceMethod's MethodInfo for List method
	info := api.MethodByName("List").Info()
	// Provide values to MethodInfo - name, HTTP method, and path.
	info.Name, info.HTTPMethod, info.Path = "listTasks", "GET", "tasks"

	// Get ServiceMethod's MethodInfo for Add method
	info = api.MethodByName("Add").Info()
	info.Name, info.HTTPMethod, info.Path = "addTask", "POST", "tasks"
	// Calls DefaultServer's HandleHttp method using default serve mux
	endpoints.HandleHTTP()
}
