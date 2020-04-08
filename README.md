# wisp-erc20
erc20 goLang 客户端


# FQ

- 启动问题
  - github.com/ethereum/go-ethereum/crypto/secp256k1
  ../vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/curve.go:42:10: fatal error: 'libsecp256k1/include/secp256k1.h' file not found
  include "libsecp256k1/include/secp256k1.h"
  
- 解决办法
  - 将 https://github.com/ethereum/go-ethereum 下载到本地
  - 拷贝 目录 go-ethereum-master/crypto/secp256k1/libsecp256k1 目录到 --> 本项目目录  vendor/github.com/ethereum/go-ethereum/crypto/secp256k1下