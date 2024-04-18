package utils

const (
	// Metrics and HTTP paths
	MetricsPath  = "/metrics"
	DefaultRoute = "/"

	// Service origins and names
	AllowedOrigins     = "*"
	AuthServiceName    = "auth service"
	JournalServiceName = "journal service"
	HelperServiceName  = "helper service"
	RecordServiceName  = "record service"

	// Logging messages
	LogFailedToRegister        = "Failed to register %s handler"
	LogFailedToStartServer     = "Failed to start HTTP server on port %s"
	LogFailedToDialService     = "Failed to dial service at URL %s"
	LogFailedToCloseConnection = "Failed to close connection"
)
