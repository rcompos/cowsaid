pipeline {
  agent {
    kubernetes {
      label 'spring-petclinic-demo'
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
  - name: boulder 
    image: docker:latest
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
             echo "Looking good..."
          """
        }
      }
    }
    stage('Deploy') {
      steps {
        container('boulder') {
          sh """
             echo "Hello. We are deploying now."
          """
        }
      }
    }
  }
}
