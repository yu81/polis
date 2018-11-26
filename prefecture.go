package polis

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Prefecture struct {
	ID      int      `json:"id"`
	Region  []Region `json:"region"`
	Name    string   `json:"name"`
	Capital City     `json:"capital"`
	Areas   []Area   `json:"areas"`
}

type Prefectures []Prefecture

func (p Prefecture) String() string {
	builder := new(strings.Builder)
	err := json.NewEncoder(builder).Encode(&p)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return builder.String()
}

func PrefecturesFromString(jsonPrefectures string) Prefectures {
	var result []Prefecture
	err := json.Unmarshal([]byte(jsonPrefectures), &result)
	if err != nil {
		var r *Prefecture
		err := json.Unmarshal([]byte(jsonPrefectures), r)
		if err == nil {
			return []Prefecture{*r}
		}
		return []Prefecture{}
	}
	return result
}

func (p Prefectures) String() string {
	builder := new(strings.Builder)
	err := json.NewEncoder(builder).Encode(&p)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return builder.String()
}

type PrefecturePair [2]Prefecture

func (p *PrefecturePair) Distance() float64 {
	c := p[1].Capital.Coordinates()
	return p[0].Capital.Distance(&c)
}

type PrefecturePairs []PrefecturePair

func (p PrefecturePairs) Len() int {
	return len([]PrefecturePair(p))
}

func (p PrefecturePairs) Less(i, j int) bool {
	ic, jc := []PrefecturePair(p)[i][1].Capital.Coordinates(), []PrefecturePair(p)[j][1].Capital.Coordinates()
	return []PrefecturePair(p)[i][0].Capital.Distance(&ic) < []PrefecturePair(p)[j][0].Capital.Distance(&jc)
}

func (p PrefecturePairs) Swap(i, j int) {
	[]PrefecturePair(p)[i], []PrefecturePair(p)[j] = []PrefecturePair(p)[j], []PrefecturePair(p)[i]
}

func FilterPrefecturesByRegions(prefectures []Prefecture, regions []Region) []Prefecture {
	if len(regions) == 0 {
		return []Prefecture{}
	}
	result := make([]Prefecture, 0, 7*len(regions))
	for _, pref := range prefectures {
	R:
		for _, region := range regions {
			for _, r := range pref.Region {
				if r == region {
					result = append(result, pref)
					break R
				}
			}
		}
	}
	return result
}
