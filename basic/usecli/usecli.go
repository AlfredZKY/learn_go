package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	// "gopkg.in/urfave/cli.v2"
	"github.com/urfave/cli"
)



func clifirst() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello %q\n", c.Args().Get(0))
		return nil
	}
	app.Run(os.Args)
}

func clisecond() {
	app := cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			fmt.Println("Boom! I say!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func cliflag() {
	app := cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the freeting",
			},
		},
		Action: func(c *cli.Context) error {
			name := "Nefertiti"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "spanish" {
				fmt.Println("hola", name)
			} else {
				fmt.Println("hello", name)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func cliflagdes(){
	var language string
	app := cli.App{
		Flags :[]cli.Flag{
			&cli.StringFlag{
				Name:"lang",
				//Value:"",
				Aliases: []string{"l"},
				Usage:"language for the freeting",
				Destination:&language,
			},
			&cli.BoolFlag{
				Name:  "genesis-miner",
				Usage: "enable genesis mining (DON'T USE ON BOOTSTRAPPED NETWORK)",
				Hidden:	true,
				//Value:true,
			},
		},
		Action:func(c*cli.Context)error{
			name := "someone"
			if c.NArg()>0{
				name = c.Args().Get(0)
			}
			// go run usecli.go  -genesis-miner 12 23 
			if c.Bool("genesis-miner"){
				fmt.Println(c.Bool("genesis-miner"))
				fmt.Println("genesis-miner")
			}
			// fmt.Println(c.String("lang"))
			if language == "spanish"{
				fmt.Println("hola",name)
			}else{
				fmt.Println("Hello",name)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}


var nameclass string
var cliFlag int 

func init(){
	flag.IntVar(&cliFlag,"flagname",1234,"Just for demo")
	flag.StringVar(&nameclass, "nameclass", "everyone", "The greeting object.")

	// flag.CommandLine = flag.NewFlagSet("",flag.ExitOnError)
	flag.CommandLine = flag.NewFlagSet("",flag.PanicOnError)
	flag.CommandLine.Usage=func(){
		fmt.Fprintf(os.Stderr,"Usage of %s:\n","question")
		flag.PrintDefaults()
	}
}

// 命令行的使用
func useFlag(){
	// 参数地址 参数名 参数的默认值 参数的含义(简短的说明)
	// flag.StringVar(&name, "name", "everyone", "The greeting object.") 对应的是地址
	var cliName = flag.String("name","nick","Input Your Name")
	var cliAge = flag.Int("age",28,"Input Your Age")
	var cliGender = flag.String("gender","male","Input Your Gender")

	// 使无参或者自定义参数 对flag.Usage重新赋值,flag.Usage的类型是func(),即一种无参数声明且无结果声明的函数类型
	// flag.Usage = func(){
	// 	fmt.Fprintf(os.Stderr,"Usage of %s:\n","question")
	// 	flag.PrintDefaults()
	// }


	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()

	fmt.Printf("Hello, %s\n",nameclass)
	fmt.Println("name=",*cliName)
	fmt.Println("age=",*cliAge)
	fmt.Println("gender=",*cliGender)
}





// -isbool    (一个 - 符号，布尔类型该写法等同于 -isbool=true)
// -age=x     (一个 - 符号，使用等号)
// -age x     (一个 - 符号，使用空格)
// --age=x    (两个 - 符号，使用等号)
// --age x    (两个 - 符号，使用空格)

func main() {
	// cli.NewApp().Run(os.Args)
	//clifirst()
	// clisecond()
	// cliflag()
	//cliflagdes()
	useFlag()
}
