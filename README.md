## Relatório sobre o Algoritmo de Assinatura Digital com Criptografia RSA

### Integrante(s)
* Jairo Gomes Evaristo - 497466
* Carlos Aldrim Freire Melo Filho - 499075

### Introdução
Este algoritmo utiliza criptografia assimétrica RSA para garantir a segurança e autenticidade das mensagens trocadas entre Bob e Alice. Ele envolve a geração de chaves, assinatura e verificação de mensagens, e cifração e decifração de dados.

### Componentes Principais

#### 1. **Geração de Pares de Chaves RSA**

A função `GenerateKeyPair` gera um par de chaves RSA (pública e privada). A chave privada é mantida em segredo, enquanto a pública é compartilhada.

```go
func GenerateKeyPair(bits int) (*rsa.PrivateKey, error) {
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        return nil, err
    }
    return privateKey, nil
}
```

A função também exporta essas chaves no formato PEM para fácil armazenamento e transferência.

```go
func ExportPublicKeyAsPEM(pubkey *rsa.PublicKey) string {
    pubKeyBytes := x509.MarshalPKCS1PublicKey(pubkey)
    pemPubKey := pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PUBLIC KEY",
        Bytes: pubKeyBytes,
    })
    return string(pemPubKey)
}
```

#### 2. **Assinatura Digital**

A função `SignMessage` assina uma mensagem utilizando a chave privada. Um hash da mensagem é criado com SHA-256 e assinado com o algoritmo RSA-PSS.

```go
func SignMessage(privateKey string, message string) ([]byte, error) {
    privateKeyParseRSA, err := parsePrivateKeyFromPEM(privateKey)
    if err != nil {
        return nil, err
    }

    hashed := sha256.Sum256([]byte(message))
    signature, err := rsa.SignPSS(rand.Reader, privateKeyParseRSA, crypto.SHA256, hashed[:], nil)
    if err != nil {
        return nil, err
    }
    return signature, nil
}
```

Isso garante a integridade e a origem da mensagem, permitindo que o destinatário verifique quem assinou a mensagem e se ela não foi modificada.

#### 3. **Verificação de Assinatura**

Para verificar a assinatura, o destinatário usa a chave pública do remetente com a função `VerifySignature`. Isso assegura que a mensagem é autêntica e não foi adulterada.

```go
func VerifySignature(publicKey string, message string, signature []byte) error {
    publicKeyParseRSA, err := parsePublicKeyFromPEM(publicKey)
    if err != nil {
        return err
    }

    hashed := sha256.Sum256([]byte(message))
    err = rsa.VerifyPSS(publicKeyParseRSA, crypto.SHA256, hashed[:], signature, nil)
    if err != nil {
        return nil
    }
    return nil
}
```

#### 4. **Cifração e Decifração de Mensagens**

A função `EncryptMessage` cifra a mensagem com a chave pública do destinatário, enquanto `DecryptMessage` usa a chave privada para decifrar a mensagem. O método RSA-OAEP é utilizado para isso, provendo segurança adicional.

```go
func EncryptMessage(publicKey string, message string) (string, error) {
    publicKeyParseRSA, err := parsePublicKeyFromPEM(publicKey)
    if err != nil {
        return "", err
    }

    label := []byte("")
    hash := sha256.New()
    ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKeyParseRSA, []byte(message), label)
    if err != nil {
        return "", err
    }
    return string(ciphertext), nil
}
```

```go
func DecryptMessage(privateKey string, ciphertext []byte) (string, error) {
    publicKeyParseRSA, err := parsePrivateKeyFromPEM(privateKey)
    if err != nil {
        return "", err
    }

    label := []byte("")
    hash := sha256.New()
    plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, publicKeyParseRSA, ciphertext, label)
    if err != nil {
        return "", err
    }
    return string(plaintext), nil
}
```

#### 5. **Medição de Tempo**

As funções que medem o tempo de execução são usadas para comparar o desempenho das operações, como geração de chaves, assinatura, verificação, cifração e decifração.

```go
func MeasureAverageTime(operationName string, iterations int, operation func() error) time.Duration {
    var totalDuration time.Duration
    for i := 0; i < iterations; i++ {
        start := time.Now()
        err := operation()
        if err != nil {
            log.Fatalf("Erro na operação %s: %v", operationName, err)
        }
        totalDuration += time.Since(start)
    }
    averageDuration := time.Duration(totalDuration.Milliseconds()/int64(iterations)) * time.Millisecond
    fmt.Printf("Tempo médio para %s: %v ms\n", operationName, averageDuration.Milliseconds())

    return averageDuration
}
```

#### Gráfico de Tempos de Execução
https://github.com/user-attachments/assets/c7d72276-817f-4dde-b41c-ccd2d23a4095

#### Gráfico de Tempos de Execução (10 vezes)
https://github.com/user-attachments/assets/c6bebd0c-6c86-4956-ad6a-386b2d6b5cd8

### Fluxo de Execução

1. **Geração de Chaves:**
   - Bob gera seu par de chaves (pública e privada).
   - A chave pública é exportada e compartilhada com Alice.

```go
bobPrivateKey, err := util.GenerateKeyPair(2048)
bobPubPEM := util.ExportPublicKeyAsPEM(&bobPrivateKey.PublicKey)
```

2. **Assinatura de Mensagem:**
   - Bob assina a mensagem com sua chave privada.

```go
signature, err := util.SignMessage(bobPrivPEM, message)
```

3. **Verificação de Assinatura:**
   - Alice verifica a assinatura da mensagem com a chave pública de Bob.

```go
err := util.VerifySignature(bobPubPEM, message, signature)
```

4. **Cifração e Decifração:**
   - Alice cifra a mensagem com a chave pública de Bob.
   - Bob decifra a mensagem com sua chave privada.

```go
encryptedMessage, err := util.EncryptMessage(bobPubPEM, message)
decryptedMessage, err := util.DecryptMessage(bobPrivPEM, []byte(encryptedMessage))
```

### Como Iniciar a Aplicação

#### Compilar e Executar o Projeto

Navegue até a pasta raiz do projeto (onde está o arquivo `go.mod`) e execute o seguinte comando para rodar o programa:

```bash
go run cmd/main.go message secreta
```

### Gravação (Link)
https://ufcbr-my.sharepoint.com/:v:/g/personal/jairoevaristo_alu_ufc_br/EZ7qJyGPKfVEkd-gxDwwXDMB1icsQ5kf4rFTiAPC00qbDA?nav=eyJyZWZlcnJhbEluZm8iOnsicmVmZXJyYWxBcHAiOiJTdHJlYW1XZWJBcHAiLCJyZWZlcnJhbFZpZXciOiJTaGFyZURpYWxvZy1MaW5rIiwicmVmZXJyYWxBcHBQbGF0Zm9ybSI6IldlYiIsInJlZmVycmFsTW9kZSI6InZpZXcifX0%3D&e=r1QqhX

### Github
https://github.com/jairoevaristo/assinatura-digital-criptografia

### Conclusão

O algoritmo provê uma solução robusta para autenticar e proteger mensagens usando criptografia RSA. O código ilustra todas as etapas necessárias para garantir confidencialidade e integridade. As medições de tempo ajudam a avaliar o desempenho das operações, fornecendo insights sobre o custo computacional das operações criptográficas.
