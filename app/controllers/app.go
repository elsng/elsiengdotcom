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
	TopImage	string
	BottomImage	string
	Href	string
	Hreftext	string
	Align	string
}

func perror(err error) {
    if err != nil {
        panic(err)
    }
}

type App struct {
	*revel.Controller
}

type elsieProfile struct {
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

type projectList struct {
	Project []projectItem
}

type projectItem struct {
	Id	int
	Name	string
	Gist	string
	Picture	string
	Brand	string
	Button	string
}

func (c App) Index() revel.Result {
    url := "http://dev.elsieng.com/data.json"

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
    
    //Profile
    myProfile := elsieProfile{}
    
	myProfile.Name = data.Elsieng.Profile.Name
	myProfile.Gist = data.Elsieng.Profile.Gist
	myProfile.Location = data.Elsieng.Profile.Location
	myProfile.Current = data.Elsieng.Profile.Current
	myProfile.Available = data.Elsieng.Profile.Available
	myProfile.Twitter = data.Elsieng.Profile.Twitter
	myProfile.Github = data.Elsieng.Profile.Github
	myProfile.Linkedin = data.Elsieng.Profile.Linkedin
	myProfile.Dribbble = data.Elsieng.Profile.Dribbble
	myProfile.Steam = data.Elsieng.Profile.Steam
    
    //Project List
    thelist :=  make([]projectItem, 0)
    
    for i := len(data.Elsieng.Project)-1; i >= 0; i-- {
    	for j, project := range data.Elsieng.Project{
    		if i == project.Id{	
        		projectItem := projectItem{}
				projectItem.Id = project.Id
				projectItem.Name = project.Name
				projectItem.Gist = project.Gist
				projectItem.Picture = project.Picture
				projectItem.Brand = project.Brand
				projectItem.Button = project.Button
				thelist = append(thelist, projectItem)
				j--
			}
        }
    }
    return c.Render(myProfile,thelist)
}

//Project Template

type currentProject struct{
	Id	int
	Name	string
	Brand	string
	Catchphrase	string
	Partner	string
	Link	string
	Gettext	string
	About	string
	Sections	[]projectSection
}

type projectSection struct{
	Id	int
	ProjectId	int
	Header	string
	Text	string
	TopImage	string
	BottomImage	string
	Href	string
	Hreftext	string
	Align	string
}

func (c App) Project(id int) revel.Result {
	url := "http://dev.elsieng.com/data.json"

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
        	cp.Id = project.Id
        	cp.Name = project.Name
        	cp.Brand = project.Brand
        	cp.Catchphrase = project.Catchphrase
        	cp.Partner = project.Partner
        	cp.Link = project.Link
        	cp.Gettext = project.Gettext
        	cp.About = project.About
        	
        	cpsections := make([]projectSection, 0)
        	
        	for j, section := range project.Section{
        		cpsection := projectSection{}
        		cpsection.Id = j
        		cpsection.Header = section.Header
        		cpsection.Text = section.Text
        		cpsection.TopImage = section.TopImage
        		cpsection.BottomImage = section.BottomImage
        		cpsection.Href = section.Href
        		cpsection.Hreftext = section.Hreftext
        		cpsection.Align = section.Align
        		cpsections = append(cpsections, cpsection)
        		j++
			}
			
			cp.Sections = cpsections
        }
    }
    
	return c.Render(cp)
}