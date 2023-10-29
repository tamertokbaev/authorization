package app

import (
	httpDelivery "diploma/authorization/internal/delivery/handlers/http"
)

func Run() {

	healthCheckFn := func() error {
		// Will be updated soon

		return nil
	}

	_ = httpDelivery.NewHandlerDelivery(
		healthCheckFn,
		"pre-order",
	)
}
