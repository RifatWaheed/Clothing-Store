import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { AuthService } from '../services/auth/auth.service';

export const adminGuard: CanActivateFn = () => {
  const auth = inject(AuthService);
  const router = inject(Router);
  if (!auth.getToken()) return router.createUrlTree(['/auth/login']);
  if (auth.getRole() !== 'admin') return router.createUrlTree(['/']);
  return true;
};
