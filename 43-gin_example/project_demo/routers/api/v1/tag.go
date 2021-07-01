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
	"project_demo/pkg/setting"
	"project_demo/pkg/util"
)

// @Summary Get multiple article tags
// @Param name query string false "Name"
// @Param state query int false "State"
// @Router /api/v1/tags [get]
func GetTags(context *gin.Context)  {
	name := context.Param("name")
	maps := make(map[string] interface{})
	data := make(map[string]interface{})

	if name != ""{
		maps["name"] = name
	}
	var state int = -1
	if arg := context.Query("state");arg != ""{
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(context),setting.PageSize,maps)
	data["total"] = models.GetTagTotal(maps)

	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":maps["data"],
	})

}
// @Summary 新增文章标签
// @Param name query string true "Name"
// @Param state query int false "State"
// @Param created_by query int false "CreatedBy"
// @Router /api/v1/tags [post]
func AddTag(context *gin.Context){
	name := context.Query("name")
	state := com.StrTo(context.DefaultQuery("state", "0")).MustInt()
	createdBy := context.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors(){
		if ! models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data": make(map[string]string),
	})

}
// @Summary 修改文章标签
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Router /api/v1/tags/{id} [put]
func EditTag(context *gin.Context){
	id := com.StrTo(context.Param("id")).MustInt()
	name := context.Query("name")
	modifiedBy := context.Query("modified_by")

	valid := validation.Validation{}
	var state int  = -1
	if arg := context.Query("state");arg != ""{
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors(){
		code = e.SUCCESS
		if models.ExistArticleById(id){
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != ""{
				data["name"] = name
			}
			if state != -1{
				data["state"] = state
			}
			models.EditTag(id,data)
		}
	}
	context.JSON(code,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data": make(map[string]string),
	})
}
// DeleteTag delete Tag
func DeleteTag(context *gin.Context){
	id := com.StrTo(context.Query("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	code := e.INVALID_PARAMS

	if !valid.HasErrors(){
		code = e.SUCCESS
		if models.ExistTagById(id){
			models.DeleteTagById(id)
		}else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}
