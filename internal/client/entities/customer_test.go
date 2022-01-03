package entities

import (
	"testing"
)

func TestCustomerFullName(t *testing.T) {
	cust := Customer{ID: 1, FirstName: "Grace", LastName: "Hopper"}

	if cust.FullName() != "Grace Hopper" {
		t.Errorf("expected: Grace Hopper, got: %s", cust.FullName())
	}
}
