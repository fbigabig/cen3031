describe('Logout', () => {

    it('properly visits the logout screen', () => {
        cy.wait(500);
        cy.visit('http://localhost:4200/logout');
    })

    it('can click logout button properly', () => {
        cy.wait(500);
        cy.visit('http://localhost:4200/logout')
        cy.wait(500);
        cy.contains('Yes').click();

    })

    it('can be properly logged out', () => {
        cy.wait(500);
        cy.visit('http://localhost:4200/logout')
        cy.wait(500);
        cy.contains('Yes').click();
        cy.contains('You are not logged in');
    })

})