package slog

import(
	"fmt"
	"time"
	"sync"
	"strings"
	"runtime"
	"os"
)

func Show(enable bool) {
	file_lock.Lock()
	defer file_lock.Unlock()
	
	stdout_enabled = enable
}

func SetBasename(name string) {
	file_name = name
}

func Info(entry string) {
	log(info, entry, false)
}

func Note(entry string) {
	log(note, entry, false)
}

func Warn(entry string) {
	log(warn, entry, false)
}

func Error(entry string) {
	log(error, entry, false)
}

func Infof(entry string) {
	log(info, entry, true)
}

func Notef(entry string) {
	log(note, entry, true)
}

func Warnf(entry string) {
	log(warn, entry, true)
}

func Errorf(entry string) {
	log(error, entry, true)
}

// ----------------------------------------------

var file_lock sync.Mutex
var file_day int = -1
var file_name string
var file* os.File
var stdout_enabled bool = true

type log_type int32

const (
	info log_type = 0
	note log_type = 1
	warn log_type = 2
	error log_type = 3
)

func init() {
	stdout_enabled = true
	file_name = get_last_slash(os.Args[0])
}

func get_last_slash(a string) string {
	if len(a) > 0 {
		cutoff := len(a) - 1
		for cutoff > 0 {
			if a[cutoff] == '/' {
				cutoff++
				break
			}
			cutoff--
		}
		return a[cutoff:]
	}
	return a
}

func get_color(t log_type) string {
	if t == info {
		return "\033[0;97m"
	} else if t == note {
		return "\033[1;96m"
	} else if t == warn {
		return "\033[1;93m"
	} else if t == error {
		return "\033[1;31m"
	}

	panic("get_color no valid argument")

	return ""
}

func reset_color() string {
	return "\033[0m"
}

func log(t log_type, entry string, file_loc bool) {
	out_str := strings.Builder{}
	out_str.Grow(256)

	file_lock.Lock()
	defer file_lock.Unlock()

	now := time.Now()
	color_start := 0
	color_end := 0

	if now.Day() != file_day {
		file_day = now.Day()
		filename := fmt.Sprint(file_name, "_", now.Day(), "-", int(now.Month()), "-", now.Year(), ".txt")
		if file != nil {
			file.Close()
		}

		f, err := os.OpenFile(filename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0755)
		if err == nil {
			file = f
			file.WriteString("\n")
		} else {
			panic(err)
		}
	}

	fmt.Fprintf(&out_str, "%02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())
	fmt.Fprintf(&out_str, ".%05d", now.Nanosecond() / 10000)

	out_str.WriteString(" ")

	color_start = out_str.Len()

	if t == info {
		out_str.WriteString("INFO")
	} else if t == note {
		out_str.WriteString("NOTE")
	} else if t == warn {
		out_str.WriteString("WARN")
	} else if t == error {
		out_str.WriteString(" ERR")
	}
	out_str.WriteString(" ")
	out_str.WriteString(entry)

	color_end = out_str.Len()

	if file_loc {
		_, file, line, ok := runtime.Caller(2)

		out_str.WriteString(" [")

		if ok {
			file = get_last_slash(file)
			out_str.WriteString(file)
			out_str.WriteString(":")
			fmt.Fprintf(&out_str, "%d", line)
		} else {
			out_str.WriteString("???")
		}

		out_str.WriteString("]")
	}

	out := out_str.String()

	file.WriteString(out)
	file.WriteString("\n")

	if stdout_enabled {
		fmt.Print(out[:color_start], get_color(t), out[color_start:color_end], reset_color(), out[color_end:])
		fmt.Println("")
	}
}
