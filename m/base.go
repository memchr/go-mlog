package m

// Read a float64 value from memory at specified position
func Read[A integer](memory string, position A) int {
	return 0
}

// Write a value to memory at specified position
func Write[A float, B integer](value A, memory string, position B) {
}

// Flush all printed statements to the provided message block
func PrintFlush(targetMessage string) {
}

// Get the linked tile at the specified address
func GetLink[A integer](address A) Link {
	return nil
}

// Retrieve a list of units that match specified conditions
//
// Conditions are combined using an `and` operation
func Radar(from Ranged, target1 RadarTarget, target2 RadarTarget, target3 RadarTarget, sortOrder bool, sort RadarSort) Unit {
	return nil
}

// Extract information indicated by sense from the provided block
func Sensor(block HealthC, sense string) float64 {
	return 0
}

func Wait[A float](time A) {
}

// String equivalent of Sensor
func SensorStr(block HealthC, sense string) string {
	return ""
}
