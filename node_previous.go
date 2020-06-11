package main

//推箱子的节点
type Node_Previous struct {
	Val      *DataModel
	Previous *Node_Previous
}
