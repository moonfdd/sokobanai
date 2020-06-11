// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"encoding/json"
	"fmt"
	"github.com/ying32/govcl/vcl"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

//::private::
type TFmMainFields struct {
	//版本号
	Version                string
	Images                 [MAXCOUNT][MAXCOUNT]*vcl.TImage
	PictureSpace           *vcl.TPicture
	PicturePerson          *vcl.TPicture
	PictureBox             *vcl.TPicture
	PictureTarget          *vcl.TPicture
	PictureWall            *vcl.TPicture
	PicturePersonAndTarget *vcl.TPicture
	PictureBoxAndTarget    *vcl.TPicture
	MaxHang                int
	MaxGe                  int

	//起点只有：墙，箱子，静止空间，运动空间
	StartDataModel *DataModel
	//终点只有：墙，箱子，静止空间，运动空间
	EndMap_Md5_DataModel map[string]*DataModel

	//StartChacheMap_Md5_DataModel map[string]*DataModel //缓存
	//EndChacheMap_Md5_DataModel   map[string]*DataModel //缓存

	StartChacheMap_Md5_Node_Previous map[string]*Node_Previous //缓存,方便去重
	EndChacheMap_Md5_Node_Next       map[string]*Node_Next     //缓存，方便去重

	StartSteps [][]*Node_Previous //方便单步遍历
	EndSteps   [][]*Node_Next     //方便单步遍历
	Steps      []*DataModel       //保存结果
	StepIndex  int
}

func (f *TFmMain) Init() {
	f.StartDataModel = nil
	f.EndMap_Md5_DataModel = nil
	f.StartChacheMap_Md5_Node_Previous = nil
	f.EndChacheMap_Md5_Node_Next = nil
	f.StartSteps = nil
	f.EndSteps = nil
	f.Steps = nil
	f.StepIndex = 0
}

func (f *TFmMain) OnFormCreate(sender vcl.IObject) {
	f.SetEnabled(false)
	go func() {
		time.Sleep(50 * time.Microsecond)
		vcl.ThreadSync(func() {
			f.PictureSpace = vcl.NewPicture()
			f.PictureSpace.LoadFromFile(ImagePath_Space)

			f.PicturePerson = vcl.NewPicture()
			f.PicturePerson.LoadFromFile(ImagePath_Person)

			f.PictureBox = vcl.NewPicture()
			f.PictureBox.LoadFromFile(ImagePath_Box)

			f.PictureTarget = vcl.NewPicture()
			f.PictureTarget.LoadFromFile(ImagePath_Target)

			f.PictureWall = vcl.NewPicture()
			f.PictureWall.LoadFromFile(ImagePath_Wall)

			f.PicturePersonAndTarget = vcl.NewPicture()
			f.PicturePersonAndTarget.LoadFromFile(ImagePath_PersonAndTarget)

			f.PictureBoxAndTarget = vcl.NewPicture()
			f.PictureBoxAndTarget.LoadFromFile(ImagePath_BoxAndTarget)

			f.Panel1.SetCaption("")
			f.PMap.SetCaption("")
			f.Version = "1.0.1"
			f.SetCaption(f.Caption() + f.Version)

			//f.ImageSelect.SetPicture(jpg)
			f.ImageSelect.SetStretch(true)
			f.OnBtnSpaceClick(nil)

			for i := 0; i < MAXCOUNT; i++ {
				for j := 0; j < MAXCOUNT; j++ {
					f.Images[i][j] = vcl.NewImage(f)
					f.Images[i][j].SetWidth(INT50)
					f.Images[i][j].SetHeight(INT50)
					f.Images[i][j].SetLeft(int32(j*INT50 + j*1))
					f.Images[i][j].SetTop(int32(i*INT50 + i*1))
					f.Images[i][j].Picture().LoadFromFile(ImagePath_Wall)
					f.Images[i][j].SetStretch(true)
					f.Images[i][j].SetParent(f.PMap)
					f.Images[i][j].SetTag(ImageTag_Wall)
					//f.Images[i][j].SetTag(i*MAXCOUNT+j)
					f.Images[i][j].SetOnClick(func(sender vcl.IObject) {
						img := vcl.AsImage(sender)
						selecttag := f.ImageSelect.Tag()
						tag := img.Tag()

						if selecttag == ImageTag_Space {
							//选中的是空间
							if tag == ImageTag_Wall {
								img.SetPicture(f.PictureSpace)
								img.SetTag(ImageTag_Space)
							} else if tag == ImageTag_Space {
								img.SetPicture(f.PictureWall)
								img.SetTag(ImageTag_Wall)
							} else {
								//fmt.Println(3)
							}
						} else if selecttag == ImageTag_Person {
							//选中的是人

							if tag == ImageTag_Space {
								img.SetPicture(f.PicturePerson)
								img.SetTag(ImageTag_Person)
							} else if tag == ImageTag_Person {
								img.SetPicture(f.PictureSpace)
								img.SetTag(ImageTag_Space)
							} else if tag == ImageTag_PersonAndTarget {
								img.SetPicture(f.PictureTarget)
								img.SetTag(ImageTag_Target)
							} else if tag == ImageTag_Target {
								img.SetPicture(f.PicturePersonAndTarget)
								img.SetTag(ImageTag_PersonAndTarget)
							}
						} else if selecttag == ImageTag_Box {
							//选中的是箱子

							if tag == ImageTag_Space {
								img.SetPicture(f.PictureBox)
								img.SetTag(ImageTag_Box)

							} else if tag == ImageTag_Box {
								img.SetPicture(f.PictureSpace)
								img.SetTag(ImageTag_Space)

							} else if tag == ImageTag_BoxAndTarget {
								img.SetPicture(f.PictureTarget)
								img.SetTag(ImageTag_Target)

							} else if tag == ImageTag_Target {
								img.SetPicture(f.PictureBoxAndTarget)
								img.SetTag(ImageTag_BoxAndTarget)
							}

						} else if selecttag == ImageTag_Target {
							//选中的是目标

							if tag == ImageTag_Space {
								img.SetPicture(f.PictureTarget)
								img.SetTag(ImageTag_Target)

							} else if tag == ImageTag_Target {
								img.SetPicture(f.PictureSpace)
								img.SetTag(ImageTag_Space)

							} else if tag == ImageTag_BoxAndTarget {
								img.SetPicture(f.PictureBox)
								img.SetTag(ImageTag_Box)

							} else if tag == ImageTag_Box {

								img.SetPicture(f.PictureBoxAndTarget)
								img.SetTag(ImageTag_BoxAndTarget)
							} else if tag == ImageTag_PersonAndTarget {

								img.SetPicture(f.PicturePerson)
								img.SetTag(ImageTag_Person)
							} else if tag == ImageTag_Person {

								img.SetPicture(f.PicturePersonAndTarget)
								img.SetTag(ImageTag_PersonAndTarget)
							}
						}
					})

				}
			}
			f.LoadFile()
			f.SetEnabled(true)
		})
	}()

}

