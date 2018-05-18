package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type tally struct {
	teamName string
	mp       int
	win      int
	draw     int
	loss     int
	points   int
}

// Tally 1
func Tally(reader io.Reader, writer io.Writer) error {
	mapTeams := make(map[string]*tally)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || line[0] == '#' {
			continue
		}
		lineSplit := strings.Split(line, ";")
		if len(lineSplit) != 3 {
			return errors.New("Invalid input string")
		}
		if lineSplit[0] == lineSplit[1] {
			return errors.New("Match must have different teams")
		}
		switch lineSplit[2] {
		case "win":
			winMatch(lineSplit[0], lineSplit[1], mapTeams)
		case "draw":
			drawMatch(lineSplit[0], lineSplit[1], mapTeams)
		case "loss":
			winMatch(lineSplit[1], lineSplit[0], mapTeams)
		default:
			return errors.New("Invalid match result")
		}

	}
	createOutput(writer, mapTeams)
	return nil
}

func winMatch(winner, loser string, mapTeams map[string]*tally) {
	if v, ok := mapTeams[winner]; ok {
		v.mp++
		v.win++
		v.points += 3
	} else {
		mapTeams[winner] = &tally{winner, 1, 1, 0, 0, 3}
	}
	if v, ok := mapTeams[loser]; ok {
		v.mp++
		v.loss++
	} else {
		mapTeams[loser] = &tally{loser, 1, 0, 0, 1, 0}
	}
}

func drawMatch(player1, player2 string, mapTeams map[string]*tally) {
	if v, ok := mapTeams[player1]; ok {
		v.mp++
		v.draw++
		v.points++
	} else {
		mapTeams[player1] = &tally{player1, 1, 0, 1, 0, 1}
	}
	if v, ok := mapTeams[player2]; ok {
		v.mp++
		v.draw++
		v.points++
	} else {
		mapTeams[player2] = &tally{player2, 1, 0, 1, 0, 1}
	}
}

func createOutput(writer io.Writer, matches map[string]*tally) {
	var matchesArray []*tally
	for _, v := range matches {
		matchesArray = append(matchesArray, v)
	}
	sort.Slice(matchesArray, func(i, j int) bool {
		if matchesArray[i].points == matchesArray[j].points {
			return matchesArray[i].teamName < matchesArray[j].teamName
		}
		return matchesArray[i].points > matchesArray[j].points
	})
	writer.Write([]byte(fmt.Sprintf("%-31s|%3s |%3s |%3s |%3s |%3s\n",
		"Team", "MP", "W", "D", "L", "P")))
	for _, v := range matchesArray {
		writer.Write([]byte(fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n",
			v.teamName, v.mp, v.win, v.draw, v.loss, v.points)))
	}
}
