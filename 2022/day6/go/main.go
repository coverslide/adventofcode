package main

import (
  "fmt"
  "os"
  "io"
)

func panicIf(err error) {
  if err != nil {
    panic(err)
  }
}

func main(){
  if len(os.Args) < 2 {
    panic("Input file argument required")
  }
  filename := os.Args[1]
  fh, err := os.Open(filename)
  panicIf(err)

  fmt.Printf("Part 1: %d\n", findUniqueInWindow(fh, 4))

  _, err = fh.Seek(0, io.SeekStart)
  panicIf(err)

  fmt.Printf("Part 2: %d\n", findUniqueInWindow(fh, 14))

}

func findUniqueInWindow(fh io.Reader, windowSize int) int{
  charMapCount := make(map[byte]int)
  charMapRepeated := make(map[byte]struct{})
  buffer := make([]byte, windowSize)
  singleBuffer := make([]byte, 1)

  bytes, err := fh.Read(buffer)
  panicIf(err)
  if bytes < windowSize {
    panic("not enough bytes")
  }
  for _, bufByte := range buffer {
    charMapCount[bufByte] += 1
    if charMapCount[bufByte] > 1 {
      charMapRepeated[bufByte] = struct{}{}
    }
  }
  count := windowSize
  for {
    if len(charMapRepeated) == 0 {
      break
    }
    count += 1
    bytes, err := fh.Read(singleBuffer)
    panicIf(err)
    if bytes < 1 {
      panic("not enough bytes")
    }
    bufByte := singleBuffer[0]
    charMapCount[bufByte] += 1
    if charMapCount[bufByte] > 1 {
      charMapRepeated[bufByte] = struct{}{}
    }
    movedByte := buffer[0]
    charMapCount[movedByte] -= 1
    if charMapCount[movedByte] < 2 {
      delete(charMapRepeated, movedByte)
    }
    if len(charMapRepeated) == 0 {
      break
    }
    buffer = append(buffer[1:], bufByte) 
  }
  return count
}
