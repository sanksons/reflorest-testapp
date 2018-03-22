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
		     echo "Mkdir go/src/github.com/sanksons"
			 sh 'mkdir -p go/src/github.com/sanksons'
             dir('go/src/github.com/sanksons') {
               // take checkout
               checkout changelog: false, poll: false, scm: [$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: 'https://github.com/sanksons/reflorest-testapp.git']]]
             }
			 dir('go/src/github.com/sanksons/reflorest-testapp') {
			   sh 'go install ./...'
			 }
           }
        }
        
    }
}
