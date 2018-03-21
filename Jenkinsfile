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
               sh 'reflorest deploy'
            }
        }
    }
}
