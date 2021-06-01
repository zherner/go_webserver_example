package main

import "os"

// the config struct for service
type cfg struct {
	address, port string
}

// config from env
func config() cfg {
	// init cfg
	c := cfg{}

	//c.address, _ = os.LookupEnv("GWE_ADDRESS")
	c.port, _ = os.LookupEnv("GWE_PORT")
	return c
}
