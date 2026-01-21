import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { Navbar } from './views/navbar/navbar';
import { Landing } from "./views/landing/landing";

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, Navbar, Landing],
  templateUrl: './app.html',
  styleUrl: './app.scss'
})
export class App {
  protected readonly title = signal('Angular');
}
