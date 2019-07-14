package controllers

// GetWorkstreamName ...
func GetWorkstreamName(ID int) string {
	switch ID {
	case 0:
		return "Air Cancel"
	case 1:
		return "Air Schedule Change"
	case 2:
		return "Shopping"
	}
	return "Workstream Not Found"
}
