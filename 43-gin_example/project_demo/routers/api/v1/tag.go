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
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"project_demo/pkg/e"
)

// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
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

	

}
// AddTag add tag
func AddTag(context *gin.Context){

}
// EditTag edit tag
func EditTag(context *gin.Context){

}
// DeleteTag delete Tag
func DeleteTag(context *gin.Context){

}
