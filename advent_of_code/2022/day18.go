package adventofcode

type Layer struct {
	a int
	b int
	c int
}

func day18Part1(input [][]int) int {
	var (
		xymap = map[Layer]bool{}
		xzmap = map[Layer]bool{}
		yzmap = map[Layer]bool{}
	)
	count := 0
	for _, shape := range input {
		x, y, z := shape[0], shape[1], shape[2]
		// firstly fit x-y
		layer1 := Layer{x, y, z - 1}
		layer2 := Layer{x, y, z}

		for _, layer := range []Layer{layer1, layer2} {
			handleShape(layer, xymap, &count)
		}
		// secondly fit x-z
		layer3 := Layer{x, z, y - 1}
		layer4 := Layer{x, z, y}
		for _, layer := range []Layer{layer3, layer4} {
			handleShape(layer, xzmap, &count)
		}

		// thirdly fit y-z
		layer5 := Layer{y, z, x - 1}
		layer6 := Layer{y, z, x}
		for _, layer := range []Layer{layer5, layer6} {
			handleShape(layer, yzmap, &count)
		}
	}
	return count
}

func handleShape(layer Layer, layerRecord map[Layer]bool, count *int) {
	if _, ok := layerRecord[layer]; !ok {
		*count += 1
		layerRecord[layer] = true
	} else {
		*count -= 1
		delete(layerRecord, layer)
	}
}

func day18Part2(input [][]int) int {
	var (
		xymap = map[Layer]bool{}
		xzmap = map[Layer]bool{}
		yzmap = map[Layer]bool{}
	)

	count := 0
	for _, shape := range input {
		x, y, z := shape[0], shape[1], shape[2]
		// firstly fit x-y
		layer1 := Layer{x, y, z - 1}
		layer2 := Layer{x, y, z}

		for _, layer := range []Layer{layer1, layer2} {
			handleShape(layer, xymap, &count)
		}
		// secondly fit x-z
		layer3 := Layer{x, z, y - 1}
		layer4 := Layer{x, z, y}
		for _, layer := range []Layer{layer3, layer4} {
			handleShape(layer, xzmap, &count)
		}

		// thirdly fit y-z
		layer5 := Layer{y, z, x - 1}
		layer6 := Layer{y, z, x}
		for _, layer := range []Layer{layer5, layer6} {
			handleShape(layer, yzmap, &count)
		}
	}
	return count
}
