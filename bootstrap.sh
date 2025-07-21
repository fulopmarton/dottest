sudo mkdir -p /etc/systemd/resolved.conf.d
echo "[Resolve]
DNS=127.0.0.1
Domains=~test" >/etc/systemd/resolved.conf.d/test-domains.conf

sudo systemctl restart systemd-resolved

# Install local CA
sudo apt install mkcert libnss3-tools -y
mkcert -install
