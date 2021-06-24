# Welcome to Simple Calculator Built with Buffalo!

I chose Buffalo as prefered framework as it gives everything out of the box for newcomers like me. This is a simple Go (Golang) calculator with API. 

Please ignore my imaturity in the coding. Any recomendations on improving the code is welcome. #LearningForEver

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## What Next?

### To Add

Pass num1 and num2 in the path followed by add/.

Example

	http://localhost:3000/add/3/4

This should give you an answer 7

### To Subtract

Pass num1 and num2 in the path followed by sub/.

Example

	http://localhost:3000/sub/3/4

This should give you an answer -1

### To Multiply

Pass num1 and num2 in the path followed by mul/.

Example

	http://localhost:3000/mul/3/4

This should give you an answer 12

### To Divide

Pass num1 and num2 in the path followed by div/.

Example

	http://localhost:3000/div/8/4

This should give you an answer 2


Good luck!

[Powered by Buffalo](http://gobuffalo.io)
