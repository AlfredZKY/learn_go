package readconfig

import (
	"fmt"
	"io/ioutil"
	"github.com/BurntSushi/toml"
	"os"
)

// LocalIP global IP
var LocalIP string

func init(){
	p,err := ReadConf("config.toml")
	if err != nil{
		fmt.Printf("%v",err)
	}
	LocalIP = p.MonitorUnit.LocalIP
}

// Person The read must config file 
type Person struct{
	MonitorUnit * MonitorUnit
}

// Friends define some files
type Friends struct{
	Age int
	Name string
}

// MonitorUnit get Local IP from config file
type MonitorUnit struct{
	LocalIP string
}

// ReadConf read the config file
func ReadConf(fname string)(p*Person,err error){
	var (
		fp *os.File
		fcontent []byte
	)
	p = new(Person)
	if fp,err = os.Open(fname);err!=nil{
		fmt.Println("open error ",err)
		return
	}

	if fcontent,err = ioutil.ReadAll(fp);err != nil{
		fmt.Println("ReadAll error ", err)
		return
	}


	if err = toml.Unmarshal(fcontent,p);err!=nil{
		fmt.Println("toml.Unmarshal error ",err)
		return 
	}
	return
}