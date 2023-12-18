package adventofcode

import "testing"

var input13 = []string{
	"##...##....##",
	"..#.##....#.#",
	"..#.##....#.#",
	"##.#.##....##",
	"#####.#..###.",
	"#....#.####.#",
	"..###.##.....",
	"..###.##.....",
	"#....#.####.#",
	"",
	"....##.......",
	"####..#######",
	"###.##.######",
	".#.#..#.#..#.",
	"#..#..#..##..",
	".#.#..#.#..#.",
	"####..#######",
	"..#....#.....",
	"#.#....#.##.#",
	"###.##.######",
	".########..##",
	"####..#######",
	"##.####.####.",
	"",
	"#.#..##....",
	"#.#...#....",
	".#.##....#.",
	".####..##..",
	"#####.....#",
	".#..####.##",
	"########.#.",
	"###..##..##",
	"#.#..####.#",
	"#..#.##.###",
	"..#...#####",
	"#....#...##",
	"#....#...##",
	"..#...#####",
	"#..#.##.###",
	"",
	"#########.##.##",
	"##.....#.####.#",
	"#.#.##...####..",
	"...#####..##...",
	"#.....#...##...",
	".##..#..#.##.#.",
	".##..#...#..#..",
	"####..##.#..#.#",
	"####..##.#..#.#",
	"",
	".#..###..##.#",
	".#..###..##.#",
	"##..#.###.###",
	"#.##...#..##.",
	"..####...#..#",
	"##..#..#.#...",
	"###..#.....##",
	"....##.##..#.",
	"....##.#...#.",
	"",
	"..#....##.#.##.",
	"...###...###..#",
	"#...##..####..#",
	"#...##..####..#",
	"...###...#.#..#",
	"..#....##.#.##.",
	".#.......###..#",
	"##.........#..#",
	"..##.###...#..#",
	".##...#........",
	"####..##.#.....",
	"....#.#.###....",
	"....#.###..####",
	".##....##.##..#",
	"#..##.#..###..#",
	".#...#######..#",
	"#.####.#####..#",
	"",
	"##.##########",
	"..#....##....",
	"..#####..####",
	"###..#....#..",
	"..#..######..",
	"....##.##.##.",
	"...##.####.##",
	"###.#.#..#.#.",
	"..##...###..#",
	"##....####...",
	"..#.##....##.",
	"",
	".##.###...#",
	"#.###....##",
	"....####.##",
	".#.##...#.#",
	".#.##...#.#",
	"....####.##",
	"#.###....##",
	".##.###...#",
	"#..##.##...",
	"#.###.#..#.",
	"#.#.....#..",
	"#...#.##.#.",
	"##...#..##.",
	"#..#.#.....",
	"##.#.#.....",
	"",
	"..##.##.##.....#.",
	"#.#...#.#..##..##",
	"####...#.####..#.",
	"####...#.####..#.",
	"#.#...#.#..##..##",
	"..##.##.##.....#.",
	".##.....#.###..#.",
	"#####.......#...#",
	"...###.#.#.###.##",
	"###..##.#..#..##.",
	"###..##....#..##.",
	"...###.#.#.###.##",
	"#####.......#...#",
	".##.....#.###..#.",
	"..##.##.##.....#.",
	"#.#...#.#..##..##",
	"####...#.####..#.",
	"",
	"#...######.##",
	".##...##.....",
	".#....##.....",
	"#...######.##",
	"####.#.#..#..",
	"..#.##.#..#..",
	".#####..#..##",
	"######..##...",
	"##....#.#####",
	"",
	"##.........",
	"#..##.....#",
	"#.#......##",
	"#.#.#####.#",
	".####...##.",
	"#.###.##.##",
	".##....#.##",
	"####..#####",
	"####..#####",
	".##....#.##",
	"#.###.##.##",
	".####...##.",
	"#.#.#####.#",
	"#.#......##",
	"#.###.....#",
	"##.........",
	"##.........",
	"",
	".##..###.##",
	"....#...#.#",
	"....#...#.#",
	".##..###.##",
	".#.#....###",
	"#.#..##.#..",
	"..#...###.#",
	"#.##..##.#.",
	"..#..#.#...",
	"###.#..##.#",
	".##########",
	".##########",
	"###.#..##.#",
	"..#..#.#.#.",
	"#.##..##.#.",
	"..#...###.#",
	"#.#..##.#..",
	"",
	"###..####.#",
	".#....#...#",
	".#.##.#...#",
	"..#..#..##.",
	"...##...#..",
	".#.##.#...#",
	"#......#...",
	"##....###.#",
	"##.##.###..",
	"..####...##",
	"...##....#.",
	"#......##..",
	".##..##.##.",
	".######.###",
	"###..###.#.",
	"###..###.#.",
	".######.##.",
	"",
	".#.#..#.#.##.",
	".##.##.##....",
	".#.#..#.#....",
	".#.#..#.#.##.",
	"#.##..##.####",
	"..#....#.....",
	".#......#....",
	"###.##.###.##",
	"..#....#..##.",
	"",
	"##.#.####.#.#",
	"...###..###..",
	"##...####...#",
	"####......###",
	"..###.##.###.",
	"..#.#..#.#.#.",
	".....#..#....",
	"##..##..##..#",
	"###.######.##",
	"...#.####.#..",
	"....######...",
	"###.#....#.##",
	"..####..####.",
	"####.#..#.###",
	"##.#.####.#.#",
	"##..#.##.#..#",
	"##..##..##..#",
	"",
	"#..####",
	"#..####",
	"#..#.##",
	".#..##.",
	"#..####",
	"####...",
	"####...",
	"#..#.##",
	".#..##.",
	"",
	".##.#.#.##.",
	"######.#..#",
	".##....##.#",
	".##...##..#",
	"#..#.#.#..#",
	"....##.....",
	"#..#.##....",
	"....#.#.##.",
	"#..####....",
	"",
	".###..##.",
	"##...#.#.",
	"#.#.#...#",
	"#.#.#...#",
	"##...#.#.",
	".###..##.",
	"####..##.",
	"",
	"#....###.##.###..",
	".####.##.##.##.##",
	"#....#........#..",
	"##..###......###.",
	"######.##..##.###",
	".#..#..........#.",
	"......##.##.##...",
	"#.##.#.##...#.#.#",
	"#######......####",
	"#######.####.####",
	"#.##.#........#.#",
	"",
	"#.#....#..#",
	"#.#....#..#",
	"########.##",
	".#.#.###.#.",
	".##.##..##.",
	".#....#..#.",
	".#....#..#.",
	".##.##..##.",
	".#.#.###.#.",
	"########.##",
	"#.#..#.#..#",
	"",
	"###.#..##...#",
	"..#...##..###",
	"...#.....#.#.",
	"#####...##...",
	"#######.....#",
	"..#..##..##.#",
	"..#..##..##.#",
	"#######.....#",
	"#####...##...",
	"...#.#...#.#.",
	"..#...##..###",
	"###.#..##...#",
	"..#.#.###.#..",
	"..###....####",
	"###..########",
	"##....##..#.#",
	"......##.##..",
	"",
	"##..##.",
	"###.##.",
	"######.",
	"##..#..",
	".....##",
	"..##.##",
	"..#####",
	"...##..",
	"..#....",
	"##.###.",
	"#####.#",
	"##..#..",
	"....#..",
	"##.##.#",
	"##.....",
	"....#.#",
	"#####.#",
	"",
	"#...#....#...",
	"#..########..",
	".####.##.####",
	"#####....####",
	"#..#......#..",
	".#####..#####",
	".##.#.##.#.##",
	".##.######.##",
	"....##..##...",
	"#.....##.....",
	".#..#.##.#..#",
	"#.....##.....",
	".....####....",
	".###.#..#.###",
	".##.#....#.##",
	"....######...",
	"###.######.##",
	"",
	"...#....#.....###",
	"####....######...",
	".#.######.#.#.#..",
	"....####.......##",
	"....####.......##",
	".#.######.#.#.#..",
	"###......#####...",
	"...#....#.....###",
	"....#..#.....##..",
	"....#..#.....####",
	"#.###..###.##..#.",
	"",
	"##..###",
	".####..",
	".#..#..",
	".####..",
	".####..",
	"#....##",
	".####..",
	".#..#..",
	"..##.##",
	"",
	".#..###..###..#",
	".###.#....#.#.#",
	"...##..##..##..",
	"#...#.####.#...",
	"##...######...#",
	"#.####.##.####.",
	"###.##....##.##",
	"###.##....##.##",
	"#.####.##.####.",
	"",
	"#........#.####",
	"###.##.###.##..",
	"###.##.######..",
	".#......#..##..",
	"###.##.###...##",
	"#.##..##.#..###",
	"####..#####.###",
	"###.##.##....##",
	"#........##.#..",
	".###..###..####",
	"..######..##.##",
	"###....###.....",
	"#.#.##.#.#.##..",
	"##..##..###.###",
	"#..#..#..##....",
	".##....##.#....",
	"#.#....#.#.####",
	"",
	".########.##.####",
	"..#..#.#..#####..",
	".#.#..#.#........",
	"#..#..#..##...###",
	".#......#.#....##",
	"##.#..#.##.##.#..",
	"....##....#.###..",
	"",
	"###.##.#.",
	"##...#...",
	"...##.#..",
	"##...#...",
	"##...#...",
	"...##.#..",
	"##...#...",
	"###.##.#.",
	"###..####",
	"....#...#",
	"#.#..###.",
	"",
	".######.####.",
	".#.##.#.#.#..",
	".........###.",
	"#.####.##.###",
	"#.#..#.#.#..#",
	"#.#..#.#.....",
	"#......#..##.",
	"..####...####",
	".........#..#",
	"##.##.##.##..",
	"###..##.###.#",
	"########.....",
	"#.#..#.##..#.",
	"#.#..#.##..#.",
	"########.....",
	"",
	".####.#.##.###.",
	"##..##...#####.",
	"##..##...##.###",
	"####.#.#.##..##",
	"##..##.##...##.",
	"#.##.##..###...",
	"#....##.##.###.",
	"######.###.#.##",
	"##..##..#######",
	"##..##..#######",
	"######.###.#.##",
	"#....##.##.###.",
	"#.##.##..###...",
	"",
	".##.####.",
	"#...#####",
	"#.######.",
	"###.#...#",
	"###.#...#",
	"#.#####..",
	"#...#####",
	".##.####.",
	"#.#...#..",
	"...##..##",
	"...##..##",
	"#.#...#..",
	".##.####.",
	"#...#####",
	"#.#####..",
	"",
	"####...####....",
	"####.....#.#...",
	".##...##.#.#.##",
	"####.....##..##",
	"#..##..##..####",
	".##.###..##.###",
	"#..#.#..##...##",
	"#..##...#.#....",
	".##.##.#.###.##",
	"####...##..#.##",
	".....####..###.",
	"",
	"...###..##.",
	".###...#.#.",
	"#..#.######",
	".#..#.#.#..",
	"#...######.",
	".##.####.##",
	".#####.###.",
	".#####.###.",
	".##.####.##",
	"#...#.####.",
	".#..#.#.#..",
	"#..#.######",
	".###...#.#.",
	"...###..##.",
	"#####..#.##",
	"####....##.",
	"####....##.",
	"",
	"####..#..###.####",
	".##..####.####..#",
	".##..###..####..#",
	"####..#..###.####",
	".##.#..#.####....",
	"####.###.#.#.....",
	"###..###..#......",
	"#...#####.#..#..#",
	"...##..###....##.",
	"",
	".#....#..#.",
	"..#..#....#",
	".##..##..##",
	".#.##.#..#.",
	".##..##..##",
	"#..##..##..",
	"#.#..#....#",
	"..####....#",
	"##....####.",
	"",
	"#.####.",
	"#.####.",
	"#......",
	".#.#..#",
	"#####.#",
	"..#...#",
	".####..",
	".####..",
	"..#...#",
	"#####.#",
	".#....#",
	"",
	"##..###...##.",
	"##..#..##.#..",
	"##..#..##.#..",
	"##.####...##.",
	"####...##.##.",
	"...####.#####",
	"..#..####..##",
	"",
	"..#..##.####.##..",
	"####.#.####...##.",
	"##....#.#...#.###",
	"....#.###..###..#",
	"##.#####.##.#..##",
	"..###....####..#.",
	"###.#..####......",
	".#.#..##.#####.##",
	"....#.....###..#.",
	"..#.#.##...##....",
	"####........##...",
	"##..#.....#.#.###",
	"##.##..#...##.###",
	"..#..##.......##.",
	"..#..##.......##.",
	"##.##..#...##.###",
	"##..#.....#.#.###",
	"",
	"..#....#..#..#.",
	".####.###..##..",
	".########.####.",
	".#......#......",
	".#.####.#.####.",
	"..######..#..#.",
	"#...##...######",
	"..######.......",
	"..######..#..#.",
	"..######..####.",
	"##..##..##....#",
	"",
	"#.#........",
	"..###.#..#.",
	"...####..##",
	"...####..##",
	"..###.#..#.",
	"#.#........",
	"..#.##.##.#",
	"....##.##.#",
	".##.##....#",
	".#...##..##",
	".##...#..#.",
	"#..#####.##",
	"#..#..#..#.",
	"#.####....#",
	"#.#.###..##",
	"",
	".##...####.",
	"....#######",
	".##.#...#..",
	".##....####",
	"#....##..#.",
	"#....##..#.",
	".##...#####",
	".##.#...#..",
	"....#######",
	".##...####.",
	"...####.##.",
	"..###..#..#",
	"..###..#..#",
	"",
	"##.###.##....##.#",
	"#..#...#.#..#.#..",
	"......###.##.###.",
	".##.#.##########.",
	"#..#...#.####.#..",
	".##...#.######.#.",
	"######.#......#.#",
	"",
	"#..#.#..##.##.#",
	".##...#####..##",
	".##...#..#....#",
	"#..#######.##.#",
	"#####.###.####.",
	"#..##....#....#",
	".....##.#.#..#.",
	"....##.##......",
	"##.##.#.##....#",
	"",
	"##...##",
	"...##..",
	"###.#..",
	".#..###",
	"###.##.",
	"###...#",
	".....##",
	"...#...",
	"##.##.#",
	"..#.#.#",
	"..#.#.#",
	"##.##.#",
	"...#...",
	".....##",
	"###...#",
	"",
	"###....#..#..#..#",
	".#....#..........",
	".#...######..####",
	"..#.##.##.####.##",
	"..####.#...##...#",
	"#.#...##.######.#",
	"....#####......##",
	"....#####......##",
	"#.#.#.##.######.#",
	"",
	".#..#..#.##",
	"##..#####..",
	"..##......#",
	".#..#..####",
	"..##..###..",
	"..##..####.",
	".#..#..####",
	"",
	".##...###.###..",
	"#..####..#.#.##",
	".#.####.#.###..",
	"..##..##.......",
	".#..##..#..##..",
	".#.#..#.#....##",
	"..#.##.#..#.#..",
	"..#.##.#..#.#..",
	"#.#....#.######",
	"....##....#..##",
	"#........###...",
	"",
	"##..#..##..##..#.",
	"...#.#.#####.##.#",
	"#.#....######..##",
	"..###.#..###.##.#",
	".####..####..##..",
	"#.########...##..",
	"####.#......#..#.",
	"#.#......##......",
	"#.#......##......",
	"####.#......#..#.",
	"#.########...##..",
	"",
	"######.#..#......",
	"######.#..#......",
	"#######......##..",
	"#..##.####...####",
	".#.....#.#.###..#",
	"#.##..#.....#...#",
	"......##.#....#.#",
	"...#.#...#..####.",
	".###..####...#.##",
	".###..####...#.##",
	"...#.#...#..####.",
	"......#..#....#.#",
	"#.##..#.....#...#",
	".#.....#.#.###..#",
	"#..##.####...####",
	"",
	".#.#...##......",
	".###..#.###.###",
	".#..#.#.#....##",
	"............###",
	"#.#.#.######.##",
	"###.#.###.#.###",
	"#.#........##..",
	"..##..#....####",
	".####...##.##..",
	"#.#.....#.#.###",
	"#.#..###.###.##",
	"#.#..###.######",
	"#.#.....#.#.###",
	"",
	"##.#.#.",
	"#.###.#",
	".....#.",
	"#.#.###",
	"##..#.#",
	".#..#.#",
	"#.#.###",
	".....#.",
	"#.###.#",
	"##.#.#.",
	"#.....#",
	"######.",
	"######.",
	"#.....#",
	"##.#.#.",
	"",
	"..#..#.....",
	"........##.",
	"........#..",
	"..#..#.....",
	"...##....#.",
	".##..##..#.",
	"........#..",
	"..#..#...##",
	"#.#..#.##.#",
	"",
	"..#..#####.",
	".###..###.#",
	".###..##.##",
	"..##.######",
	".....####.#",
	".#.##.#.#..",
	"##....#..#.",
	"##....#..#.",
	".#.##...#..",
	".....####.#",
	"..##.######",
	".###..##.##",
	".###..###.#",
	"..#..#####.",
	"..#..#####.",
	"",
	".##.#.##.#.#.#..#",
	".#######..#.###.#",
	"..#..##...#...#.#",
	"..#..##....###.##",
	"..#..##....###.##",
	"..##.##...#...#.#",
	".#######..#.###.#",
	".##.#.##.#.#.#..#",
	"#####..##..##...#",
	".#.#...#####.....",
	"..###.#.#..####..",
	".##.#...##...#..#",
	".##.#...##...#..#",
	"",
	".###.#...#.",
	"#......####",
	"#..##...###",
	".##.##..###",
	"#.#..##.##.",
	"..###..#.##",
	"..#.##.##..",
	"##.##.###..",
	"..#....###.",
	".##....###.",
	"##.##.###..",
	"##.##.###..",
	".##....###.",
	"..#....###.",
	"##.##.###..",
	"",
	"##...##",
	"#####..",
	"..#..#.",
	"...#.##",
	"####..#",
	"..####.",
	"#####..",
	"##...#.",
	"####.#.",
	"###..#.",
	"##...#.",
	"#####..",
	"..####.",
	"####..#",
	"...#.##",
	"",
	".##.####.",
	"#.#.#..#.",
	"..#####.#",
	"....####.",
	".########",
	"#####..##",
	"##.######",
	"##.######",
	".##..##..",
	"###......",
	"###......",
	".##..##..",
	"##.######",
	"##.######",
	"#####..##",
	"",
	"..#...#.....###..",
	"##.#....#...#....",
	"#..##......#..#..",
	"###..#....#.##.##",
	"#.#..#....#.##.##",
	"#..##......#..#..",
	"##.#....#...#....",
	"..#...#.....###..",
	"#..###..#.#......",
	".#...###.####.#..",
	".#........###....",
	"###..#..##.....##",
	"...#.#..##..#....",
	"",
	"#.#..###.",
	".#.#####.",
	"#......#.",
	"...#.####",
	".....####",
	"#......#.",
	".#.#####.",
	"#.#..###.",
	"#.##...#.",
	"..#.#....",
	"..#.#....",
	"",
	".....#..#",
	"...#..###",
	"#####..##",
	"##..###..",
	"########.",
	"###...#..",
	"##.##.#.#",
	"..##.##.#",
	"##.....#.",
	"...#.####",
	"...#.####",
	"##.#...#.",
	"..##.##.#",
	"",
	"..#.#..#.#...",
	"##.#.##.#.###",
	".#..####..#..",
	"..#..##..#...",
	"#.##.##.##.##",
	".....##......",
	"#...#..#...##",
	"##.######.###",
	"...##..##....",
	"#..##..##..##",
	"...#....#....",
	"#.#.#..#.#.##",
	"###..##..#.##",
	"#.##.##.##.##",
	"####....#####",
	"",
	"##.###.##.....#.#",
	"##..##.####.####.",
	"..###.#.###.##.##",
	"........#...##...",
	".##....##.#..####",
	"##.#.###...###.##",
	"#####.#.##....##.",
	"###..##.#..#...##",
	"...##.#.#..#####.",
	"....#####....##..",
	"...##..#.#.#..##.",
	"..#.###...#..##.#",
	"..#.###...#..##.#",
	"...##..#.#.#..##.",
	"....#####....##..",
	"",
	"...##..##.##.",
	"...##..##.##.",
	"###.##.#.#...",
	"...###.##...#",
	"####.##.###.#",
	"##.#.#..##..#",
	"..#.#####...#",
	"######....#.#",
	".....###.#.#.",
	"#..#.##..##.#",
	".....#...#.#.",
	"",
	".#........#.#",
	"#..#.##.#..##",
	"##.##..##.###",
	"..###..###...",
	"..###..###...",
	"##.##..##.###",
	"#..#.##.#..##",
	".#........#.#",
	"#####...#####",
	".###.##.###..",
	"####....####.",
	"",
	"#...#.#..###...",
	"##..#.#..###...",
	"#.#..#.#####.#.",
	"#.#..#.#####.#.",
	"##..#.#..###...",
	"#...#.#..###...",
	"#..#.....#..###",
	"##....##..##.##",
	"..##.##....##.#",
	"",
	"######..#",
	"...#.#..#",
	".#...####",
	"####.#..#",
	"#.###.##.",
	"..#...##.",
	".....#..#",
	"...#.....",
	"###.#.##.",
	".##..#..#",
	".#.##....",
	".#.##....",
	".##.##..#",
	"###.#.##.",
	"...#.....",
	"",
	"..##.##.#....",
	".##########..",
	".##########..",
	"..##.##.#.#..",
	".######..##..",
	"###..#.####..",
	"#...#.#..##..",
	".#.#...#..###",
	"#.##.#.####..",
	"..#####...###",
	"..#..##....##",
	".#.####.##...",
	"#.########.##",
	"",
	"...#..#.#..#.###.",
	"...#..#.#..#..##.",
	"..#....#..#.##..#",
	".#.##.#..##....#.",
	"###....#..#.#..##",
	"##.#.##.#...#...#",
	"##......##...#.##",
	"#.....#..#.....##",
	"..##..#####..###.",
	"##.##.##.#....##.",
	"##.##.##.#....##.",
	"..##..#####..###.",
	"#.....#..#.....##",
	"",
	"##.####",
	"...##.#",
	"#.#.##.",
	"#.#.##.",
	"...##.#",
	"##.####",
	"##.###.",
	"###..##",
	".##..##",
	"",
	"#..###..##.#.#.",
	"##.##.#####.#..",
	"##.#.#.#....###",
	"##.#.#.#....###",
	"##.##.#####.#..",
	"#..###..##.#.#.",
	"...##.#..#..#..",
	"#..###.#####..#",
	".##..#..#.##...",
	".##..#..#.##...",
	"#..###.#####..#",
	"...##.#..#..#..",
	"#..###..#..#.#.",
	"##.##.#####.#..",
	"##.#.#.#....###",
	"",
	"##.##....",
	"..##.##.#",
	"..#.#..##",
	"..#.##...",
	"...##.###",
	"..##...#.",
	"#####.##.",
	"###..###.",
	"###..###.",
	"#####.##.",
	"..##...#.",
	"...##.###",
	"..#.##.#.",
	"..#.#..##",
	"..##.##.#",
	"##.##....",
	"###.#####",
	"",
	"###.##.",
	"###.##.",
	"#####.#",
	"..#.##.",
	"#..#..#",
	".##....",
	"....##.",
	"",
	"###.#.#.##.##..#.",
	"..#...#..#..##...",
	"..##.#...####...#",
	"..#..####.#..##..",
	"..###.#....##....",
	"#####..#..#..###.",
	"##.#.#....######.",
	"##......#..#..#..",
	"......#.##....#..",
	"###...##...##.#.#",
	"####.#.#.#.##.##.",
	"....#.####....#..",
	"....#.#####...#..",
	"####.#.#.#.##.##.",
	"###...##...##.#.#",
	"",
	".######.........#",
	"#.#.#..#......#..",
	"#######..####..##",
	"#...#.#........#.",
	"#...##.#.#..#.#.#",
	"...#.##.######.##",
	"...#.##.######.##",
	"#...##.#.#..#.#.#",
	"#...#.#........#.",
	"",
	"...#.##.##...##..",
	"......###..#..###",
	".##...##..###....",
	".####..###.#.##..",
	"#..#..##.#.##.#..",
	"#.######.###.##..",
	"##.###...#.###.##",
	"##.###...#.###.##",
	"#.#####..###.##..",
	"",
	"#......#....###.#",
	"#......#....###.#",
	"#.#..#.#..####.##",
	"........#.##...#.",
	"#..##..#.#...#...",
	"###.####......##.",
	".######...###.#..",
	"",
	".#.####.#.##.##",
	"#........#.#...",
	"..##..##..##...",
	"..##..##..##...",
	"#........#.#...",
	"##.####.####.##",
	".#..##..#...#.#",
	".#.####.#.#....",
	"###.##.####.##.",
	"###....###.#..#",
	"###....###.#.##",
	"",
	"##...####..",
	"#....#..#..",
	"###.##..##.",
	"#.##.#..#.#",
	"....######.",
	"....######.",
	"#.##.#..#.#",
	"###.##..##.",
	"#....#..#..",
	"##...####..",
	"#.##.#..#.#",
	"##.##....#.",
	".##.#.##.#.",
	"",
	"##.#..#.#",
	"##.#..#.#",
	".#####.#.",
	"##.#...##",
	".##...#..",
	"#..######",
	".##.#..#.",
	"..##.##..",
	"..##.##..",
	".##.#..#.",
	"#..######",
	".##...#..",
	"##.#...##",
	"..####.#.",
	"##.#..#.#",
	"",
	"#####.#...###..",
	".##.......##..#",
	"########.#.##..",
	"....###.#..##..",
	"....#...##.##..",
	".........##...#",
	"#..##.####.###.",
	"#..##.####.###.",
	".........##...#",
	"....#...##.##..",
	"....###.#..#...",
	"########.#.##..",
	".##.......##..#",
	"",
	"##..##..#",
	"##..###.#",
	"...##.#..",
	"#.#.##.##",
	"#.#.##.##",
	"...##.#..",
	"##..###.#",
	"##..##..#",
	"#.##.####",
	"",
	"#..####.##.#.",
	"#..####.##.#.",
	"#..##........",
	".##...##.#.##",
	"#######..####",
	".##.#.###.#.#",
	".##..####...#",
	"###.#......##",
	"....#...#.##.",
	"#..#######.##",
	".##.##...##.#",
	".##.##.#...##",
	"#..###...####",
	".##.##...##..",
	".##.##.##.#..",
	"....#.##.####",
	".##.#....##.#",
	"",
	"#####.#.##.",
	"###.#.#....",
	"##.##...##.",
	"##...#..##.",
	"...#..#.##.",
	"#####.#####",
	"...#..##..#",
	"..####.#..#",
	"....#...###",
	"...#..##..#",
	"##.#.#.####",
	"........##.",
	"##.####.##.",
	"",
	"..#.###.#.#....#.",
	"###..##.#...##...",
	".##..##.#...##...",
	"..#.###.#.#....#.",
	".##..#.###.#..#.#",
	"###.#.##.#......#",
	"##.#..###.#....#.",
	"",
	"#..##..##....##.#",
	"###..####...###..",
	"#..##..##.###.#.#",
	"#..##..##.###.#.#",
	"###..####...###..",
	"#..##..##....##.#",
	".#........#.#...#",
	"..#..#.......###.",
	"##########.###..#",
	".#....#..##.###..",
	"##########..##.#.",
	"",
	".#..#....#.",
	".####.##.##",
	"#....#..#..",
	"..##......#",
	"##..######.",
	".####....##",
	"######..###",
	"####.#..#.#",
	".#..#....#.",
	"",
	".#...##.#...#",
	".....##.#...#",
	"..###.#..####",
	"####.#..####.",
	"....##.#..#..",
	"#.###..#.....",
	"##..##.###.##",
	".##..#.#.#...",
	"..#...#..#.#.",
	".......####..",
	"#######...#..",
	".##.#.####.##",
	".##.#.####.##",
	"",
	"##.##..##.###",
	"..#......#...",
	".##..##..##..",
	"..##....##...",
	"...#....#....",
	".##.#..#.##..",
	"#..........##",
	"###.####.####",
	"...##..##....",
	"...#.##.#....",
	".###.##.###..",
	"#..........##",
	".....##....##",
	"",
	"##.#.###..#.#.#",
	"##.###.#...####",
	"####..##..#..##",
	"####.###..#..##",
	"##.###.#...####",
	"##.#.###..#.#.#",
	"....#..#...#.#.",
	"##.#..#.##.#.#.",
	"####..#..###..#",
	"..###.##....##.",
	"###...##.....#.",
	"",
	"#.#.#.##..#",
	"#.#.#.##..#",
	".#..#..#.##",
	"#..#.#.....",
	"...#..####.",
	"########.##",
	"########.##",
	"...#..####.",
	"#..#.#.....",
	".#..#..#.##",
	"#.#.####..#",
	"",
	".#..#.....##.....",
	"#..#...##....##..",
	"..#.##.#.####.#.#",
	".#.#.#..##..##..#",
	"...#..##..##..##.",
	"....#.###....##..",
	"###..............",
	"###...###.##.###.",
	".....#....##....#",
	"...#.##.##..##.##",
	"...#.##.##..##.##",
	"",
	".#....###.#",
	"#.########.",
	"#.########.",
	".#....###.#",
	".####.....#",
	"#.#...##.##",
	"##.#.......",
	"##.#.#.#.#.",
	"##.#.#.#.#.",
	"##.#.#.....",
	"#.#...##.##",
	"",
	"...##.....#",
	"..##..#####",
	"##..#..####",
	"##..#..####",
	"..##...####",
	"...##.....#",
	"....###....",
	".#.....#.#.",
	"##...####.#",
	"##...####.#",
	".#.....#.#.",
	"....###....",
	"...##.....#",
	"",
	"#.##..#..",
	"......#..",
	"...##..#.",
	".##.#..#.",
	"..#......",
	"..#......",
	".##.#..#.",
	"...##..#.",
	"......#..",
	"#.##..#..",
	"#####.#..",
	"..##.##..",
	"....###..",
	"....###..",
	"..##.##..",
	"#####.#..",
	"#.##..#.#",
	"",
	"##.##..##",
	"...######",
	"..#......",
	"###.#..#.",
	".....##..",
	".###.##.#",
	"..#..##..",
	"",
	"..##...##...#",
	"##..#..##..#.",
	"#..#.#.##.#.#",
	"#..#.#.##.#.#",
	"##..#..##..#.",
	"..##...##...#",
	"#.###########",
	"#...#......#.",
	"##.#..####..#",
	"##..###...##.",
	".##..##..##..",
	"",
	".###..#.#...###",
	"...##..##.#..#.",
	"...#.##.#.##.#.",
	"##.#..#..#..#.#",
	"#..###..##..#.#",
	"#.#.#......##..",
	"....#..#..#.#.#",
	"..#.##.#.####.#",
	"..#.##.#.####.#",
	"....#..#..#.#..",
	"#.#.#......##..",
	"#..###..##..#.#",
	"##.#..#..#..#.#",
	"...#.##.#.##.#.",
	"...##..##.#..#.",
	".###..#.#...###",
	".###..#.#...###",
	"",
	"##.#.#.....####..",
	"##.#.#.....####..",
	".#..###.#..#..#..",
	"#.....##.##.##.##",
	"##..#..#.#......#",
	"#...#.##.###..###",
	"######...#.####..",
	"#.....#..#......#",
	"#.##.#..#.##..##.",
	"#.#..##....#..#..",
	"##.#.#.###.####.#",
	"#.###............",
	".##.####.###..###",
	".##.##....######.",
	"##...#####.#..#.#",
	"",
	".##....#..####..#",
	".##..###..####..#",
	".##..#...#....#..",
	".##..#.##..##..##",
	".##...#...#..#...",
	"#..##.#..#.##.#..",
	"#..##...##.##.##.",
	"########.##..##.#",
	".##.#.#.###..###.",
	"#..#.##..##..##..",
	"#..#.##.#..##..#.",
	"#..#.####.#....##",
	"#..#..#.#.#..#.#.",
	".......##......##",
	".##.##.#.#.##.#.#",
}

func TestDay13Part1(t *testing.T) {
	expectedRes := 35521
	if res := day13Part1(input13); res != expectedRes {
		t.Errorf("expected res is %v,what we get is %v", expectedRes, res)
	}
}
