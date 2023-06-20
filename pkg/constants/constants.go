package constants

const (
	// name of service
	ServiceName = "StripeScripts"
	// port service wants to listen
	Port = "port"

	// default port service wants to listen when user doesn't adding port in config file
	PortDefault = 9090
	// PathWorkingDirecotry path of the configs
	PathWorkingDirecotry = "."
	// ConfigFileName is the name of configuration file.
	ConfigFileName = "config"

	// ConfigFileType is the type of configuration file.
	ConfigFileType = "yml"

	StripePrefix = "stripe"
	StripeKey    = StripePrefix + ".key"
)
