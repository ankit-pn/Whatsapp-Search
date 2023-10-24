package main

import ()

func checkGroupedSection(section string) string {
	sectionGroupMap := map[string]string{
		"wsurveyor1": "up_data_collection",
		"wsurveyor2": "up_data_collection",
		"wsurveyor3": "up_data_collection",
		"wsurveyor4": "up_data_collection",
		"wsurveyor5": "up_data_collection",
		"wsurveyor6": "up_data_collection",
		"ved":        "up_data_collection",
	}

	if group, exists := sectionGroupMap[section]; exists {
		return group
	}
	return section
}
