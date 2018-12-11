package main

import (
	"fmt"
	"image"
	"log"
	"strings"
	"testing"
)

func TestDay6(t *testing.T) {
	log.SetFlags(0)
	f := func(t *testing.T, input string, p2thres, expect1, expect2 int) {
		one, two, err := day6(input, p2thres)
		if err != nil {
			t.Fatal(err)
		}
		if one != expect1 {
			t.Fatalf("got %d, expected %d", one, expect1)
		}
		if two != expect2 {
			t.Fatalf("got %d, expected %d", two, expect2)
		}
	}
	s := t.Run("example", func(t *testing.T) { f(t, day6Example, 32, 17, 16) })
	if s {
		t.Run("part2", func(t *testing.T) { f(t, day6Input, 10000, 3449, 44868) })
	}
}

func day6(input string, p2thres int) (int, int, error) {
	lines := strings.Split(input, "\n")
	coords := make([]image.Point, len(lines))
	var maxX, maxY int
	id := "A"
	names := make(map[image.Point]string)
	names[tie] = "."
	for i, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d, %d", &x, &y)
		coords[i].X = x
		coords[i].Y = y
		names[coords[i]] = id
		log.Print("found ", coords[i], id, " in '", line, "' with ", x, y)
		id = string(byte(id[0]) + 1)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	for j := 0; j <= maxY; j++ {
		row := ""
	row:
		for i := 0; i <= maxX; i++ {
			for _, p := range coords {
				if p.X == i && p.Y == j {
					row += names[p]
					continue row
				}
			}
			row += "."
		}
		log.Print(row)
	}
	counts := make(map[image.Point]int)
	max := 0
	var foundPoint image.Point
	borderpoints := findBorderPoints(coords, maxX, maxY)
	for j := 0; j <= maxY; j++ {
		row := ""
		for i := 0; i <= maxX; i++ {
			p := image.Point{X: i, Y: j}
			closest := findClosestPoint(p, coords)
			row += names[closest]
			counts[closest]++
			if _, ok := borderpoints[closest]; !ok && counts[closest] > max {
				max = counts[closest]
				foundPoint = closest
			}
		}
		log.Println(row)
	}
	log.Println(foundPoint, names[foundPoint], " is the max area with ", max)

	// part 2
	// What is the size of the region containing all locations which have a
	// total distance to all given coordinates of less than p2thres?
	area := 0
	for j := 0; j <= maxY; j++ {
		row := ""
		for i := 0; i <= maxX; i++ {
			p := image.Point{X: i, Y: j}
			d := distToAllCoords(p, coords)
			if d < p2thres {
				area++
			}
			if n, ok := names[p]; ok {
				row += n
			}
			if d < p2thres {
				row += "#"
				continue
			}
			row += "."
		}
	}
	return max, area, nil
}

func distToAllCoords(p image.Point, pts []image.Point) int {
	d := 0
	for i := range pts {
		d += manhattanDistance(p, pts[i])
	}
	return d
}

var tie = image.Point{X: -1, Y: -1}

func findClosestPoint(p image.Point, pts []image.Point) image.Point {
	minDist := 10000000 // a bignumber.
	closestPoint := tie
	for i := range pts {
		d := manhattanDistance(p, pts[i])
		if d == minDist {
			closestPoint = tie
		}
		if d < minDist {
			minDist = d
			closestPoint = pts[i]
		}
	}
	return closestPoint
}

func findBorderPoints(pts []image.Point, maxX, maxY int) map[image.Point]struct{} {
	bs := make(map[image.Point]struct{})
	for i := 0; i < maxX; i++ {
		p := image.Point{X: i, Y: 0}
		closest := findClosestPoint(p, pts)
		bs[closest] = struct{}{}
		p = image.Point{X: i, Y: maxY}
		closest = findClosestPoint(p, pts)
		bs[closest] = struct{}{}
	}
	for j := 0; j < maxY; j++ {
		p := image.Point{X: 0, Y: j}
		closest := findClosestPoint(p, pts)
		bs[closest] = struct{}{}
		p = image.Point{X: maxX, Y: j}
		closest = findClosestPoint(p, pts)
		bs[closest] = struct{}{}
	}
	return bs
}

func manhattanDistance(p1, p2 image.Point) int {
	xd := abs(p1.X - p2.X)
	yd := abs(p1.Y - p2.Y)
	return xd + yd
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

var day6Example = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

var day6Input = `46, 246
349, 99
245, 65
241, 253
127, 128
295, 69
205, 74
167, 72
103, 186
101, 242
256, 75
122, 359
132, 318
163, 219
87, 309
283, 324
164, 342
255, 174
187, 305
145, 195
69, 266
137, 239
241, 232
97, 319
264, 347
256, 214
217, 47
109, 118
244, 120
132, 310
247, 309
185, 138
215, 323
184, 51
268, 188
54, 226
262, 347
206, 260
213, 175
302, 277
188, 275
352, 143
217, 49
296, 237
349, 339
179, 309
227, 329
226, 346
306, 238
48, 163`
