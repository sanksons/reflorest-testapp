pipeline {
    agent {
      dockerfile true
    }
    options {
       skipDefaultCheckout()
    }
    stages {
        stage('test') {
            steps {
                echo "Running tests"
            }
        }
        stage('prebuild') {
            steps {
              echo "mkdir /gospace/src"
              sh 'mkdir /gospace/src'
              echo "copy application"
              sh 'cp -r . /gospace/src/myapplication/*'
              echo "change cwd"
               
              echo "CWD:"
              sh 'export GOPATH=$(pwd)'
              echo "gopath is : $GOPATH"
            }
        }
        stage('deploy') {
            steps {
               echo "deploying"
               sh 'go install ./...'
               echo "Finished deploying"
            }
        }
    }
}
