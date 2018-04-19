pipeline {
    agent {
        docker { image 'elt/xmeter:jm3.2' }
    }
    stages {
        stage('Test') {
            steps {
                sh 'hostname'
                sh 'ls -lA /opt/jmeter'
            }
        }
    }
}
