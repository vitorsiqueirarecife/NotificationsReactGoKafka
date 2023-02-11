# NotificationsReactGoKafka

Example messaging system using react js, golang, docker and kafka. The system's proposal is to simulate the sending of notifications based on categories selected by the user.

The architecture uses kafka as a messaging system, receiving messages from BFF(API) in 3 topics that are read by 3 different services (SMS, EMAIL and PUSH NOTIFICATION).

The frontend uses react, react hook forms, react query, styled components and styled systems.

Open 5 terminals and
- run in ./backend: docker-compose up
- run in ./backend/bff: go run main.go
- run in ./backend/sender: : go run main.go sms
- run in ./backend/sender: : go run main.go email
- run in ./backend/sender: : go run main.go push
- run in ./frontend/webapp: npm run start

The logs are: ./backend/sender/logs/