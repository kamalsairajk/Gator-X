describe('Create Review', function () {

    it("successfully opens homepage", function(){
        cy.visit('localhost:3000')
    })

    it("creates new review", function () {
        
        cy.contains('Add New Review').click()
        cy.url().should('include','/create')
        cy.get('.form-control1').type('fake title').should('have.value', 'fake title')
        cy.contains('Create').click()
    })
    
})

describe('Delete Review', function(){

    it("successfully deletes a review", function(){
        
    cy.visit('localhost:3000')
    cy.contains('Delete').click()
    cy.contains('Yes, delete').click()
    cy.contains('OK').click()

    })
})

describe('Edit Review', function(){

    it("successfully edit a review", function(){
        
    cy.visit('localhost:3000')
    cy.contains('Edit').click()
    cy.url().should('include','/edit')
    cy.get('.form-control1').type('fake title')
    cy.contains('Save changes').click()
    cy.contains('OK').click()

    })
})

describe('Display all reviews', function(){

    it("successfully displays list of reviews", function(){
        
    cy.visit('localhost:3000/create')
    cy.contains('List All Reviews').click()
   
    })
})

describe('Display single review', function(){

    it("successfully displays all details of a single review", function(){
        
       cy.contains('Show').click()
       cy.url().should('include','/post')
   
    })

    it("contains time stamp", function(){
        
        cy.contains('Create At')
       
    
     })
})

