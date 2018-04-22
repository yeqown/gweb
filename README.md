# gweb

golang web frame, maybe a frame ~, mainly to less preparing-works while starting a web server

* [Using](#using)
* [Validation](#validation)
* [Decodeer](#decodeer)

## using

```shell
git clone git@github.com:yeqown/gweb.git
cd gweb
gvt restore # if not installed, go get github.com/FiloSottile/gvt

# do some change, then

cd mainC
go run *
```

## validation

> ref to github.com/astaxie/beego/validation

```golang

```

## decoder

> ref to github.com/gorilla/schema

```golang
```

## file-tree

```shell
# tree -L 2
	.
	├── LICENSE
	├── README.md
	├── constant
	│   ├── code.go
	│   └── code_test.go
	├── controllers
	│   ├── default.go
	│   └── hello.go
	├── logger
	│   └── logger.go
	├── logs
	│   ├── app.log
	│   └── request.log
	├── mainC
	│   ├── conf.go
	│   ├── db.go
	│   ├── main.go
	│   └── web.go
	├── models
	│   ├── connect.go
	│   ├── mysql_demo.go
	│   ├── postgres_demo.go
	│   └── redis_demo.go
	├── router
	│   ├── middleware
	│   └── router.go
	├── sample
	│   └── main.go
	├── services
	├── utils
	│   └── uuid.go
	└── vendor
	    ├── github.com
	    └── manifest
```

## manifest

> go deps managed by gvt

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
			"importpath": "github.com/denisenkom/go-mssqldb",
			"repository": "https://github.com/denisenkom/go-mssqldb",
			"vcs": "git",
			"revision": "e32faac87a2220f9342289f2c3b567d1424b8ec5",
			"branch": "master",
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