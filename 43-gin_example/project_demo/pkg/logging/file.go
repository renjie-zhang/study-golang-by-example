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
package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var(
	LogSavePath = "runtime/logs"
	LogSaveName = "log"
	LogFileExt = "log"
	TimeFormat = "20060102"
)

func getLogFilePath() string{
	return fmt.Sprintf("%s",LogSavePath)
}

func getLogFileFullPath() string{
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File{
	_,err := os.Stat(filePath)
	switch  {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission:%v",err)
	}
	handle,err := os.OpenFile(filePath,os.O_APPEND| os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil{
		log.Fatalf("Fail to OpenFile :%v",err)

	}
	return handle
}

func mkDir(){
	dir,_ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(),os.ModePerm)
	if err != nil{
		panic(err)
	}
}