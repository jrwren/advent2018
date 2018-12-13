package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestDay7(t *testing.T) {
	log.SetFlags(0)
	f := func(t *testing.T, input string, delay, workers int, expect1 string, expect2 int) {
		one, two, err := day7(input, delay, workers)
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
	s := t.Run("example", func(t *testing.T) { f(t, day7Example, 0, 2, "CABDFE", 15) })
	if s {
		t.Run("real", func(t *testing.T) { f(t, day7Input, 60, 5, "JDEKPFABTUHOQSXVYMLZCNIGRW", 1048) })
	}
}

func day7(input string, delay, workers int) (string, int, error) {
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
		//		if _, ok := metdeps[met]; !ok {
		order += met
		metdeps[met] = struct{}{}
		//		}
		log.Println(order)
	}
	log.Println("--")
	p2 := day7part2(deps, steps, delay, workers)
	return order, p2, nil
}

func day7part2(deps map[string][]string, steps map[string]struct{}, delay, workers int) int {

	pickedup := make(map[string]struct{})
	metdeps := make(map[string]struct{})
	order := ""
	s := 0
	workerJobs := make(map[int]string)
	jobTime := make(map[string]int)
	for len(order) < len(steps) {
		works := ""
		r := nextDeps(deps, steps, metdeps)
		sort.Sort(sort.StringSlice(r))
		log.Println(s, " next:", r, "met:", metdeps, "picked:", pickedup)
		for i := 0; i < workers; i++ {
			log.Println("WTF", workerJobs, jobTime, "next", r)
			job := workerJobs[i]
			if job == "" {
				log.Println("worker", i, "ready for job")
				r := nextDeps(deps, steps, metdeps)
				sort.Sort(sort.StringSlice(r))
				if len(r) == 0 {
					works += ".  "
					continue
				}
				met := r[0]
				// is another worker already working on this?
				found := false
				for i := range r {
					if _, ok := pickedup[r[i]]; ok {
						continue
					}
					found = true
					met = r[i]
					break
				}
				if !found {
					works += ".  "
					continue
				}
				log.Println("worker", i, "picking up job ", met)
				pickedup[met] = struct{}{}
				workerJobs[i] = met
				jobTime[met] = 1
				works += met + "  " + strconv.Itoa(jobTime[met]) + "  "
				continue
			}

			jobTime[job]++

			works += job + " " + strconv.Itoa(jobTime[job]) + "  "

			//log.Println("END", workerJobs, jobTime, "next", r)
		}
		log.Println(s, "  ", works, "   |", order)

		// now that all the work is assigned see if any completed.
		for i := 0; i < workers; i++ {
			job := workerJobs[i]
			if job == "" {
				continue
			}
			if steptime(job)+delay == jobTime[job] {
				// its done!
				order += job
				delete(workerJobs, i)
				delete(jobTime, job)
				metdeps[job] = struct{}{}
				log.Println("adding met dep", job)
				continue
			}
		}

		s++
	}
	return s // doing it wrong shows me 89 ... is the real answer 90? _NO! 208 is too low.
}

func steptime(step string) int {
	n := int(byte(step[0]))
	return n - 64
}

func nextDeps(deps map[string][]string, steps, metdeps map[string]struct{}) (ready []string) {
step:
	for step := range steps {
		//		log.Print("step ", step, " deps ", deps[step])
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
