// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package misc

import (
	"net/http"
	"time"

	"gitea.dev/modules/structs"
	"gitea.dev/services/context"
)

var startTime = time.Now()

// Health returns the health status and uptime of the Gitea server
func Health(ctx *context.APIContext) {
	ctx.JSON(http.StatusOK, &structs.HealthStatus{
		Status:        "ok",
		UptimeSeconds: int64(time.Since(startTime).Seconds()),
	})
}
