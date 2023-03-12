# A Golang BookStore CRUD  management system


### lessons learnt
**Problem** : Couldn't correctly import mysql driver lib. 

**The Wayout :**  

The mysql driver lib wasn't correctly imported. The underscore `-`is called blank identifier in Go.
you'll need to add it in front of that import in the case of side-effect imports. for side-effect imports, you wouldn't be calling any function by yourself. you only need to import it. The underscore  would then prevent Go formatter from removing it (because naturally go removes all unused imports). Go database driver libraries use this pattern a lot.
[Read this StackOverFlow Blog](https://stackoverflow.com/questions/21220077/what-does-an-underscore-in-front-of-an-import-statement-mean#:~:text=Underscore%20is%20a%20special%20character,compiler%20will%20simply%20ignore%20it.)
and [this](https://v1.gorm.io/docs/connecting_to_the_database.html#MySQL)


## Run the Program 

To run the program on your machine, you need to:
* fork and clone the repo ().Then you need to install any API development tool POST MAN, Rapid Api etc ...

* goto to `pkg/config/app.go` file and edit the mysql database info i.e(database username, password, your Database name)

* open your terminal and `run cmd/main/main.go` to run the program.
* while you get the `server now running on PORT:` message, head to your api development and test tool and test each endpoint.
* use ctrl-c to stop the server.