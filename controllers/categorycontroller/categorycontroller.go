package categorycontroller

import (
	"crud-go/entities"
	"crud-go/models/categorymodel"
	"html/template"
	"net/http"
	"time"
)


func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}



func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		temp, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}


	if r.Method == "POST"{
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreatedAt = time.Now()
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Create(category); !ok{
			temp, err := template.ParseFiles("views/category/create.html")
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			temp.Execute(w , nil)
		}
		http.Redirect(w, r, "/category", http.StatusSeeOther)
	}
}


func Edit(w http.ResponseWriter, r *http.Request) {
	
}

func Delete(w http.ResponseWriter, r *http.Request) {
	
}