describe('Database', () => {

    it('properly visits the database', () => {
        cy.wait(500);
        cy.visit('http://localhost:4200/database');
    })

    it('can load games', () => {
        cy.wait(500);
        cy.visit('http://localhost:4200/database');
        cy.wait(500);
        cy.contains('Portal 2');

    })

})