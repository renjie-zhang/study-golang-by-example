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
package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

// mqtt回调方法
var fun mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Topic : %s\n", msg.Topic())
	fmt.Printf("Message :%s\n", msg.Payload())
}

func main() {
	// 创建链接参数
	clientPara := mqtt.NewClientOptions().AddBroker("tcp://192.168.6.130:1883").SetClientID("mqtt-test-1")
	clientPara.SetKeepAlive(2 * time.Second)
	clientPara.SetDefaultPublishHandler(fun)
	clientPara.SetPingTimeout(1 * time.Second)

	// 创建client
	client := mqtt.NewClient(clientPara)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅
	if token := client.Subscribe("renjie/#", 0, fun); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 发送信息到topic中
	client.Publish("renjie/test1", 0, false, "hello mqtt broker,this from renjie/test1")

	client.Publish("renjie/test2", 0, false, "hello mqtt broker,this from renjie/test2")

	time.Sleep(20 * time.Second)

	// 取消订阅
	if token := client.Unsubscribe("renjie/#"); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 断开连接
	client.Disconnect(200)
	time.Sleep(1 * time.Second)

	/*
		Topic : renjie/test1
		Message :hello mqtt broker,this from renjie/test1
		Topic : renjie/test2
		Message :hello mqtt broker,this from renjie/test2
		Topic : renjie/test1
	*/
}
