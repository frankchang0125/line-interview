package mazeshortestpath

import (
    "math"
)

func goMaze(maze [][]byte, W, H int) int {
    var startX, startY int
    var endX, endY int

    for y := range maze {
        for x := range maze[y] {
            if maze[y][x] == 'S' {
                startX = x
                startY = y
            } else if maze[y][x] == 'G' {
                endX = x
                endY = y 
            }
        }
    }

    return shortestPath(maze, W, H, startX, startY, endX, endY)
}

func shortestPath(maze [][]byte, W, H, startX, startY, targetX, targetY int) int {
    visited := make([][]bool, H)
    for i := 0; i < H; i++ {
        visited[i] = make([]bool, W)
    }

    return dfs(maze, visited, W, H, targetX, targetY, startX, startY, 0)
}

// l: Length of path.
func dfs(maze [][]byte, visited [][]bool, W, H, targetX, targetY, x, y, l int) int {
    // Found target.
    if x == targetX && y == targetY {
        return l
    }

    // Out of bound.
    if x < 0 || x >= W {
        return math.MaxInt32
    }

    // Out of bound.
    if y < 0 || y >= H {
        return math.MaxInt32
    }

    // Cannot go.
    if maze[y][x] == '1' {
        return math.MaxInt32
    }

    // Already visited.
    if visited[y][x] {
        return math.MaxInt32
    }

    // Mark current point as visited.
    visited[y][x] = true

    left := dfs(maze, visited, W, H, targetX, targetY, x-1, y, l+1)  // Go left
    right := dfs(maze, visited, W, H, targetX, targetY, x+1, y, l+1) // Go right
    down := dfs(maze, visited, W, H, targetX, targetY, x, y+1, l+1)  // Go down
    up := dfs(maze, visited, W, H, targetX, targetY, x, y-1, l+1)    // Go up

    // Reset current point to prevent affecting other DFS.
    visited[y][x] = false

    p := min(left, right, down, up)
    return p
}

func min(w, x, y, z int) int {
    return _min(w, _min(x, _min(y, z)))
}

func _min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
