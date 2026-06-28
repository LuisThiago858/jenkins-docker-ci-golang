pipeline {
    agent any

    triggers {
        cron('H/2 * * *')
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