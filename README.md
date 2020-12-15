# home24
This provides the basic information of the webpage. 

I have deployed this in HEROKU <a href="https://home24.herokuapp.com/">home24</a> https://home24.herokuapp.com/ app to access it in an easier way, Also enabled the automatic deployement when there is a push to main branch

## What is this app?

This is a web application written in Go. 
This application expects the valid URL for which it gets the following details for you. 
  1. Page Title.
  2. HTML Version. 
  3. Internal Links Count. 
  4. External Links Count.
  5. InAccessible Links Count.
  6. Whether WebPage has Login Form or not.
  7. Each Heading Tag Count.
  
## How to run?

### Pre-requisites 
 Language: Golang Installed

### Steps

  1. Clone the repository. 
  2. Run the app using the command `go run .` Make sure that no errors found in console.
  3. Now run the http://localhost:3000 in your browser to open the app. 
 
 The port 3000 can be changed in `.env` file. 
  
  
## Where to Go from here?
The roadmap for this app are 

  1. To crawl the web page and give the in-depth details.

Any other suggestions are welcome... :) 
