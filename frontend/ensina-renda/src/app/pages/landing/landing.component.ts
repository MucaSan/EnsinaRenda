import { Component } from '@angular/core';
import AOS from 'aos';
import { HeaderService } from '../../services/header.service';
import { FooterComponent } from '../../shared/footer/footer.component';
import { HeaderComponent } from '../../shared/header/header.component';
import { FaqComponent } from './components/faq/faq.component';

@Component({
  selector: 'app-landing',
  standalone: true,
  imports: [FaqComponent, FooterComponent, HeaderComponent],
  templateUrl: './landing.component.html',
  styleUrl: './landing.component.css'
})
export class LandingComponent {
  constructor(private headerService: HeaderService) { }

  ngOnInit() {
    this.headerService.updateHeader({
      showLogin: true,
      showRegister: true,
      showNavigation: true
    })

    AOS.init({
      duration: 800,
      once: false,
      easing: 'ease-in-out',
    });
  }
}
