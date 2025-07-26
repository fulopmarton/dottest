echo "[*] Configuring systemd-resolved for local DNS..."

sudo mkdir -p /etc/systemd/resolved.conf.d
echo "[Resolve]
DNS=127.0.0.1
Domains=~test" >/etc/systemd/resolved.conf.d/test-domains.conf

sudo systemctl restart systemd-resolved

echo "[*] Installing mkcert root CA..."
mkcert -install || true

dottest install
