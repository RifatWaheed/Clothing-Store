import { Injectable, signal, computed, inject, PLATFORM_ID } from '@angular/core';
import { isPlatformBrowser } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { tap } from 'rxjs/operators';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private http = inject(HttpClient);
  private platformId = inject(PLATFORM_ID);
  private isBrowser = isPlatformBrowser(this.platformId);

  isLoggedIn = signal(this.isBrowser ? !!localStorage.getItem('token') : false);
  isAdmin = computed(() => this.isLoggedIn() && this.getRole() === 'admin');

  login(email: string, password: string) {
    return this.http.post<{ token: string; role: string }>('/api/auth/login', { email, password }).pipe(
      tap(res => {
        if (this.isBrowser) localStorage.setItem('token', res.token);
        this.isLoggedIn.set(true);
      })
    );
  }

  sendOTP(email: string) {
    return this.http.post('/api/auth/request-sendOTP', { email });
  }

  validateOTP(email: string, otp: string) {
    return this.http.post('/api/auth/request-validateOTP', { email, otp });
  }

  register(email: string, password: string, name: string) {
    return this.http.post('/api/auth/register', { email, password, name });
  }

  logout() {
    if (this.isBrowser) localStorage.removeItem('token');
    this.isLoggedIn.set(false);
  }

  getToken(): string | null {
    return this.isBrowser ? localStorage.getItem('token') : null;
  }

  getRole(): string | null {
    const token = this.getToken();
    if (!token) return null;
    try {
      const payload = JSON.parse(atob(token.split('.')[1]));
      return payload.role ?? null;
    } catch {
      return null;
    }
  }
}
