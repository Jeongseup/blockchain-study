# Advanced Encryption Standard, CBC모드 연습
from Crypto.Cipher import AES
from Crypto import Random
import numpy as np

# 대칭키 만들기
secretKey128 = b'0123456701234567'
secretKey192 = b'012345670123456701234567'
secretKey256 = b'01234567012345670123456701234567'

secretKey = secretKey128
plaintText ='This is plain text'
print("\n\n")


