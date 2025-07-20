
# DotTest
dottest is a local DNS server and reverse proxy that lets you map multiple custom domains to different ports on your local machine. This is especially useful for development environments where you want to access various applications running on different ports using easy-to-remember domain names.

## Setup

1. Run the bootstrap script:
   ```sh
   ./bootstrap.sh
   ```

2. Update the `config.yaml` file with your domain-to-port mappings. Example:
   ```yaml
   mappings:
     "sitea.test": 5173
     "siteb.test": 8080
   ```

3. Start the server (requires sudo for privileged ports):
   ```sh
   sudo go run .
   ```

Now you can access your local apps at `http://sitea.test` and `http://siteb.test`.

## Roadmap
- [x] Basic domain-to-port mapping
- [ ] HTTPS support for local domains
- [ ] Web UI for managing domain mappings