func (f *TFmMain) OnBtnSpaceClick(sender vcl.IObject) {
	f.ImageSelect.SetPicture(f.PictureSpace)
	f.ImageSelect.SetTag(ImageTag_Space)
}

func (f *TFmMain) OnBtnTargetClick(sender vcl.IObject) {
	f.ImageSelect.SetPicture(f.PictureTarget)
	f.ImageSelect.SetTag(ImageTag_Target)
}

//func (f *TFmMain) OnBtnGoClick(sender vcl.IObject) {
//	f.MovePosition()
//}

//上一步
func (f *TFmMain) OnBtnPreviousClick(sender vcl.IObject) {
	f.StepIndex--
	f.UpdateBtn()
	f.UpdateStepUI()
}

func (f *TFmMain) OnPMapClick(sender vcl.IObject) {

}

func (f *TFmMain) OnBtnPersonClick(sender vcl.IObject) {
	f.ImageSelect.SetPicture(f.PicturePerson)
	f.ImageSelect.SetTag(ImageTag_Person)
}

func (f *TFmMain) OnBtnBoxClick(sender vcl.IObject) {
	f.ImageSelect.SetPicture(f.PictureBox)
	f.ImageSelect.SetTag(ImageTag_Box)
}

//func (f *TFmMain) OnBtnBackClick(sender vcl.IObject) {
//	f.MovePosition()
//	f.LoadStartData()
//}

//func (f *TFmMain) OnBtnTwoWayClick(sender vcl.IObject) {
//
//}

//下一步
func (f *TFmMain) OnBtnNextClick(sender vcl.IObject) {
	f.StepIndex++
	f.UpdateBtn()
	f.UpdateStepUI()
}

//整理位置
func (f *TFmMain) MovePosition() {
	x0 := MAXCOUNT
	y0 := MAXCOUNT
	x1 := 0
	y1 := 0
	tag := 0
	for i := 0; i < MAXCOUNT; i++ {
		for j := 0; j < MAXCOUNT; j++ {
			tag = f.Images[i][j].Tag()
			if tag != ImageTag_Wall {
				//fmt.Println("i = ", i, ",j = ", j)
				if x0 > i {
					x0 = i
				}
				if x1 < i {
					x1 = i
				}
				if y0 > j {
					y0 = j
				}
				if y1 < j {
					y1 = j
				}
			}
		}
	}
	f.MaxHang = x1 - x0 + 1
	f.MaxGe = y1 - y0 + 1
	for i := 0; i < MAXCOUNT; i++ {
		for j := 0; j < MAXCOUNT; j++ {
			if i < f.MaxHang && j < f.MaxGe {
				f.Images[i][j].SetPicture(f.Images[i+x0][j+y0].Picture())
				f.Images[i][j].SetTag(f.Images[i+x0][j+y0].Tag())
			} else {
				f.Images[i][j].SetPicture(f.PictureWall)
				f.Images[i][j].SetTag(ImageTag_Wall)
			}
		}
	}
}

//保存到文件
func (f *TFmMain) SaveFile() {
	dm := f.NewData()
	fdm := new(FileDataModel)
	fdm.MaxHang = f.MaxHang
	fdm.MaxGe = f.MaxGe
	fdm.Data = dm
	for i := 0; i < f.MaxHang; i++ {
		for j := 0; j < f.MaxGe; j++ {
			dm[i][j] = ImageTag(f.Images[i][j].Tag())
		}
	}
	data, _ := json.Marshal(fdm)
	//fmt.Printf("%s\n", string(data))
	//	//fmt.Fprintln("", 1)
	os.Remove("data.bin")
	ioutil.WriteFile("data.bin", data, 0x777)
}

//从文件中读取
func (f *TFmMain) LoadFile() {
	if data, err := ioutil.ReadFile("data.bin"); err == nil {
		var fdm *FileDataModel
		//fmt.Println(string(data))
		if json.Unmarshal(data, &fdm) == nil {
			//fmt.Println("dm0 = ", fdm)
			f.MaxHang = fdm.MaxHang
			f.MaxGe = fdm.MaxGe
			for i := 0; i < f.MaxHang; i++ {
				for j := 0; j < f.MaxGe; j++ {
					f.Images[i][j].SetTag(int(fdm.Data[i][j]))
					if fdm.Data[i][j] == ImageTag_Space {
						f.Images[i][j].SetPicture(f.PictureSpace)
					} else if fdm.Data[i][j] == ImageTag_Person {
						f.Images[i][j].SetPicture(f.PicturePerson)
					} else if fdm.Data[i][j] == ImageTag_Box {
						f.Images[i][j].SetPicture(f.PictureBox)
					} else if fdm.Data[i][j] == ImageTag_Target {
						f.Images[i][j].SetPicture(f.PictureTarget)
					} else if fdm.Data[i][j] == ImageTag_Wall {
						f.Images[i][j].SetPicture(f.PictureWall)
					} else if fdm.Data[i][j] == ImageTag_PersonAndTarget {
						f.Images[i][j].SetPicture(f.PicturePersonAndTarget)
					} else if fdm.Data[i][j] == ImageTag_BoxAndTarget {
						f.Images[i][j].SetPicture(f.PictureBoxAndTarget)
					}
				}
			}
		}
	}
}

