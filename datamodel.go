package main

import (
	"crypto/md5"
	"fmt"
)

type DataModel struct {
	Data [][]ImageTag
	MD5  string
}

//将符合条件的静止空间变成运动空间
func (d *DataModel) SpaceToSpaceActive() {
	MaxHang := len(d.Data)
	MaxGe := len(d.Data[0])
	flag := true
	for flag {
		flag = false
		for i := 0; i < MaxHang; i++ {
			for j := 0; j < MaxGe; j++ {
				if d.Data[i][j] == ImageTag_SpaceActive {
					if i-1 >= 0 && d.Data[i-1][j] == ImageTag_Space {
						d.Data[i-1][j] = ImageTag_SpaceActive
						flag = true
					}
					if j-1 >= 0 && d.Data[i][j-1] == ImageTag_Space {
						d.Data[i][j-1] = ImageTag_SpaceActive
						flag = true
					}
					if i+1 < MaxHang && d.Data[i+1][j] == ImageTag_Space {
						d.Data[i+1][j] = ImageTag_SpaceActive
						flag = true
					}
					if j+1 < MaxGe && d.Data[i][j+1] == ImageTag_Space {
						d.Data[i][j+1] = ImageTag_SpaceActive
						flag = true
					}
				}
			}
		}
	}
}

func (d *DataModel) SpaceActiveToSpace() {
	MaxHang := len(d.Data)
	MaxGe := len(d.Data[0])
	for i := 0; i < MaxHang; i++ {
		for j := 0; j < MaxGe; j++ {
			if d.Data[i][j] == ImageTag_SpaceActive {
				d.Data[i][j] = ImageTag_Space
			}
		}
	}
}

//更新MD5，做相等判断用的
func (d *DataModel) UpdateMD5() {
	MaxHang := len(d.Data)
	MaxGe := len(d.Data[0])
	bytes := make([]byte, MaxHang*MaxGe)
	for i := 0; i < MaxHang; i++ {
		for j := 0; j < MaxGe; j++ {
			bytes[i*MaxGe+j] = byte(d.Data[i][j])
		}
	}
	d.MD5 = fmt.Sprintf("%X", md5.Sum(bytes)) //将[]byte转成16进制
}
func (d *DataModel) Copy() *DataModel {
	if d == nil {
		return nil
	}
	MaxHang := len(d.Data)
	MaxGe := len(d.Data[0])
	rettemp := make([][]ImageTag, MaxHang)
	for i := 0; i < MaxHang; i++ {
		rettemp[i] = make([]ImageTag, MaxGe)
	}

	for i := 0; i < MaxHang; i++ {
		for j := 0; j < MaxGe; j++ {
			rettemp[i][j] = d.Data[i][j]
		}
	}

	return &DataModel{Data: rettemp}
}
