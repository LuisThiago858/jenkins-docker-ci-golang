# Jenkins Docker CI Golang

Projeto desenvolvido para a **Terceira Atividade Prática - Usando containers com Integração Contínua (2026-1)**.

O objetivo do projeto é demonstrar o uso do **Jenkins com Docker em um pipeline de Integração Contínua**, realizando a compilação dos fontes em um container Docker e a execução dos testes em outro container Docker isolado.

## Tema

**Jenkins com Docker**

## Objetivo da atividade

Adaptar o Jenkins para executar uma pipeline na qual:

* o código-fonte é obtido de um repositório no GitHub;
* o build da aplicação é executado dentro de um container Docker;
* os testes automatizados são executados em outro container Docker;
* diferentes cenários de execução são demonstrados no Jenkins.

## Tecnologias utilizadas

* Golang 1.26
* Docker
* Jenkins
* Git
* GitHub

## Descrição do projeto

O projeto implementa um conversor de temperatura em Golang.

Foram criadas funções para realizar conversões entre as escalas Celsius e Fahrenheit:

* Fahrenheit para Celsius;
* Celsius para Fahrenheit.

Além disso, foram implementados testes automatizados para validar o comportamento esperado dos métodos de conversão.

## Estrutura do projeto

```text
jenkins-docker-ci-golang/
├── Jenkinsfile
├── go.mod
├── cmd/
│   └── app/
│       └── main.go
└── internal/
    └── temperature/
        ├── converter.go
        ├── service.go
        └── service_test.go
```

## Organização do código

A implementação foi organizada de forma modular:

* `cmd/app/main.go`: ponto de entrada da aplicação;
* `internal/temperature/converter.go`: definição dos tipos, entidade de temperatura e interface de conversão;
* `internal/temperature/service.go`: implementação das regras de conversão;
* `internal/temperature/service_test.go`: testes automatizados da regra de negócio.

Embora Golang não utilize orientação a objetos baseada em classes, o projeto aplica princípios de organização e design por meio de `structs`, métodos, interfaces e encapsulamento por pacote.

## Como executar localmente com Docker

Não é necessário ter o Go instalado diretamente na máquina. A execução pode ser feita utilizando a imagem oficial do Golang no Docker.

### Verificar a versão do Go no container

```powershell
docker run --rm -v "${PWD}:/app" -w /app golang:1.26 go version
```

### Organizar dependências

```powershell
docker run --rm -v "${PWD}:/app" -w /app golang:1.26 go mod tidy
```

### Compilar o projeto

```powershell
docker run --rm -v "${PWD}:/app" -w /app golang:1.26 go build -v ./...
```

### Executar os testes

```powershell
docker run --rm -v "${PWD}:/app" -w /app golang:1.26 go test -v ./...
```

## Pipeline Jenkins

A pipeline foi definida no arquivo `Jenkinsfile`.

O Jenkins realiza as seguintes etapas:

1. Checkout do código-fonte a partir do GitHub;
2. Build da aplicação em um container Docker;
3. Execução dos testes em outro container Docker;
4. Indicação do resultado final da pipeline.

## Jenkinsfile

```groovy
pipeline {
    agent any

    triggers {
        cron('H 2 * * *')
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build em container Docker') {
            steps {
                bat '''
                    echo Container de build:
                    docker run --rm ^
                      -v "%WORKSPACE%:/app" ^
                      -v go-mod-cache:/go/pkg/mod ^
                      -w /app ^
                      golang:1.26 ^
                      sh -c "hostname && go version && go mod download && go build -v ./..."
                '''
            }
        }

        stage('Testes em outro container Docker') {
            steps {
                catchError(buildResult: 'UNSTABLE', stageResult: 'UNSTABLE') {
                    bat '''
                        echo Container de testes:
                        docker run --rm ^
                          -v "%WORKSPACE%:/app" ^
                          -v go-mod-cache:/go/pkg/mod ^
                          -w /app ^
                          golang:1.26 ^
                          sh -c "hostname && go version && go test -v ./..."
                    '''
                }
            }
        }
    }

    post {
        success {
            echo 'Pipeline executado com sucesso.'
        }

        unstable {
            echo 'Pipeline ficou instável porque algum teste falhou.'
        }

        failure {
            echo 'Pipeline falhou durante o build.'
        }
    }
}
```

