package main

import (
    "container/heap"
    "fmt"
)

// Node represents a node in the Huffman tree
type Node struct {
    value    rune
    freq     int
    left     *Node
    right    *Node
}

// PriorityQueue implements heap.Interface and holds Nodes
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].freq < pq[j].freq
}
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*Node)
    *pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

// BuildHuffmanTree constructs the Huffman tree from a map of character frequencies
func BuildHuffmanTree(freq map[rune]int) *Node {
    var pq PriorityQueue
    for char, f := range freq {
        pq = append(pq, &Node{char, f, nil, nil})
    }
    heap.Init(&pq)
    for pq.Len() > 1 {
        left := heap.Pop(&pq).(*Node)
        right := heap.Pop(&pq).(*Node)
        parent := &Node{0, left.freq + right.freq, left, right}
        heap.Push(&pq, parent)
    }
    return heap.Pop(&pq).(*Node)
}

// BuildCodes generates Huffman codes for each character
func BuildCodes(root *Node, code string, codes map[rune]string) {
    if root == nil {
        return
    }
    if root.left == nil && root.right == nil {
        codes[root.value] = code
    }
    BuildCodes(root.left, code+"0", codes)
    BuildCodes(root.right, code+"1", codes)
}

// Encode encodes a string using Huffman coding
func Encode(s string) (encoded string, codes map[rune]string) {
    freq := make(map[rune]int)
    for _, char := range s {
        freq[char]++
    }
    root := BuildHuffmanTree(freq)
    codes = make(map[rune]string)
    BuildCodes(root, "", codes)
    for _, char := range s {
        encoded += codes[char]
    }
    return encoded, codes
}

// Decode decodes a string using Huffman coding
func Decode(encoded string, root *Node) (decoded string) {
    current := root
    for _, bit := range encoded {
        if bit == '0' {
            current = current.left
        } else {
            current = current.right
        }
        if current.left == nil && current.right == nil {
            decoded += string(current.value)
            current = root
        }
    }
    return decoded
}

func main() {
    message := "hello world"
    encoded, codes := Encode(message)
    fmt.Println("Original message:", message)
    fmt.Println("Encoded message:", encoded)
    fmt.Println("Huffman codes:")
    for char, code := range codes {
        fmt.Printf("%c: %s\n", char, code)
    }
    root := BuildHuffmanTree(map[rune]int{'h': 1, 'e': 1, 'l': 3, 'o': 2, ' ': 1, 'w': 1, 'r': 1, 'd': 1})
    decoded := Decode(encoded, root)
    fmt.Println("Decoded message:", decoded)
}

