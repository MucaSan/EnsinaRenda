import { Component } from '@angular/core';
import AOS from 'aos';
import { FooterComponent } from '../../shared/footer/footer.component';
import { FaqComponent } from './components/faq/faq.component';

@Component({
  selector: 'app-landing',
  standalone: true,
  imports: [FaqComponent, FooterComponent],
  templateUrl: './landing.component.html',
  styleUrl: './landing.component.css'
})
export class LandingComponent {
  ngOnInit() {
    AOS.init({
      duration: 800,
      once: false,
      easing: 'ease-in-out',
    });
  }
}
