describe('Navigation', () => {
  
    it('should navigate to the home page', () => {
      cy.visit('http://localhost:9000/')
    })
    it('should navigate to login page',() => {
      cy.get('a[href*="login"]').click({force: true})
      cy.url().should('include', '/login')
    })
    it('should navigate to register page',() => {
      cy.visit('http://localhost:9000/')
      cy.get('a[href*="register"]').click({force: true})
      cy.url().should('include', '/register')
      
    })
    
  })