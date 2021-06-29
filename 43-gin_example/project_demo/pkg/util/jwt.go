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
package util

import (
	jwt "github.com/dgrijalva/jwt-go"
	"project_demo/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.Claims
}

func GenerateToken(username,password string) (string,error){
	now := time.Now()
	expireTime := now.Add(2*time.Hour)
	claims := Claims{
		Username: username,
		Password: password,
		Claims:   jwt.StandardClaims{
			ExpiresAt : expireTime.Unix(),
			Issuer : "gin-blog",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256,claims)
	return tokenClaims.SignedString(jwtSecret)
}

func ParseToken(token string) (*Claims,error){
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}