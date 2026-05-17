import { Component, signal, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule, FormBuilder, Validators } from '@angular/forms';
import { Router, RouterLink } from '@angular/router';
import { MatIconModule } from '@angular/material/icon';
import { AuthService } from '../../../services/auth/auth.service';

@Component({
  selector: 'app-register',
  imports: [CommonModule, ReactiveFormsModule, RouterLink, MatIconModule],
  templateUrl: './register.html',
  styleUrl: './register.scss',
})
export class Register {
  private fb = inject(FormBuilder);
  private router = inject(Router);
  private authService = inject(AuthService);

  step = signal<1 | 2 | 3>(1);
  error = signal('');
  showPw = signal(false);
  protected verifiedEmail = '';

  emailForm = this.fb.group({
    email: ['', [Validators.required, Validators.email]],
  });

  otpForm = this.fb.group({
    otp: ['', [Validators.required, Validators.minLength(6), Validators.maxLength(6)]],
  });

  registerForm = this.fb.group(
    {
      name: ['', [Validators.required, Validators.minLength(2)]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      confirmPassword: ['', Validators.required],
    },
    { validators: [this.passwordsMatchValidator] }
  );

  get passwordsMismatch(): boolean {
    return !!this.registerForm.errors?.['mismatch'];
  }

  passwordsMatchValidator(group: any) {
    const pw = group.get('password')?.value;
    const cpw = group.get('confirmPassword')?.value;
    return pw && cpw && pw !== cpw ? { mismatch: true } : null;
  }

  submitEmail() {
    if (this.emailForm.invalid) { this.emailForm.markAllAsTouched(); return; }
    const email = this.emailForm.getRawValue().email!;
    this.authService.sendOTP(email).subscribe({
      next: () => {
        this.verifiedEmail = email;
        this.error.set('');
        this.step.set(2);
      },
      error: (err: any) => this.error.set(err.error?.error ?? 'Failed to send OTP'),
    });
  }

  submitOTP() {
    if (this.otpForm.invalid) { this.otpForm.markAllAsTouched(); return; }
    const otp = this.otpForm.getRawValue().otp!;
    this.authService.validateOTP(this.verifiedEmail, otp).subscribe({
      next: () => {
        this.error.set('');
        this.step.set(3);
      },
      error: (err: any) => this.error.set(err.error?.error ?? 'Invalid OTP'),
    });
  }

  submitRegister() {
    if (this.registerForm.invalid) { this.registerForm.markAllAsTouched(); return; }
    const { name, password } = this.registerForm.getRawValue();
    this.authService.register(this.verifiedEmail, password!, name!).subscribe({
      next: () => this.router.navigate(['/auth/login']),
      error: (err: any) => this.error.set(err.error?.error ?? 'Registration failed'),
    });
  }
}
