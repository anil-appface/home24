# home24
This provides the basic information of the webpage. 

I have deployed this in HEROKU <a href="https://home24.herokuapp.com/">home24</a> https://home24.herokuapp.com/ app to access it in an easier way, Also enabled the automatic deployement when there is a push to main branch

## What is this app?

This is a web application written in Go. 
This application expects the valid URL for which it gets the following details for you. 
  1. Page Title.
  
  The text inside the title tag considers the page title.
  
  2. HTML Version. 
  
  The HTML version is determined by the doctype tag. 
  
  3. Internal Links Count & list of all links
  
  The Links which are internal to the application, and to which address cannot be found are treated as Internal links
  
  4. External Links Count & list of all external links
  
  The Links which are external, and those which has valid address are treated as External links.
  
  5. InAccessible Links Count.
  
  If the response of the head request for the external links falls under 200 and above 300 are considered as Inaccessible links.
  
  6. Whether WebPage has Login Form or not.
  
  Finding this subjective to developer. 
  But in this app, I'm deciding whether the app has login page or not based on the `input type=password`. <br/>
  TODO::  <br/>
    i .Can check whether the form has method post <br/>
    ii. Should have one input with type submit. (Again this one is not necessary) 

7. Each Heading Tag Count.
  
  Count the H1, H2, H3, H4, H5, H6 Tags in the page.
  
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
