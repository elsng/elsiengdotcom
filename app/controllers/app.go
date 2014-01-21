package controllers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/robfig/revel"
)


type Elsie struct {
    Elsieng Elsieng_info
}

type Elsieng_info struct {
    Profile  Profile_info
    Project []Project_info
}

type Project_info struct {
    Id	int
    Name	string
    Gist	string
    Catchphrase	string
    Picture	string
    Partner	string
    Button	string
    Brand	string
    Link	string
    Gettext	string
    About	string
    Role	string
    Section []Section_info
}

type Profile_info struct {
    Name    string
    Gist    string
    Location string
    Current    string
    Available	bool
    Twitter	string
    Github	string
    Linkedin	string
    Dribbble	string
    Steam	string
}

type Section_info struct {
	Header	string
	Text	string
	Image	string
}

func perror(err error) {
    if err != nil {
        panic(err)
    }
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
    url := "http://elsieng.com/data.json"

    res, err := http.Get(url)
    perror(err)
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    perror(err)

    var data Elsie
    err = json.Unmarshal(body, &data)
    if err != nil {
        fmt.Printf("%T\n%s\n%#v\n",err, err, err)
        switch v := err.(type){
            case *json.SyntaxError:
                fmt.Println(string(body[v.Offset-40:v.Offset]))
        }
    }
    
    //for i, project := range data.Elsieng.Project{	
        //for j, section := range project.Section{
        	//fmt.Println(j, section.Text)
        //}
    //}
    return c.Render(data)
}

type currentProject struct{
	Name	string
	Catchphrase	string
	Partner	string
	Link	string
	Gettext	string
	About	string
	Sections	[]projectSection
}

type projectSection struct{
	Header	string
	Text	string
	Image	string
}

func (c App) Project(id int) revel.Result {
	url := "http://www.elsieng.com/data.json"

    res, err := http.Get(url)
    perror(err)
    defer res.Body.Close()
    body, err := ioutil.ReadAll(res.Body)
    perror(err)

    var data Elsie
    err = json.Unmarshal(body, &data)
    if err != nil {
        fmt.Printf("%T\n%s\n%#v\n",err, err, err)
        switch v := err.(type){
            case *json.SyntaxError:
                fmt.Println(string(body[v.Offset-40:v.Offset]))
        }
    }
    
	cp := currentProject{}
    
    for i, project := range data.Elsieng.Project{
    	
        if (i == id){
        	cp.Name = project.Name
        	cp.Catchphrase = project.Catchphrase
        	
        	cpsections := make([]projectSection, 0)
        	
        	for j, section := range project.Section{
        		cpsection := projectSection{}
        		cpsection.Header = section.Header
        		cpsections = append(cpsections, cpsection)
        		j++
			}
			
			cp.Sections = cpsections
        }
    }
    
	return c.Render(cp)
}