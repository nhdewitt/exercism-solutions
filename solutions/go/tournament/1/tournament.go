package tournament

import (
    "bufio"
    "fmt"
    "io"
    "sort"
    "strings"
)

type Team struct {
    name	string
    wins	int
    draws	int
    losses	int
}

func (t *Team) Points() int {
    return (3 * t.wins) + t.draws
}

func (t *Team) MatchesPlayed() int {
    return t.wins + t.draws + t.losses
}

func Tally(reader io.Reader, writer io.Writer) error {
	seenTeams := make(map[string]int)
    standings := make([]Team, 0)

    scanner := bufio.NewScanner(reader)
    for scanner.Scan() {
        line := scanner.Text()
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }
        
        parts := strings.Split(line, ";")
        if len(parts) != 3 {
            return fmt.Errorf("invalid result format: %s", parts)
        }

        homeIdx := getOrCreateTeam(parts[0], seenTeams, &standings)
        awayIdx := getOrCreateTeam(parts[1], seenTeams, &standings)

        if err := updateResults(&standings[homeIdx], &standings[awayIdx], parts[2]); err != nil {
            return err
        }
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    sort.Slice(standings, func(i, j int) bool {
        if standings[i].Points() != standings[j].Points() {
            return standings[i].Points() > standings[j].Points()
        }
        return standings[i].name < standings[j].name
    })

    return writeStandings(writer, standings)
}

func updateResults(home, away *Team, result string) error {
    switch result {
    case "win":
        home.wins++
        away.losses++
    case "loss":
        home.losses++
        away.wins++
    case "draw":
        home.draws++
        away.draws++
    default:
        return fmt.Errorf("unknown result: %s", result)
    }
    return nil
}

func getOrCreateTeam(name string, seenTeams map[string]int, standings *[]Team) int {
    if idx, ok := seenTeams[name]; ok {
        return idx
    }
    idx := len(*standings)
    seenTeams[name] = idx
    *standings = append(*standings, Team{name: name})
    return idx
}

func writeStandings(writer io.Writer, standings []Team) error {
    if _, err := fmt.Fprintf(writer, "%-30s | %2s | %2s | %2s | %2s | %2s\n",
              "Team", "MP", "W", "D", "L", "P"); err != nil {
        return err
    }

    for _, team := range standings {
        if _, err := fmt.Fprintf(writer, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
                  team.name, team.MatchesPlayed(), team.wins,
                  team.draws, team.losses, team.Points()); err != nil {
            return err
        }
    }
    return nil
}