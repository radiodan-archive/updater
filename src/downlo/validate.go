package downlo

import(
  "log"
  "os"
  "io"
  // "io/ioutil"
  "crypto/sha1"
  "fmt"
)

func IsUpdateValid(path string, expectedHash string) (isValid bool) {
  isValid = false

  file, err := os.Open(path)

  if err != nil {
    log.Printf("Error reading file: %s \n", path)
  }

  hash := sha1.New()
  io.Copy(hash, file)
  fileHash := hash.Sum([]byte(""))

  fileHashString := fmt.Sprintf("%x", fileHash)

  if fileHashString == expectedHash {
    isValid = true
  } else {
    log.Printf("Invalid hash (expected %s, found %s)", expectedHash, fileHashString)
  }

  return
}