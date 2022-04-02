Frontend Team:

1) Anviksha Sharma 
 
2) Sai Siddharth Upadhyayula



Frontend Updates:

In addition to the existing features, we have added features which include:

1) Create and delete post.
2) Display existing posts on the posts page.
3) The category of the review, author name and review creation date are also shown along with the review.

We have successfully run cypress tests on all the newly added features.

The pictures below show the added features and the cypress tests:

1) The posts page:

![Unknown](https://user-images.githubusercontent.com/46457398/161357855-517d2be6-88be-41b0-b3b0-3b721099fcd2.png)

2) Listing all the reviews 
 
![Unknown-1](https://user-images.githubusercontent.com/46457398/161357868-f06c2eb5-1a85-45b9-b9ef-e59aecf7d08d.png)

3) Create posts
 
![Unknown-2](https://user-images.githubusercontent.com/46457398/161357873-5788f95b-5836-4b26-8eb7-203db24be038.png)

4) Delete post
 
![46f60973-4eeb-420c-b814-2b6721464ab1](https://user-images.githubusercontent.com/46457398/161358441-176a4c6e-5fb6-4fee-90da-e0d254140b58.jpg)

5) Confirm Deletion

![d1ed621d-b3f1-4328-a3fa-66e5837310cb](https://user-images.githubusercontent.com/46457398/161358478-9a919bd3-82cf-4f6f-969f-64fd97427457.jpg)

6) Update the review
 
![7c3cc465-b3fd-40e2-adf5-735e0ebddc86](https://user-images.githubusercontent.com/46457398/161358811-04b929ef-5dbe-4ba6-a568-22dbade3ba33.jpg)

7) Succesful Updation
 
![a6fe384c-024a-4093-aa08-a4d7257619a2](https://user-images.githubusercontent.com/46457398/161358843-5da0775c-95f2-4783-ba9e-63df06176359.jpg)


8) Cypress Tests for the added features

![Unknown-3](https://user-images.githubusercontent.com/46457398/161357943-776523a6-aaad-4987-b6ba-7460164905c2.png)


Backend Team:

1) Kamal Sai Raj K
 
2) Yuva Roshit Maddipati



Backend Updates:

The following additional views have been added:
1. GetPlacebyIDView
2. EditplaceView
3. DeleteplaceView
4. GetreviewsbyuserView
5. GetreviewsbyplaceView

1) GetPlacebyIDView:
•	Logged in users can retrieve a place corresponding to an ID. The API receives a GET request with the place id of the corresponding place.
•	Place details like location, name, etc could be viewed.
![image](https://user-images.githubusercontent.com/38933993/161362920-6ab90ea6-a201-49c8-a1a7-deb40a25b81e.png)

2) EditplaceView
•	Logged in users can edit the place details corresponding to a place ID. The API receives a PATCH request with the place id of the corresponding place.
•	A patch request with a format similar to the create place request gets sent as a json file with changes made to whichever fields the user desires.
•	All the previously created places in the database can be edited.
![image](https://user-images.githubusercontent.com/38933993/161362950-be31adab-f8c2-4673-a9db-eda4b23400cf.png)

3) DeleteplaceView
•	Logged in users can delete an existing place from the database corresponding to a place ID. The API receives a DELETE request with the place id of the corresponding place.
•	Any of the previously created places in the database can be deleted by appending the place id to the APIs url.
![image](https://user-images.githubusercontent.com/38933993/161362964-85f785f5-a7db-49d3-a8c6-7bb020095924.png)

4) GetreviewsbyuserView
•	Logged in users can retrieve all reviews for all places. The API receives a GET request and send all the reviews data stored in the database
•	Review details like location, name, place name, title and rating could be viewed.
•	All the previously posted reviews in the database can be viewed.
![image](https://user-images.githubusercontent.com/38933993/161362980-9e532d8c-58e9-496f-88eb-03a41f810822.png)

5) GetreviewsbyplaceView
•	Logged in users can retrieve all reviews for a particular place corresponding to an ID. The API receives a GET request along with the place id and sends all the reviews data stored in the database for that specific place.
•	Review details like location, name, place name, title and rating could be viewed.
•	All the previously posted reviews for the corresponding place id in the database can be viewed.
![image](https://user-images.githubusercontent.com/38933993/161362993-0934c467-e4f7-43d8-b989-f94e510a2534.png)
