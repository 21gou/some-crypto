# Auth cryptographic algorithms
## HMAC
Some notes:
* hmac = H([key ^ opad] H([key ^ ipad] text)) 
* key is zero padded to the block size of the hash function
* ipad = 0x36 and opad = 0x5c byte repeated for key length
    * Maximize Hamming distance

## AES 
* Galois field GF(2^128) 
* Polynomial: x^128 + x^7 + x^2 + x + 1 
