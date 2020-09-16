pipeline {
  agent {
    kubernetes {
      label 'argocd-client'
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
  - name: joulder 
    image: rcompos/ubuntu-argocd:latest
    command:
    - cat
    tty: true
    stdin: true
    stdout: true
    stderr: true
"""
}
  }
  stages {
    stage('Build') {
      steps {
        container('cowsaid') {
          sh """
              echo "Greetings and salutations!"          
          """
        }
      }
    }
    stage('Test') {
      steps {
        container('cowsaid') {
          sh """
             echo "Looking good 6:50"
          """
        }
      }
    }
    stage('Deploy') {
      steps {
        container('joulder') {
          sh """
             echo "Nothing can stop our deploying now!"
			 /usr/local/bin/argocd version
          """
        }
      }
    }
  }
}
