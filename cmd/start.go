err := godotenv.Load()
log.Printf("Default config: %v", config.DefaultConfig)
if err != nil {
	log.Fatal("Error loading .env file")
}
godotenv.Load(".env.local")
go dns.StartDNSServer()
proxy.StartReverseProxy()

