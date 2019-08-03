package helpers

import (
	"math"

	"github.com/zenofbeer/go-zen-velocity/configuration"
)

var config = configuration.GetConfig()

// CalculateCapacityAsPercentage calculates the engineer's capacity and returns that value as a percentage
func CalculateCapacityAsPercentage(currentAvailability int, lastAvailability int) int {
	// (current/last)*100
	retVal := 0
	if lastAvailability != retVal {
		retVal = (currentAvailability / lastAvailability) * 100
	}

	return retVal
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

// CalculatePercentageOfTargetAchieved the value of this sprint/last sprint.
// returns 0 if last sprint is 0.
func CalculatePercentageOfTargetAchieved(
	thisSprintCompleted int, lastSprintCompleted int) float64 {
	retVal := float64(0)
	thisSprint := float64(thisSprintCompleted)
	lastSprint := float64(lastSprintCompleted)
	if retVal < lastSprint {
		retVal = thisSprint / lastSprint
	}
	return retVal
}

// CalculateProductivity calculates the productivity score dividing points
// completed by working days
func CalculateProductivity(thisSprintCompleted int, workingDays int) float64 {
	retVal := float64(0)
	ts := float64(thisSprintCompleted)
	wd := float64(workingDays)
	if retVal < wd {
		retVal = ts / wd
	}
	return retVal
}

// Calculate productivity change by previous from current
func CalculateProductivityChange(
	currentProductivity float64, previousProductivity float64) float64 {
		return currentProductivity = previousProductivity;
}

// RoundToTwoDecimals utility function to round a float64 to two decimal points
func RoundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}
