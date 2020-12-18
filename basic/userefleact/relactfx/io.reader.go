package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"go.uber.org/fx"
)
//nolint:golint
func fxUse() {
	var reader io.Reader

	app := fx.New(
		// io.reader的应用
		// 提供构造函数
		fx.Provide(func() io.Reader {
			return strings.NewReader("hello world")
		}),
		fx.Populate(&reader), // 通过依赖注入完成变量与具体类的映射
	)
	app.Start(context.Background())
	defer app.Stop(context.Background())

	fmt.Printf("%T\n", reader)
	// 使用
	// reader变量已与fx.Provide注入的实现类关联了
	bs, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Panic("read occur error, ", err)
	}
	fmt.Printf("the result is '%s' \n", string(bs))
}

func fxComplexUse() {
	type t3 struct {
		Name string
	}

	type t4 struct {
		Age int
	}

	var (
		v1 *t3
		v2 *t4
	)

	app := fx.New(
		fx.Provide(func() *t3 {
			return &t3{"hello everybody!!!"}
		}),
		fx.Provide(func() *t4 {
			return &t4{2019}
		}),
		fx.Populate(&v1),
		fx.Populate(&v2),
	)

	app.Start(context.Background())
	defer app.Stop(context.Background())
	fmt.Printf("The result is %v, %v\n", v1.Name, v2.Age)
}

func fxComplexUseGroup() {
	type t3 struct {
		Name string
	}

	// 使用group标签
	type result struct {
		fx.Out
		V1 *t3 `group:"g"`
		V2 *t3 `group:"g"`
	}

	targets := struct {
		fx.In
		Group []*t3 `group:"g"`
	}{}

	app := fx.New(
		fx.Provide(func() result {
			return result{
				V1: &t3{"hello-000"},
				V2: &t3{"world-www"},
			}
		}),
		fx.Populate(&targets),
	)
	
	app.Start(context.Background())
	defer app.Stop(context.Background())
	for _, t := range targets.Group {
		fmt.Printf("the result is %v\n", t.Name)
	}

}

func main() {
	fxUse()
	fxComplexUse()
	fxComplexUseGroup()
}
