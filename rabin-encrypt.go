package main

import (
  "fmt"
  "io/ioutil"
  "math/big"
  "os"
  "crypto/sha256"
)

func main() {

  if len(os.Args) != 3 {
    fmt.Println(" \n Follow command line specification \n ./rabin-encrypt" +
      "<publickey-file-name> <message to be encrypted, in decimal>\n")

  } else {

  file_name := os.Args[1]
  MessageInString := os.Args[2]
  Message := ConvertMessageToBigInt(MessageInString)

  N := ExtractDetailsFromPublicKeyFile(file_name)

  Ciphertext := Encrypt(Message, N)
  CipherTextInString := Ciphertext.String()

  hashofMessageInString := getMessageHashInString(MessageInString)

  output := CipherTextInString + hashofMessageInString
  
  // I've concatenated the SHA256 hash of the message with the ciphertext
  fmt.Println(" The ciphertext is ", output)

  }
}

func getMessageHashInString(MessageInString string)(string) {

  sum := sha256.Sum256([]byte(MessageInString))
  sumInHex := fmt.Sprintf("%x", sum)
  return sumInHex

}

func Encrypt(Message *big.Int, N *big.Int) (*big.Int) {

  exponentationComponent := big.NewInt(2)
  Ciphertext := squareAndMultiple(Message, exponentationComponent, N)
  return Ciphertext

}

func ExtractDetailsFromPublicKeyFile(file_name string) (*big.Int) {

  // In Rabin's crypto-system, N is the public key
  FileContent, err := ioutil.ReadFile(file_name)
  N := big.NewInt(0)

  if err != nil {
    fmt.Println(" Error readng data from the file")
  } else {

  NinString := string(FileContent)
  // Below statements to remove left and right bracket from the string
  NinString = NinString[1:(len(NinString) - 1)]


  boolError := false
  N, boolError = N.SetString(NinString,10)
  if boolError != true {
    fmt.Println(" Error in Set String")
    }

  }
  return N
}

func ConvertMessageToBigInt(MessageInString string) (*big.Int) {

  boolError := false
  Message := big.NewInt(0)

  Message, boolError = Message.SetString(MessageInString, 10)
  if boolError != true {
    fmt.Println(" Error in Set String")
    }

  return Message
}


func squareAndMultiple(a *big.Int, b *big.Int, c *big.Int) (*big.Int) {

  binExp := fmt.Sprintf("%b", b)
  binExpLength := len(binExp)

  initialValue := big.NewInt(0)
  initialValue = initialValue.Mod(a,c)

  result := big.NewInt(0)
  result = result.Set(initialValue)

  for i := 1; i < binExpLength; i++ {
    interMediateResult := big.NewInt(0)
    interMediateResult = interMediateResult.Mul(result,result)
    result = result.Mod(interMediateResult, c)

    if byte(binExp[i]) == byte(49) {
      interResult := big.NewInt(0)
      interResult = interResult.Mul(result,initialValue)
      result = result.Mod(interResult, c)
    }
  }
  return result

}
