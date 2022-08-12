# movie-subtitle-auto
A command line program to automatically rename your movies and their subtitles. Works best with series.

# How It Works 
You run the program with paths to the movie series folder using "-moviepath" flag and subtitles folder using "-subtitlepath" flag and provide a name for the movies using "-name" flag and run the program. It will remain all movies and their respective subtitles for you  

# All Flags
-moviepath: movies directory  *required*

-subtitlepath: subtitles directory *required*

-name: name of the series or preferred name *required*

-s: season of the series *defaults to 1*

-d: destination or folder to put the processed movie and subtitles in *defaults to current movie folder*

-c: boolean whether to copy or move files to destination folder *defaults to false*

-template: the format to name movies and episodes *eg. S0\*E0 means moviename-S0nE0n* where n is the season and espisode number which is caculated automatically *defaults to S0\*E0*
