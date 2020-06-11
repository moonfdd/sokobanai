package main

const (
	//空间路径
	ImagePath_Space = "./images/空间.png"
	//人路径
	ImagePath_Person = "./images/人.png"
	//箱子路径
	ImagePath_Box = "./images/箱子.png"
	//目标路径
	ImagePath_Target = "./images/目标.png"
	//墙路径
	ImagePath_Wall = "./images/墙.png"
	//人目标路径
	ImagePath_PersonAndTarget = "./images/人目标.png"
	//箱子目标路径
	ImagePath_BoxAndTarget = "./images/箱子目标.png"
	//最大行数和每行的最大个数
	MAXCOUNT = 15
	INT50    = 40
	//静止空间
	ImageTag_Space = 1 //起点和终点只有：墙，箱子，静止空间，运动空间
	//活动空间
	ImageTag_SpaceActive = 8 //起点和终点只有：墙，箱子，静止空间，运动空间
	//人
	ImageTag_Person = 2
	//箱子
	ImageTag_Box = 3 //起点和终点只有：墙，箱子，静止空间，运动空间
	//目标
	ImageTag_Target = 4
	//墙
	ImageTag_Wall = 5 //起点和终点只有：墙，箱子，静止空间，运动空间
	//人目标
	ImageTag_PersonAndTarget = 6
	//箱子目标
	ImageTag_BoxAndTarget = 7
	//成功
	StepState_Suc = 1
	//失败
	StepState_Failed = 2
	//继续
	StepState_Continue = 3
)

//一个格子的状态：静止空间，活动空间，人，箱子，目标，墙，人目标，箱子目标。
type ImageTag byte

//推箱子或者拉箱子状态：成功，失败，继续
type StepState byte
