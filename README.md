minimal-docker-golang
=====================

This builds an executable with no external dependencies (thanks Golang!) and runs it in a sandboxed
environment (thanks Docker!). 

To build/run:

- Have docker installed
- Have golang installed
- Run `go build`
  - This should produce an executable called `minimal-docker-golang`
- Run `docker build -t hello .`
  - This should make an image called `hello` (`docker images` will show it)
- Run `docker run --rm hello`
  - You should see the output from the executable
  
