pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'echo "Stage BUILD"'
                sh 'echo "Hello World"'
                sh '''
                    echo "Multiline shell steps works too"
                    ls -lah
                '''
            }
        }
        stage('Test') {
            steps {
                sh 'echo "Stage TEST"'
                sh 'pwd'
                sh 'ls -AlF'
            }
        }
    }
}
