pipeline {
  agent any
  environment {
    serverContainer = "todo-server"
    frontendContainer = "todo-frontend"
    nginxContainer = "todo-nginx"

    serverRegistry = "arigatopix/todoapp-server:latest"
    frontendRegistry = "arigatopix/todoapp-frontend:latest"
    nginxRegistry = "arigatopix/todoapp-nginx:latest"

    prodCredentials = "prod-credentials"
    server = "ubuntu@ec2-18-141-24-224.ap-southeast-1.compute.amazonaws.com"
    DOCKERHUB_CREDENTIALS = credentials('dockerhub-cred-arigatopix')
  }

  stages {
    stage('clean old image') {
      steps {
        sh 'docker rm -f ' + serverContainer
        sh 'docker rm -f ' + frontendContainer
        sh 'docker rm -f ' + nginxContainer
        sh 'docker image rm -f ' + serverRegistry 
        sh 'docker image rm -f ' + frontendRegistry
        sh 'docker image rm -f ' + nginxRegistry
      }
    }

    stage('build images latest') {
      steps {
        sh 'docker build -t ' + serverRegistry + ' -f server/Dockerfile .'
        sh 'docker build -t ' + frontendRegistry + ' -f frontend/Dockerfile .'
        sh 'docker build -t ' + nginxRegistry + ' -f nginx/Dockerfile .'
      }
    }

    stage('login dockerhub') {
      steps {
        sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
      }
    }  

    stage('push to docker hub') {
      steps {
				sh 'docker push ' + serverRegistry
				sh 'docker push ' + frontendRegistry
				sh 'docker push ' + nginxRegistry
			}
    }

    stage('copy docker-compose to production') {
      steps {
        sshagent([prodCredentials]) {
          sh 'ssh -o StrictHostKeyChecking=no ' + server + ' sudo docker compose -f /home/ubuntu/app/tasks-app/docker-compose.yml down'
          sh 'scp -o StrictHostKeyChecking=no docker-compose.yml ' + server + ':/home/ubuntu/app/tasks-app/docker-compose.yml'
          sh 'ssh -o StrictHostKeyChecking=no ' + server + ' sudo docker compose -f /home/ubuntu/app/tasks-app/docker-compose.yml pull'
          sh 'ssh -o StrictHostKeyChecking=no ' + server + ' sudo docker compose -f /home/ubuntu/app/tasks-app/docker-compose.yml up -d'
        }
      }
    }

    stage('clean up from jenkins') {
      steps {
        sh 'docker rm -f ' + serverContainer
        sh 'docker rm -f ' + frontendContainer
        sh 'docker rm -f ' + nginxContainer
        sh 'docker image rm -f ' + serverRegistry 
        sh 'docker image rm -f ' + frontendRegistry
        sh 'docker image rm -f ' + nginxRegistry
      }
    }
    
  }

  post {
		always {
			sh 'docker logout'
		}
	}
}