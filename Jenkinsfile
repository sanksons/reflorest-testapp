pipeline {
    agent {
      dockerfile true
    }
    stages {
        stage('test') {
            steps {
                echo "Running tests"
            }
        }
        stage('debug') {
            steps {
              echo "CWD:"
              sh 'pwd'
              echo "user:"
              sh 'whoami'
              echo "files:"
              sh 'ls -lahrt' 
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
