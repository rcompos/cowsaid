pipeline {
    agent any
    stages {
        stage('setenv') {
          steps {
            script {
              env.GIT_BRANCH_NAME=sh(returnStdout: true, script: "git rev-parse --abbrev-ref HEAD").trim()
              env.GIT_REF=sh(returnStdout: true, script: "git rev-parse HEAD").trim()
            }
          }
        }
        stage('build') {
            steps {
                sh 'echo "***  STAGE BUILD  ***"'
                sh '''
                    echo "Beginning build"
                    pwd
                    ls -lah
                    which docker
                    docker build -t rcompos/cowsaid:${env.GIT_REF} .
                '''
            }
        }
        stage('test') {
            steps {
                sh 'echo "***  STAGE TEST  ***"'
                sh 'ls -AlF'
            }
        }
        stage('publish') {
            steps {
                sh 'echo "***  STAGE PUBLISH  ***"'
            }
        }
    }
}