//加载startdata 墙，箱子，静止空间，运动空间
func (f *TFmMain) LoadStartData() {
	//fmt.Println(f.MaxHang, "行")
	//fmt.Println(f.MaxGe, "个")
	f.StartDataModel = new(DataModel)
	f.StartDataModel.Data = f.NewData()
	for i := 0; i < f.MaxHang; i++ {
		for j := 0; j < f.MaxGe; j++ {
			if f.Images[i][j].Tag() == ImageTag_Space {
				f.StartDataModel.Data[i][j] = ImageTag(ImageTag_Space)
			} else if f.Images[i][j].Tag() == ImageTag_Target {
				f.StartDataModel.Data[i][j] = ImageTag(ImageTag_Space)
			} else if f.Images[i][j].Tag() == ImageTag_Wall {
				f.StartDataModel.Data[i][j] = ImageTag(ImageTag_Wall)
			} else if f.Images[i][j].Tag() == ImageTag_Person {
				f.StartDataModel.Data[i][j] = ImageTag(ImageTag_SpaceActive)
			} else if f.Images[i][j].Tag() == ImageTag_PersonAndTarget {
				f.StartDataModel.Data[i][j] = ImageTag(ImageTag_SpaceActive)
			} else if f.Images[i][j].Tag() == ImageTag_Box {
				f.StartDataModel.Data[i][j] = ImageTag(ImageTag_Box)
			} else if f.Images[i][j].Tag() == ImageTag_BoxAndTarget {
				f.StartDataModel.Data[i][j] = ImageTag(ImageTag_Box)
			}
		}
	}
	f.StartDataModel.SpaceToSpaceActive()
	f.StartDataModel.UpdateMD5()
	//if false {
	//	fmt.Println("----f.StartDataModel.Data = ", f.StartDataModel.Data)
	//	fmt.Println("----f.StartDataModel.MD5 = ", f.StartDataModel.MD5)
	//}
}

//创建二维数组
func (f *TFmMain) NewData() [][]ImageTag {
	ret := make([][]ImageTag, f.MaxHang)
	for i := 0; i < f.MaxHang; i++ {
		ret[i] = make([]ImageTag, f.MaxGe)
	}
	return ret
}

//加载listenddata 墙，箱子，静止空间，运动空间
func (f *TFmMain) LoadEndData() {
	f.EndMap_Md5_DataModel = make(map[string]*DataModel)
	tempdata := f.NewData()
	fulldata := f.NewData()
	var enddata [][]ImageTag = nil
	var enddatamodel *DataModel = nil
	//f.StartDataModel.Data = make([][]ImageTag, f.MaxHang)

	for i := 0; i < f.MaxHang; i++ {
		for j := 0; j < f.MaxGe; j++ {
			if f.Images[i][j].Tag() == ImageTag_Space {
				tempdata[i][j] = ImageTag_Space
				fulldata[i][j] = ImageTag_Space
			} else if f.Images[i][j].Tag() == ImageTag_Target {
				tempdata[i][j] = ImageTag_Box
				fulldata[i][j] = ImageTag_Box
			} else if f.Images[i][j].Tag() == ImageTag_Wall {
				tempdata[i][j] = ImageTag_Wall
				fulldata[i][j] = ImageTag_Wall
			} else if f.Images[i][j].Tag() == ImageTag_Person {
				tempdata[i][j] = ImageTag_Space
				fulldata[i][j] = ImageTag_Space
			} else if f.Images[i][j].Tag() == ImageTag_PersonAndTarget {
				tempdata[i][j] = ImageTag_Box
				fulldata[i][j] = ImageTag_Box
			} else if f.Images[i][j].Tag() == ImageTag_Box {
				tempdata[i][j] = ImageTag_Space
				fulldata[i][j] = ImageTag_Space
			} else if f.Images[i][j].Tag() == ImageTag_BoxAndTarget {
				tempdata[i][j] = ImageTag_Box
				fulldata[i][j] = ImageTag_Box
			}
		}
	}
	isforfor := true
	hangindex := 0
	geindex := 0
	//fmt.Println("fulldata = ", fulldata)
	for true {
		for i := 0; i < f.MaxHang; i++ {
			for j := 0; j < f.MaxGe; j++ {
				if fulldata[i][j] == ImageTag_Space {
					hangindex = i
					geindex = j
					isforfor = false
					break
				}
			}
			if !isforfor {
				break
			}
		}
		//fmt.Println(hangindex, "----1111111111111111111111----", geindex)
		if !isforfor {
			//fmt.Println("222222222222222222")
			isforfor = true
			enddata = f.NewData()
			enddatamodel = new(DataModel)
			enddatamodel.Data = enddata
			for i := 0; i < f.MaxHang; i++ {
				for j := 0; j < f.MaxGe; j++ {
					if hangindex == i && geindex == j {
						enddata[i][j] = ImageTag_SpaceActive
					} else {
						enddata[i][j] = tempdata[i][j]
					}
				}
			}
			enddatamodel.SpaceToSpaceActive()
			enddatamodel.UpdateMD5()
			//fmt.Println("enddatamodel = ", enddatamodel)
			f.EndMap_Md5_DataModel[enddatamodel.MD5] = enddatamodel
			for i := 0; i < f.MaxHang; i++ {
				for j := 0; j < f.MaxGe; j++ {
					if enddata[i][j] == ImageTag_SpaceActive {
						fulldata[i][j] = ImageTag_SpaceActive
					}
				}
			}
		} else {
			break
		}
	}
	//fmt.Println("LoadEndData =", len(f.EndMap_Md5_DataModel))
}

