# errors

## Defining new Errors

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
