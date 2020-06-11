package main

//拉箱子的节点
type Node_Next struct {
	Val  *DataModel
	Next *Node_Next
}
