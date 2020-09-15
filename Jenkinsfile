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
    stage('Push') {
      steps {
        container('docker') {
          sh """
             #docker build -t spring-petclinic-demo:$BUILD_NUMBER .
             docker images
          """
        }
      }
    }
    stage('Deploy') {
      steps {
        container('testing') {
          sh """
             echo "Deploying with ArgoCD"
             #argocd version
          """
        }
      }
    }
  }
}
