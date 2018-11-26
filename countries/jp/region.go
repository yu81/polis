package jp

const (
	Hokkaido   = "北海道"
	Tohoku     = "東北"
	Kanto      = "関東"
	Koshinetsu = "甲信越"
	Tokai      = "東海"
	Chubu      = "中部"
	Hokuriku   = "北陸"
	Kinki      = "近畿"
	Chugoku    = "中国"
	Shikoku    = "四国"
	Kyushu     = "九州"
	Okinawa    = "沖縄"
)

func Regions() []string {
	return []string{
		Hokkaido,
		Tohoku,
		Kanto,
		Koshinetsu,
		Tokai,
		Chubu,
		Hokuriku,
		Kinki,
		Chugoku,
		Shikoku,
		Kyushu,
		Okinawa,
	}
}
