import { Routes } from '@angular/router';
import { CadastroComponent } from './pages/cadastro/cadastro.component';
import { LandingComponent } from './pages/landing/landing.component';

export const routes: Routes = [
    { path: '', title: 'Ensina Renda', component: LandingComponent },
    { path: 'cadastro', title: 'Cadastre-se', component: CadastroComponent },
];
