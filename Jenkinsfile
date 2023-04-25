pipeline {
  // Run on an agent where we want to use Go
  agent any

  // Ensure the desired Go version is installed for all stages,
  // using the name defined in the Global Tool Configuration

  stages {
    stage('Build') {
      steps {
        // Output will be something like "go version go1.19 darwin/arm64"
        sh 'apt install golang-go'
        sh 'go version'
      }
    }
  }
}