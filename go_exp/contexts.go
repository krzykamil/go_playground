package main

import (
    "context";
    "fmt";
    "log";
    "time"
)

func main() {
    start := time.Now()
    ctx := context.WithValue(context.Background(), "foo", "bar")
    userID := 10
    val, err := fetchUserData(ctx, userID)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Result: ", val)
    fmt.Println("took: ", time.Since(start))
}

type Response struct {
    val int
    err error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
    val := ctx.Value("foo")
    fmt.Println(val.(string))
    ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 200)
    defer cancel()

    respch := make(chan Response)

    go func() {
        val, err := fetchSlowStuff()
        respch <- Response{val, err}
    }()

    for {
        select {
            case <-ctx.Done():
                return 0, fmt.Errorf("timeout")
            case resp := <-respch:
                return resp.val, resp.err
        }
    }
}

func fetchSlowStuff() (int, error) {
  time.Sleep(time.Millisecond * 100)
  return 42, nil
}
