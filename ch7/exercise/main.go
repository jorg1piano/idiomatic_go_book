package main

import (
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if _, ok := l.Teams[team1]; !ok {
		return
	}

	if _, ok := l.Teams[team2]; !ok {
		return
	}

	if score1 == score2 {
		return
	}

	if score1 > score2 {
		l.Wins[team1]++
	} else {
		l.Wins[team2]++
	}
}

func (l League) Ranking() []string {

	/// Return slice of team names in order of wins
	names := []string{}
	for k := range l.Teams {
		names = append(names, k)
	}

	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})

	return names

}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	result := r.Ranking()

	for _, v := range result {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Norway": {
				Name:    "Norway",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Sweden": {
				Name:    "Sweden",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}

	l.MatchResult("Italy", 50, "Sweden", 70)
	l.MatchResult("India", 85, "Norway", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("Sweden", 100, "Norway", 110)
	l.MatchResult("Italy", 65, "Norway", 70)
	l.MatchResult("Sweden", 95, "India", 80)

	RankPrinter(l, os.Stdout)
}