//加载到缓存
func (f *TFmMain) LoadToChache() {
	//if false {
	//	fmt.Println("2----f.StartDataModel.Data = ", f.StartDataModel.Data)
	//	fmt.Println("2----f.StartDataModel.MD5 = ", f.StartDataModel.MD5)
	//}
	f.StartChacheMap_Md5_Node_Previous = make(map[string]*Node_Previous)
	f.StartChacheMap_Md5_Node_Previous[f.StartDataModel.MD5] = &Node_Previous{Val: f.StartDataModel}
	f.StartSteps = make([][]*Node_Previous, 0)
	f.StartSteps = append(f.StartSteps, make([]*Node_Previous, 0))
	f.StartSteps[0] = append(f.StartSteps[0], &Node_Previous{Val: f.StartDataModel})
	//fmt.Println("f.StartDataModel = ", f.StartDataModel)

	f.EndSteps = make([][]*Node_Next, 0)
	f.EndSteps = append(f.EndSteps, make([]*Node_Next, 0))
	f.EndChacheMap_Md5_Node_Next = make(map[string]*Node_Next)
	for k, v := range f.EndMap_Md5_DataModel {
		temp := &Node_Next{Val: v}
		f.EndChacheMap_Md5_Node_Next[k] = temp
		f.EndSteps[0] = append(f.EndSteps[0], temp)
	}
	//fmt.Printf("f.StartSteps = %+v\r\n", f.StartSteps)
	//fmt.Printf("f.EndSteps = %+v\r\n", f.EndSteps)

	//if true {
	//	for i := len(f.StartSteps) - 1; i >= 0; i-- {
	//		for j := len(f.StartSteps[i]) - 1; j >= 0; j-- {
	//			fmt.Println("f.StartSteps[", i, "][", j, "].Val.Data = ", f.StartSteps[i][j].Val)
	//		}
	//	}
	//
	//	for i := len(f.EndSteps) - 1; i >= 0; i-- {
	//		for j := len(f.EndSteps[i]) - 1; j >= 0; j-- {
	//			fmt.Println("f.EndSteps[", i, "][", j, "].Val.Data = ", f.EndSteps[i][j].Val)
	//		}
	//	}
	//}
}

//推箱子ai
func (f *TFmMain) OnBtnGoAIClick(sender vcl.IObject) {
	f.Init()
	f.MovePosition()
	f.SaveFile()
	f.LoadStartData()
	f.LoadEndData()
	f.LoadToChache()
	var d0 time.Duration = 0
	var ret StepState = StepState_Continue
	for ret == StepState_Continue {
		now := time.Now()
		ret = f.PushOneStep()
		d0 += time.Now().Sub(now)
		//if true {
		//	fmt.Println("ret = ", ret)
		//}
	}
	if true {
		//fmt.Println("完成")
		//fmt.Println("正向：", d0)

		msg := fmt.Sprintln("正向：", d0)
		msg += fmt.Sprintln("")
		msg += fmt.Sprintln("总数据：", len(f.StartChacheMap_Md5_Node_Previous)+len(f.EndChacheMap_Md5_Node_Next))
		msg += fmt.Sprintln("正向数据：", len(f.StartChacheMap_Md5_Node_Previous))
		msg += fmt.Sprintln("逆向数据：", len(f.EndChacheMap_Md5_Node_Next))

		msg += fmt.Sprintln("")
		msg += fmt.Sprintln("正向步数：", len(f.StartSteps)-1)
		vcl.ShowMessage(msg)
		//for i := 0; i < len(f.Steps); i++ {
		//	fmt.Println(f.Steps[i])
		//}
	}
	f.StepIndex = 0
	f.UpdateBtn()
}

//拉箱子ai
func (f *TFmMain) OnBtnBackAIClick(sender vcl.IObject) {
	f.Init()
	f.MovePosition()
	f.SaveFile()
	f.LoadStartData()
	f.LoadEndData()
	f.LoadToChache()
	var d1 time.Duration = 0
	var ret StepState = StepState_Continue
	for ret == StepState_Continue {
		now := time.Now()
		ret = f.PullOneStep()
		d1 += time.Now().Sub(now)
		//if false {
		//	fmt.Println("ret = ", ret)
		//}
	}
	if true {
		//fmt.Println("完成")
		//fmt.Println("逆向：", d1)
		msg := fmt.Sprintln("逆向：", d1)
		msg += fmt.Sprintln("")
		msg += fmt.Sprintln("总数据：", len(f.StartChacheMap_Md5_Node_Previous)+len(f.EndChacheMap_Md5_Node_Next))
		msg += fmt.Sprintln("正向数据：", len(f.StartChacheMap_Md5_Node_Previous))
		msg += fmt.Sprintln("逆向数据：", len(f.EndChacheMap_Md5_Node_Next))

		msg += fmt.Sprintln("")
		msg += fmt.Sprintln("逆向步数：", len(f.EndSteps)-1)
		vcl.ShowMessage(msg)
		//for i := 0; i < len(f.Steps); i++ {
		//	fmt.Println(f.Steps[i])
		//}
	}
	f.StepIndex = 0
	f.UpdateBtn()
}

//生成步骤
func (f *TFmMain) GenSteps(nodeprevious *Node_Previous, nodenext *Node_Next) {
	f.Steps = make([]*DataModel, 0)

	//遍历
	for {
		f.Steps = append([]*DataModel{nodeprevious.Val}, f.Steps...)
		nodeprevious = nodeprevious.Previous
		if nodeprevious == nil {
			break
		}
	}

	//遍历
	for {
		f.Steps = append(f.Steps, nodenext.Val)

		nodenext = nodenext.Next
		if nodenext == nil {
			break
		}
	}
	//fmt.Println("总数量：", len(f.StartChacheMap_Md5_Node_Previous)+len(f.EndChacheMap_Md5_Node_Next))
}

//更新按钮界面
func (f *TFmMain) UpdateBtn() {
	if f.StepIndex == 0 {
		f.BtnPrevious.SetEnabled(false)
		f.BtnNext.SetEnabled(true)
	} else if f.StepIndex >= len(f.Steps)-1 {
		f.BtnPrevious.SetEnabled(true)
		f.BtnNext.SetEnabled(false)
	} else {
		f.BtnPrevious.SetEnabled(true)
		f.BtnNext.SetEnabled(true)
	}
	f.TxtMsg.SetCaption(strconv.Itoa(f.StepIndex) + "/" + strconv.Itoa(len(f.Steps)-1))
}

