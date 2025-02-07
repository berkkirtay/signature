## A Simple PKCS-Based File Signer 
A simple digital signature processor program written in Go. Built to provide authenticity for my documents, papers and code.

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