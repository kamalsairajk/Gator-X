Front-end Documentation:

Gator X:

Gator X is meant is to be an on-point guide for the UF students where they get accurate reviews which are close to reality. For the first sprint, we have created a basic skeleton-like frontend structure which includes three pages. The technologies used in building the client-side application are React.js, HTML, CSS and JavaScript.

1st page:

The first page is the home page of the website which contains the name of the website and the featured reviews of various places in Gainesville which will be added as one of the features in the upcoming sprints. The page contains three links which are Home, Register and Login. The Home link directs the user to the home page. The register link allows the user to create an account in the website. The login link allows the user to login the website and add any reviews if he wishes to. 

2nd page:

The second page is the register page for the user where he has to provide details like his name, email address and a password for the account. It can be observed that the user cannot create an account if he does not include ‘@’ symbol and ‘.com’ in the email address he is providing.

3rd page:

The third page contains the login links for the user. The third page allows the website to authenticate the user so that the user’s account is secure and not misused. If the user logs into the website, the user will be able to react to the reviews so that the reviews replicate the reality and the user will be able to add any reviews in addition to reading the reviews to ensure he gets the best experience.

Here is a sample video of the User Interaction with the website:

https://user-images.githubusercontent.com/46457398/152624104-eb63f9f5-9283-4494-a50b-a35ff8cd8db7.mp4

Back-end Documentation:

For the first sprint, we have generated models for various places that will be displayed and reviews for one particular user for a given place using the sqllite3 as database. Gorm for the database connection and Gin for the web framework in Golang. We have developed CRUD functionality for one particular user reviews for any place and place generation and deletion. We have to add various other validations and features along the way but currently these functionality works at a basic sense.

Also the REST api is generated for the present functionality which can be utilised by the frontend in order to retrieve results. User management will be covered in the upcoming sprint this would also include the session management and many more.Let me explain the files and file structure. 

First comes models folder which consists of base file with 2 models places and base reviews for a user. Next, is the views folder which consists of views pertaining to places and reviews. These views covers the CRUD functionality. And the main file which acts as a controller consisting of middleware, server execution, database generation or usage and API for the various functionality. 

Here is a sample video of the backend working with the database and generation of JSON objects for frontend to work with demonstrated in POSTMAN.


https://user-images.githubusercontent.com/30821679/152626664-24c57efd-5add-4bda-a6b9-41aa408a46cf.mp4



