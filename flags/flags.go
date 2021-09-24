package flags
import (
  "flag"
)

var MoviePath string
var SubTitlePath string
var Suffix string
var Name string
var Destination string
var C bool
var Season string


func init(){
 flag.StringVar(&MoviePath,"moviepath","","path to movie files")
 flag.StringVar(&SubTitlePath,"subtitlepath","","path to subtitle files")
 flag.StringVar(&Suffix,"suffix","S0*E0*","example Seoson01Episode01 = Season0*Episode0*")
 flag.StringVar(&Name,"name","","movie name")
 flag.StringVar(&Destination,"d","","destination folder to put both movie files and subtitle")
 flag.StringVar(&Season,"s","1","movie season")
 flag.BoolVar(&C,"c",false,"move files or copy")
 flag.Parse()
}
