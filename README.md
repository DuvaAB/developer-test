# Developer Candidate Test

Briefly write down how you solved the tasks and commit each step in a git repo.  
You can complete the steps in any order you want. If you need more information, don’t hesitate to ask.  

### Project 0 - Python

This program takes one argument which is the path to a csv file. This csv file is structured as follows: id (unique positive int64), x (int64), y (int64), z (int64).

The program parses this file and returns the rows which meets these rules to __stdout__. The rows that did not pass is returned to __stderr__.

Example csv files are located in the __test_files__ folder.

__Step 0a__: Solve the bug/s and make optimizations.  
__Step 0b__: Add a feature where, if x, y or z is negative. Replace its value with the absolute.  


### Project 1 - PHP

This file checks if the password is correct and sets the session key __user_logged_in__ to 1, but it has a quite serious security flaw that will allow an unauthorized user to login without knowing what the password is.

__Step 1a__: Fix the security flaw.  
__Step 1b__: Implement a form of password hashing  


### Project 2 - GoLang

Reads a csv from stdin that contains the output from __Project 0__.

The developer that made this program does not work here anymore, read the documentation ( _proj1/README.txt_ ) and investigate why it doesn’t work.

__Step 2a__: Solve the bug/s and make optimizations.  
__Step 2b__: Calculate the area of the triangle with the side lengths x, y and z. Add that as a fifth column in the csv called area.  


### Project 3 - Your choice

These steps can be made in the programming language of your choice.

__Step 3a__: Write a function that reverses the characters of a string.  
__Step 3b__: Write a function that reverses the words(separated with spaces) of a string.  
__Step 3c__: Write a function that prints the numbers from 1 to 100. But for multiples of 3 print “Fizz” and for the multiples of 5 print “Buzz”. For numbers which are multiples of both 3 and 5 print “FizzBuzz”.  
__Step 3d__: Write a function that calculates the angle between the hour and minute hands on a clock.  
