package main

const (
	ImagePath_Space           = "./images/空间.png"
	ImagePath_Person          = "./images/人.png"
	ImagePath_Box             = "./images/箱子.png"
	ImagePath_Target          = "./images/目标.png"
	ImagePath_Wall            = "./images/墙.png"
	ImagePath_PersonAndTarget = "./images/人目标.png"
	ImagePath_BoxAndTarget    = "./images/箱子目标.png"
	MAXCOUNT                  = 15
	INT50                     = 40
	ImageTag_Space            = 1 //1
	ImageTag_SpaceActive      = 8 //1
	ImageTag_Person           = 2
	ImageTag_Box              = 3 //1
	ImageTag_Target           = 4
	ImageTag_Wall             = 5 //1
	ImageTag_PersonAndTarget  = 6
	ImageTag_BoxAndTarget     = 7
	StepState_Suc             = 1
	StepState_Failed          = 2
	StepState_Continue        = 3
)

type ImageTag byte

type StepState byte
