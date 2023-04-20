import { LoginComponent } from "./login.component"
import { HttpClientModule } from "@angular/common/http";



describe('LoginComponent', () => {

    it('can mount', () => {
        cy.mount(LoginComponent, {
            imports: [HttpClientModule],
            declarations: [LoginComponent],
            providers: []
        })
    })

    it('can load text', () => {
        cy.mount(LoginComponent, {
            imports: [HttpClientModule],
            declarations: [LoginComponent],
            providers: []
        })
        cy.contains('Enter Username:');
    })

    it('can detect the username box', () => {
        cy.mount(LoginComponent, {
            imports: [HttpClientModule],
            declarations: [LoginComponent],
            providers: []
        })
        cy.wait(1000)
        cy.get("input[placeholder=\"Username\"]");
    })

    it('can detect the password box', () => {
        cy.mount(LoginComponent, {
            imports: [HttpClientModule],
            declarations: [LoginComponent],
            providers: []
        })
        cy.wait(1000)
        cy.get("input[placeholder=\"Password\"]");
    })

    it('can sign in', () => {
        cy.mount(LoginComponent, {
            imports: [HttpClientModule],
            declarations: [LoginComponent],
            providers: []
        })
        cy.wait(500)
        cy.get("input[placeholder=\"Username\"]").type("pablo");
        cy.wait(500)
        cy.get("input[placeholder=\"Password\"]").type("bueno");
        cy.wait(500)
        cy.contains('Sign in').click();
    })
})