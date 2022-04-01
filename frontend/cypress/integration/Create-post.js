describe('Post creation', function () {

    it("should create the post", function () {
        cy.visit('https://localhost:9000/')
    })
    
    it("should enable the user create the post", funtion () {
        cy.get('#create').click()
    })
})