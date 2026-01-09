package main

import (
	"fmt"
	"sort"
)

type brainrotmeme struct {
	name       string
	views      float64
	trendlevel int
	category   string
}

var memes = []brainrotmeme{
	{"meme1", 120.5, 9, "Surreal"},
	{"meme2", 85.3, 8, "Sigma"},
	{"meme3", 45.2, 6, "Sigma"},
	{"meme4", 210.7, 10, "Economy"},
	{"meme5", 60.8, 7, "Sigma"},
	{"meme6", 30.1, 4, "Chad"},
	{"meme7", 95.4, 8, "Surreal"},
}

func findtoptrending(memes []brainrotmeme, minviews float64) []brainrotmeme {
	var filtered []brainrotmeme
	for _, meme := range memes {
		if meme.views > minviews {
			filtered = append(filtered, meme)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].trendlevel > filtered[j].trendlevel
	})

	return filtered
}

func calculateimpact(memes []brainrotmeme) map[string]float64 {
	impact := make(map[string]float64)

	for _, meme := range memes {
		impact[meme.category] += meme.views
	}

	return impact
}

func filtercondition(memes []brainrotmeme) []string {
	var result []string

	for _, meme := range memes {
		if meme.trendlevel >= 7 || (meme.views > 50 && meme.category == "Sigma") {
			result = append(result, meme.name)
		}
	}

	return result
}

func main() {
	for i, meme := range memes {
		fmt.Printf("%d. %s Просмотры: %.1fM Уровень тренда: %d Категория: %s\n",
			i+1, meme.name, meme.views, meme.trendlevel, meme.category)
	}

	topTrending := findtoptrending(memes, 60.0)
	for i, meme := range topTrending {
		fmt.Printf("%d. %s Trend: %d Views: %.1fM\n",
			i+1, meme.name, meme.trendlevel, meme.views)
	}

	categoryimpact := calculateimpact(memes)
	for category, totalViews := range categoryimpact {
		fmt.Printf("%s: %.1fM просмотров\n", category, totalViews)
	}

	fmt.Println("(trendLevel ≥ 7 или views > 50 и category = 'sigma')")
	filteredMemes := filtercondition(memes)
	for i, name := range filteredMemes {
		fmt.Printf("%d. %s\n", i+1, name)
	}

	fmt.Printf("всего мемов %d\n", len(memes))
	fmt.Printf("попадают в сложное условие %d\n", len(filteredMemes))
	fmt.Printf("категорий %d\n", len(categoryimpact))

	var topcategory string
	var maxviews float64
	for category, views := range categoryimpact {
		if views > maxviews {
			maxviews = views
			topcategory = category
		}
	}
	fmt.Printf("самая популярная категория %s (%.1fм просмотров)\n", topcategory, maxviews)
}
