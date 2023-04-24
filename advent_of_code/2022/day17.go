package adventofcode

import (
	"fmt"
	"strings"
)

var rock1 = []string{"####"}

var rock2 = []string{".#.", "###", ".#."}

var rock3 = []string{"..#", "..#", "###"}

var rock4 = []string{"#", "#", "#", "#"}

var rock5 = []string{"##", "##"}

var rocks = [][]string{rock1, rock2, rock3, rock4, rock5}

type Position struct {
	X int
	Y int
}

type Rock struct {
	Shape       []string
	startLeftX  int // when it's firstly put into game, left-most position
	startRightY int // when it's done,left-most position
}

// Generate all positions that are rendered with '#'
func (r Rock) generateAllValidPositions() []Position {
	res := []Position{}
	for i := 0; i < len(r.Shape); i++ {
		for j := 0; j < len(r.Shape[0]); j++ {
			if r.Shape[i][j] == '#' {
				res = append(res, Position{
					X: r.startLeftX + j,
					Y: r.startRightY + i,
				})
			}
		}
	}

	return res
}

func (r Rock) isThereColisionCheck(visitedPosition map[Position]bool) bool {
	for _, position := range r.generateAllValidPositions() {
		if _, ok := visitedPosition[position]; ok {
			return true
		}
		if position.Y < 0 || position.X < 0 || position.X > 6 {
			return true
		}
	}

	return false
}

func (r Rock) updateVisitedPosition(visitedPosition *map[Position]bool) {
	for _, position := range r.generateAllValidPositions() {
		(*visitedPosition)[position] = true
	}
}

func (r Rock) Print(prevMaxY int, visitedPosition map[Position]bool, rock Rock) {
	res := []string{}

	for i := 0; i < rock.startRightY-len(rock.Shape)+1; i++ {
		var temp strings.Builder
		for j := 0; j < 7; j++ {
			p := Position{j, i}
			if _, ok := visitedPosition[p]; ok {
				temp.WriteString("#")
			} else {
				temp.WriteString(".")
			}
		}
		res = append(res, temp.String())
	}

	fuck := []string{}
	for i := 0; i < len(rock.Shape); i++ {
		var temp strings.Builder
		for j := 0; j < 7; j++ {
			if j >= 2 && j < 2+len(rock.Shape[0]) {
				if rock.Shape[j-2][i] == '#' {
					temp.WriteString("@")
					continue
				}
			}
			temp.WriteString(".")
		}
		fuck = append(fuck, temp.String())
	}

	for i := 0; i < len(fuck); i++ {
		fmt.Println(fuck[i])
	}
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Println(res[i])
	}
	fmt.Println()
}

func day17Part1(actions string) int {
	var (
		visitedPosition = map[Position]bool{}
	)

	prevMaxY := 0
	actionIndex := 0
	// iterate times
	for i := 0; i < 6; i++ {
		fmt.Println(prevMaxY)
		if i == 2 {
			fmt.Println("===")
		}
		rock := Rock{
			Shape:      rocks[i%5],
			startLeftX: 2,
		}
		rock.startRightY = prevMaxY + 3 + len(rock.Shape) - 1

		for {
			xAction := 0
			switch actions[actionIndex%(len(actions))] {
			case '<':
				xAction = -1
			case '>':
				xAction = 1
			}

			rock.startLeftX = rock.startLeftX + xAction
			if rock.isThereColisionCheck(visitedPosition) {
				// seems simpler if we just store a copy of previous positions...
				rock.startLeftX = rock.startLeftX - xAction
			}
			actionIndex += 1

			// it drops down, so y will decrease by 1
			yAction := -1

			rock.startRightY = rock.startRightY + yAction
			if rock.isThereColisionCheck(visitedPosition) {
				// seems simpler if we just store a copy of previous positions...
				rock.startRightY = rock.startRightY - yAction
				rock.updateVisitedPosition(&visitedPosition)
				prevMaxY = rock.startRightY + len(rock.Shape)
				if i == 2 {
					rock.Print(prevMaxY, visitedPosition, rock)
				}
				break
			}
			if i == 2 {
				rock.Print(prevMaxY, visitedPosition, rock)
			}

		}
	}

	return prevMaxY

}
