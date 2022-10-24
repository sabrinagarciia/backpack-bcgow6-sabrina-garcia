package testing

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func RestarBien(t *testing.T) {
	num1 := 0
	num2 := 3

	errorEsperado := fmt.Sprintf("a no puede ser: %d", num1)

	_, err := Restar(num1, num2)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, errorEsperado)
}