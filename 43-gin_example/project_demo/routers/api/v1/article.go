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
	"log"
	"net/http"
	"project_demo/models"
	"project_demo/pkg/e"
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
			data = models.DeleteArticleById(id)
			code = e.SUCCESS
		}else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}else {
		for _,err := range valid.Errors{
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}
