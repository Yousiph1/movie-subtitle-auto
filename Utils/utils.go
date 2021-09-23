package utils
import (
  "path"
  "os"
  "errors"
  "strings"
  "fmt"
  "io/ioutil"
  "log"
  "strconv"

 "github.com/Yousiph1/movie-subtitle-auto/flags"
)

var validSubTitleFormats = [7]string{".srt",".ssa",".ttml",".sbv",".dfxp",".vvt",".txt"}
var validVideoFormats = [7]string{".mp4",".wmv",".mkv",".avchd",".mov",".avi",".webm"}

func isSubTitle(s string) bool {
     extName := path.Ext(s)
    for i := 0; i < len(validSubTitleFormats); i++ {
       if extName == validSubTitleFormats[i] {
         return true
       }
    }
   return false
}

func isVideo(s string) bool {
  extName := path.Ext(s)
  for i := 0; i < len(validVideoFormats); i++ {
     if extName == validVideoFormats[i] {
       return true
     }
  }
 return false
}

func isValidDirectory(s,t string) error {
    _, err := os.Open(s)
   if err != nil {
     if os.IsNotExist(err) {
       return errors.New(t +" directory does not exist")
     }
     if os.IsPermission(err) {
       return errors.New(t +" directory permission denied")
     }
      return errors.New("couldn't open " + t + " directory. Check path name and try again")
   }

   return nil
}


func getPrefix(prefix,name,s string) string {
   if strings.Contains(prefix,"*") && len(strings.Split(prefix,"*")) == 2 {
     arr := strings.Split(prefix,"*")
     sPrefix, ePrefix := arr[0],arr[1]
     return fmt.Sprintf("%s-%s%s%s",name,sPrefix,s,ePrefix)
   }
   return fmt.Sprintf("%s-S0%sE0",name,s)
}

func rename(pathPrefix, oldName,newName string) error {
      extName := path.Ext(oldName)
      oldName = path.Join(pathPrefix, oldName)
      newName = path.Join(pathPrefix, newName + extName)
      err := os.Rename(oldName,newName)
      return err
}

func renameSub(pathPrefix, moviepath, oldName, newName string) error {
        extName := path.Ext(oldName)
        oldName = path.Join(pathPrefix, oldName)
        newName = path.Join(moviepath, newName + extName)
        err := os.Rename(oldName,newName)
        return err
}

func HandleMovies(){
   err := isValidDirectory(flags.MoviePath,"movies")
   if err != nil {
    log.Fatal(err)
   }
   movies, err := ioutil.ReadDir(flags.MoviePath)
   i := 1;
   prefix := getPrefix(flags.Suffix,flags.Name, flags.Season)
   for _ , movie := range movies {
       if isVideo(movie.Name()) && !movie.IsDir() {
            nprefix := prefix + strconv.Itoa(i)
           err := rename(flags.MoviePath,movie.Name(), nprefix)
           if err != nil {
             log.Fatalf("could not rename movie: %s",movie.Name())
           }
           i++
       }else {
         log.Println(movie.Name(),"is not a valid video file and will be skipped")
       }
   }
}

func HandleSubTitles()  {
  err := isValidDirectory(flags.SubTitlePath,"subtitles")
  if err != nil {
   log.Fatal(err)
  }
  subtitles, err := ioutil.ReadDir(flags.SubTitlePath)
  i := 1;
  prefix := getPrefix(flags.Suffix,flags.Name, flags.Season)
  for _, subtitle := range subtitles {
    if isSubTitle(subtitle.Name()) && ! subtitle.IsDir() {
      nprefix := prefix + strconv.Itoa(i)
      err := renameSub(flags.SubTitlePath,flags.MoviePath,subtitle.Name(), nprefix)
      if err != nil {
        log.Fatalf("could not rename subtitle: %s",subtitle.Name())
      }
      i++
    }else {
      log.Println(subtitle.Name(),"is not a valid subtitle file and will be skipped")
    }
  }
}
