package gastar

import "github.com/ichinaski/gastar/pqueue"

type Position interface{}// Node ID, representing its position within the Graph

type State struct {
    id Position
    parentId Position
    cost int
}

type Graph interface {
    GetChildren(position Position) map[Position]int
    GetHeuristicCost(position, goal Position) int
}

/**
 * A Star implementation
 */
func FindPath(graph Graph, start, goal Position) []Position {
    openList := &pqueue.PQueue{}// Nodes not visited yet

    closedList := make(map[Position]Position)// Visited nodes, and their parents

    state := State { start, nil, 0 }

    openList.Push(state, 0)

    found := false

    for openList.Len() > 0 {
        item, _ := openList.Pop()
        state = item.(State)

        position, cost := state.id, state.cost

        if _, ok := closedList[position]; ok {
            // Position already expanded.
            continue
        }

        // Store this position in the closed list, with a reference to its parent
        closedList[position] = state.parentId

        // TODO: will '==' serve for any state comparison? Probably not
        if position == goal {
            found = true
            break
        }

        // Add the positions that are not in the closed list yet
        children := graph.GetChildren(position)
        for childPosition := range children {
            if _, ok := closedList[childPosition]; !ok {
                childState := State { childPosition, position, cost + children[childPosition] }
                openList.Push(childState, childState.cost + graph.GetHeuristicCost(childPosition, goal))
            }
        }

    }

    path := make([]Position, 0, len(closedList))

    if found {
        position := goal
        for position != nil {
            path = append(path, position)
            position = closedList[position]
        }

        // Reverse the slice
        for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
            path[i], path[j] = path[j], path[i]
        }

    }

    return path
}
