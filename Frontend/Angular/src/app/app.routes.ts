import { Routes } from '@angular/router';
import { AboutUs } from './views/about-us/about-us';
import { PublicLayout } from './views/public-layout/public-layout';
import { Landing } from './views/landing/landing';
import { AdminShell } from './views/admin/admin-shell/admin-shell';
import { AdminDashboard } from './views/admin/admin-dashboard/admin-dashboard';
import { Orders } from './views/admin/orders/orders';
import { Products } from './views/admin/products/products';
import { Inventory } from './views/admin/inventory/inventory';

export const routes: Routes = [
  {
    path: '',
    component: PublicLayout,
    children: [{ path: '', component: Landing }],
  },
  { path: 'about-us', component: AboutUs },

  {
    path: 'admin',
    component: AdminShell,
    children: [
      { path: '', pathMatch: 'full', redirectTo: 'dashboard' },
      { path: 'admin-dashboard', component: AdminDashboard },
      { path: 'orders', component: Orders },
      { path: 'products', component: Products },
      { path: 'inventory', component: Inventory },
    ],
  },
];
