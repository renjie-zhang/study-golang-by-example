/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}
// GetTags 获取Tags
func GetTags(pageNum int,pageSize int,maps interface{}) (tags []Tag){
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}
// GetTagTotal 获取Tag数量
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// ExistTagByName 通过Name查找Tag
func ExistTagByName(name string) bool{
	var tag Tag
	db.Model(&Tag{}).Where("name = ?" ,name).First(&tag)
	if tag.ID > 0{
		return true
	}
	return false
}
// AddTag 添加Tag
func AddTag(name string,state int,createBy string) bool{
	tag := &Tag{
		Model:      Model{},
		Name:       name,
		CreatedBy:  createBy,
		ModifiedBy: "",
		State:      state,
	}
	db.Create(tag)
	return true
}
// EditTag 编辑Tag
func EditTag(id int,data interface{}) bool{
	db.Model(&Tag{}).Where("id = ?",id).Update(data)
	return true
}
// DeleteTagById 通过Id进行删除
func DeleteTagById(id int) bool{
	db.Model(&Tag{}).Where("id = ?",id).Delete(&Tag{})
	return true
}

// ExistTagById 通过Id查找是否存在
func ExistTagById(id int) bool{
	var tag Tag
	db.Model(&Tag{}).Where("id=?",id).First(&tag)
	if tag.ID >0 {
		return true
	}
	return false
}


func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}