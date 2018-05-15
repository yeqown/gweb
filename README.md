# Gweb
	golang web frame, maybe a frame ~, mainly to less preparing-works while starting a web server

doc ref to: [https://godoc.org/github.com/yeqown/gweb](https://godoc.org/github.com/yeqown/gweb)
 
*[Using](#using)*
*[Validation](#validation)*
*[Decoder](#decoder)*

## Samples

Link to [gweb-usage-samples](https://github.com/yeqown/gweb-usage-smaples)

## Validation

> ref to github.com/astaxie/beego/validation

```golang
// https://beego.me/docs/mvc/controller/validation.md

import (
    "log"
    "strings"

    "github.com/astaxie/beego/validation"
)

// 验证函数写在 "valid" tag 的标签里
// 各个函数之间用分号 ";" 分隔，分号后面可以有空格
// 参数用括号 "()" 括起来，多个参数之间用逗号 "," 分开，逗号后面可以有空格
// 正则函数(Match)的匹配模式用两斜杠 "/" 括起来
// 各个函数的结果的 key 值为字段名.验证函数名
type user struct {
    Id     int
    Name   string `valid:"Required;Match(/^Bee.*/)"` // Name 不能为空并且以 Bee 开头
    Age    int    `valid:"Range(1, 140)"` // 1 <= Age <= 140，超出此范围即为不合法
    Email  string `valid:"Email; MaxSize(100)"` // Email 字段需要符合邮箱格式，并且最大长度不能大于 100 个字符
    Mobile string `valid:"Mobile"` // Mobile 必须为正确的手机号
    IP     string `valid:"IP"` // IP 必须为一个正确的 IPv4 地址
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *user) Valid(v *validation.Validation) {
    if strings.Index(u.Name, "admin") != -1 {
        // 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
        v.SetError("Name", "名称里不能含有 admin")
    }
}

func main() {
    valid := validation.Validation{}
    u := user{Name: "Beego", Age: 2, Email: "dev@beego.me"}
    b, err := valid.Valid(&u)
    if err != nil {
        // handle error
    }
    if !b {
        // validation does not pass
        // blabla...
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }
}
```

StructTag 可用的验证函数：

* `Required` 不为空，即各个类型要求不为其零值
* `Min(min int)` 最小值，有效类型：`int`，其他类型都将不能通过验证
* `Max(max int)` 最大值，有效类型：`int`，其他类型都将不能通过验证
* `Range(min, max int)` 数值的范围，有效类型：`int`，他类型都将不能通过验证
* `MinSize(min int)` 最小长度，有效类型：`string slice`，其他类型都将不能通过验证
* `MaxSize(max int)` 最大长度，有效类型：`string slice`，其他类型都将不能通过验证
* `Length(length int)` 指定长度，有效类型：`string slice`，其他类型都将不能通过验证
* `Alpha` alpha字符，有效类型：`string`，其他类型都将不能通过验证
* `Numeric` 数字，有效类型：`string`，其他类型都将不能通过验证
* `AlphaNumeric` alpha 字符或数字，有效类型：`string`，其他类型都将不能通过验证
* `Match(pattern string)` 正则匹配，有效类型：`string`，其他类型都将被转成字符串再匹配(fmt.Sprintf("%v", obj).Match)
* `AlphaDash` alpha 字符或数字或横杠 `-_`，有效类型：`string`，其他类型都将不能通过验证
* `Email` 邮箱格式，有效类型：`string`，其他类型都将不能通过验证
* `IP` IP 格式，目前只支持 IPv4 格式验证，有效类型：`string`，其他类型都将不能通过验证
* `Base64` base64 编码，有效类型：`string`，其他类型都将不能通过验证
* `Mobile` 手机号，有效类型：`string`，其他类型都将不能通过验证
* `Tel` 固定电话号，有效类型：`string`，其他类型都将不能通过验证
* `Phone` 手机号或固定电话号，有效类型：`string`，其他类型都将不能通过验证
* `ZipCode` 邮政编码，有效类型：`string`，其他类型都将不能通过验证


## Decoder

> ref to github.com/gorilla/schema

```golang
// http://www.gorillatoolkit.org/pkg/schema
// website demo usage:

type Person struct {
    Name  string `schema:"name"`  // custom name
    Phone string `schema:"phone"` // custom name
    Admin bool   `schema:"-"`     // this field is never set
}
```

## File tree

```shell
# tree -L 2
.
├── LICENSE
├── README.md
├── conf.go
├── http_handler.go
├── logger
│   └── logger.go
├── middleware
│   ├── param.go
│   ├── readme.md
│   ├── reqlog.go
│   ├── response.go
│   └── safe.go
├── router.go
├── utils
│   ├── code.go
│   ├── code_test.go
│   ├── string.go
│   ├── string_test.go
│   ├── uuid.go
│   └── uuid_test.go
├── vendor
│   ├── github.com
│   └── manifest
└── web.go
```

## manifest

go deps managed by gvt

```json
{
	"version": 0,
	"dependencies": [
		{
			"importpath": "github.com/astaxie/beego/validation",
			"repository": "https://github.com/astaxie/beego",
			"vcs": "git",
			"revision": "f16688817aa428d10361394015b40d096b680542",
			"branch": "master",
			"path": "/validation",
			"notests": true
		},
		{
			"importpath": "github.com/gorilla/schema",
			"repository": "https://github.com/gorilla/schema",
			"vcs": "git",
			"revision": "d0e4c24cff97ae983e9847e0ed5a02dc10013d41",
			"branch": "master",
			"notests": true
		},
		{
			"importpath": "github.com/jinzhu/gorm",
			"repository": "https://github.com/jinzhu/gorm",
			"vcs": "git",
			"revision": "6842b49a1ad0feb6b93be830fe63a682cf853ada",
			"branch": "master",
			"notests": true
		},
		{
			"importpath": "github.com/yeqown/log",
			"repository": "https://github.com/yeqown/log",
			"vcs": "git",
			"revision": "7063b5e0e64b332ff54845f426b54629f13110fd",
			"branch": "master",
			"notests": true
		}
	]
}
```