import "strings"

/**
There is a robot starting at position (0, 0), the origin, on a 2D plane. Given
a sequence of its moves, judge if this robot ends up at (0, 0) after it completes
its moves.

The move sequence is represented by a string, and the character moves[i] represents
its ith move. Valid moves are R (right), L (left), U (up), and D (down). If the robot
returns to the origin after it finishes all of its moves, return true. Otherwise,
return false.

Example 1:

Input: "UD"
Output: true
Explanation: The robot moves up once, and then down once. All moves have the
same magnitude, so it ended up at the origin where it started. Therefore, we return true.

**/
func judgeCircle(moves string) bool {
	if moves == "" {
		return true
	}
	left_count, right_count := strings.Count(moves, "L"), strings.Count(moves, "R")
	if left_count != right_count {
		return false
	}
	up_count, down_count := strings.Count(moves, "U"), strings.Count(moves, "D")
	if up_count != down_count {
		return false
	}
	return true

}