package robotname

import (
    "fmt"
    "math/rand"
    "sync"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var (
    availableNames	[]string
    usedNames		map[string]struct{}
    mu 				sync.Mutex
)

func init() {
    availableNames = make([]string, 0, 26*26*1000)
    for i := 0; i < 26; i++ {
        for j := 0; j < 26; j++ {
            for k := 0; k < 1000; k++ {
                availableNames = append(availableNames,
                                       fmt.Sprintf("%c%c%03d", charset[i], charset[j], k))
            }
        }
    }

    rand.Shuffle(len(availableNames), func(i, j int) {
        availableNames[i], availableNames[j] = availableNames[j], availableNames[i]
    })

    usedNames = make(map[string]struct{})
}

type Robot struct {
    name string
}

func assignName(r *Robot) error {
    if len(availableNames) == 0 {
        return fmt.Errorf("max robot names reached")
    }

    idx := rand.Intn(len(availableNames))
    r.name = availableNames[idx]
    availableNames[idx] = availableNames[len(availableNames)-1]
    availableNames = availableNames[:len(availableNames)-1]
    usedNames[r.name] = struct{}{}
    return nil
}

func (r *Robot) Name() (string, error) {
    mu.Lock()
    defer mu.Unlock()

    if r.name != "" {
        return r.name, nil
    }

    if err := assignName(r); err != nil {
        return "", err
    }
    return r.name, nil
}

func (r *Robot) Reset() {
	mu.Lock()
    defer mu.Unlock()

    oldName := r.name
    if oldName != "" {
        delete(usedNames, oldName)
        r.name = ""
    }

    assignName(r)

    if oldName != "" {
        availableNames = append(availableNames, oldName)
    }
}
