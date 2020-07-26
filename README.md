# GoSetup
⚡️ GoSetup is the fastest way to initialize your next project with [Go](https://github.com/golang/go), [MongoDB](https://github.com/mongodb/mongo-go-driver), [Gin-Gonic](https://github.com/gin-gonic/gin) and [Docker](https://www.docker.com/) !  
Have you ever spent a lot of time re-initializing your web framework at the beginning of a new project? Setup the project layout, the docker configuration, the database connection and the associated structure, your API router & the matching controllers, etc.  
GoSetup will perform a basic initialization for a quick and easy start of your next project.

<br />

## ✨ Features
♻️    -    Automatic code generation for your need with a minimal configuration file. Your models, your route handlers, your controller actions are generated automatically.   

📋    -    Follow the **[Standard Go Project Layout](https://github.com/golang-standards/project-layout)** to provide an solid start, ready to evolve and to be adapted for all your needs.

🥃    -    Use the **[GinGonic](https://github.com/gin-gonic/gin)** web framework to setup your API, with a logger, the CORS configuration, and all the basic routes.  

🌿    -    Integrate with **[MongoDB](https://github.com/mongodb/mongo-go-driver)** for an easy modern webapp integration, by initializing the connection and the multiple collections.

🐳    -    Provide some basic ready-to-use deployment configuration with **[Docker](https://www.docker.com/)**, with a `DockerFile`, a specific `docker-compose.yml` for your dev environment and one for your prod.


## ⚙️ How does this work ?
Based on a minimal configuration file, GoSetup will generate the needed code to setup your API routing with your database connection.  
The generated code is defined in the `templates` directory. Let's take a simple example with the following `config.json` :

```json
{
	"name": "bookSetup",
	"root": "github.com/tbouder/bookSetup",
	"collections": [
		{
			"name": "books",
			"dbName": "books",
			"modelName": "book",
			"model": [
				{"name": "name", "type": "*string"},
				{"name": "author", "type": "*string"},
				{"name": "category", "type": "*string"},
				{"name": "pages", "type": "*int"},
				{"name": "price", "type": "*float32"}
			]
		}
	]
}
```
This config file means :
> We want to create a new project, `bookSetup`, which is an API with a database, where we have one `books` collection with 5 fields : `name`, `author`, `category`, `pages` and `price`.

When running the GoSetup tool, we will have a `bookSetup` directory created, with a `Docker` setup, a `MongoDB` setup and a router with `Gin`.
To interact with the API, 6 routes will be automatically generated with the associated functions :
```go
c := booksController{}
router.GET(`book/find/:publicID`, c.FindID) //Find a book based on it's publicID
router.POST(`book/find`, c.Find) //Find a book based on the body params
router.POST(`books/list`, c.List) //Find some books based on the body params
router.POST(`book/post`, c.Post) //Add a new book the the database
router.PUT(`book/update`, c.Update) //Update an existing book
router.DELETE(`book/delete`, c.Delete) //Delete an existing book
```
The associated functions will proceed to some check before querying the `model` package in order to execute the corresponding action.

```go
//Find will retreive one element in the database
func (y booksController) Find(c *gin.Context) {
    var requestFilter models.RequestFilter
    if err := c.ShouldBindBodyWith(&requestFilter, binding.JSON); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{`error`: `bad request`})
        return
    }

    result, err := models.NewBook().Find(requestFilter)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{`error`: `could not retreive element`})
        return
    }
    c.JSON(http.StatusOK, result)
}
```

In the `model` package, a new `Book` go type has been created (struct with functions) in order to interact with the database :
```go
//Book represent a element of type book
type Book struct {
	ID        *primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	PublicID  *string             `json:"publicID" bson:"publicID"`
	CreatedAt *time.Time          `json:"-" bson:"createdAt"`
	UpdatedAt *time.Time          `json:"-" bson:"updatedAt"`
	DeletedAt *time.Time          `json:"-" bson:"deletedAt"`
	Name      *string             `json:"name" bson:"name"`
	Author    *string             `json:"author" bson:"author"`
	Category  *string             `json:"category" bson:"category"`
	Pages     *int                `json:"pages" bson:"pages"`
	Price     *float32            `json:"price" bson:"price"`
}

//Find will perform a search in the Books collection to find the element matching the filters
func (x *Book) Find(rf RequestFilter) (*Book, error) {
	document := db.Books.FindOne(
		context.Background(),
		rf.Filter,
	)
	element := Book{}
	err := document.Decode(&element)
	return &element, err
}
```

All in one, *assuming your database is already populated*, you can easily access some element with a simple curl command.

```sh
curl --request POST \
  --url {{YOUR_SERVER_BASE_URI}}/book/find \
  --header 'content-type: application/json' \
  --data '{
  "filter": {
	  "name": "Mars trilogy"
  },
  "skip": 0
}'
```

## Road Map
- [ ] Provide installation process
- [ ] Provide full example
- [ ] Auto-generated repository example
- [ ] Detailled `config` file
- [ ] Auto-generated .env file
- [ ] Middleware integration
- [ ] JWT integration
