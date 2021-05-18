package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

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

func indexOf(element string, data []string) (int) {
   for k, v := range data {
       if element == v {
           return k
       }
   }
   return -1    //not found.
}

func main() {
  var added_lines []string
  file_name := "notifications.txt"

  // read line by line
  lines, err := readLines(file_name)
  if err != nil {
    log.Fatalf("readLines: %s", err)
  }
  // print file contents
  num_lines := len(lines) - 1
  last_line := lines[num_lines]

  for {
    // read line by line
    lines, err := readLines(file_name)
    if err != nil {
      log.Fatalf("readLines: %s", err)
    }

    current_num_lines := len(lines) - 1
    current_last_line := lines[num_lines]

    if num_lines != current_num_lines {
      index := indexOf(last_line, lines)
      added_lines = lines[index+1:]
      for k_new, v_new := range added_lines {
	fmt.Println("index:", k_new)
	fmt.Println("content:", v_new)
      }
      added_lines = nil
      num_lines = current_num_lines
      last_line = current_last_line
    }
  }
}

