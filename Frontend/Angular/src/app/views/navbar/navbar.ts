import { Component, inject } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { MatDividerModule } from '@angular/material/divider';
import { Router, RouterLink, RouterLinkActive } from '@angular/router';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-navbar',
  imports: [
    MatIconModule,
    MatMenuModule,
    MatDividerModule,
    RouterLink,
    RouterLinkActive,
    FormsModule,
  ],
  templateUrl: './navbar.html',
  styleUrl: './navbar.scss',
})
export class Navbar {
  searchOpen = false;
  menuOpen = false;
  private router = inject(Router);

  toggleSearch() {
    this.searchOpen = !this.searchOpen;
  }

  protected showLoginPage() {
    this.router.navigate(['/auth/login']);
  }

  protected showSignUpPage() {
    this.router.navigate(['/auth/register']);
  }
}
