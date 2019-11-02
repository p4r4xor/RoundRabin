# Rabin-Cryptosystem
Encryption, Decryption and Keygeneration of Rabin Cryptosystem

### References Used
* Applied Cryptography by Bruce Schneier
* http://einspem.upm.edu.my/journal/fullpaper/vol11sapril/2.%20Zahari.pdf
  
### Build Details
```
git clone https://github.com/p4r4xor/rabin-cryptosystem.git
cd integer-factorization/  
go build rabin-decrypt.go
go build rabin-encrypt.go
go build rabin-keygenerate.go
```
### Usage
```
./rabin-encrypt <public_key_file_name> <message_in_decimal_format>
./rabin-decrypt <private_key_file_name> <ciphertext_in_decimal_format>
./rabin-keygenerate.go <public_key_file_name> <private_key_file_name>
```
Running the following prompts you to enter a number of your choice.  
Keyboard Interrupt stops the program.
