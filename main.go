package main 

import ("fmt"
	"net/http"
	"html/template"
	"io/ioutil"
	"strings"
	"os"
)
 type handleE struct {
	Status string
	Code int

}
var eachWord string 
var tpl *template.Template 
func main(){

	tpl, _ = template.ParseGlob("templates/*.html")
	http.HandleFunc("/" ,indexHandler)
	
	http.HandleFunc("/ascii-art" ,asciiHandler)
	//http.HandleFunc("/error" ,ErrorHandler)
	
	fmt.Println("serving on port 8080.....")
	http.ListenAndServe(":8080" , nil)
	 
}



func indexHandler(w http.ResponseWriter, r *http.Request) {
	res := handleE{}
	if r.Method == http.MethodGet {
		if r.URL.Path != "/" {
			res.Code = 404
			res.Status = "not found" 
			ErrorHandler(w,r,&res)
			return 	
		}
		tpl.ExecuteTemplate(w,"index.html" , nil)

	}else {
		res.Code = 400
		res.Status = "bad request"
		ErrorHandler(w,r,&res)
			return 	

	}
	
	
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, res *handleE) {
	w.WriteHeader(res.Code)
	err := tpl.ExecuteTemplate(w, "error.html", res)
	if err != nil {
		fmt.Println("Error with error.html")
		os.Exit(2)
	}
}


func asciiHandler(w http.ResponseWriter, r *http.Request) {
	res := handleE{}
	if r.Method == http.MethodPost{
		if r.URL.Path != "/ascii-art" { 
			res.Code = 404
			res.Status = "not found"
			ErrorHandler(w,r,&res)
			return 
		}
	var text string 
	text = r.FormValue("text")
	for _, g:= range text {
		if !(g >= '\a' && g <= '~') {
			res.Code = 400 
		res.Status = "bad request"
		ErrorHandler(w,r,&res)
			return 
		}
	}
	var realWord []string 
	realWord = strings.Split(text,"\r\n")
	tpl.ExecuteTemplate(w,"index.html" , text)
	fmt.Fprintf(w, "<pre>")
	for i:=0 ; i < len(realWord) ; i++ {
		temp ,err:= PrintAscii(w,realWord[i],r)
		if err != "" {
			res := handleE{}
			res.Code = 500 
			res.Status = err 
			ErrorHandler(w,r,&res)
			return 
		}
		fmt.Fprintf(w,temp )
		fmt.Fprintf(w, "\n")
		
	}
	
	fmt.Fprintf(w, "</pre>")
}
}

	
func PrintAscii(w http.ResponseWriter , word string , r *http.Request) (string ,string ){
		res := handleE{}
	var bannerFile string
	var  banner string 
	var eachLine []string
	var eachlineOfTheWord string 
	banner = r.FormValue("typeBanner")
	if banner == "standard" {
		bannerFile = "banners/standard.txt"
	}
	if banner == "shadow" {
		bannerFile = "banners/shadow.txt"
	}
	if banner == "thinkertoy" {
		bannerFile = "banners/thinkertoy.txt"
	}
		read , err := ioutil.ReadFile(bannerFile)
	if err != nil {
		res.Code = 404
			res.Status = "not found" 
			ErrorHandler(w,r,&res)
			
		return "" , ""
	// 	}
		fmt.Println("I can't read the file")
		//http.HandleFunc("/error" ,errorHandler)
	}
	// here first we read the file and split it by the new lines and then we put it in variable called ascii 
	ascii:= strings.Split(string(read),"\n")
	// here we put each line in the ascii variable in array called eachline 
	for i:= 0 ; i < len(ascii) ; i++ {
		eachLine = append(eachLine, ascii[i])
	}
	// we will use this array to put the ascii characters 
	var asciiCH []string 
	// here we are puting the ascii characters in the array called asciich
	for i:= ' ' ; i <= '~' ; i++ {
		asciiCH = append(asciiCH, string(i))
	}
	var startLine []int  
	counter:= 2
	for i:= 0 ; i <= 94 ; i++ {
		startLine = append(startLine, counter)
		//we add nine beacause the next ascii char starts after 9 lines 
		counter += 9 
	}

	// we will use this variable to do the incrementation we need to get the next line of each character 
	count:=0
	// this variable to put the content horizantally on each line  
	st:= ""
	// this variable to put the line we need 
	line:=0
	// this loop for adding each line and the going to the next line until we raech eight lines 
	for i:= 1 ; i <= 8 ; i++ {
		// here when we are in the first line we will do a loop on each character of the argument to se if it match 
		for t:= 0 ; t < len(word) ; t++ {
			
			for k:= 0 ; k<len(asciiCH) ; k++{
				
				if string(word[t]) == asciiCH[k] {
					// we will put the line of the charcter -1 beacause the array of eachline starts after one line 
					line= startLine[k]+count-1 
				}
			}
			// here we add the ascii art char horizantally of each character 
			st = st + eachLine[line]
		}
		
		eachlineOfTheWord += "\n"
		eachlineOfTheWord += st 
		st= ""
		count++	
	}
	return eachlineOfTheWord , ""
	}

