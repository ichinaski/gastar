package pqueue

type entry struct {
    item interface{}
    priority int
}

type PQueue []entry

func (h *PQueue) Len() int { return len(*h) }

func (h PQueue) priority(index int) int { return h[index].priority }

func (h *PQueue) swap (i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *PQueue) Push(item interface{}, priority int) {
    entry := entry{
        item, priority,
    }

    // Stick the element as the end of the last level
    *h = append(*h, entry)

    // Bubble up to restore 'heap' property
    index := h.Len() - 1
    parent := int((index - 1) / 2)

    for parent >= 0 && h.priority(index) < h.priority(parent) {
        h.swap(index, parent)

        index = parent
        parent = int((index - 1) / 2)
    }
}

func (h *PQueue) Pop() (interface{}, int) {
    size := h.Len()

    // Move last leaf to root
    h.swap(size - 1, 0)

    entry := (*h)[size - 1]// Item to return

    old := *h
    *h = old[0 : size-1]// Resize the slice

    // Bubble down to restore the heap property
    index := 0
    childL, childR := 2 * index + 1, 2 * index + 2

    for h.Len() > childL {
        child := childL
        if h.Len() > childR && h.priority(childR) < h.priority(childL) {
            child = childR
        }

        if h.priority(index) > h.priority(child) {
            h.swap(index, child)

            index = child
            childL, childR = 2 * index + 1, 2 * index + 2
        } else {
            break
        }
    }

    return entry.item, entry.priority
}