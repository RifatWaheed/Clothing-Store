import { Injectable, signal, inject, PLATFORM_ID } from '@angular/core';
import { isPlatformBrowser } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { tap } from 'rxjs/operators';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private http = inject(HttpClient);
  private platformId = inject(PLATFORM_ID);
  private isBrowser = isPlatformBrowser(this.platformId);

  isLoggedIn = signal(this.isBrowser ? !!localStorage.getItem('token') : false);

  login(email: string, password: string) {
    return this.http.post<{ token: string }>('/api/auth/login', { email, password }).pipe(
      tap(res => {
        if (this.isBrowser) localStorage.setItem('token', res.token);
        this.isLoggedIn.set(true);
      })
    );
  }

  logout() {
    if (this.isBrowser) localStorage.removeItem('token');
    this.isLoggedIn.set(false);
  }

  getToken(): string | null {
    return this.isBrowser ? localStorage.getItem('token') : null;
  }
}
