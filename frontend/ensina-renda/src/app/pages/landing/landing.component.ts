import { Component } from '@angular/core';
import { FaqComponent } from './components/faq/faq.component';

@Component({
  selector: 'app-landing',
  standalone: true,
  imports: [FaqComponent],
  templateUrl: './landing.component.html',
  styleUrl: './landing.component.css'
})
export class LandingComponent {

}
