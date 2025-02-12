## A Simple PKCS-Based File Signer 
A simple digital signature processor program written in Go. Built to provide authenticity for my documents, papers and code. 

### How it works
#### Signing a file
To create a new digital signature, the program simply creates a SHA256 hash from the data that combines the file content and a new timestamp. Then signs this hash with the private key of the user.

	H(x) = SHA256(x + timestamp)
 	S(priv_key, x) = Encrypt(H(x), priv_key)

 	``Where x is the given text(file).``

#### Verification of a signature
For verification, the program creates a SHA256 hash in the same way as above and decrypts the signature with the public key of the user. Then compares the two hashes to verify the authenticity of the signature.

	H(x) = SHA256(x + timestamp)
 	V(pub_key, x) = Decrypt(S, pub_key)
    Verification is successful if V equals to H 

 	``Where x is the given text(file) and S is the signature of the text.``

### Usage
You can compile the code and use the following commands with the executable:
- --generate
- --all
- --sign {private_key} {file}
- --verify {public_key} {signature} {file}
- --help

Where {*} represents the file paths for the key and signature files.

---
### Example
This above text is [signed](example/signature.pem) with my private key. You can verify it with my [public key](example/public_key.pem).
