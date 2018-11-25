package polis

import "github.com/yu81/polis/jp"

type Region string

func Regions(country string) []Region {
	switch country {
	case "jp":
		return stringsToRegions(jp.Regions())
	default:
		return []Region{}
	}
}

func stringsToRegions(sRegions []string) []Region {
	if len(sRegions) == 0 {
		return []Region{}
	}
	result := make([]Region, 0, len(sRegions))
	for _, r := range sRegions {
		result = append(result, Region(r))
	}
	return result
}
