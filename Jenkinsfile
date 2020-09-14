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
  - name: maven
    image: maven:latest
    command:
    - cat
    tty: true
    //volumeMounts:
    //  - mountPath: "/root/.m2"
    //    name: m2
  - name: docker
    image: docker:latest
    command:
    - cat
    tty: true
    volumeMounts:
    - mountPath: /var/run/docker.sock
      name: docker-sock
  volumes:
    - name: docker-sock
      hostPath:
        path: /var/run/docker.sock
    //- name: m2
    //  persistentVolumeClaim:
    //    claimName: m2
"""
}
   }
  stages {
    stage('Build') {
      steps {
        container('maven') {
          sh """
                        mvn package -DskipTests
                                                """
        }
      }
    }
    stage('Test') {
      steps {
        container('maven') {
          sh """
             mvn test
          """
        }
      }
    }
    stage('Push') {
      steps {
        container('docker') {
          sh """
             docker build -t spring-petclinic-demo:$BUILD_NUMBER .
          """
        }
      }
    }
  }
}
