package adventofcode

type BeamPosition struct {
	dir string
	x   int
	y   int
}

func day16Part1(grid []string) int {
	BPs := []BeamPosition{}
	BPs = append(BPs, BeamPosition{"e", 0, 0})
	visited := map[BeamPosition]bool{}

	for len(BPs) > 0 {
		newDir := BPs[0]
		BPs = BPs[1:]
		handleLight(&BPs, newDir.dir, grid, newDir.x, newDir.y, visited)
	}

	extractVisited := map[location]bool{}
	for k := range visited {
		extractVisited[location{k.x, k.y}] = true
	}
	return len(extractVisited)

}

func handleLight(bps *[]BeamPosition, dir string, grid []string, x, y int, visited map[BeamPosition]bool) {
	if x < 0 || x > len(grid)-1 || y < 0 || y > len(grid[0])-1 {
		return
	}

	beamPosition := BeamPosition{dir, x, y}
	if _, isFound := visited[beamPosition]; isFound {
		return
	}
	// if it's not visited, then we should handle different grids

	visited[beamPosition] = true

	/**
	       N
	    W     E
	       S
		**/

	// 1. for '.', same dir with change of position
	if grid[x][y] == '.' {
		switch dir {
		case "w":
			y -= 1
		case "e":
			y += 1
		case "n":
			x -= 1
		case "s":
			x += 1
		}
		*bps = append(*bps, BeamPosition{dir, x, y})
		return
	}

	// 2. for '/'
	if grid[x][y] == '/' {
		newDir := ""
		switch dir {
		case "w":
			x += 1
			newDir = "s"
		case "e":
			x -= 1
			newDir = "n"
		case "n":
			y += 1
			newDir = "e"
		case "s":
			y -= 1
			newDir = "w"
		}
		*bps = append(*bps, BeamPosition{newDir, x, y})
		return
	}

	// 3. for '\'
	if grid[x][y] == '\\' {
		newDir := ""
		switch dir {
		case "w":
			x -= 1
			newDir = "n"
		case "e":
			x += 1
			newDir = "s"
		case "n":
			y -= 1
			newDir = "w"
		case "s":
			y += 1
			newDir = "e"
		}
		*bps = append(*bps, BeamPosition{newDir, x, y})
		return
	}

	// 3.  pointy end of a splitter
	if grid[x][y] == '|' && (dir == "n" || dir == "s") {
		if dir == "n" {
			x -= 1
		}
		if dir == "s" {
			x += 1
		}
		*bps = append(*bps, BeamPosition{dir, x, y})
		return
	}

	if grid[x][y] == '-' && (dir == "w" || dir == "e") {
		if dir == "w" {
			y -= 1
		}
		if dir == "e" {
			y += 1
		}
		*bps = append(*bps, BeamPosition{dir, x, y})
		return
	}

	// 4. flat side of a splitter
	if grid[x][y] == '|' {
		*bps = append(*bps, BeamPosition{"n", x - 1, y})
		*bps = append(*bps, BeamPosition{"s", x + 1, y})
		return
	}
	if grid[x][y] == '-' {
		*bps = append(*bps, BeamPosition{"w", x, y - 1})
		*bps = append(*bps, BeamPosition{"e", x, y + 1})
		return
	}
}
