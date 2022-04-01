import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OutstandingordersComponent } from './outstandingorders.component';

describe('OutstandingordersComponent', () => {
  let component: OutstandingordersComponent;
  let fixture: ComponentFixture<OutstandingordersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ OutstandingordersComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(OutstandingordersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
