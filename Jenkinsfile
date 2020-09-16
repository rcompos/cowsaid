pipeline {
  agent {
    kubernetes {
      label 'argocd-client-alpha'
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
  - name: ubuntu-argocd
    image: rcompos/ubuntu-argocd:latest
    command:
    - cat
    tty: true
"""
}
  }
  stages {
    stage('Build') {
      steps {
        container('cowsaid') {
          checkout scm
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
        container('ubuntu-argocd') {
          sh """
             echo "Nothing can stop our deploying now!"
			 ls -AlF /usr/local/bin/argocd 
          """
        }
      }
    }
  }
}
