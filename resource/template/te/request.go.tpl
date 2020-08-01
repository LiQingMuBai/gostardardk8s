package request

import "gostardardk8s/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}