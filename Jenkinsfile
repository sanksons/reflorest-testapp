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
			 sh 'mkdir -m 777 -p go/src/github.com/sanksons'
			 sh 'mkdir -m 777 -p go/src/github.com/sanksons@tmp'
             dir('go/src/github.com/sanksons') {
               // take checkout
			   sh 'git clone https://github.com/sanksons/reflorest-testapp.git'
			   
			 }
			 dir('go/src/github.com/sanksons/reflorest-testapp') {
			   sh 'mkdir -m 777 -p reflorest-testapp@tmp'
			   sh 'git checkout master'
			   sh 'go install ./...'
			 }
			 dir('go') {
             }			 
           }
        }
        
    }
	post {
	  always {
            echo 'One way or another, I have finished'
            deleteDir() /* clean up our workspace */
      }
	}
}
