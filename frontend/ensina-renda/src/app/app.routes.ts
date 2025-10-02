import { Routes } from '@angular/router';
import { LandingComponent } from './pages/landing/landing.component';
import { LoginComponent } from './pages/login/login.component';
import { SignupComponent } from './pages/signup/signup.component';

export const routes: Routes = [
    { path: '', title: 'Ensina Renda', component: LandingComponent },
    { path: 'signup', title: 'Cadastre-se', component: SignupComponent },
    { path: 'login', title: 'Bem-vindo de volta', component: LoginComponent },
];
