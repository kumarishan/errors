# errors

```
Err := New("some error")
```

equivalent to 
```
class Err extends Exception {
    public Err() {
        super("some error")
    }
}
```