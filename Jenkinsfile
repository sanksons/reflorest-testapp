pipeline {
    agent any
    environment {
        // Some environment variables to be used later
        ARTIFACTPATH = "output"
        OUTPUT = "bundle.tar.gz"
    }
    stages {
	    stage ('checkout') {
		  steps {
                git branch: 'master', url: 'https://github.com/sanksons/reflorest-testapp.git'
            }
		}
	    stage('prebuild') {
		   agent {
                docker {
	               image 'sanksons/go1.9-ubuntu:initial'
		           args '-u 0:0 -v /home/sab/.jenkins/workspace/reflotest-test2:/gospace/src/github.com/sanksons/reflorest-testapp:rw,z  -w /gospace/src/github.com/sanksons/reflorest-testapp'
		           reuseNode true
	            }
            }
            steps {    
                   	sh 'pwd'
				     sh 'ls -lahrt'
					 // need to see why dir not works??
					 sh 'cd /gospace/src/github.com/sanksons/reflorest-testapp && reflorest deploy'
					 sh 'pwd'
					 echo "See whats in bin"
					 sh 'ls -lahrt /gospace/bin'
					 sh 'rm -rf $ARTIFACTPATH'
                     sh 'mkdir -p $ARTIFACTPATH'
					 sh 'cp -r /gospace/bin/conf $ARTIFACTPATH/ && cp /gospace/bin/reflorest-testapp $ARTIFACTPATH/'
					 
				   }
		}
		stage('bundle') {
		   sh 'cd output && tar czf $OUTPUT *'
		   archiveArtifacts "${env.ARTIFACTPATH}/*"
		}
		
        
    }
}
