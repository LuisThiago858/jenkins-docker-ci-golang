pipeline {
    agent none

    triggers {
        cron('H 2 * * *')
    }

    stages {
        stage('Checkout') {
            agent any

            steps {
                checkout scm
                stash name: 'source-code', includes: '**/*'
            }
        }

        stage('Build em container Docker') {
            agent {
                docker {
                    image 'golang:1.26'
                    args '-v go-mod-cache:/go/pkg/mod'
                }
            }

            steps {
                unstash 'source-code'

                sh '''
                    echo "Container de build:"
                    hostname
                    go version
                    go mod download
                    go build -v ./...
                '''

                stash name: 'build-output', includes: '**/*'
            }
        }

        stage('Testes em outro container Docker') {
            agent {
                docker {
                    image 'golang:1.26'
                    args '-v go-mod-cache:/go/pkg/mod'
                }
            }

            steps {
                unstash 'build-output'

                catchError(buildResult: 'UNSTABLE', stageResult: 'UNSTABLE') {
                    sh '''
                        echo "Container de testes:"
                        hostname
                        go version
                        go test -v ./...
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