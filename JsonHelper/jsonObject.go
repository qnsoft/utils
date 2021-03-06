package JsonHelper

import (
	"bytes"
	"reflect"
	"strings"
	"sort"
)

type JsonObject struct {
	Attributes map[string]*JsonObject
	Value      interface{}
	VType      reflect.Kind
}

// 获取属性对象
func (j *JsonObject) GetJsonObject(key string) *JsonObject {
	if j == nil {
		return nil
	}
	return j.Attributes[key]
}

// 获取string值
func (j *JsonObject) GetString() string {
	if j == nil && j.VType != reflect.String {
		return ""
	}
	return j.Value.(string)
}

// 获取数组
func (j *JsonObject) GetJsonArray() []*JsonObject {
	if j == nil && j.VType != reflect.Slice {
		return nil
	}
	return j.Value.([]*JsonObject)
}

// 获取bool值
func (j *JsonObject) GetBool() bool {
	if j == nil && j.VType != reflect.Bool {
		return false
	}
	return j.Value.(bool)
}

// 获取浮点值
func (j *JsonObject) GetFloat64() float64 {
	if j == nil {
		return 0
	}
	switch j.VType {
	case reflect.Int64:
		return float64(j.Value.(int64))
	case reflect.Float64:
		return j.Value.(float64)
	}
	return 0
}

// 获取整型值
func (j *JsonObject) GetInt64() int64 {
	if j == nil {
		return 0
	}
	switch j.VType {
	case reflect.Int64:
		return j.Value.(int64)
	case reflect.Float64:
		return int64(j.Value.(float64))
	}
	return 0
}

// 获取整型值
func (j *JsonObject) GetInterface() interface{} {
	if j == nil && j.VType != reflect.Struct {
		return nil
	}
	return j.Value
}

// 获取整型值
func (j *JsonObject) GetDefaultCoding() string {
	if j == nil {
		return ""
	}
	buffer := bytes.Buffer{}
	buildCoding("AutoGenerated", j, &buffer)
	return buffer.String()
}
// 获取整型值
func (j *JsonObject) GetCoding(strcutName string) string {
	if j == nil {
		return ""
	}
	buffer := bytes.Buffer{}
	buildCoding(strcutName, j, &buffer)
	return buffer.String()
}

func buildCoding(key string, root *JsonObject, buffer *bytes.Buffer) {
	structsExist := map[string]bool{}
	switch root.VType {
	case reflect.Struct:
		if len(root.Attributes) > 0 {
			for subKey, subObject := range root.Attributes {
				subKey = key + "_" + subKey
				buildCoding(subKey, subObject, buffer)
			}
		}
		if !structsExist[key] {
			getObjectCoding(key, root, buffer)
			structsExist[key]=true
		}
	case reflect.Slice:
		object := root.Value.([]*JsonObject)[0]
		buildCoding(key, object, buffer)
	default:
		//fmt.Println(root.VType,root.Value)
	}

}
type attribute struct {
	name string   `json:"name"`
	jsonName string
	attrType string
	origType reflect.Kind
}

type attributeSlice []*attribute

func (c attributeSlice) Len() int {
	return len(c)
}
func (c attributeSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c attributeSlice) Less(i, j int) bool {
	return c[i].name < c[j].name
}

func getObjectCoding(structName string, object *JsonObject, buffer *bytes.Buffer) {

	buffer.WriteString("\ntype "+camelString(structName)+" struct{\n")
	attrs := attributeSlice{}
	nameMaxLength := 0
	attrTypeMaxLength := 0
	for key,attrObject:=range object.Attributes{
		attrName := camelString(key)
		attr := &attribute{
			name:attrName,
			jsonName:key,
			origType:attrObject.VType,
			attrType: func()string {
				switch attrObject.VType {
				case reflect.Struct:
					if len(attrObject.Attributes) >0 {
						return "*"+camelString(getSubAtrrName(structName,key))
					}else {
						return reflect.Interface.String()+"{}"
					}
				case reflect.Slice:
					subObject := attrObject.Value.([]*JsonObject)[0]
					switch subObject.VType {
					case reflect.Struct:
						if len(subObject.Attributes) >0 {
							return "[]*"+camelString(getSubAtrrName(structName,key))
						}else {
							return "[]"+reflect.Interface.String()+"{}"
						}
					case reflect.Slice:
						subObject := attrObject.Value.([]*JsonObject)[0]
						switch subObject.VType {
						case reflect.Struct:
							if len(subObject.Attributes) >0 {
								return "[][]*"+camelString(getSubAtrrName(structName,key))
							}else {
								return "[][]"+reflect.Interface.String()+"{}"
							}
						case reflect.Slice:
							return "[][][]*"
						default:
							return "[][]"+subObject.VType.String()
						}
					default:
						return "[]"+subObject.VType.String()
					}

				case reflect.Interface:
					return reflect.Interface.String()+"{}"
				default:
					return attrObject.VType.String()
				}
			}(),
		}
		if len(attr.name) >nameMaxLength{
			nameMaxLength = len(attr.name)
		}
		if len(attr.attrType) > attrTypeMaxLength{
			attrTypeMaxLength = len(attr.attrType)
		}
		attrs = append(attrs,attr)
	}
	sort.Sort(attrs)
	for _,attr :=range attrs{
		buffer.WriteString("\t"+attr.name+getblanks(attr.name,nameMaxLength)+" "+
			attr.attrType+getblanks(attr.attrType,attrTypeMaxLength)+" "+
				"`json:\""+attr.jsonName+"\"`"+" \n")
	}
	buffer.WriteString("}\n")
	/*

func (this *NihaoComponents) GetCityName() string {
	if this == nil {
		return ""
	}
	return this.CityName
}
	*/
	for _,attr :=range attrs{
		// get方法
		buffer.WriteString("func (this *"+camelString(structName)+") Get"+attr.name+"() "+attr.attrType+" {\n")
		buffer.WriteString("\tif this == nil {\n")
		switch attr.origType {
		case reflect.Int32,reflect.Int64,reflect.Float64:
			buffer.WriteString("\t\treturn 0\n")
		case reflect.Bool:
			buffer.WriteString("\t\treturn false\n")
		case reflect.String:
			buffer.WriteString("\t\treturn \"\"\n")
		default:
			buffer.WriteString("\t\treturn nil\n")
		}
		buffer.WriteString("\t}\n")
		buffer.WriteString("\treturn this."+attr.name+"\n")
		buffer.WriteString("}\n")
		// set方法
		paramName:=getParamName(attr.name)
		buffer.WriteString("func (this *"+camelString(structName)+") Set"+attr.name+"("+paramName+" "+attr.attrType+") {\n")
		buffer.WriteString("\tif this == nil {\n")
		buffer.WriteString("\t\treturn\n")
		buffer.WriteString("\t}\n")
		buffer.WriteString("\tthis."+attr.name+" = "+paramName+"\n")
		buffer.WriteString("}\n")
	}
	/*


func (this *NihaoComponents) SetCityName(cityName string) {
	if this == nil {
		return
	}
	this.CityName = cityName
}
	 */
}
func getParamName(name string) string {
	if len(name)>1{
		return strings.ToLower(string(name[0]))+name[1:]
	}
	return name
}

func getSubAtrrName(father ,sub string) string {
	return father+"_"+sub
}

func getblanks(key string,length int) string {
	length = length-len(key)
	if length<= 0{
		return ""
	}
	buffer := bytes.Buffer{}
	for i:=0 ;i<length;i++{
		buffer.WriteString(" ")
	}
	return buffer.String()
}
// snake string, XxYy to xx_yy , XxYY to xx_yy
func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func camelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