//更新界面
func (f *TFmMain) UpdateStepUI() {
	var enddatamodel *DataModel = nil
	for _, v := range f.EndMap_Md5_DataModel {
		enddatamodel = v
		break
	}
	pi := -1
	pj := -1
	if f.StepIndex == 0 {
		//当前 下一个
		pi1 := -1
		pj1 := -1
		pi2 := -1
		pj2 := -1
		for i := 0; i < f.MaxHang; i++ {
			for j := 0; j < f.MaxGe; j++ {
				if f.Steps[f.StepIndex].Data[i][j] == ImageTag_Box &&
					(f.Steps[f.StepIndex+1].Data[i][j] == ImageTag_Space ||
						f.Steps[f.StepIndex+1].Data[i][j] == ImageTag_SpaceActive) {
					pi1 = i
					pj1 = j
					//fmt.Println("pi1 = ", pi1)
					//fmt.Println("pj1 = ", pj1)
				} else if f.Steps[f.StepIndex+1].Data[i][j] == ImageTag_Box &&
					(f.Steps[f.StepIndex].Data[i][j] == ImageTag_Space ||
						f.Steps[f.StepIndex].Data[i][j] == ImageTag_SpaceActive) {
					pi2 = i
					pj2 = j
					//fmt.Println("pi2 = ", pi2)
					//fmt.Println("pj2 = ", pj2)
				}
			}
		}
		if pi1 < pi2 {
			//fmt.Println("下推")
			pi = pi1 - 1
			pj = pj1
		} else if pi1 > pi2 {
			//fmt.Println("上推")
			pi = pi1 + 1
			pj = pj1
		} else {
			if pj1 < pj2 {
				//fmt.Println("右推")
				pi = pi1
				pj = pj1 - 1
			} else if pj1 > pj2 {
				//fmt.Println("左推")
				pi = pi1
				pj = pj1 + 1
			}
		}
		//fmt.Println("pi = ", pi)
		//fmt.Println("pj = ", pj)
	} else {
		//上一个 当前
		pi1 := -1
		pj1 := -1
		pi2 := -1
		pj2 := -1
		for i := 0; i < f.MaxHang; i++ {
			for j := 0; j < f.MaxGe; j++ {
				if f.Steps[f.StepIndex-1].Data[i][j] == ImageTag_Box &&
					(f.Steps[f.StepIndex].Data[i][j] == ImageTag_Space ||
						f.Steps[f.StepIndex].Data[i][j] == ImageTag_SpaceActive) {
					pi1 = i
					pj1 = j
					//fmt.Println("pi1 = ", pi1)
					//fmt.Println("pj1 = ", pj1)
				} else if f.Steps[f.StepIndex].Data[i][j] == ImageTag_Box &&
					(f.Steps[f.StepIndex-1].Data[i][j] == ImageTag_Space ||
						f.Steps[f.StepIndex-1].Data[i][j] == ImageTag_SpaceActive) {
					pi2 = i
					pj2 = j
					//fmt.Println("pi2 = ", pi2)
					//fmt.Println("pj2 = ", pj2)
				}
			}
		}
		if pi1 < pi2 {
			//fmt.Println("下推")
			pi = pi2 - 1
			pj = pj2
		} else if pi1 > pi2 {
			//fmt.Println("上推")
			pi = pi2 + 1
			pj = pj2
		} else {
			if pj1 < pj2 {
				//fmt.Println("右推")
				pi = pi2
				pj = pj2 - 1
			} else if pj1 > pj2 {
				//fmt.Println("左推")
				pi = pi2
				pj = pj2 + 1
			}
		}
		//fmt.Println("pi = ", pi)
		//fmt.Println("pj = ", pj)
	}
	for i := 0; i < f.MaxHang; i++ {
		for j := 0; j < f.MaxGe; j++ {

			if f.Steps[f.StepIndex].Data[i][j] == ImageTag_Space ||
				f.Steps[f.StepIndex].Data[i][j] == ImageTag_SpaceActive {
				f.Images[i][j].SetPicture(f.PictureSpace)
				if i == pi && j == pj {
					f.Images[i][j].SetPicture(f.PicturePerson)
				}
				if enddatamodel.Data[i][j] == ImageTag_Box {
					f.Images[i][j].SetPicture(f.PictureTarget)
					if i == pi && j == pj {
						f.Images[i][j].SetPicture(f.PicturePersonAndTarget)
					}
				}
			} else if f.Steps[f.StepIndex].Data[i][j] == ImageTag_Box {
				f.Images[i][j].SetPicture(f.PictureBox)
				if enddatamodel.Data[i][j] == ImageTag_Box {
					f.Images[i][j].SetPicture(f.PictureBoxAndTarget)
				}
			}

		}
	}
}

