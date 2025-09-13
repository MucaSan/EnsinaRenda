import { Component } from '@angular/core';
import { HeaderService } from '../../services/header.service';
import { HeaderComponent } from '../../shared/header/header.component';

@Component({
  selector: 'app-cadastro',
  standalone: true,
  imports: [HeaderComponent],
  templateUrl: './cadastro.component.html',
  styleUrl: './cadastro.component.css'
})
export class CadastroComponent {
  constructor(private HeaderService: HeaderService) {}

  ngOnInit() {
    this.HeaderService.updateHeader({
      showLogin: true,
      showRegister: true,
      showNavigation: false
    })
  }
}
