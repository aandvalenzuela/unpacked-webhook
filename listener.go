package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "time"
)

func watchFile(filePath string) error {
    initialStat, err := os.Stat(filePath)
    if err != nil {
        return err
    }

    for {
        stat, err := os.Stat(filePath)
        if err != nil {
            return err
        }

        if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
            break
        }

        time.Sleep(1 * time.Second)
    }

    return nil
}

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func main() {
  file_name := "notifications.txt"
  for {
    doneChan := make(chan bool)

    go func(doneChan chan bool) {
        defer func() {
            doneChan <- true
        }()

        err := watchFile(file_name)
        if err != nil {
            fmt.Println(err)
        }

        fmt.Println("New webhook notification received")
    }(doneChan)

    <-doneChan

    // read line by line
    lines, err := readLines(file_name)
    if err != nil {
      log.Fatalf("readLines: %s", err)
    }
    // print file contents
    num_lines := len(lines)
    last_line := lines[num_lines-1]
    fmt.Println(last_line)
  }
}

