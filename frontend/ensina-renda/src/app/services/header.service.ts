import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

export interface HeaderConfig {
  showLogin: boolean;
  showRegister: boolean;
  showProfile: boolean;
  showLogout: boolean;
  showNavigation: boolean;
}

@Injectable({
  providedIn: 'root'
})
export class HeaderService {
  private headerConfigSubject = new BehaviorSubject<HeaderConfig>({
    showLogin: true,
    showRegister: true,
    showProfile: false,
    showLogout: false,
    showNavigation: true
  });

  headerConfig$ = this.headerConfigSubject.asObservable();

  constructor() { }

  updateHeader(config: Partial<HeaderConfig>) {
    const currentConfig = this.headerConfigSubject.getValue();
    const newConfig = { ...currentConfig, ...config };
    this.headerConfigSubject.next(newConfig);
  }
}