package dividir

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func DividirMal(t *testing.T) {
	num1 := 4
	num2 := 0

	errorEsperado := fmt.Sprintf("denominador no puede ser 0")

	_, err := Dividir(num1, num2)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, errorEsperado)
}