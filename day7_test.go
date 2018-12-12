package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"testing"
)

func TestDay7(t *testing.T) {
	log.SetFlags(0)
	f := func(t *testing.T, input string, expect1 string, expect2 int) {
		one, two, err := day7(input)
		if err != nil {
			t.Fatal(err)
		}
		if one != expect1 {
			t.Fatalf("got %s, expected %s", one, expect1)
		}
		if two != expect2 {
			t.Fatalf("got %d, expected %d", two, expect2)
		}
	}
	s := t.Run("example", func(t *testing.T) { f(t, day7Example, "CABDFE", 0) })
	if s {
		t.Run("part2", func(t *testing.T) { f(t, day7Input, "JDEKPFABTUHOQSXVYMLZCNIGRW", 0) })
	}
}

func day7(input string) (string, int, error) {
	lines := strings.Split(input, "\n")
	deps := make(map[string][]string)
	steps := make(map[string]struct{})
	for _, line := range lines {
		var s, d string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &d, &s)
		depfors, ok := deps[s]
		if !ok {
			deps[s] = []string{}
			depfors = deps[s]
		}
		depfors = append(depfors, d)
		deps[s] = depfors
		steps[s] = struct{}{}
		steps[d] = struct{}{}
	}
	metdeps := make(map[string]struct{})
	order := ""
	for len(order) < len(steps) {
		r := nextDeps(deps, steps, metdeps)
		sort.Sort(sort.StringSlice(r))
		log.Println("next:", r, "met:", metdeps)
		met := r[0]
		if _, ok := metdeps[met]; !ok {
			order += met
			metdeps[met] = struct{}{}
		}
		log.Println(order)
	}
	return order, 0, nil
}

func nextDeps(deps map[string][]string, steps, metdeps map[string]struct{}) (ready []string) {
step:
	for step := range steps {
		log.Print("step ", step, " deps ", deps[step])
		for _, need := range deps[step] {
			if _, ok := metdeps[need]; !ok {
				continue step
			}
		}
		if _, ok := metdeps[step]; !ok {
			ready = append(ready, step)
		}
	}
	return
}

var day7Example = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`
var day7Input = `Step J must be finished before step E can begin.
Step X must be finished before step G can begin.
Step D must be finished before step A can begin.
Step K must be finished before step M can begin.
Step P must be finished before step Z can begin.
Step F must be finished before step O can begin.
Step B must be finished before step I can begin.
Step U must be finished before step W can begin.
Step A must be finished before step R can begin.
Step E must be finished before step R can begin.
Step H must be finished before step C can begin.
Step O must be finished before step S can begin.
Step Q must be finished before step Y can begin.
Step V must be finished before step W can begin.
Step T must be finished before step N can begin.
Step S must be finished before step I can begin.
Step Y must be finished before step W can begin.
Step Z must be finished before step C can begin.
Step M must be finished before step L can begin.
Step L must be finished before step W can begin.
Step N must be finished before step I can begin.
Step I must be finished before step G can begin.
Step C must be finished before step G can begin.
Step G must be finished before step R can begin.
Step R must be finished before step W can begin.
Step Z must be finished before step R can begin.
Step Z must be finished before step N can begin.
Step G must be finished before step W can begin.
Step L must be finished before step G can begin.
Step Y must be finished before step R can begin.
Step P must be finished before step I can begin.
Step C must be finished before step W can begin.
Step T must be finished before step G can begin.
Step T must be finished before step R can begin.
Step V must be finished before step Z can begin.
Step L must be finished before step C can begin.
Step K must be finished before step I can begin.
Step J must be finished before step I can begin.
Step Q must be finished before step C can begin.
Step F must be finished before step A can begin.
Step H must be finished before step Y can begin.
Step M must be finished before step N can begin.
Step P must be finished before step H can begin.
Step M must be finished before step C can begin.
Step V must be finished before step Y can begin.
Step O must be finished before step V can begin.
Step O must be finished before step Q can begin.
Step A must be finished before step G can begin.
Step T must be finished before step Z can begin.
Step K must be finished before step R can begin.
Step H must be finished before step O can begin.
Step O must be finished before step Y can begin.
Step O must be finished before step C can begin.
Step K must be finished before step P can begin.
Step P must be finished before step F can begin.
Step E must be finished before step M can begin.
Step M must be finished before step I can begin.
Step T must be finished before step W can begin.
Step P must be finished before step L can begin.
Step A must be finished before step O can begin.
Step X must be finished before step V can begin.
Step S must be finished before step G can begin.
Step A must be finished before step Y can begin.
Step J must be finished before step R can begin.
Step K must be finished before step F can begin.
Step J must be finished before step A can begin.
Step P must be finished before step C can begin.
Step E must be finished before step N can begin.
Step F must be finished before step Y can begin.
Step J must be finished before step D can begin.
Step H must be finished before step Z can begin.
Step U must be finished before step H can begin.
Step J must be finished before step T can begin.
Step V must be finished before step G can begin.
Step Z must be finished before step I can begin.
Step H must be finished before step W can begin.
Step B must be finished before step R can begin.
Step F must be finished before step B can begin.
Step X must be finished before step C can begin.
Step L must be finished before step R can begin.
Step F must be finished before step U can begin.
Step D must be finished before step N can begin.
Step P must be finished before step O can begin.
Step B must be finished before step O can begin.
Step F must be finished before step C can begin.
Step H must be finished before step L can begin.
Step O must be finished before step N can begin.
Step J must be finished before step Y can begin.
Step H must be finished before step N can begin.
Step O must be finished before step L can begin.
Step I must be finished before step W can begin.
Step J must be finished before step H can begin.
Step D must be finished before step Z can begin.
Step F must be finished before step W can begin.
Step X must be finished before step W can begin.
Step Y must be finished before step M can begin.
Step T must be finished before step M can begin.
Step U must be finished before step G can begin.
Step L must be finished before step I can begin.
Step N must be finished before step W can begin.
Step E must be finished before step C can begin.`