//拉箱子一次
func (f *TFmMain) PullOneStep() StepState {
	lastindex := len(f.EndSteps) - 1
	//fmt.Println("lastindex = ", lastindex)
	//fmt.Printf("len(f.EndSteps[lastindex]) = %+v\r\n", len(f.EndSteps[lastindex]))

	currentindex := lastindex + 1
	//f.EndSteps = make([][]*Node_Next, 0)
	f.EndSteps = append(f.EndSteps, make([]*Node_Next, 0))
	lastlen := len(f.EndSteps[lastindex])
	if lastlen > 0 {
		//var lasttempmd5 string
		var lasttemp *Node_Next
		var tempdatamodel *Node_Next
		for k := 0; k < lastlen; k++ {
			lasttemp = f.EndSteps[lastindex][k]
			//lasttempmd5 = lasttemp.Val
			for i := 0; i < f.MaxHang; i++ {
				for j := 0; j < f.MaxGe; j++ {
					if f.EndChacheMap_Md5_Node_Next[lasttemp.Val.MD5].Val.Data[i][j] == ImageTag_Box {
						//往上拉
						if i-2 >= 0 &&
							lasttemp.Val.Data[i-2][j] == ImageTag_SpaceActive &&
							lasttemp.Val.Data[i-1][j] == ImageTag_SpaceActive {
							//fmt.Println("上拉")
							tempdatamodel = &Node_Next{Val: lasttemp.Val.Copy(), Next: lasttemp} //复制
							tempdatamodel.Val.SpaceActiveToSpace()                               // 将所有运动空间变成静止空间
							tempdatamodel.Val.Data[i-2][j] = ImageTag_SpaceActive
							tempdatamodel.Val.Data[i-1][j] = ImageTag_Box
							tempdatamodel.Val.Data[i][j] = ImageTag_Space
							tempdatamodel.Val.SpaceToSpaceActive() //将符合条件的静止空间变成运动空间
							tempdatamodel.Val.UpdateMD5()
							if _, ok := f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5]; ok { //说明找到起点了
								//获得结果
								f.GenSteps(f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5], lasttemp)
								//fmt.Println("成功")
								return StepState_Suc
							} else {
								if _, ok := f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5]; ok { //说明已经有缓存，不需要添加
								} else { //需要添加到缓存
									f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5] = tempdatamodel
									f.EndSteps[currentindex] = append(f.EndSteps[currentindex], &Node_Next{Val: tempdatamodel.Val, Next: lasttemp})
								}
							}
						}
						//往下拉
						if i+2 < f.MaxHang &&
							lasttemp.Val.Data[i+2][j] == ImageTag_SpaceActive &&
							lasttemp.Val.Data[i+1][j] == ImageTag_SpaceActive {
							//fmt.Println("下拉")
							tempdatamodel = &Node_Next{Val: lasttemp.Val.Copy(), Next: lasttemp} //复制
							tempdatamodel.Val.SpaceActiveToSpace()                               // 将所有运动空间变成静止空间
							tempdatamodel.Val.Data[i+2][j] = ImageTag_SpaceActive
							tempdatamodel.Val.Data[i+1][j] = ImageTag_Box
							tempdatamodel.Val.Data[i][j] = ImageTag_Space
							tempdatamodel.Val.SpaceToSpaceActive() //将符合条件的静止空间变成运动空间
							tempdatamodel.Val.UpdateMD5()
							if _, ok := f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5]; ok { //说明找到起点了
								//获得结果
								f.GenSteps(f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5], lasttemp)
								//fmt.Println("成功")
								return StepState_Suc
							} else {
								if _, ok := f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5]; ok { //说明已经有缓存，不需要添加
								} else { //需要添加到缓存
									f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5] = tempdatamodel
									f.EndSteps[currentindex] = append(f.EndSteps[currentindex], &Node_Next{Val: tempdatamodel.Val, Next: lasttemp})
								}
							}
						}
						//往左拉
						if j-2 >= 0 &&
							lasttemp.Val.Data[i][j-2] == ImageTag_SpaceActive &&
							lasttemp.Val.Data[i][j-1] == ImageTag_SpaceActive {
							//fmt.Println("")
							//fmt.Println("左拉")
							tempdatamodel = &Node_Next{Val: lasttemp.Val.Copy(), Next: lasttemp} //复制
							tempdatamodel.Val.SpaceActiveToSpace()                               // 将所有运动空间变成静止空间
							tempdatamodel.Val.Data[i][j-2] = ImageTag_SpaceActive
							tempdatamodel.Val.Data[i][j-1] = ImageTag_Box
							tempdatamodel.Val.Data[i][j] = ImageTag_Space
							tempdatamodel.Val.SpaceToSpaceActive() //将符合条件的静止空间变成运动空间
							tempdatamodel.Val.UpdateMD5()
							if _, ok := f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5]; ok { //说明找到起点了
								//fmt.Println("左拉0")
								//获得结果
								f.GenSteps(f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5], lasttemp)
								//fmt.Println("成功")
								return StepState_Suc
							} else {
								//fmt.Println("左拉1")
								if _, ok := f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5]; ok { //说明已经有缓存，不需要添加
								} else { //需要添加到缓存
									f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5] = tempdatamodel
									f.EndSteps[currentindex] = append(f.EndSteps[currentindex], &Node_Next{Val: tempdatamodel.Val, Next: lasttemp})
								}
							}
						}
						//往右拉
						if j+2 < f.MaxGe &&
							lasttemp.Val.Data[i][j+2] == ImageTag_SpaceActive &&
							lasttemp.Val.Data[i][j+1] == ImageTag_SpaceActive {
							//fmt.Println("")
							//fmt.Println("右拉")
							tempdatamodel = &Node_Next{Val: lasttemp.Val.Copy(), Next: lasttemp} //复制
							tempdatamodel.Val.SpaceActiveToSpace()                               // 将所有运动空间变成静止空间
							tempdatamodel.Val.Data[i][j+2] = ImageTag_SpaceActive
							tempdatamodel.Val.Data[i][j+1] = ImageTag_Box
							tempdatamodel.Val.Data[i][j] = ImageTag_Space
							tempdatamodel.Val.SpaceToSpaceActive() //将符合条件的静止空间变成运动空间
							tempdatamodel.Val.UpdateMD5()
							if _, ok := f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5]; ok { //说明找到起点了
								//获得结果
								f.GenSteps(f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5], lasttemp)
								//fmt.Println("成功")
								return StepState_Suc
							} else {
								if _, ok := f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5]; ok { //说明已经有缓存，不需要添加
								} else { //需要添加到缓存
									f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5] = tempdatamodel
									f.EndSteps[currentindex] = append(f.EndSteps[currentindex], &Node_Next{Val: tempdatamodel.Val, Next: lasttemp})
								}
							}
						}
					}
				}
			}
		}
		return StepState_Continue
	} else {
		return StepState_Failed
	}
}

