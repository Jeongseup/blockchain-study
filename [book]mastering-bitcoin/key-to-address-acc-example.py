from __future__ import print_function
# bitcoin tools
import cryptos

# Base switching
code_strings = {
    2: '01',
    10: '0123456789',
    16: '0123456789abcdef',
    32: 'abcdefghijklmnopqrstuvwxyz234567',
    58: '123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz',
    256: ''.join([chr(x) for x in range(256)])
}


# 각 진수별 가능한 문자열들
def get_code_string(base):
    if base in code_strings:
        return code_strings[base]
    else:
        raise ValueError("Invalid base!")


# 인코드 개인키 encode(private_key, 256, 32)
# 밸류를 256비트로 32미니멈 렌스만큼?
def encode(val, base, minlen=0):
    base, minlen = int(base), int(minlen)
    code_string = get_code_string(base)
    result = ""

    while val > 0:
        result = code_string[val % base] + result
        val //= base

    return code_string[0] * max(minlen - len(result), 0) + result

'''
def bin_to_b58check(inp, magicbyte=0):
    # 블록체인 네트워크에 따라 vtype version 정보 추가
    if magicbyte == 0:
        inp = '\x00' + inp
    while magicbyte > 0:
        inp = chr(int(magicbyte % 256)) + inp
        magicbyte //= 256
    leadingzbytes = len(re.match('^\x00*', inp).group(0))
    # 더블 해싱
    checksum = bin_dbl_sha256(inp)[:4]
    return '1' * leadingzbytes + changebase(inp+checksum, 256, 58)
'''

# Generate a random private key
valid_private_key = False
valid_private_key = False

# 1 부터 cryptos.N 사의 숫자 랜덤한 숫자 하나를 추출
while not valid_private_key:
    private_key = cryptos.random_key()
    decoded_private_key = cryptos.decode_privkey(private_key, 'hex')
    valid_private_key = 0 < decoded_private_key < cryptos.N

print("Private Key (decimal) is: ", decoded_private_key)
print("Private Key (hex) is: ", private_key)

# Convert private key to WIF format
# WIF 포맷이란 Wallet Import Format의 약자로 개인키를 보다 쉽게 표현하기 위한 표준 포맷

# encode input은 십진수의 프라이빗 키 값
# 이제 고걸.. 256비트 인코더에 넣고
wif_encoded_private_key = cryptos.encode_privkey(decoded_private_key, 'wif')
# print("Private Key (WIF) by decoded key is: ", wif_encoded_private_key)


''' WIF code


def encode_privk(priv, feyormt, vbyte=0):
    if not isinstance(priv, int_types):
        return encode_privkey(decode_privkey(priv), formt, vbyte)
    elif formt == 'wif':
        return bin_to_b58check(encode(priv, 256, 32), 128+int(vbyte))
    elif formt == 'wif_compressed':
        return bin_to_b58check(encode(priv, 256, 32)+b'\x01', 128+int(vbyte))
    else: raise Exception("Invalid format!")
    
    def encode(val, base, minlen=0):
        base, minlen = int(base), int(minlen)
        code_string = get_code_string(base)
        result = ""
        while val > 0:
            result = code_string[val % base] + result
            val //= base
        return code_string[0] * max(minlen - len(result), 0) + result    
'''


# Add suffix "01" to indicate a compressed private key
# compressed_private_key = private_key + '01'
# print("Private Key Compressed (hex) is: ", compressed_private_key)

# Generate a WIF format from the compressed private key (WIF-compressed)
# wif_compressed_private_key = cryptos.encode_privkey(
    # cryptos.decode_privkey(compressed_private_key, 'hex_compressed'), 'wif_compressed')
# print("Private Key (WIF-Compressed) is: ", wif_compressed_private_key)

# Multiply the EC generator point G with the private key to get a public key point
# public_key = cryptos.fast_multiply(cryptos.G, decoded_private_key)
# print("Public Key (x,y) coordinates is:", public_key)
'''

# Encode as hex, prefix 04
hex_encoded_public_key = cryptos.encode_pubkey(public_key, 'hex')
print("Public Key (hex) is:", hex_encoded_public_key)

# Compress public key, adjust prefix depending on whether y is even or odd
(public_key_x, public_key_y) = public_key
compressed_prefix = '02' if (public_key_y % 2) == 0 else '03'
hex_compressed_public_key = compressed_prefix + (cryptos.encode(public_key_x, 16).zfill(64))
print("Compressed Public Key (hex) is:", hex_compressed_public_key)

# Generate Bitcoin address from public key
print("Bitcoin Address (b58check) is:", cryptos.pubkey_to_address(public_key))

# Generate compressed Bitcoin address from compressed public key
print("Compressed Bitcoin Address (b58check) is:",
      cryptos.pubkey_to_address(hex_compressed_public_key))
'''