## Configuração do Jenkins

O job foi configurado como uma pipeline conectada ao repositório no GitHub.

Configuração utilizada:

```text
Tipo do job: Pipeline
Definition: Pipeline script from SCM
SCM: Git
Repository URL: https://github.com/LuisThiago858/jenkins-docker-ci-golang.git
Branch Specifier: */main
Script Path: Jenkinsfile
```

## Cenários da atividade

### Cenário 1 — Tudo certo

Neste cenário, o job é executado manualmente no Jenkins.

Resultado esperado:

```text
Build: sucesso
Testes: sucesso
Resultado da pipeline: SUCCESS
```

Neste cenário, o Jenkins executa o build em um container Docker e depois executa os testes em outro container Docker. Ambos finalizam com sucesso.

### Cenário 2 — Deu ruim

Neste cenário, é criada uma falha proposital de compilação no código-fonte.

Exemplo de erro utilizado:

```go
func (c ConverterService) ToCelsius(temperature Temperature) Temperature {
    if temperature.Unit() == Celsius {
        return temperature
    }

    value := (temperature.Value() - 32) * 5 / 9

    return NewTemperature(value, Celsius)
```

Nesse exemplo, uma chave de fechamento é removida, fazendo com que o código não compile.

Resultado esperado:

```text
Build: falha
Testes: não são executados corretamente
Resultado da pipeline: FAILURE
```

### Cenário 3 — Tá instável

Neste cenário, o código compila corretamente, mas uma regra de conversão é alterada de forma incorreta para provocar falha nos testes.

Exemplo de alteração:

```go
value := (temperature.Value() - 32) * 9 / 5
```

Com essa alteração, a compilação é concluída com sucesso, mas os testes automatizados identificam erro no comportamento da aplicação.

Resultado esperado:

```text
Build: sucesso
Testes: falham
Resultado da pipeline: UNSTABLE
```

### Cenário 4 — Nightly

Neste cenário, o Jenkins executa a pipeline automaticamente em horário agendado, sem intervenção manual.

A configuração utilizada no `Jenkinsfile` foi:

```groovy
triggers {
    cron('H 2 * * *')
}
```

Essa configuração agenda a execução da pipeline diariamente no período das 2h.

Para fins de demonstração em vídeo, pode ser utilizada temporariamente uma configuração com intervalo menor:

```groovy
triggers {
    cron('H/2 * * * *')
}
```

Essa configuração permite demonstrar a execução automática sem precisar aguardar o horário real do nightly.

Resultado esperado:

```text
Build: sucesso
Testes: sucesso
Resultado da pipeline: SUCCESS
```

## Evidência de containers separados

Durante a execução da pipeline, o Jenkins exibe no console o identificador do container utilizado em cada etapa.

Exemplo:

```text
Container de build:
1b161a3d5a1d

Container de testes:
bfde1f4ea6c3
```

Como os identificadores são diferentes, é possível comprovar que o build e os testes foram executados em containers Docker separados.

## Como visualizar os containers durante a execução

Durante a execução do job no Jenkins, é possível abrir um terminal PowerShell e executar:

```powershell
while ($true) {
    cls
    docker ps
    Start-Sleep -Seconds 2
}
```

Esse comando permite acompanhar os containers Docker ativos durante a execução da pipeline.

## Links dos vídeos

* Cenário 1 — Tudo certo: inserir link do vídeo aqui
* Cenário 2 — Deu ruim: inserir link do vídeo aqui
* Cenário 3 — Tá instável: inserir link do vídeo aqui
* Cenário 4 — Nightly: inserir link do vídeo aqui

## Conclusão

O projeto demonstrou a integração entre Jenkins, GitHub e Docker em uma pipeline de Integração Contínua.

A estratégia adotada permite que o build e os testes sejam executados em ambientes isolados e reproduzíveis, reduzindo a dependência de configurações locais da máquina. Dessa forma, o Jenkins atua como orquestrador do processo, enquanto o Docker fornece os ambientes necessários para compilação e validação da aplicação.
