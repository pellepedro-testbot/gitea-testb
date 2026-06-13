#!/usr/bin/env bash
set -euo pipefail
U="${GITEA_ADMIN:-gitea-admin}"; P="${GITEA_PASS:-Password123!}"
curl -sf -u "$U:$P" -X POST "http://localhost:3001/api/v1/users/$U/tokens" \
  -H 'Content-Type: application/json' \
  -d "{\"name\":\"testbot-$(date +%s%N)\",\"scopes\":[\"all\"]}" \
  | python3 -c "import json,sys;print(json.load(sys.stdin)['sha1'])"
