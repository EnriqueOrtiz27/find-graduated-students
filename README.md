# find-student

I created this program to better understand go routines and channels, so bear in mind it's somewhat primitive.

The program works as follows: you will receive a prompt to enter a student's name and last name,
and this code will look for students who graduated from [ITAM's](https://www.itam.mx) most popular careers.

To run it, you need to have [Go](https://go.dev) installed.

Download the repository and then execute 

```shell
go get
go build main.go
./main
```

You should see something like this:

````text
Do not use accents please!
Enter the first name of the student you're looking for: 
Pablo
Enter the last name of the student you're looking for: 
Alvarez

Searching...
Searching...
Searching...
Searching...
Searching...
Searching...

These are the names most similar to "Pablo Alvarez":

Alvarez Ortega Pablo: Graduated from Finance in 2021
Alvarez De Los Cobos Pablo: Graduated from Business Administration in 1995
Alvarez Magaña Pablo: Graduated from Business Administration in 2017
Iglesias Alvarez Pablo: Graduated from Business Administration in 1994
````

Potential improvements:
* Allow accents in user's input and remove them
* Handle cases where last names have more than one space (as in "De  Toledo")
* Create a web app where user can select from a given set of careers to speed up the process


Note: You can also run the repository unit tests with:
```shell
go test -v ./src/tests
```
