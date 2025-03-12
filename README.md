# go_projects

## go_learning

GO Fundamentals from :

https://quii.gitbook.io/learn-go-with-tests

### Hello, World
Writing tests

Declaring functions, with arguments and return types

if, const and switch

Declaring variables and constants

TDD workflow

### Integers
Integers, addition

Writing better documentation so users of our code can understand its usage quickly

More practice of the TDD workflow

### Iteration
Learned for

Learned how to write benchmarks

More TDD practice

### Arrays and slices
Arrays

Slices

- The various ways to make them

- How they have a fixed capacity but you can create new slices from old ones using append

- How to slice, slices!

len to get the length of an array or slice

Test coverage tool

reflect.DeepEqual and why it's useful but can reduce the type-safety of your code

slices.Equal 

### Structs, methods & interfaces

Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer

Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)

Adding methods so you can add functionality to your data types and so you can implement interfaces

Table driven tests to make your assertions clearer and your test suites easier to extend & maintain

### Pointers & errors

Pointers

- Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.

- The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).

nil

- Pointers can be nil

- When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.

- Useful for when you want to describe a value that could be missing

Errors

- Errors are the way to signify failure when calling a function/method.

- By listening to our tests we concluded that checking for a string in an error would result in a flaky test. So we refactored our implementation to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.

- This is not the end of the story with error handling, you can do more sophisticated things but this is just an intro. Later sections will cover more strategies.

- Don’t just check errors, handle them gracefully

Create new types from existing ones

- Useful for adding more domain specific meaning to values

- Can let you implement interfaces

### Maps

In this section, we covered a lot. We made a full CRUD (Create, Read, Update and Delete) API for our dictionary. Throughout the process we learned how to:

- Create maps

- Search for items in maps

- Add new items to maps

- Update items in maps

- Delete items from a map

- Learned more about errors

- How to create errors that are constants

- Writing error wrappers

### Dependency Injection
Our first round of code was not easy to test because it wrote data to somewhere we couldn't control.

Motivated by our tests we refactored the code so we could control where the data was written by injecting a dependency which allowed us to:

    Test our code If you can't test a function easily, it's usually because of dependencies hard-wired into a function or global state. If you have a global database connection pool for instance that is used by some kind of service layer, it is likely going to be difficult to test and they will be slow to run. DI will motivate you to inject in a database dependency (via an interface) which you can then mock out with something you can control in your tests.

    Separate our concerns, decoupling where the data goes from how to generate it. If you ever feel like a method/function has too many responsibilities (generating data and writing to a db? handling HTTP requests and doing domain level logic?) DI is probably going to be the tool you need.

    Allow our code to be re-used in different contexts The first "new" context our code can be used in is inside tests. But further on if someone wants to try something new with your function they can inject their own dependencies.
