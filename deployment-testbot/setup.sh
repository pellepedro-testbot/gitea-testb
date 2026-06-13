#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"
echo "Building Gitea from source + starting (slow)..."
docker compose -f docker-compose.yml up -d --build
echo "Waiting for health (up to 5 min)..."
for i in $(seq 1 100); do
  if curl -sf -m3 http://localhost:3001/api/healthz | grep -q pass; then break; fi
  sleep 3; [ "$i" = 100 ] && { echo "unhealthy"; docker compose logs --tail 80; exit 1; }
done
# create admin (idempotent)
docker compose exec -T -u git gitea gitea admin user create \
  --username gitea-admin --password 'Password123!' --email admin@gitea.test \
  --admin --must-change-password=false >/dev/null 2>&1 || true
echo "Setup complete"