//推箱子一次
func (f *TFmMain) PushOneStep() StepState {

	lastindex := len(f.StartSteps) - 1
	currentindex := lastindex + 1
	f.StartSteps = append(f.StartSteps, make([]*Node_Previous, 0))
	lastlen := len(f.StartSteps[lastindex])
	if lastlen > 0 {
		var lasttemp *Node_Previous
		var tempdatamodel *Node_Previous
		for k := 0; k < lastlen; k++ {
			lasttemp = f.StartSteps[lastindex][k]
			for i := 0; i < f.MaxHang; i++ {
				for j := 0; j < f.MaxGe; j++ {
					if f.StartChacheMap_Md5_Node_Previous[lasttemp.Val.MD5].Val.Data[i][j] == ImageTag_Box {
						//往上推
						if i-1 >= 0 && i+1 < f.MaxHang &&
							(lasttemp.Val.Data[i-1][j] == ImageTag_SpaceActive || lasttemp.Val.Data[i-1][j] == ImageTag_Space) && //静止空间和活动空间
							lasttemp.Val.Data[i+1][j] == ImageTag_SpaceActive { //活动空间
							//fmt.Println("上推")
							tempdatamodel = &Node_Previous{Val: lasttemp.Val.Copy(), Previous: lasttemp} //复制
							tempdatamodel.Val.SpaceActiveToSpace()                                       // 将所有运动空间变成静止空间
							tempdatamodel.Val.Data[i-1][j] = ImageTag_Box                                //变成箱子
							tempdatamodel.Val.Data[i][j] = ImageTag_SpaceActive                          //变成活动空间
							tempdatamodel.Val.SpaceToSpaceActive()                                       //将符合条件的静止空间变成运动空间
							tempdatamodel.Val.UpdateMD5()
							if _, ok := f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5]; ok { //说明找到终点了
								//获得结果
								f.GenSteps(lasttemp, f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5])
								//fmt.Println("成功")
								return StepState_Suc
							} else {
								//fmt.Println("未找到")
								if _, ok := f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5]; ok { //说明已经有缓存，不需要添加
								} else { //需要添加到缓存
									if f.IsPushDeadLock(tempdatamodel) {

									} else {
										f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5] = tempdatamodel
										f.StartSteps[currentindex] = append(f.StartSteps[currentindex], &Node_Previous{Val: tempdatamodel.Val, Previous: lasttemp})
									}
								}
							}
						}

						//往下推
						if i-1 >= 0 && i+1 < f.MaxHang &&
							(lasttemp.Val.Data[i+1][j] == ImageTag_SpaceActive || lasttemp.Val.Data[i+1][j] == ImageTag_Space) && //静止空间和活动空间
							lasttemp.Val.Data[i-1][j] == ImageTag_SpaceActive { //活动空间
							//fmt.Println("下推")
							tempdatamodel = &Node_Previous{Val: lasttemp.Val.Copy(), Previous: lasttemp} //复制
							tempdatamodel.Val.SpaceActiveToSpace()                                       // 将所有运动空间变成静止空间
							tempdatamodel.Val.Data[i+1][j] = ImageTag_Box                                //变成箱子
							tempdatamodel.Val.Data[i][j] = ImageTag_SpaceActive                          //变成活动空间
							tempdatamodel.Val.SpaceToSpaceActive()                                       //将符合条件的静止空间变成运动空间
							tempdatamodel.Val.UpdateMD5()
							if _, ok := f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5]; ok { //说明找到终点了
								//获得结果
								f.GenSteps(lasttemp, f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5])
								//fmt.Println("成功")
								return StepState_Suc
							} else {
								//fmt.Println("未找到")
								if _, ok := f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5]; ok { //说明已经有缓存，不需要添加
								} else { //需要添加到缓存
									if f.IsPushDeadLock(tempdatamodel) {

									} else {
										f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5] = tempdatamodel
										f.StartSteps[currentindex] = append(f.StartSteps[currentindex], &Node_Previous{Val: tempdatamodel.Val, Previous: lasttemp})
									}
								}
							}
						}

						//往左推
						if j-1 >= 0 && j+1 < f.MaxGe &&
							(lasttemp.Val.Data[i][j-1] == ImageTag_SpaceActive || lasttemp.Val.Data[i][j-1] == ImageTag_Space) && //静止空间和活动空间
							lasttemp.Val.Data[i][j+1] == ImageTag_SpaceActive { //活动空间
							//fmt.Println("左推")
							tempdatamodel = &Node_Previous{Val: lasttemp.Val.Copy(), Previous: lasttemp} //复制
							tempdatamodel.Val.SpaceActiveToSpace()                                       // 将所有运动空间变成静止空间
							tempdatamodel.Val.Data[i][j-1] = ImageTag_Box                                //变成箱子
							tempdatamodel.Val.Data[i][j] = ImageTag_SpaceActive                          //变成活动空间
							tempdatamodel.Val.SpaceToSpaceActive()                                       //将符合条件的静止空间变成运动空间
							tempdatamodel.Val.UpdateMD5()
							if _, ok := f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5]; ok { //说明找到终点了
								//获得结果
								f.GenSteps(lasttemp, f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5])
								//fmt.Println("成功")
								return StepState_Suc
							} else {
								//fmt.Println("未找到")
								if _, ok := f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5]; ok { //说明已经有缓存，不需要添加
								} else { //需要添加到缓存
									if f.IsPushDeadLock(tempdatamodel) {

									} else {
										f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5] = tempdatamodel
										f.StartSteps[currentindex] = append(f.StartSteps[currentindex], &Node_Previous{Val: tempdatamodel.Val, Previous: lasttemp})
									}
								}
							}
						}

						//往右推
						if j-1 >= 0 && j+1 < f.MaxGe &&
							(lasttemp.Val.Data[i][j+1] == ImageTag_SpaceActive || lasttemp.Val.Data[i][j+1] == ImageTag_Space) && //静止空间和活动空间
							lasttemp.Val.Data[i][j-1] == ImageTag_SpaceActive { //活动空间
							//fmt.Println("右推")
							tempdatamodel = &Node_Previous{Val: lasttemp.Val.Copy(), Previous: lasttemp} //复制
							tempdatamodel.Val.SpaceActiveToSpace()                                       // 将所有运动空间变成静止空间
							tempdatamodel.Val.Data[i][j+1] = ImageTag_Box                                //变成箱子
							tempdatamodel.Val.Data[i][j] = ImageTag_SpaceActive                          //变成活动空间
							tempdatamodel.Val.SpaceToSpaceActive()                                       //将符合条件的静止空间变成运动空间
							tempdatamodel.Val.UpdateMD5()
							if _, ok := f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5]; ok { //说明找到终点了
								//获得结果
								f.GenSteps(lasttemp, f.EndChacheMap_Md5_Node_Next[tempdatamodel.Val.MD5])
								//fmt.Println("成功")
								return StepState_Suc
							} else {
								//fmt.Println("未找到")
								if _, ok := f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5]; ok { //说明已经有缓存，不需要添加
								} else { //需要添加到缓存
									if f.IsPushDeadLock(tempdatamodel) {

									} else {
										f.StartChacheMap_Md5_Node_Previous[tempdatamodel.Val.MD5] = tempdatamodel
										f.StartSteps[currentindex] = append(f.StartSteps[currentindex], &Node_Previous{Val: tempdatamodel.Val, Previous: lasttemp})
									}
								}
							}
						}
					}
				}
			}
		}
		return StepState_Continue
	} else {
		return StepState_Failed
	}
}

