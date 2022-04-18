Backend API documentation
Developers:
1. Kamal Sai Raj Kuncha
2. Yuva Roshit Maddipati
APIs: 
1.	User Management (Register, login, etc.)
2.	View/Add places
3.	View/Post/Edit Reviews

User Management API:
Description:
•	This API allows the user to register by entering their credentials and getting added to the database.
•	Once the user is registered, they can login into the website whenever, using their email address and password
•	Registering and logging in is mandatory for any user trying to post any new reviews for the visible locations on the site
Acceptance Criteria:
•	For registering, the following fields are required to be entered by the user: Name, Email, Phone and Password.
•	If any/all of these fields weren’t entered, the API throws an error message.
•	Once all the inputs are properly entered by the user, the API throws a confirmation message: “User added to the database”
•	If the same user tries registering again using the same credentials, the API throws the message: “User already exists”
•	Similarly for login, the API should throw an error message, if the user entered either a wrong email or password.
Sample Requests and Responses:
![image](https://user-images.githubusercontent.com/38933993/163862284-5f819808-3189-43fc-b5ba-0c8350378ad3.png)
![image](https://user-images.githubusercontent.com/38933993/163862304-33e669b8-bafc-4eda-8bc5-979d03c10cbe.png)
