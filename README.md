# 프로젝트 이름

**[뱃지나 프로젝트에 관한 이미지들이 이 위치에 들어가면 좋습니다]**  
프로젝트의 전반적인 내용에 대한 요약을 여기에 적습니다

## 주요이력

20XX.09 프로젝트 시작

20XX.11 신규기능 추가

## 어떻게 시작하나요?

이 곳에서 설치에 관련된 이야기를 해주시면 좋습니다.

### 선행 조건

아래 사항들이 설치가 되어있어야합니다.

```
예시
```

### 빌드
~~~
$ cd src
$ go build -o ../bin/tls_server tls-server.go 
$ go build -o ../bin/tls_client tls-client.go
$ go build -o ../bin/tls_relay  tls-relay.go 
~~~

### 설치

아래 사항들로 현 프로젝트에 관한 모듈들을 설치할 수 있습니다.

```
예시
```

## 테스트의 실행

~~~
# Key considerations for algorithm "RSA" ≥ 2048-bit
$ openssl genrsa -out server.key 2048

# Key considerations for algorithm "ECDSA" ≥ secp384r1
# List ECDSA the supported curves (openssl ecparam -list_curves)
$ openssl ecparam -genkey -name secp384r1 -out server.key

Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)
$ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
~~~

### tls_server 실행
~~~
$ ./bin/tls_server -c ./config-server.yaml
~~~

### tls_relay 실행
~~~
$ ./bin/tls_relay -c ./config-relay.yaml
~~~

### tls_client 실행
~~~
$ ./bin/tls_client -c ./config-client.yaml  -f 0 -s "sender"  send_file.dat
~~~


왜 이렇게 동작하는지, 설명합니다

```
예시
```

### 테스트는 이런 식으로 작성하시면 됩니다

```
예시
```

## 배포

시스템을 배포하는 방법

## 누구랑 만들었나요?

* [이름](링크) - 무엇 무엇을 했어요
* [이름](링크) - 무엇 무엇을 했어요

## 라이센스

이 프로젝트는 MIT 라이센스로 라이센스가 부여되어 있습니다. 자세한 내용은 LICENSE.md 파일을 참고하세요.

## 기타

* 기타 사항을 작성합니다.
* 기타 사항을 작성합니다.
