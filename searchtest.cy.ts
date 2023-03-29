describe('Search', () => {

    it('searches properly', () => {
      cy.wait(1000)
      cy.visit('http://localhost:4200/search')
      cy.get("input[placeholder=\"Search many games\"]").type("All Games");
      cy.wait(1000)
      cy.contains('Search').click()
    })
  })