//推箱子是否死锁
func (f *TFmMain) IsPushDeadLock(tempdatamodel *Node_Previous) bool {
	dm := tempdatamodel.Val.Copy()
	dm.SpaceActiveToSpace()

	//获取箱子列表，方便遍历
	nodelist := make([]*BoxNode, 0)
	for i := 0; i < f.MaxHang; i++ {
		for j := 0; j < f.MaxGe; j++ {
			if dm.Data[i][j] == ImageTag_Box {
				nodelist = append(nodelist, &BoxNode{X: i, Y: j})
			}
		}
	}

	//消灭能推动的箱子
	diecount := 1
	nodelistlen := len(nodelist)
	for nodelistlen > 0 && diecount > 0 {
		diecount = 0
		for i := nodelistlen - 1; i >= 0; i-- {
			if (nodelist[i].X-1 >= 0 && nodelist[i].X+1 < f.MaxHang &&
				dm.Data[nodelist[i].X-1][nodelist[i].Y] == ImageTag_Space &&
				dm.Data[nodelist[i].X+1][nodelist[i].Y] == ImageTag_Space) ||

				(nodelist[i].Y-1 >= 0 && nodelist[i].Y+1 < f.MaxGe &&
					dm.Data[nodelist[i].X][nodelist[i].Y-1] == ImageTag_Space &&
					dm.Data[nodelist[i].X][nodelist[i].Y+1] == ImageTag_Space) { //能推动，必须消灭
				dm.Data[nodelist[i].X][nodelist[i].Y] = ImageTag_Space //二维数组中消灭
				diecount++                                             //消灭次数
				nodelist = append(nodelist[:i], nodelist[i+1:]...)     //箱子列表中消灭
			}

		}
		nodelistlen = len(nodelist)
	}

	//判断箱子是否在目标位置。不在的话直接返回true；在的话，继续
	if nodelistlen > 0 {
		//EndMap_Md5_DataModel
		var enddatamodel *DataModel
		for _, v := range f.EndMap_Md5_DataModel {
			enddatamodel = v
			break
		}
		for i := nodelistlen - 1; i >= 0; i-- {

			if dm.Data[nodelist[i].X][nodelist[i].Y] == enddatamodel.Data[nodelist[i].X][nodelist[i].Y] {
			} else {
				return true
			}
		}
	}
	return false
}

//双向ai
func (f *TFmMain) OnBtnTwoWayAIClick(sender vcl.IObject) {
	f.Init()
	f.MovePosition()
	f.SaveFile()
	f.LoadStartData()
	f.LoadEndData()
	f.LoadToChache()
	var ret StepState = StepState_Continue
	var d0 time.Duration = 0
	var d1 time.Duration = 0
	for ret == StepState_Continue {
		//双向13-22919  21-19680 20-20716
		//if 194*len(f.EndChacheMap_Md5_Node_Next) > 309*len(f.StartChacheMap_Md5_Node_Previous) {
		//if len(f.EndChacheMap_Md5_Node_Next) > len(f.StartChacheMap_Md5_Node_Previous) {

		startlen := len(f.StartSteps[len(f.StartSteps)-1])
		endlen := len(f.EndSteps[len(f.EndSteps)-1])

		//fmt.Println("startlen = ", startlen)
		//fmt.Println("endlen = ", endlen)
		//vcl.ShowMessage("1")

		//if 194*endlen > 309*startlen {
		//if 1*endlen > 1*startlen {
		if endlen > startlen<<1 {
			//if 194*endlen > 309*startlen {
			now := time.Now()
			ret = f.PushOneStep()
			d0 += time.Now().Sub(now)
		} else {
			now := time.Now()
			ret = f.PullOneStep()
			d1 += time.Now().Sub(now)
		}
		//if false {
		//	fmt.Println("ret = ", ret)
		//}
	}
	if true {
		//fmt.Println("完成")
		//fmt.Println("总共：", d0+d1)
		//fmt.Println("正向：", d0)
		//fmt.Println("逆向：", d1)
		//for i := 0; i < len(f.Steps); i++ {
		//	fmt.Println(f.Steps[i])
		//}
		msg := fmt.Sprintln("总共：", d0+d1)
		msg += fmt.Sprintln("正向：", d0)
		msg += fmt.Sprintln("逆向：", d1)
		msg += fmt.Sprintln("")
		msg += fmt.Sprintln("总数据：", len(f.StartChacheMap_Md5_Node_Previous)+len(f.EndChacheMap_Md5_Node_Next))
		msg += fmt.Sprintln("正向数据：", len(f.StartChacheMap_Md5_Node_Previous))
		msg += fmt.Sprintln("逆向数据：", len(f.EndChacheMap_Md5_Node_Next))

		msg += fmt.Sprintln("")
		msg += fmt.Sprintln("正向步数：", len(f.StartSteps)-1)
		msg += fmt.Sprintln("逆向步数：", len(f.EndSteps)-1)
		vcl.ShowMessage(msg)

	}
	f.StepIndex = 0
	f.UpdateBtn()
}

//清空
func (f *TFmMain) OnBtnClearClick(sender vcl.IObject) {
	for i := 0; i < MAXCOUNT; i++ {
		for j := 0; j < MAXCOUNT; j++ {
			f.Images[i][j].SetPicture(f.PictureWall)
			f.Images[i][j].SetTag(ImageTag_Wall)
		}
	}
}
