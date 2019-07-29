package helpers

import (
	"math"

	"github.com/zenofbeer/go-zen-velocity/configuration"
)

var config = configuration.GetConfig()

// CalculateCapacityAsPercentage calculates the engineer's capacity and returns that value as a percentage
func CalculateCapacityAsPercentage(currentAvailability int, lastAvailability int) int {
	// (current/last)*100
	return (currentAvailability / lastAvailability) * 100
}

// CalculateTargetPoints calculates the target points for this sprint
func CalculateTargetPoints(lastCompletedPoints int, currentCapacity int) float64 {
	// l == last sprint completed points
	// c == current capacity
	// vic == velocity increase constant
	// round up ((lsp * c + ((lsp * c) * vic)), 2 decimals) + 1
	// return math.Round(rawNumber*100) / 100
	l := float64(lastCompletedPoints)
	c := float64(currentCapacity)
	i := float64(config.SprintSettings.VelocityIncreaseGoalConstant)

	rawTarget := RoundToTwoDecimals((l*c + ((l * c) * i)))

	return rawTarget + 1
}

// RoundToTwoDecimals utility function to round a float64 to two decimal points
func RoundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}
