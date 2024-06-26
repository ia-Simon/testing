# Test-crypto

This project is a test script to test the cryptographic functions related to DES, 3DES and their usage by the Mastercard.

## Definitions

**DES:** A cryptographic algorithm which uses an 8-byte key to perform encryption/decryption of a block of the same length.

**3DES:** A cryptographic algorithm consisting of performing DES 3 times in a roll wih 3 keys. The key used by this algorithm may be 16
or 24 bytes long, and will be split in 3 keys of 8 bytes each. If the provided key has only 16 bytes, the first 8 bytes will be reused as 
the third key.

## References

1. [Key component generator](https://emvlab.org/keyshares/?combined=94A6F81D22222E4645134C410C7DF8AF&combined_kcv=482598&one=CCAA04EDACE52FBD3C5D0BA89B86C416&one_kcv=346288&two=5FCE656D978D348FBE49D5A05BC7F944&two_kcv=82ED85&three=07C2999D194A3574C7079249CC3CC5FD&three_kcv=AB5BCA&numcomp=three&parity=ignore&action=Generate+128+bit)
2. [XOR calculator](https://xor.pw)
3. [DES Calculator](https://paymentcardtools.com/basic-calculators/des-calculator)
4. [Cryptographic algorithms](https://tecnologiadarede.webnode.com.br/news/noticia-aos-visitantes/)
5. [AWS Payment Cryptography key import](https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-import.html#keys-import-rsaunwrap)