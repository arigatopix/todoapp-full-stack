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


    stage('login dockerhub') {
      steps {
        sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
      }
    }  

    stage('build images latest') {
      steps {
        sh 'docker buildx build --push -t ' + serverRegistry + ' --platform=linux/amd64 -f server/Dockerfile .'
        sh 'docker buildx build --push -t ' + frontendRegistry + ' --platform=linux/amd64 -f frontend/Dockerfile .'
        sh 'docker buildx build --push -t ' + nginxRegistry + ' --platform=linux/amd64 -f nginx/Dockerfile .'
      }
    }

    stage('copy docker-compose to production') {
      steps {
        sshagent([prodCredentials]) {
          sh 'ssh -o StrictHostKeyChecking=no ' + server + ' mkdir -p app/todoapp'
          sh 'ssh -o StrictHostKeyChecking=no ' + server + ' sudo docker compose -f /home/ubuntu/app/todoapp/docker-compose.yml down'
          sh 'scp -o StrictHostKeyChecking=no docker-compose.yml ' + server + ':/home/ubuntu/app/todoapp/docker-compose.yml'
          sh 'ssh -o StrictHostKeyChecking=no ' + server + ' sudo docker compose -f /home/ubuntu/app/todoapp/docker-compose.yml pull'
          sh 'ssh -o StrictHostKeyChecking=no ' + server + ' sudo docker compose -f /home/ubuntu/app/todoapp/docker-compose.yml up -d'
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