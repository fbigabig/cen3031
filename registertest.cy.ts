describe('Register', () => {

  it('registers a user', () => {
    
    cy.visit('http://localhost:4200') // change URL to match your dev UR
    cy.visit('http://localhost:4200/register')
    cy.get("input[placeholder=\"Username\"]").type("pablo");
    cy.wait(1000)
    cy.get("input[placeholder=\"Password\"]").type("bueno");
    cy.wait(1000)
    cy.contains('Submit').click()
  })
})