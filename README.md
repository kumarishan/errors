# errors

## Sentinel Errors

### Creating or Extending Errors

```golang
var ErrInvalidInput := New("Invalid Input")
var ErrMissingIdParameter := Extend(ErrInvalidInput, "Missing Id Parameter")
```

`New` creates a new error. And `Extend` creates a new error which satisfies `Is` relationship with the error it extends.

```golang
errors.Is(ErrMissingIdParameter, ErrInvalidInput) // will return true
```

### Returning errors

```golang
func Return(err error, cause error, msg, errType string) error
```

Return returns an error of type `err`. You pass the `cause` of the error and override the error message of the error.

Return internally captures the call stack which lets you print the stack trace.

For eg,

```golang
import (
	"fmt"

	"github.com/kumarishan/errors"
)

var ErrInternalError = errors.New("internal error")
var ErrInvalidInput = errors.New("invalid input")
var ErrMissingIdParameter = errors.Extend(ErrInvalidInput, "missing id parameter")

func A() error {
	return errors.Return(ErrInvalidInput, nil, "")
}

func B() error {
	err := A()

	if errors.Is(err, ErrInvalidInput) {
		return errors.Return(ErrMissingIdParameter, err, "overriden error message")
	}

	return errors.Return(ErrInternalError, nil, "")

}

func main() {
	err := B()
	fmt.Printf("Got error: %+v\n", err)
}
```

the output is

```
Got error: overriden error message
	main.B(./errors/example/main.go:21) (0x112223b)
	main.main(./errors/example/main.go:29) (0x11222d8)
Caused by: invalid input
	main.A(./errors/example/main.go:14) (0x11221ff)
	main.B(./errors/example/main.go:18) (0x11221f8)
```

---

_todo_

## Similarities from Java

### Using New

```golang
Err := New("some error")
```

is equivalent to

```java
class Err extends Exception {
    public Err() {
        super("some error")
    }

    public Err(String message, Throwable cause) {
        super("some error, " + message, cause)
    }

    public Err(Throwable cause) {
        super("some error", cause)
    }
}
```

### Using Extend

```golang
ErrE := Extend(Err, "derieved error")
```

is equivalent to

```java
class ErrE extends Err {
    private String message = "derieved error";

    public ErrE(Throwable cause) {
        super(cause)
    }

    public ErrE() {
        super()
    }

    public toString() String {
        return super.toString() + ":" + message;
    }

}
```

### Using Struct Embedding

```golang

type EErr struct {
    errors.Error
    int a
    string b
}

func F() error {
    return &EErr({nil, nil, "some error"}, 10, "some value")
}
```

is equivalent to

```java

class EErr extends Exception {
    private int a;
    private String b;

    public EErr(int a, String b, string msg) {
        super(msg);
        this.a = a;
        this.b = b;
    }

}

class Example {
    public static void F() throws Exception {
        throw new EErr(10, "some value", "some error")
    }
}
```

## Returning Error

In golang the error is not thrown, but instead returned from the function and handled by the caller.

```golang

func F1() error {
    return errors.Return(Err, nil, "")
}

func F2() error {
    return errors.Return(Err, nil, "custom message")
}

func F3() error {
    cause := errors.New("some cuase")
    return errors.Return(Err, cause, "custom message")
}
```

is equivalent to

```java

class Example {

    class SomeCause extends Exception {
        public SomeCause() {
            super("some cause")
        }
    }

    public static void F1() throws Exception {
        throw new Err();
    }

    public static void F2() throws Exception {
        throw new Err("custom message")
    }

    public static void F3() throws Exception {
        throw new Err("cu)
    }
}

```
