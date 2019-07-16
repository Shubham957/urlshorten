package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)
var ss string="abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var db *gorm.DB

func init() {
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:123@/urlshort?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&shortit{})
}

func shubham(c *gin.Context){
	c.HTML(http.StatusOK, "index.html",gin.H{

	})

}
func solve(s string,b uint) string{
	var ans string=""
	for b!=0 {
		ans+=string(ss[b%60])
		b/=60

	}
	return ans


return "ssss"

}
func shortener(c *gin.Context) {
	var todo shortit
	todo.Url=c.PostForm("url")
	count:=0

	db.Model(&shortit{}).Where("Url = ?", c.PostForm("url")).Count(&count)


	if count >= 1{

		db.Where("Url = ?", c.PostForm("url")).First(&todo)
		fmt.Print("In if")
		c.HTML(http.StatusOK, "short.html", gin.H{
			"complete" : todo.Url,
			"short": todo.ShortUrl,
		})
		return
	}

	db.Save(&todo)

 	todo.ShortUrl =solve(c.PostForm("url"),todo.ID)
	db.Save(&todo)
	c.HTML(http.StatusOK, "short.html", gin.H{
		"complete" : todo.Url,
		"short": todo.ShortUrl,
	})

//	var a shortit
//	a.Url=c.PostForm("url")
//	//db.Create(a)
//a.ShortUrl=solve(a.Url)
//db.Save(&a)

}
func main() {

	router := gin.Default()

	router.LoadHTMLGlob("views/*")
	router.Use(static.Serve("/", static.LocalFile("./views", true)))
//router.Post("/acton",createTodo)
router.POST("/action",shortener)

router.Run(":8088")

//	v1 := router.Group("/short")
//	{
//		v1.POST("/action", shortener)
//
////		//v1.GET("/:id", fetchSingleTodo)
////		//v1.PUT("/:id", updateTodo)
////		//v1.DELETE("/:id", deleteTodo)
//	}
//	router.Run()

}


type (
	// todoModel describes a todoModel type
	shortit struct {
		gorm.Model
		Url     string `json:"url"`
		ShortUrl string    `json:"Short_url"`
 	//	Id int `json:"id"`
		Hash int `json:"hash"`
	}


)

// createTodo add a new todo
//func createTodo(c *gin.Context) {
//	completed, _ := strconv.Atoi(c.PostForm("completed"))
//	todo := todoModel{Title: c.PostForm("title"), Completed: completed}
//	db.Save(&todo)
//	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
//}

// fetchAllTodo fetch all todos
//func fetchAllTodo(c *gin.Context) {
//	var todos []todoModel
//	//var _todos []transformedTodo
//
//	db.Find(&todos)
//
//	if len(todos) <= 0 {
//		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
//		return
//	}
//
//	//transforms the todos for building a good response
//	for _, item := range todos {
//		completed := false
//		if item.Completed == 1 {
//			completed = true
//		} else {
//			completed = false
//		}
//		_todos = append(_todos, transformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
//	}
//	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
//}
//
//// fetchSingleTodo fetch a single todo
//func fetchSingleTodo(c *gin.Context) {
//	var todo todoModel
//	todoID := c.Param("id")
//
//	db.First(&todo, todoID)
//
//	if todo.ID == 0 {
//		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
//		return
//	}
//
//	completed := false
//	if todo.Completed == 1 {
//		completed = true
//	} else {
//		completed = false
//	}
//
//	_todo := transformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
//	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
//}
//
//// updateTodo update a todo
//func updateTodo(c *gin.Context) {
//	var todo todoModel
//	todoID := c.Param("id")
//
//	db.First(&todo, todoID)
//
//	if todo.ID == 0 {
//		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
//		return
//	}
//
//	db.Model(&todo).Update("title", c.PostForm("title"))
//	completed, _ := strconv.Atoi(c.PostForm("completed"))
//	db.Model(&todo).Update("completed", completed)
//	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
//}
//
//// deleteTodo remove a todo
//func deleteTodo(c *gin.Context) {
//	var todo todoModel
//	todoID := c.Param("id")
//
//	db.First(&todo, todoID)
//
//	if todo.ID == 0 {
//		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
//		return
//	}
//
//	db.Delete(&todo)
//	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
//}

//
//
//ALTER USER 'root'@123'localhost' IDENTIFIED WITH mysql_native_password BY '123';
