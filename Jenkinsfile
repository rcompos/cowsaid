pipeline {
    agent {
        docker { image 'node:14-alpine' }
    }
    stages {
        stage('Test') {
            steps {
                sh 'apt-get update && apt-get install -y docker.io'
                sh 'node --version'
            }
        }
    }
}
