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

type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"`
	Tag Tag `json:"tag"`

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int `json:"state"`
}
// ExistArticleById 根据ID查找存在的文章
func ExistArticleById(id int) bool{
	var article Article
	db.Select("id").Where("id = ?",id).Find(&article)
	if article.ID > 0{
		return true
	}
	return false
}
// GetArticleTotal 获取所有的文章总数
func GetArticleTotal(maps interface{})(count int){
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

// GetArticles 分页获取文章
func GetArticles(pageNum int,pageSize int,maps interface{}) (articles []Article){
	db.Preloads("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}
// GetArticle 根据ID获取文章
func GetArticle(id int) (article Article){
	db.Where("id = ? ",id).First(article)
	db.Model(&article).Related(&article.Tag)
	return
}
// EditArticle 编辑文章
func EditArticle(id int,data interface{}) bool{
	db.Model(&Article{}).Where("id = ?",id).Update(data)
	return true
}
// AddArticle 添加文章
func AddArticle(data map[string]interface{}) bool{
	db.Create(&Article {
		TagID : data["tag_id"].(int),
		Title : data["title"].(string),
		Desc : data["desc"].(string),
		Content : data["content"].(string),
		CreatedBy : data["created_by"].(string),
		State : data["state"].(int),
	})

	return true
}
// DeleteArticleById 根据Id删除文章
func DeleteArticleById(id int) bool{
	db.Where("id = ?",id).Delete(Article{})
	return true
}


func (article *Article) BeforeCreate(scope *gorm.Scope) error{
	return scope.SetColumn("CreatedOn",time.Now().Unix())
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error{
	return scope.SetColumn("ModifiedOn", time.Now().Unix())
}
