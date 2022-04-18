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

API for Viewing/Adding places:
Description:
•	This API sends a get request whenever the user wishes to view all the places available to review.
•	A post request is sent whenever the user wishes to add a new place to the database. The following are the required parameters for sending a post request to add a new place: placename, location, type, avgrating.
Acceptance Criteria:
•	The API should throw an error message whenever any of the above mentioned required parameters aren’t specified.
•	Whenever a get request is sent the API should send a response with the list of places along with the values of all the other place parameters.
Sample Requests and Responses:
![image](https://user-images.githubusercontent.com/38933993/163862518-ac209298-2c44-498e-896d-5fed39dad9dc.png)
![image](https://user-images.githubusercontent.com/38933993/163862535-1f949a81-c555-4e07-bf2a-a0a72ed7ae05.png)

API for Viewing, Posting or Editing reviews:
Description:
•	Logged in users can add a review to a corresponding place. The API receives a POST request with the place id of the corresponding place.
•	Reviews from other users for all the places can be viewed
•	Previously posted reviews can be edited.
Acceptance Criteria:
•	Whenever the user is not logged in and tries to write a review, the API should throw an error message saying: “User not logged in”
•	The API should throw an error message when any of the following required fields are not filled, when trying to post a review: reviewtitle, review, rating, placeid, reviewerid
•	The review id should be included in the URL when trying to edit an old review.
•	Whenever a get request is sent, the API should send data having all the previously posted reviews that are stored on the database.
•	When a delete review request is sent to the API along with the review id, the API should remove the corresponding review from the database.

Sample Requests and Responses:
![image](https://user-images.githubusercontent.com/38933993/163862699-1c2b9460-aed5-4e20-a51e-ed49d976ab17.png)
![image](https://user-images.githubusercontent.com/38933993/163862720-7275fba5-a6eb-42d9-adb0-ebe9e8073861.png)

Sprint 3:
GetPlaceByIdView:
Description:
•	Logged in users can retrieve a place corresponding to an ID. The API receives a GET request with the place id of the corresponding place.
•	Place details like location, name, etc could be viewed.
Acceptance Criteria:
•	The place id should be included in the URL when trying to get the place details.
•	Whenever a get request is sent, the API should send data having all the place details from the database.
![image](https://user-images.githubusercontent.com/38933993/163862840-17944e3b-e9a2-43c7-8a07-edca911aa14c.png)

EditplaceView:
Description:
•	Logged in users can edit the place details corresponding to a place ID. The API receives a PATCH request with the place id of the corresponding place.
•	A patch request with a format similar to the create place request gets sent as a json file with changes made to whichever fields the user desires.
•	All the previously created places in the database can be edited.
Acceptance Criteria:
•	Whenever the user is not logged in and tries to edit a place, the API should throw an error message saying: “User not logged in”
•	The API should throw an error message when any of the following required fields are not filled, when trying to edit a place: placename, location, type, avgrating.
•	The place id should be included in the URL when trying to edit a previously created place.
•	Whenever a patch request is sent, the API should send back an appropriate acknowledgement message.
![image](https://user-images.githubusercontent.com/38933993/163863186-71ca5bbf-a085-4c20-97c2-25924017ed27.png)

DeleteplaceView:
Description:
•	Logged in users can delete an existing place from the database corresponding to a place ID. The API receives a DELETE request with the place id of the corresponding place.
•	Any of the previously created places in the database can be deleted by appending the place id to the APIs url.
Acceptance Criteria:
•	The API should throw an error message when the place id in the url is not valid.
•	The place id should be included in the URL when trying to delete a previously created place.
•	Whenever a delete request is sent, the API should send back an appropriate acknowledgement message.
![image](https://user-images.githubusercontent.com/38933993/163863281-3c0901f2-c091-4fad-9b07-48b041b31c69.png)




