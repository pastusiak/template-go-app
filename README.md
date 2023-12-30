<img src="https://go.dev/images/gophers/biplane.svg" height="200px" />

# GO LANG App Template

### Created by Przemyslaw Pastusiak
> https://github.com/pastusiak/template-go-app

---
### Usage
```
go run main.go

go build main.go
```

### Console arguments
Option list / Help
``` -? | --help```

Autorun option
``` -o <option_name> ```

Shutdown application after option done
``` -q ```

### Examples
```
go run main.go -o ex1
go run main.go -q -o ex1
go run main.go -o ex2 -q
go run main.go -o ex2 -o ex1
go run main.go -o ex2 -o ex1 -q -o ex3
```