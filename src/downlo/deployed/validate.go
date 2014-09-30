package deployed

import(
  "log"
  "os"
  "io"
  "crypto/sha1"
  "fmt"
  "downlo"
)

func IsValidRelease(release downlo.Release) (isValid bool) {
  isValid = false

  path := release.Source
  expectedHash := release.Hash

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