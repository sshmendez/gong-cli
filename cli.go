package main
  
 
import (
	"github.com/sshmendez/gong"
	"flag"
	"fmt"
	"net/http"
)

func main(){
	var filepath string
	flag.StringVar(&filepath, "f","", "Config File Path")
	flag.Parse()

	if filepath == ""{
		fmt.Printf("config path required: try -f <configpath>\n")
		return
	}

	muxconf, err := gong.BuildMuxConfigFromFile(filepath)
	
	if err != nil {
		fmt.Printf("config at %s failed to build\n", filepath)
		return
	}

	mux := gong.BuildMux(muxconf)

	http.ListenAndServe(fmt.Sprintf(":%d", muxconf.Port), mux)
	
	
	
}
