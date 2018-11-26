package polis

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yu81/polis/countries/jp"
	"math"
	"sort"
	"testing"
)

func TestPrefecturePair_Distance(t *testing.T) {
	cases := allJPPrefPairs()
	sort.Sort(PrefecturePairs(cases))
	const numberOfCombination = 47 * 23
	assert.Len(t, cases, numberOfCombination)

	const epsilon = 0.001
	// 0	大津市	京都市	10.441607 km
	assert.True(t, math.Abs(cases[0].Distance()-10.441607) < epsilon)
	assert.True(t, (cases[0][0].Capital.Name == "大津市" && cases[0][1].Capital.Name == "京都市") || (cases[0][0].Capital.Name == "京都市" && cases[0][1].Capital.Name == "大津市"))

	// 1080	札幌市	那覇市	2245.960949 km
	farestIndex := numberOfCombination - 1
	assert.True(t, math.Abs(cases[farestIndex].Distance()-2245.960949) < epsilon)
	assert.True(t, (cases[farestIndex][0].Capital.Name == "札幌市" && cases[farestIndex][1].Capital.Name == "那覇市") || (cases[farestIndex][0].Capital.Name == "那覇市" && cases[farestIndex][1].Capital.Name == "札幌市"))

	//for i, p := range cases {
	//	fmt.Printf("%d\t%s\t%s\t%f km\n", i, p[0].Capital.Name, p[1].Capital.Name, p.Distance())
	//}
}
func BenchmarkPrefecturesByRegions(b *testing.B) {
	regions := Regions("jp")
	prefs := PrefecturesFromString(jp.Prefectures)
	for _, r := range regions {
		b.Run(string(r), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FilterPrefecturesByRegions(prefs, []Region{r})
			}
		})
	}
}

func allJPPrefPairs() []PrefecturePair {
	result := make(PrefecturePairs, 0, 47*23)
	prefs := PrefecturesFromString(jp.Prefectures)
	for i := 0; i < len(prefs); i++ {
		for j := i + 1; j < len(prefs); j++ {
			result = append(result, PrefecturePair{prefs[i], prefs[j]})
		}
	}
	return result
}

func TestPrefectureCapitalDistance(t *testing.T) {
	for _, v := range allJPPrefPairs() {
		fmt.Printf("Distance between %s and %s%f km\n", v[0].Capital.Name, v[1].Capital.Name, v[0].Capital.Coordinates.Distance(&v[1].Capital.Coordinates))
	}
}

func Benchmark_allPrePairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		allJPPrefPairs()
	}
}
