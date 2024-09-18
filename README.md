## Relatório sobre o Algoritmo de Assinatura Digital com Criptografia RSA

### Introdução
Este algoritmo utiliza criptografia assimétrica RSA para garantir a segurança e autenticidade das mensagens trocadas entre Bob e Alice. Ele envolve a geração de chaves, assinatura e verificação de mensagens, e cifração e decifração de dados, além de simular o envio de uma chave pública.

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

#### 5. **Envio de Chave Pública via E-mail**

A chave pública de Bob é enviada a Alice por e-mail usando a API `Resend`. A função `Send` prepara o e-mail e envia um anexo com a chave pública em formato de texto.

```go
func (r *ResendEmail) Send(to []string, attachments []byte) error {
    params := &resend.SendEmailRequest{
        From:    "Chatbot <jairoevaristodev@gmail.com>",
        To:      to,
        Html:    "<p>Olá, essa é sua chave publica para troca de mensagens</p>",
        Subject: "Sua chave pública chegou!",
        Attachments: []*resend.Attachment{
            {
                Content:  attachments,
                Filename: "public-key.txt",
            },
        },
    }

    sent, err := r.client.Emails.Send(params)
    if err != nil {
        return err
    }

    fmt.Println("Email successfully send", sent.Id)
    return nil
}
```

#### 6. **Medição de Tempo**

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

### Fluxo de Execução

1. **Geração de Chaves:**
   - Bob gera seu par de chaves (pública e privada).
   - A chave pública é exportada e enviada para Alice.

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

5. **Envio de Chave Pública:**
   - Bob envia a chave pública para Alice por e-mail.

```go
err := resendEmail.Send([]string{"alice@example.com"}, []byte(bobPubPEM))
```

Para iniciar a aplicação e realizar as operações de criptografia, assinatura digital, e envio de e-mail, siga os passos descritos abaixo. Vamos focar em como rodar o projeto Go.

### Como Iniciar a Aplicação

#### Compilar e Executar o Projeto

Navegue até a pasta raiz do projeto (onde está o arquivo `go.mod`) e execute o seguinte comando para rodar o programa:

```bash
go run cmd/main.go message secreta
```

#### Conclusão

O algoritmo provê uma solução robusta para autenticar e proteger mensagens usando criptografia RSA. O código ilustra todas as etapas necessárias para garantir confidencialidade e integridade, integrando também o envio da chave pública por e-mail. As medições de tempo ajudam a avaliar o desempenho das operações, fornecendo insights sobre o custo computacional das operações criptográficas.
