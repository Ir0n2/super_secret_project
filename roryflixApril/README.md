This is the new and improved version of RoryFlix. 
Let's call it verison 4.20 because I have not been keeping track.

What we added:

-imdb id search. When you go to imdb you'll notice the url has a peice of text such as 'tt28093628'
You can search that id on our website and ~90% of the time you'll find where to watch it. It just works by concatenating that id to a url. 

-saved playlists. You can save lists of urls, image urls, and the titles of the film or show you're saving to a txt file. You can also set the file name when you save it. It's pretty cool

-downloaded videos. You can watch movies and shows you downloaded on Roryflix. The 'Videos' tab has an excellent way to browse movies or shows, or any videos you want saved to it. Maybe I should add another page to download movies. That's illegal but pretty cool. To use this feature you copy or move the videos you want to add to the 'videos' folder in the RoryFlix folder.

The file structure of this project goes like this:

-roryflix 

    -currentRoryflix(copy)
    -videos
    -README
    -server.go

to run the program.

type "go run server.go"

open your browser, preferably firefox with adblock, and type "localhost:8080"

go has this great feature to run straight from file without needing to compile it. Makes it very easy to distribute code. currentRoryflix(copy) just holds the save files and webapp html files. If you don't know what you're doing, don't fuck with it unless you're okay with fucking up your install of the app. If you do know what you're doing, you're free to modify shit as you like.

the Videos section of the webapp is written in the go server using html/templates. Why? becuase it's mostly code stolen from chatpt and stack overflow. But it works. And that's what's important.
