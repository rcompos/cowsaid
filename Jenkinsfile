pipeline {
  agent {
    kubernetes {
      label 'cowsaid-agent-00003'
      defaultContainer 'jnlp'
      yaml """
apiVersion: v1
kind: Pod
metadata:
labels:
  component: ci
spec:
  # Use service account that can deploy to all namespaces
  serviceAccountName: jenkins
  containers:
  - name: cowsaid 
    image: rcompos/cowsaid:latest
  - name: docker
    image: docker:latest
    command:
    - cat
    tty: true
    volumeMounts:
    - mountPath: /var/run/docker.sock
      name: docker-sock
  - name: testing
    image: rcompos/ubuntu-argocd:latest
    command:
    - cat
    tty: true
  volumes:
    - name: docker-sock
      hostPath:
        path: /var/run/docker.sock
"""
}
   }
  stages {
    stage('Pre') {
      steps {
        //checkout scm
        script {
          env.GIT_BRANCH_NAME=sh(returnStdout: true, script: "git rev-parse --abbrev-ref HEAD").trim()
          env.GIT_REF=sh(returnStdout: true, script: "git rev-parse HEAD").trim()
        }
        //slackSend (
        //  message: "Deploy (${env.GIT_BRANCH_NAME}: ${env.GIT_REF}) - Starting deploy...",
        //  color: "#6067f1"
        //)
      }
    }
    stage('Build') {
      steps {
        container('docker') {
          sh """
             echo "Build docker temp image"          
             docker version
             #docker tag hellobloom/my-image.temp hellobloom/my-image.temp:$GIT_REF
             echo "Push docker temp image"          
             #docker image push hellobloom/my-image.temp:$GIT_REF
          """
        }
      }
    }
    stage('Test') {
      steps {
        container('cowsaid') {
          sh """
             echo "Looking good..."
          """
        }
      }
    }
    stage('Deploy') {
      steps {
        container('docker') {
          sh """
             echo "Tag image"
             #docker tag hellobloom/my-image hellobloom/my-image:$GIT_REF
             echo "Push image"
             #docker push hellobloom/my-image:$GIT_REF
          """
        }
      }
    }
    stage('Post') {
      steps {
        container('docker') {
          sh """
             echo "Delete temp images."
             #docker image rm hellobloom/my-image.temp:$GIT_REF
          """
        }
      }
    }
  }
}
