# Sprint 4 

## Project Description:
Want to eat out but can’t decide where? Want to tour new places but can’t figure out what to do?

Gator-X is here to your rescue!

Exploring new places is fun but figuring out what to do can be perplexing. Gator-X is designed with the intention of elaborately expressing one's views about a café or museum. A user can easily browse through the reviews about a place that helps them to make an informed choice and thus enhance their experience. Gator-X stands for Gator-Experience! 

The primary location kept in mind while designing this website is Gainesville, the home of Gators. Though Gainesville is mostly seen as a town of students, it has a diverse population distribution catering needs of various age groups. It often gets tough for people to identify the places that suit their taste especially if they are new to the city. Gator-X provides an opportunity for visitors/residents to find new adventures and meet like-minded people in the process. A user can post, edit, delete and view reviews posted on the app. They can search and filter the posts according to different categories. Hence, Gator-X helps users save time and make the most of their time spent in Gator-land.

## Tasks accomplished in Sprint 4:
### 1. Frontend:
* **Added Upload image feature to reviews**: The user will now be able to add images in addition to the text, category of the place and the author name. This shows the users the real condition of a place versus the advertised pictures which will help users in making the right choice to pick a place.

* **Added SearchBar component to navbar**: The user will now be able to search to find exactly what they want. They need not waste their precious time in browsing.

* **Added Ratings feature**: The user will now be able to add rating to a review in the range of 1-5 which showcases the degree of the experience a user had at a place which correlates to the overall experience of a user including hospitality, customer service and the place offering.

* **End-to-end testing using cypress**: All the new features added to the frontend have been end-to-end tested and the tests can be found at [Tests](https://github.com/kamalsairajk/Gator-X/tree/main/frontend/cypress)

### 2. Backend:

* **Image upload feature**: Upload images to any particular review that a user creates or place which is created by admin. Images are saved in the server under separate folders which correspond to places and reviews. This feature is accessible when the user or admin creates a review or place respectively. Also, for space efficiency whenever a review or place is deleted then the file corresponding to it is also deleted. Also, edit review or place also changes the file if provided. This made changes on how data is exchanged as this involves form file and form data along with JSON.
* **Creation of Admin Layer** : Admin layer is incorporated so that places can be created by the admin layer after verification of the place. Along with creation, update and delete places have an admin verification layer so that all users cannot access it. Since admin users must be created a new function i.e register admin is created for this purpose. But other features of users such as login, logout, edit and delete which were created for earlier sprints are still useful in this layer.
* **Comments and some minor fixes**: All functions have a comment description which covers the code flow and structure. This is throughout the project structure. Minor fixes such as few ID related i.e ID not present, average rating in edit view and other validations.
* **Tests**: Several unit test cases have been provided to validate the code's reliability. Unit test cases have been developed for the new functionalities such as the image functionality, admin layer with the register admin feature and many more that have been added during this sprint. Since some functions were edited in order to incorporate the image functionality so the unit tests that were used for these functions are also edited to use this functionality.

## Link to API Documentation:
[API Doc](https://github.com/kamalsairajk/Gator-X/blob/main/backend/API%20Doc.MD)

## Link to project board: 
[Sprint 4 board](https://github.com/kamalsairajk/Gator-X/projects/4)

