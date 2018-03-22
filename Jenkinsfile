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
             dir('/gospace/src') {
               // some blocki
               checkout changelog: false, poll: false, scm: [$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: 'https://github.com/sanksons/reflorest-testapp.git']]]
             }
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
