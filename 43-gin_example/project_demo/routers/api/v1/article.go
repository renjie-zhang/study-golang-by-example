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
package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"project_demo/models"
	"project_demo/pkg/e"
	"project_demo/pkg/logging"
	"project_demo/pkg/setting"
	"project_demo/pkg/util"
)
// GetArticle 根据Id获得文章
func GetArticle(context *gin.Context){
	id := com.StrTo(context.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id,1,"id").Message("ID必须大于0")
	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors(){
		if models.ExistArticleById(id){
			data = models.GetArticle(id)
			code = e.SUCCESS
		}else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}else {
		for _,err := range valid.Errors{
			logging.Error(err.Key, err.Message)
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

// GetArticles 获取多个文章
func GetArticles(context *gin.Context)  {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := context.Query("state");arg != ""{
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	var tagId int = -1
	if arg := context.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors(){
		code = e.SUCCESS
		data["lists"] = models.GetArticles(util.GetPage(context),setting.PageSize,maps)
		data["count"] = models.GetArticleTotal(maps)
	}else {
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func AddArticle(context *gin.Context){
	tagId := com.StrTo(context.Query("tag_id")).MustInt()
	title := context.Query("title")
	desc := context.Query("desc")
	content := context.Query("content")
	createdBy := context.Query("created_by")
	state := com.StrTo(context.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistTagById(tagId) {
			data := make(map[string]interface {})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})
}

func DeleteArticle(context *gin.Context){
	id := com.StrTo(context.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleById(id) {
			models.DeleteArticleById(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}

func EditArticle(context *gin.Context){
	valid := validation.Validation{}

	id := com.StrTo(context.Param("id")).MustInt()
	tagId := com.StrTo(context.Query("tag_id")).MustInt()
	title := context.Query("title")
	desc := context.Query("desc")
	content := context.Query("content")
	modifiedBy := context.Query("modified_by")

	var state int = -1
	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleById(id) {
			if models.ExistTagById(tagId) {
				data := make(map[string]interface {})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != "" {
					data["title"] = title
				}
				if desc != "" {
					data["desc"] = desc
				}
				if content != "" {
					data["content"] = content
				}

				data["modified_by"] = modifiedBy

				models.EditArticle(id, data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}
