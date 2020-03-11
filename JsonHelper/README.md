# go-jsonObject
json字符串校验、反序列化JSON字符串成Golang对象、反序列化JSON字符串成Golang代码（JSON string validation, deserializing JSON strings into Golang objects, deserializing JSON strings into Golang code）

## 安装
1. 使用go get
```bash
go get github.com/qnsoft/utils
```
2. 导入包
```go
import (
	"github.com/qnsoft/utils/JsonHelper"
)
```

## 主要特性

* json字符串校验（JSON string validation）
* 反序列化JSON字符串成Golang对象（deserializing JSON strings into Golang objects）
* 反序列化JSON字符串成Golang结构体、get、set方法（deserializing JSON strings into Golang code）
## 性能

[测试文件](./test/json_test.go)

## 示例
#### 一、json字符串有效性校验
```go
func Test_Validator(t *testing.T) {
	data := `{"id":524042,"name":"酷旅-mob-otv-2","male":true,"other":null}`
	err := JsonHelper.CheckValid([]byte(data))
	if err != nil {
		fmt.Println(err)
	}
}
```
#### 二、json字符串转Golang对象
```go
func Test_json(t *testing.T) {
	data := `{"id":524042,"name":"酷旅-mob-otv-2","male":true,"other":null}`
	object, err := JsonHelper.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		jsonBytes, _ := json.Marshal(object)
		fmt.Println(string(jsonBytes))
	}
}
```
#### 三、json字符串内嵌数组转Golang对象
```go
func Test_json_array_parse(t *testing.T) {
	data := `{
			  "id": [
				-524042.5,
				2312314444
			  ],
			  "name": "酷旅-mob-otv-2",
			  "male": true,
			  "other": null
			}`
	object, err := JsonHelper.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	idArray:=object.GetJsonObject("id").GetJsonArray()
	for i,v:=range idArray{
		fmt.Println(i,v.GetFloat64())
		fmt.Println(i,v.GetInt64())
	}
	name:=object.GetJsonObject("name").GetString()
	fmt.Println(name)
	male:=object.GetJsonObject("male").GetBool()
	fmt.Println(male)
	other:=object.GetJsonObject("other").GetInterface()
	fmt.Println(other)
}
```
#### 四、json字符串转golang代码
```go
func Test_json_array111(t *testing.T) {
	data := `{
	  "id": [
		-524042.5,
		231231444445555555
	  ],
	  "name": "酷旅-mob-otv-2",
	  "male": true,
	  "other": null
	}`
	object, err := JsonHelper.FromBytes([]byte(data))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(object.GetCoding("travel"))
	}
}
```
结果如下：
```go
type Travel struct{
	Id    []float64   `json:"id"` 
	Male  bool        `json:"male"` 
	Name  string      `json:"name"` 
	Other interface{} `json:"other"` 
}
func (this *Travel) GetId() []float64 {
	if this == nil {
		return nil
	}
	return this.Id
}
func (this *Travel) SetId(id []float64) {
	if this == nil {
		return
	}
	this.Id = id
}
func (this *Travel) GetMale() bool {
	if this == nil {
		return false
	}
	return this.Male
}
func (this *Travel) SetMale(male bool) {
	if this == nil {
		return
	}
	this.Male = male
}
func (this *Travel) GetName() string {
	if this == nil {
		return ""
	}
	return this.Name
}
func (this *Travel) SetName(name string) {
	if this == nil {
		return
	}
	this.Name = name
}
func (this *Travel) GetOther() interface{} {
	if this == nil {
		return nil
	}
	return this.Other
}
func (this *Travel) SetOther(other interface{}) {
	if this == nil {
		return
	}
	this.Other = other
}
```
