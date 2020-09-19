pipeline {
  agent {
    kubernetes {
      label 'argocd-client-beta'
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
             echo "Tesing:  success"
          """
        }
      }
    }
    stage('Deploy') {
      steps {
        container('ubuntu-argocd') {
          sh """
             echo "Checking if nothing can stop our deploying now?"
          """
          withKubeConfig(caCertificate: '', clusterName: '', contextName: '', credentialsId: 'jenkins-kubernetes-cli', namespace: 'cowsaid', serverUrl: 'https://kubernetes.default') {
            // some block
            sh "kubectl version"
          }
          sh """
             echo "Nothing can stop our deploying now!"
             ls -AlF /usr/local/bin/argocd 
          """
        }
      }
    }
  }
}
