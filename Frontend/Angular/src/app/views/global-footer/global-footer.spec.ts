import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GlobalFooter } from './global-footer';

describe('GlobalFooter', () => {
  let component: GlobalFooter;
  let fixture: ComponentFixture<GlobalFooter>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GlobalFooter]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GlobalFooter);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
