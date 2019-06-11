package types

import (
	"context"
	"time"

	corev1 "k8s.io/api/core/v1"
)

var (
	// MaxPendingPeriod is the max waiting time for starting a Pod.
	MaxPendingPeriod float64         = 5 * 60 // 5 minutes
	ContextRoot      context.Context = context.Background()
	POStop           chan bool
	// PendingWatchInterval is the interval to watch Pending Pods.
	PendingWatchInterval                 = 5 * time.Minute
	statusPhase          corev1.PodPhase = "status.phase="
	// StatusPhasePending means Pod has been accepted by the system, but one or more of the containers
	// has not been started. This includes time before being bound to a node, as well as time spent
	// pulling images onto the host.
	StatusPhasePending = statusPhase + corev1.PodPending
	MainStop           chan bool
)
