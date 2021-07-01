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
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"project_demo/pkg/logging"
	"project_demo/pkg/setting"
)

var db *gorm.DB


type Model struct{
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init(){
	var (
		err error
		dbType,dbName,user,password,host,tablePrefix string
	)
	sec,err := setting.Cfg.GetSection("database")
	if err != nil{
		logging.Error(2,"Fail to get section database %v",err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db,err = gorm.Open(dbType,fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil{
		logging.Error(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName;
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxIdleConns(20)

}


func CloseDb(){
	defer db.Close()
}