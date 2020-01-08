package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/urfave/cli.v2"
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


func main() {
	// cli.NewApp().Run(os.Args)
	//clifirst()
	// clisecond()
	// cliflag()
	cliflagdes()
}
