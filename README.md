# movie-subtitle-auto
A command line program to automatically rename your movies and their subtitles. Works best with series.

# How It Works 
You run the program with paths to the movie series folder using "-moviepath" flag and subtitles folder using "-subtitlepath" flag and provide a name for the movies using "-name" flag and run the program. It will remain all movies and their respective subtitles for you  

# All Flags
-moviepath: path to the movies folder  *required*

-subtitlepath: path to the subtitle folder *required*

-name: name of the movie series *required*

-s: season of the movie series *defaults to 1*

-d: destination or folder to put the processed movie and subtitles in *defaults to current movie folder*

-c: boolean whether to copy or move files to destination folder *defaults to false*

-template: the format to name movies and episodes *eg. S0\*E0 means moviesname-S0nE0n* where n is the espisode number which is caculated automatically *defaults to S0\*E0*
