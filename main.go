package main

import (
  "log"
  "github.com/Yousiph1/movie-subtitle-auto/Utils"
)


func main()  {
     utils.HandleMovies()
     utils.HandleSubTitles()
     log.Println("success: renamed files successfully")
}
