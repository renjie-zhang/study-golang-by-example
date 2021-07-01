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
package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"project_demo/models"
	"project_demo/pkg/e"
	"project_demo/pkg/util"
)

type auth struct{
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

// GetAuth 获得auth
func GetAuth(context *gin.Context){
	username := context.Query("username")
	password := context.Query("password")

	valid := validation.Validation{}
	t := auth{
		Username: username,
		Password: password,
	}
	ok,_ := valid.Valid(&t)
	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok{
		isExist := models.CheckAuth(username,password)
		if isExist{
			token,err := util.GenerateToken(username,password)
			if err != nil{
				code = e.ERROR_AUTH_TOKEN
			}else {
				data["token"] = token
				code = e.SUCCESS
			}
		}else {
			code = e.ERROR_AUTH
		}

	}else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	context.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})
}