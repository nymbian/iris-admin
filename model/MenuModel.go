package model

import (
	"errors"
	"iris-admin/libs"
	"strings"

	"github.com/jinzhu/gorm"
)

var MenuListTree []Menu

type Menu struct {
	gorm.Model
	Name     string `gorm:"not null;VARCHAR(100);"validate:"required"`
	ParentId int    `gorm:"default:'0';not null;"validate:"number,min=0"`
	Sort     int    `gorm:"default:'0';not null;"validate:"number,min=0"`
	Level    int    `gorm:"-"`
}

func (this *Menu) List() []Menu {
	var data = []Menu{}
	db := libs.DB

	err := db.Order("sort desc").Find(&data).Error
	if err != nil {
		//log.Fatalln(err)
	}
	return data
}

func (this *Menu) MenuInfo(id uint) (Menu, error) {
	var menu Menu
	db := libs.DB

	if db.Where("id = ?", id).First(&menu).RecordNotFound() {
		return Menu{}, errors.New("分类未找到")
	}
	return menu, nil
}

func (this *Menu) MenuMoreInfo(ids string) ([]Menu, error) {
	var data = []Menu{}
	db := libs.DB

	if db.Where("id in (?)", strings.Split(ids, ",")).Find(&data).RecordNotFound() {
		return []Menu{}, errors.New("分类未找到")
	}
	return data, nil
}

func (this *Menu) MenuAdd(postValues map[string][]string) error {
	var menu Menu
	db := libs.DB

	if err := libs.FormDecode(&menu, postValues); err != nil {
		return err
	}
	if err := libs.Validate(menu); err != nil {
		return err
	}
	if !db.Where("name = ? ", menu.Name).First(&Menu{}).RecordNotFound() {
		return errors.New("该名称已经存在")
	}
	if err := db.Create(&menu).Error; err != nil {
		return err
	}
	return nil
}

func (this *Menu) MenuUpdate(postValues map[string][]string) error {
	var menu Menu
	db := libs.DB

	if err := libs.FormDecode(&menu, postValues); err != nil {
		return err
	}
	if err := libs.Validate(menu); err != nil {
		return err
	}

	if !db.Where("name = ? and id != ?", menu.Name, menu.ID).Find(&Menu{}).RecordNotFound() {
		return errors.New("该名称已经存在")
	}
	if db.Where("id = ? ", menu.ID).Find(&Menu{}).RecordNotFound() {
		return errors.New("未查询到分类id")
	}
	if err := db.Save(&menu).Error; err != nil {
		return err
	}
	return nil
}

func (this *Menu) MenuDel(id uint) error {
	var menu Menu
	db := libs.DB

	if !db.Where("parent_id = ?", id).Find(&menu).RecordNotFound() {
		return errors.New("该分类下存在子级分类，请先删除子级分类")
	}
	if err := db.Where("id = ?", id).Delete(&menu).Error; err != nil {
		return err
	}
	return nil
}

func (this *Menu) GetTree(list []Menu, pid int, level int) []Menu {
	for _, v := range list {
		if v.ParentId == pid {
			v.Level = level
			v.Name = strings.Repeat("----", v.Level) + v.Name
			MenuListTree = append(MenuListTree, v)
			/*if len(list) == 0 {
				list = []Category{}
			} else {
				list = append(list[:index], list[index+1:]...)
			}*/
			this.GetTree(list, int(v.Model.ID), v.Level+1)
		}
	}
	return MenuListTree
}
