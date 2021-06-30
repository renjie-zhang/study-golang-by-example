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
package jwt

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project_demo/pkg/e"
	"project_demo/pkg/util"
	"time"
)

func JWT() gin.HandlerFunc{
	return func(context *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := context.Query("token")
		if token == ""{
			code = e.INVALID_PARAMS
		}else {
			claims,err := util.ParseToken(token)
			if err != nil{
				code  = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}else if time.Now().Unix() > claims.ExpiresAt{
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
			if code != e.SUCCESS{
				context.JSON(http.StatusUnauthorized,gin.H{
					"code":code,
					"msg":e.GetMsg(code),
					"data":data,
				})
				context.Abort()
				return
			}
		}
		context.Next()
	}
}
