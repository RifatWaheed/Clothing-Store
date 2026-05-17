import { Routes } from '@angular/router';
import { AboutUs } from './views/about-us/about-us';
import { PublicLayout } from './views/public-layout/public-layout';
import { Landing } from './views/landing/landing';
import { AdminShell } from './views/admin/admin-shell/admin-shell';
import { AdminDashboard } from './views/admin/admin-dashboard/admin-dashboard';
import { Orders } from './views/admin/orders/orders';
import { Products } from './views/admin/products/products';
import { Inventory } from './views/admin/inventory/inventory';
import { Register } from './views/auth/register/register';
import { Login } from './views/auth/login/login';
import { adminGuard } from './guards/admin.guard';

export const routes: Routes = [
  {
    path: '',
    component: PublicLayout,
    children: [
      { path: '', component: Landing },
      { path: 'about-us', component: AboutUs },
      { path: 'auth/register', component: Register },
      { path: 'auth/login', component: Login },
    ],
  },


  {
    path: 'admin',
    component: AdminShell,
    canActivate: [adminGuard],
    children: [
      { path: '', pathMatch: 'full', redirectTo: 'admin-dashboard' },
      { path: 'admin-dashboard', component: AdminDashboard },
      { path: 'orders', component: Orders },
      { path: 'products', component: Products },
      { path: 'inventory', component: Inventory },
    ],
  },
];
