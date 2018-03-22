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
        stage('deploy') {
            steps {
               echo "deploying"
               sh 'go install ./...'
               echo "Finished deploying"
            }
        }
    }
}
