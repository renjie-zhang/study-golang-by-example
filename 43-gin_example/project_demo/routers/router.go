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
package routers

import (
	"github.com/gin-gonic/gin"
	"project_demo/pkg/setting"
	v1 "project_demo/routers/api/v1"
)

func InitRouter() *gin.Engine{
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags",v1.GetTags)
		apiv1.POST("/tags",v1.AddTag)
		apiv1.PUT("/tags/:id",v1.EditTag)
		apiv1.DELETE("/tags/:id",v1.DeleteTag)
		apiv1.GET("/articles",v1.GetArticles)
		apiv1.POST("/articles",v1.AddArticle)
		apiv1.PUT("/articles/:id",v1.EditArticle)
		apiv1.DELETE("/articles/:id",v1.DeleteArticle)
	}
	return r
}