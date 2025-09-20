package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	a := []string{"Girl", "Boy", "Smart", "Dark", "Silent", "Happy", "Shadow", "Cyber", "Moon", "Storm", "Bliss", "Dream", "Fire", "Wolf", "Rogue", "Ghost", "Magic", "Wild"}
	b := []string{"Rose", "Heart", "Maker", "Hunter", "Rider", "Gamer", "Dragon", "Angel", "Song", "Sky", "Flame", "Star", "Light", "Fox", "Storm", "Wing", "Stone", "Dawn"}
	types := []string{"@gmail.com", "@outlook.com", "@yahoo.com"}
	chars := "abcdefghijklmnopqrstuvwxyz"
	nums := "0123456789"
	special := "!@#$%^&*"

	total := len(a) * len(b) * len(chars) * len(nums) * len(special) * len(types)
	pool := make([]string, 0, total)

	for _, f := range a {
		for _, s := range b {
			for _, c := range chars {
				for _, n := range nums {
					for _, m := range special {
						for _, t := range types {
							email := fmt.Sprintf("%s_%s%c%c%c%s", f, s, c, n, m, t)
							pool = append(pool, email)
						}
					}
				}
			}
		}
	}

	r.Shuffle(len(pool), func(i, j int) { pool[i], pool[j] = pool[j], pool[i] })

	for _, e := range pool {
		fmt.Println("Email:", e)
		time.Sleep(100 * time.Millisecond)
	}
}
