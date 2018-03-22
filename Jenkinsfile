pipeline {
    agent {
      dockerfile {
        args  '-u 0:0'
      }
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
              //sh 'mkdir /gospace/src'
              sh 'mkdir /gospace/src/myapplication'
              echo "copy application"
              sh 'cp -r . /gospace/src/myapplication/*'
              echo "change cwd"
              sh 'cd /gospace/src/myapplication'
              echo "PWD:"
              sh 'pwd'
              //sh 'export GOPATH=$(pwd)'
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
