# todoapi
A todo list RESTful API wannabe

This is a personal, unfinished, ongoing implementation of a todo list as a RESTful API written in Go.
It's far from being production ready yet.

I'm writing it as I learn about network affairs in Go. I started with the blog post [Making a RESTful JSON API in Go](http://thenewstack.io/make-a-restful-json-api-go/) by Cory Lanou, and slowly drifted away from it in a few approaches due to didactic purposes. Major bullets are:

 - Stick with the default http Go library.
 - Test a [Keep](http://godoc.org/github.com/coolparadox/go/storage/keep) database as the backend.

Keep databases are a recent creation of mine, targeting storage of typed Go data maintaining a minimal memory footprint. Their system of automatic generation of integer keys fits well when applied to management of resource ids in RESTful APIs.

See commit logs for further details.

Cheers,

Rafael Lorandi
