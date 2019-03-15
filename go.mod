module sas

replace (
	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8
	google.golang.org/genproto v0.0.0-20180831171423-11092d34479b => github.com/google/go-genproto v0.0.0-20190307195333-5fe7a883aa19
)

replace (
	golang.org/x/sys v0.0.0-20180830151530-49385e6e1522 => github.com/golang/sys v0.0.0-20180830151530-49385e6e1522
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 => github.com/golang/sys v0.0.0-20190222072716-a9d3bda3a223
)

replace (
	golang.org/x/net v0.0.0-20180724234803-3673e40ba225 => github.com/golang/net v0.0.0-20180724234803-3673e40ba225
	golang.org/x/net v0.0.0-20180826012351-8a410e7b638d => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/net v0.0.0-20180906233101-161cd47e91fd => github.com/golang/net v0.0.0-20180906233101-161cd47e91fd
	golang.org/x/net v0.0.0-20190213061140-3a22650c66bd => github.com/golang/net v0.0.0-20190213061140-3a22650c66bd
)

replace (
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sync v0.0.0-20181108010431-42b317875d0f => github.com/golang/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sync v0.0.0-20190227155943-e225da77a7e6 => github.com/golang/sync v0.0.0-20190227155943-e225da77a7e6
)

replace golang.org/x/exp v0.0.0-20190121172915-509febef88a4 => github.com/golang/exp v0.0.0-20190121172915-509febef88a4

replace (
	golang.org/x/tools v0.0.0-20190114222345-bf090417da8b => github.com/golang/tools v0.0.0-20190114222345-bf090417da8b
	golang.org/x/tools v0.0.0-20190226205152-f727befe758c => github.com/golang/tools v0.0.0-20190226205152-f727befe758c
)

replace google.golang.org/grpc v1.19.0 => github.com/grpc/grpc-go v1.19.0

replace (
	golang.org/x/lint v0.0.0-20181026193005-c67002cb31c3 => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/lint v0.0.0-20190227174305-5b3e6a55c961 => github.com/golang/lint v0.0.0-20190227174305-5b3e6a55c961
)

replace golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be

replace (
	google.golang.org/appengine v1.1.0 => github.com/golang/appengine v1.1.0
	google.golang.org/appengine v1.4.0 => github.com/golang/appengine v1.4.0
)

replace golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0

replace cloud.google.com/go v0.26.0 => github.com/googleapis/google-cloud-go v0.26.0

require (
	github.com/Unknwon/goconfig v0.0.0-20181105214110-56bd8ab18619
	github.com/braintree/manners v0.0.0-20160418043613-82a8879fc5fd
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/golang/protobuf v1.3.0 // indirect
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go/codec v0.0.0-20190309163734-c4a1c341dc93 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)
