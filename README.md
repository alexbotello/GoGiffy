# GoGiffy: Automate viewing of GIFs/Videos

GoGiffy is small command line tool for lazy reddit surfing.
It automatically crawls a subreddit and collects all Gfycat/Imgur links and displays them in a headless chrome browser

<img src="/gogiffy.gif?raw=true" width="600px" height="381px">

## Installing
Install in the usual Go way:
```
go get -u github.com/alexbotello/GoGiffy
```
Build the exectuable:
```
go build github.com/alexbotello/GoGiffy
```

## Usage
```
./GoGiffy SUBREDDIT DURATION
```
Subreddit: The name of the subreddit to open, e.g. `r/gifs` or `r/funny`

Duration: The number of seconds to wait before continuing to next GIF