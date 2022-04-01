The detailed description of the activities in the Sprint-2 are recorded below: 

Frontend Team(Anviksha Sharma & Sai Siddharth Upadhyayula):

1) Frontend-Backend Integration:
The frontend team was able to modify the User Interaction (UI) aspect of the website by integrating the frontend React code with the Golang Backend through the use of Axios library and rendered each function to the backend.
 
2) Cypress Tests:
The frontend team unit tested the working of various small components of the frontend which includes the support, authentication of the user, checking the routing of the responses through fixtures and check the working of the plugins which the website uses to incorporate various small cases to perform a bigger operation and deliver to the user by using Cypress.

Cypress Tests: https://github.com/kamalsairajk/Gator-X/tree/main/frontend/cypress

Backend Team(Kamal Sai Raj K & Yuva Roshith Maddipati):

1)	User Management:
API created to handle users’ data. User can register onto the app by entering the required credentials like name, email, password, phone. The API receives a post request with the data and the user gets created and their data is added to the database.
  
2)	Session Handling:
Registered users can now login and stay in session. The API gets a request with the user’s email and password and a session is created, once the credentials match. Similarly, a logout request could be sent to end the session. The session data with the relevant login and logout times gets stored on the database.  

3)	Tests:
Several unit test cases have been added to check the reliability of the code. These tests are always expected to pass whenever any new changes are made to the backend code.

4)	Session Validation layers added to relevant pages/functionalities:
Parts of the application that require a user to be logged in for access have been added an extra layer of code to verify the current session to check if the user is logged in. For example: Users can view reviews as a guest but login is required if they need to add/edit a review for a certain place.
 
Backend API Doc: https://github.com/kamalsairajk/Gator-X/blob/main/backend/Backend%20API%20Doc.pdf
