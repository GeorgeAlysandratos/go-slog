
Hey~~~

This is me learning go!!! LETS GOOOO!!!!!

There exists better loggers out there, but I thought it would be a good idea to learn go by making one (that I had already made in the past for C++)

The following program

```
package main

import (
    "github.com/GeorgeAlysandratos/go-slog"
)

func main() {
    slog.SetBasename("slog_test")

    slog.Info("Heyyy")
    slog.Note("point me out")
    slog.Warn("eeee")
    slog.Error("adasd")

    slog.Infof("Heyyy")
    slog.Notef("point me out")
    slog.Warnf("eeee")
    slog.Errorf("adasd")
}
```

Should give you the following output

![Alt text](canvas.png?raw=true "Title")