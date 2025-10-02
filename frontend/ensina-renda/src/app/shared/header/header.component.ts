import { AsyncPipe, NgIf } from "@angular/common";
import { Component } from '@angular/core';
import { RouterLink } from "@angular/router";
import { Observable } from 'rxjs';
import { HeaderConfig, HeaderService } from '../../services/header.service';

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [NgIf, AsyncPipe, RouterLink],
  templateUrl: './header.component.html',
  styleUrl: './header.component.css'
})
export class HeaderComponent {
  headerConfig$: Observable<HeaderConfig>;

  constructor(private headerService: HeaderService) {
    this.headerConfig$ = this.headerService.headerConfig$;
   }

  ngOnInit(): void {
    this.headerConfig$ = this.headerService.headerConfig$;
  }
}
