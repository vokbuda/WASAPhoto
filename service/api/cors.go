package api

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// applyCORSHandler applies a CORS policy to the router. CORS stands for Cross-Origin Resource Sharing: it's a security
// feature present in web browsers that blocks JavaScript requests going across different domains if not specified in a
// policy. This function sends the policy of this API server.
func applyCORSHandler(h http.Handler) http.Handler {
	return handlers.CORS(

		// handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin:*", "Authorization", "Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"}),
		handlers.AllowedOrigins([]string{"http://127.0.0.1", "http://localhost"}),
		handlers.AllowCredentials(),
		handlers.AllowedHeaders([]string{"Authorization", "authorization"}),
		// handlers.AllowedHeaders([]string{"Content-taype"}),
	)(h)

}
