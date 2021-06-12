module Study-Golang-By-Example

go 1.14

require (
	github.com/eclipse/paho.mqtt.golang v1.3.0
	github.com/gin-gonic/gin v1.7.2
	github.com/golang/protobuf v1.4.3
	github.com/smartystreets/goconvey v1.6.4
	google.golang.org/protobuf v1.25.0

)

replace (
	github.com/renjie-zhang/Study-Golang-By-Example/43-gin_example/project_demo/conf => ../Study-Golang-By-Example/43-gin_example/project_demo/pkg/conf
	github.com/renjie-zhang/Study-Golang-By-Example/43-gin_example/project_demo/middleware => ../Study-Golang-By-Example/43-gin_example/project_demo/middleware
	github.com/renjie-zhang/Study-Golang-By-Example/43-gin_example/project_demo/models => ../Study-Golang-By-Example/43-gin_example/project_demo/models
	github.com/renjie-zhang/Study-Golang-By-Example/43-gin_example/project_demo/pkg/setting => ../Study-Golang-By-Example/43-gin_example/project_demo/pkg/setting
	github.com/renjie-zhang/Study-Golang-By-Example/43-gin_example/project_demo/routers => ../Study-Golang-By-Example/43-gin_example/project_demo/routers
)
