pipeline {
    agent any
    tools {
        go 'go-1.18'
    }
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Test') {
            steps {
                go test -v
            }
        }
        stage('build') {
            steps {
                go build
            }
        }
    }
}