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
    Picture	string
    Partner	string
    Style	string
    Link	string
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
    
    for i, project := range data.Elsieng.Project{
    
        fmt.Println(i, project.Name, project.Gist)
        
        for j, section := range project.Section{
        	fmt.Println(j, section.Text)
        }
    }
    
    return c.Render(data)
}
