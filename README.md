# fizbuz
Fizbuz as a service, written in Go.  Mainly just a place to play with go.

# Installing
```
go get github.com/poynting/fizbuz
go install $GOPATH/github.com/poynting/fizbuz
```

# Usage

By default, `fizbuz` runs on port 8080.
```
fizbuz
```

To change the port, set your PORT environment variable and run the program:
```
PORT=8081 fizbuz
```

Now point your browser to http://localhost:8081 and try urls such as /1, /3, /5, and /15.
