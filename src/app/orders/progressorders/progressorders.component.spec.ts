import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ProgressordersComponent } from './progressorders.component';

describe('ProgressordersComponent', () => {
  let component: ProgressordersComponent;
  let fixture: ComponentFixture<ProgressordersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ProgressordersComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ProgressordersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
