pipeline {
    environment { 
        registry = "rcompos/cowsaid" 
        //registryCredential = 'dockerhub_id' 
        dockerImage = '' 
    }

    agent any
    //agent { dockerfile true }

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

            //agent {
            //    // Equivalent to "docker build -f Dockerfile.build --build-arg version=1.0.2 ./build/
            //    dockerfile {
            //        filename 'Dockerfile.build'
            //        dir 'build'
            //        label 'my-defined-label'
            //        additionalBuildArgs  '--build-arg version=1.0.2'
            //        args '-v /tmp:/tmp'
            //    }
            //}

            //agent {
            //    // Equivalent to "docker build -f Dockerfile ."
            //    dockerfile {
            //        filename 'Dockerfile'
            //        dir '.'
            //    }
            //}

                script { 
                    dockerImage = docker.build registry + ":$BUILD_NUMBER" 
                }

                sh 'echo "***  STAGE BUILD  ***"'
                sh '''
                    echo "Beginning build"
                    pwd
                    ls -lah
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
