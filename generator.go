package main

import (
	"fmt"
	"gitee.com/eden-framework/plugins"
	"path"
)

var Plugin GenerationPlugin

type GenerationPlugin struct {
}

func (g *GenerationPlugin) GenerateEntryPoint(opt plugins.Option, cwd string) string {
	globalPkgPath := path.Join(opt.PackageName, "internal/global")
	globalFilePath := path.Join(cwd, "internal/global")
	tpl := fmt.Sprintf(`,
		{{ .UseWithoutAlias "gitee.com/eden-framework/eden-framework/pkg/application" "" }}.WithConfig(&{{ .UseWithoutAlias "%s" "%s" }}.KafkaConfig)`, globalPkgPath, globalFilePath)
	return tpl
}

func (g *GenerationPlugin) GenerateFilePoint(opt plugins.Option, cwd string) []*plugins.FileTemplate {
	file := plugins.NewFileTemplate("global", path.Join(cwd, "internal/global/kafka.go"))
	file.WithBlock(`
var KafkaConfig = struct {
	KafkaProducer *{{ .UseWithoutAlias "gitee.com/eden-framework/plugin-kafka/kafka" "" }}.Producer
	KafkaConsumer *{{ .UseWithoutAlias "gitee.com/eden-framework/plugin-kafka/kafka" "" }}.Consumer
}{
	KafkaProducer: &{{ .UseWithoutAlias "gitee.com/eden-framework/plugin-kafka/kafka" "" }}.Producer{
		Host:  "localhost",
		Port:  9092,
	},
	KafkaConsumer: &{{ .UseWithoutAlias "gitee.com/eden-framework/plugin-kafka/kafka" "" }}.Consumer{
		Brokers: []string{"localhost:9092"},
	},
}
`)

	return []*plugins.FileTemplate{file}
}
