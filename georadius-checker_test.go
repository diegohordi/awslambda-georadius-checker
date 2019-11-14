/**
 * Geo Radius Checker test.
 */
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// From Marquês de Pombal square to Centro Cultural de Cascais (~ 23.71km)
func Test_IsInsideRadius(t *testing.T) {
	t.Run("IS_INSIDE_RADIUS", func(t *testing.T) {
		aLatitude := 38.6938
		aLongitude := -9.4204
		bLatitude := 38.7254
		bLongitude := -9.1502
		radius := 23710
		result := IsInsideRadius(aLatitude, aLongitude, bLatitude, bLongitude, radius)
		assert.Equal(t, true, result)
	})
	t.Run("ISNT_INSIDE_RADIUS", func(t *testing.T) {
		aLatitude := 38.6938
		aLongitude := -9.4204
		bLatitude := 38.7254
		bLongitude := -9.1502
		radius := 23700
		result := IsInsideRadius(aLatitude, aLongitude, bLatitude, bLongitude, radius)
		assert.Equal(t, false, result)
	})
}
