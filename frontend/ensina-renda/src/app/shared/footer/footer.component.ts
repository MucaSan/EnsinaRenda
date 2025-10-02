import { AsyncPipe, NgIf } from '@angular/common';
import { Component } from '@angular/core';
import { RouterLink } from '@angular/router';
import { Observable } from 'rxjs';
import { FooterConfig, FooterService } from '../../services/footer.service';

@Component({
  selector: 'app-footer',
  standalone: true,
  imports: [NgIf, AsyncPipe, RouterLink],
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.css']
})
export class FooterComponent {
  footerConfig$ = new Observable<FooterConfig>();

  constructor(private footerService: FooterService) {
    this.footerConfig$ = this.footerService.footerConfig$;
  }

  ngOnInit(): void {
    this.footerConfig$ = this.footerService.footerConfig$;
  }
}
