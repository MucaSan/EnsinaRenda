import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

export interface FooterConfig { 
  showGoTop: boolean;
}

@Injectable({
  providedIn: 'root'
})
export class FooterService {
  private footerConfigSubject = new BehaviorSubject<FooterConfig>({
    showGoTop: true
  });

  footerConfig$ = this.footerConfigSubject.asObservable();

  constructor() { }

  updateFooter(config: Partial<FooterConfig>) {
    const currentConfig = this.footerConfigSubject.getValue();
    const newConfig = { ...currentConfig, ...config };
    this.footerConfigSubject.next(newConfig);
  }
}
